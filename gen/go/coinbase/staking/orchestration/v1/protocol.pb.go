// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: coinbase/staking/orchestration/v1/protocol.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A Protocol resource (e.g. ethereum_kiln, solana etc.).
type Protocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the Protocol.
	// Format: protocols/{protocolName}
	// Ex: protocols/ethereum_kiln
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Protocol) Reset() {
	*x = Protocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_orchestration_v1_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Protocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Protocol) ProtoMessage() {}

func (x *Protocol) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_orchestration_v1_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Protocol.ProtoReflect.Descriptor instead.
func (*Protocol) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_orchestration_v1_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *Protocol) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request message for ListProtocols.
type ListProtocolsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProtocolsRequest) Reset() {
	*x = ListProtocolsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_orchestration_v1_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProtocolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProtocolsRequest) ProtoMessage() {}

func (x *ListProtocolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_orchestration_v1_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProtocolsRequest.ProtoReflect.Descriptor instead.
func (*ListProtocolsRequest) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_orchestration_v1_protocol_proto_rawDescGZIP(), []int{1}
}

// The response message for ListProtocols.
type ListProtocolsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of protocols.
	Protocols []*Protocol `protobuf:"bytes,1,rep,name=protocols,proto3" json:"protocols,omitempty"`
}

func (x *ListProtocolsResponse) Reset() {
	*x = ListProtocolsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_orchestration_v1_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProtocolsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProtocolsResponse) ProtoMessage() {}

func (x *ListProtocolsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_orchestration_v1_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProtocolsResponse.ProtoReflect.Descriptor instead.
func (*ListProtocolsResponse) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_orchestration_v1_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *ListProtocolsResponse) GetProtocols() []*Protocol {
	if x != nil {
		return x.Protocols
	}
	return nil
}

var File_coinbase_staking_orchestration_v1_protocol_proto protoreflect.FileDescriptor

var file_coinbase_staking_orchestration_v1_protocol_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6d, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x3a, 0x4d, 0xea, 0x41, 0x4a, 0x0a, 0x1d, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x7d, 0x2a, 0x09, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x32, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22,
	0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x49, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x42, 0x48, 0x5a, 0x46, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x62, 0x68, 0x71, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_staking_orchestration_v1_protocol_proto_rawDescOnce sync.Once
	file_coinbase_staking_orchestration_v1_protocol_proto_rawDescData = file_coinbase_staking_orchestration_v1_protocol_proto_rawDesc
)

func file_coinbase_staking_orchestration_v1_protocol_proto_rawDescGZIP() []byte {
	file_coinbase_staking_orchestration_v1_protocol_proto_rawDescOnce.Do(func() {
		file_coinbase_staking_orchestration_v1_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_staking_orchestration_v1_protocol_proto_rawDescData)
	})
	return file_coinbase_staking_orchestration_v1_protocol_proto_rawDescData
}

var file_coinbase_staking_orchestration_v1_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coinbase_staking_orchestration_v1_protocol_proto_goTypes = []interface{}{
	(*Protocol)(nil),              // 0: coinbase.staking.orchestration.v1.Protocol
	(*ListProtocolsRequest)(nil),  // 1: coinbase.staking.orchestration.v1.ListProtocolsRequest
	(*ListProtocolsResponse)(nil), // 2: coinbase.staking.orchestration.v1.ListProtocolsResponse
}
var file_coinbase_staking_orchestration_v1_protocol_proto_depIdxs = []int32{
	0, // 0: coinbase.staking.orchestration.v1.ListProtocolsResponse.protocols:type_name -> coinbase.staking.orchestration.v1.Protocol
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_coinbase_staking_orchestration_v1_protocol_proto_init() }
func file_coinbase_staking_orchestration_v1_protocol_proto_init() {
	if File_coinbase_staking_orchestration_v1_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_staking_orchestration_v1_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Protocol); i {
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
		file_coinbase_staking_orchestration_v1_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProtocolsRequest); i {
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
		file_coinbase_staking_orchestration_v1_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProtocolsResponse); i {
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
			RawDescriptor: file_coinbase_staking_orchestration_v1_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_staking_orchestration_v1_protocol_proto_goTypes,
		DependencyIndexes: file_coinbase_staking_orchestration_v1_protocol_proto_depIdxs,
		MessageInfos:      file_coinbase_staking_orchestration_v1_protocol_proto_msgTypes,
	}.Build()
	File_coinbase_staking_orchestration_v1_protocol_proto = out.File
	file_coinbase_staking_orchestration_v1_protocol_proto_rawDesc = nil
	file_coinbase_staking_orchestration_v1_protocol_proto_goTypes = nil
	file_coinbase_staking_orchestration_v1_protocol_proto_depIdxs = nil
}
