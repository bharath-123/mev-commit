// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: rpc/searcherapi/v1/searcherapi.proto

package searcherapiv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxnHash     string `protobuf:"bytes,1,opt,name=txn_hash,json=txnHash,proto3" json:"txn_hash,omitempty"`
	BidAmt      int64  `protobuf:"varint,2,opt,name=bid_amt,json=bidAmt,proto3" json:"bid_amt,omitempty"`
	BlockNumber int64  `protobuf:"varint,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_searcherapi_v1_searcherapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_searcherapi_v1_searcherapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_rpc_searcherapi_v1_searcherapi_proto_rawDescGZIP(), []int{0}
}

func (x *Bid) GetTxnHash() string {
	if x != nil {
		return x.TxnHash
	}
	return ""
}

func (x *Bid) GetBidAmt() int64 {
	if x != nil {
		return x.BidAmt
	}
	return 0
}

func (x *Bid) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

type Commitment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid                 *Bid   `protobuf:"bytes,1,opt,name=bid,proto3" json:"bid,omitempty"`
	BidHash             []byte `protobuf:"bytes,4,opt,name=bid_hash,json=bidHash,proto3" json:"bid_hash,omitempty"`
	Signature           []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	DataHash            []byte `protobuf:"bytes,6,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	CommitmentSignature []byte `protobuf:"bytes,7,opt,name=commitment_signature,json=commitmentSignature,proto3" json:"commitment_signature,omitempty"`
}

func (x *Commitment) Reset() {
	*x = Commitment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_searcherapi_v1_searcherapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commitment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commitment) ProtoMessage() {}

func (x *Commitment) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_searcherapi_v1_searcherapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commitment.ProtoReflect.Descriptor instead.
func (*Commitment) Descriptor() ([]byte, []int) {
	return file_rpc_searcherapi_v1_searcherapi_proto_rawDescGZIP(), []int{1}
}

func (x *Commitment) GetBid() *Bid {
	if x != nil {
		return x.Bid
	}
	return nil
}

func (x *Commitment) GetBidHash() []byte {
	if x != nil {
		return x.BidHash
	}
	return nil
}

func (x *Commitment) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Commitment) GetDataHash() []byte {
	if x != nil {
		return x.DataHash
	}
	return nil
}

func (x *Commitment) GetCommitmentSignature() []byte {
	if x != nil {
		return x.CommitmentSignature
	}
	return nil
}

var File_rpc_searcherapi_v1_searcherapi_proto protoreflect.FileDescriptor

var file_rpc_searcherapi_v1_searcherapi_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x5c, 0x0a, 0x03, 0x42, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x62,
	0x69, 0x64, 0x5f, 0x61, 0x6d, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x69,
	0x64, 0x41, 0x6d, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xbf, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x03, 0x62, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x62, 0x69, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0x50, 0x0a, 0x08, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x69, 0x64,
	0x12, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x64, 0x1a, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x46, 0x5a, 0x44, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x76,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6d, 0x65, 0x76, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_searcherapi_v1_searcherapi_proto_rawDescOnce sync.Once
	file_rpc_searcherapi_v1_searcherapi_proto_rawDescData = file_rpc_searcherapi_v1_searcherapi_proto_rawDesc
)

func file_rpc_searcherapi_v1_searcherapi_proto_rawDescGZIP() []byte {
	file_rpc_searcherapi_v1_searcherapi_proto_rawDescOnce.Do(func() {
		file_rpc_searcherapi_v1_searcherapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_searcherapi_v1_searcherapi_proto_rawDescData)
	})
	return file_rpc_searcherapi_v1_searcherapi_proto_rawDescData
}

var file_rpc_searcherapi_v1_searcherapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_searcherapi_v1_searcherapi_proto_goTypes = []interface{}{
	(*Bid)(nil),        // 0: rpc.seacherapi.v1.Bid
	(*Commitment)(nil), // 1: rpc.seacherapi.v1.Commitment
}
var file_rpc_searcherapi_v1_searcherapi_proto_depIdxs = []int32{
	0, // 0: rpc.seacherapi.v1.Commitment.bid:type_name -> rpc.seacherapi.v1.Bid
	0, // 1: rpc.seacherapi.v1.Searcher.SendBid:input_type -> rpc.seacherapi.v1.Bid
	1, // 2: rpc.seacherapi.v1.Searcher.SendBid:output_type -> rpc.seacherapi.v1.Commitment
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_searcherapi_v1_searcherapi_proto_init() }
func file_rpc_searcherapi_v1_searcherapi_proto_init() {
	if File_rpc_searcherapi_v1_searcherapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_searcherapi_v1_searcherapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_searcherapi_v1_searcherapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commitment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_searcherapi_v1_searcherapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_searcherapi_v1_searcherapi_proto_goTypes,
		DependencyIndexes: file_rpc_searcherapi_v1_searcherapi_proto_depIdxs,
		MessageInfos:      file_rpc_searcherapi_v1_searcherapi_proto_msgTypes,
	}.Build()
	File_rpc_searcherapi_v1_searcherapi_proto = out.File
	file_rpc_searcherapi_v1_searcherapi_proto_rawDesc = nil
	file_rpc_searcherapi_v1_searcherapi_proto_goTypes = nil
	file_rpc_searcherapi_v1_searcherapi_proto_depIdxs = nil
}
