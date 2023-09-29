package node

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/primevprotocol/mev-commit/pkg/apiserver"
	"github.com/primevprotocol/mev-commit/pkg/debugapi"
	"github.com/primevprotocol/mev-commit/pkg/discovery"
	"github.com/primevprotocol/mev-commit/pkg/p2p"
	"github.com/primevprotocol/mev-commit/pkg/p2p/libp2p"
	"github.com/primevprotocol/mev-commit/pkg/register"
	"github.com/primevprotocol/mev-commit/pkg/topology"
)

type Options struct {
	Version   string
	PrivKey   *ecdsa.PrivateKey
	Secret    string
	PeerType  string
	Logger    *slog.Logger
	P2PPort   int
	HTTPPort  int
	Bootnodes []string
}

type Node struct {
	closers []io.Closer
}

func NewNode(opts *Options) (*Node, error) {
	reg := register.New()

	minStake, err := reg.GetMinimumStake()
	if err != nil {
		return nil, err
	}

	srv := apiserver.New(opts.Version, opts.Logger)

	closers := make([]io.Closer, 0)
	peerType := p2p.FromString(opts.PeerType)

	p2pSvc, err := libp2p.New(&libp2p.Options{
		PrivKey:        opts.PrivKey,
		Secret:         opts.Secret,
		PeerType:       peerType,
		Register:       reg,
		MinimumStake:   minStake,
		Logger:         opts.Logger,
		ListenPort:     opts.P2PPort,
		MetricsReg:     srv.MetricsRegistry(),
		BootstrapAddrs: opts.Bootnodes,
	})
	if err != nil {
		return nil, err
	}
	closers = append(closers, p2pSvc)

	topo := topology.New(p2pSvc, opts.Logger)
	disc := discovery.New(topo, p2pSvc, opts.Logger)
	closers = append(closers, disc)

	// Set the announcer for the topology service
	topo.SetAnnouncer(disc)
	// Set the notifier for the p2p service
	p2pSvc.SetNotifier(topo)

	// Register the discovery protocol with the p2p service
	p2pSvc.AddProtocol(disc.Protocol())

	debugapi.RegisterAPI(srv, topo, p2pSvc, opts.Logger)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", opts.HTTPPort),
		Handler: srv.Router(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			opts.Logger.Error("failed to start server", "err", err)
		}
	}()
	closers = append(closers, server)

	return &Node{closers: closers}, nil
}

func (n *Node) Close() error {
	var err error
	for _, c := range n.closers {
		err = errors.Join(err, c.Close())
	}

	return err
}
