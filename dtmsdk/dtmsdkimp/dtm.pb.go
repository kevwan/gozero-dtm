// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: dtmsdk/dtmsdkimp/dtm.proto

package dtmsdkimp

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DtmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid           string   `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	TransType     string   `protobuf:"bytes,2,opt,name=TransType,proto3" json:"TransType,omitempty"`
	TransOptions  string   `protobuf:"bytes,3,opt,name=TransOptions,proto3" json:"TransOptions,omitempty"`
	CustomedData  string   `protobuf:"bytes,4,opt,name=CustomedData,proto3" json:"CustomedData,omitempty"`
	BinPayloads   [][]byte `protobuf:"bytes,5,rep,name=BinPayloads,proto3" json:"BinPayloads,omitempty"`     // for MSG/SAGA branch payloads
	QueryPrepared string   `protobuf:"bytes,6,opt,name=QueryPrepared,proto3" json:"QueryPrepared,omitempty"` // for MSG
	Steps         string   `protobuf:"bytes,7,opt,name=Steps,proto3" json:"Steps,omitempty"`
}

func (x *DtmRequest) Reset() {
	*x = DtmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtmsdk_dtmsdkimp_dtm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmRequest) ProtoMessage() {}

func (x *DtmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dtmsdk_dtmsdkimp_dtm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmRequest.ProtoReflect.Descriptor instead.
func (*DtmRequest) Descriptor() ([]byte, []int) {
	return file_dtmsdk_dtmsdkimp_dtm_proto_rawDescGZIP(), []int{0}
}

func (x *DtmRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *DtmRequest) GetTransType() string {
	if x != nil {
		return x.TransType
	}
	return ""
}

func (x *DtmRequest) GetTransOptions() string {
	if x != nil {
		return x.TransOptions
	}
	return ""
}

func (x *DtmRequest) GetCustomedData() string {
	if x != nil {
		return x.CustomedData
	}
	return ""
}

func (x *DtmRequest) GetBinPayloads() [][]byte {
	if x != nil {
		return x.BinPayloads
	}
	return nil
}

func (x *DtmRequest) GetQueryPrepared() string {
	if x != nil {
		return x.QueryPrepared
	}
	return ""
}

func (x *DtmRequest) GetSteps() string {
	if x != nil {
		return x.Steps
	}
	return ""
}

var File_dtmsdk_dtmsdkimp_dtm_proto protoreflect.FileDescriptor

var file_dtmsdk_dtmsdkimp_dtm_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x74, 0x6d, 0x73, 0x64, 0x6b, 0x2f, 0x64, 0x74, 0x6d, 0x73, 0x64, 0x6b, 0x69,
	0x6d, 0x70, 0x2f, 0x64, 0x74, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x74,
	0x6d, 0x73, 0x64, 0x6b, 0x69, 0x6d, 0x70, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x0a, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x47, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x69, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x0b, 0x42, 0x69, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x24, 0x0a,
	0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x65, 0x70, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x65, 0x70, 0x73, 0x32, 0x43, 0x0a, 0x06, 0x44, 0x74, 0x6d,
	0x53, 0x76, 0x63, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x15, 0x2e,
	0x64, 0x74, 0x6d, 0x73, 0x64, 0x6b, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2f,
	0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x76,
	0x77, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x7a, 0x65, 0x72, 0x6f, 0x2d, 0x64, 0x74, 0x6d, 0x2f, 0x64,
	0x74, 0x6d, 0x73, 0x64, 0x6b, 0x2f, 0x64, 0x74, 0x6d, 0x73, 0x64, 0x6b, 0x69, 0x6d, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dtmsdk_dtmsdkimp_dtm_proto_rawDescOnce sync.Once
	file_dtmsdk_dtmsdkimp_dtm_proto_rawDescData = file_dtmsdk_dtmsdkimp_dtm_proto_rawDesc
)

func file_dtmsdk_dtmsdkimp_dtm_proto_rawDescGZIP() []byte {
	file_dtmsdk_dtmsdkimp_dtm_proto_rawDescOnce.Do(func() {
		file_dtmsdk_dtmsdkimp_dtm_proto_rawDescData = protoimpl.X.CompressGZIP(file_dtmsdk_dtmsdkimp_dtm_proto_rawDescData)
	})
	return file_dtmsdk_dtmsdkimp_dtm_proto_rawDescData
}

var (
	file_dtmsdk_dtmsdkimp_dtm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
	file_dtmsdk_dtmsdkimp_dtm_proto_goTypes  = []interface{}{
		(*DtmRequest)(nil),    // 0: dtmsdkimp.DtmRequest
		(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
	}
)

var file_dtmsdk_dtmsdkimp_dtm_proto_depIdxs = []int32{
	0, // 0: dtmsdkimp.DtmSvc.Submit:input_type -> dtmsdkimp.DtmRequest
	1, // 1: dtmsdkimp.DtmSvc.Submit:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dtmsdk_dtmsdkimp_dtm_proto_init() }
func file_dtmsdk_dtmsdkimp_dtm_proto_init() {
	if File_dtmsdk_dtmsdkimp_dtm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dtmsdk_dtmsdkimp_dtm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmRequest); i {
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
			RawDescriptor: file_dtmsdk_dtmsdkimp_dtm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dtmsdk_dtmsdkimp_dtm_proto_goTypes,
		DependencyIndexes: file_dtmsdk_dtmsdkimp_dtm_proto_depIdxs,
		MessageInfos:      file_dtmsdk_dtmsdkimp_dtm_proto_msgTypes,
	}.Build()
	File_dtmsdk_dtmsdkimp_dtm_proto = out.File
	file_dtmsdk_dtmsdkimp_dtm_proto_rawDesc = nil
	file_dtmsdk_dtmsdkimp_dtm_proto_goTypes = nil
	file_dtmsdk_dtmsdkimp_dtm_proto_depIdxs = nil
}
