package libp2p

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	BlockedPeerCount             prometheus.Counter
	FailedIncomingHandshakeCount prometheus.Counter
	FailedOutgoingHandshakeCount prometheus.Counter
}

func newMetrics(registry prometheus.Registerer, namespace string) *metrics {
	subsystem := "libp2p"

	m := &metrics{
		BlockedPeerCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "blocked_peer_count",
			Help:      "Number of blocked peers.",
		}),
		FailedIncomingHandshakeCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "failed_incoming_handshake_count",
			Help:      "Number of failed incoming handshake count.",
		}),
		FailedOutgoingHandshakeCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "failed_outgoing_handshake_count",
			Help:      "Number of failed outgoing handshake count.",
		}),
	}

	registry.MustRegister(
		m.BlockedPeerCount,
		m.FailedIncomingHandshakeCount,
		m.FailedOutgoingHandshakeCount,
	)

	return m
}
