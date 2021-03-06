// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: fasta.proto

package bioproto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Fasta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HEADER    string `protobuf:"bytes,1,opt,name=HEADER,proto3" json:"HEADER,omitempty"`
	ACCESSION string `protobuf:"bytes,2,opt,name=ACCESSION,proto3" json:"ACCESSION,omitempty"`
	VERSION   string `protobuf:"bytes,3,opt,name=VERSION,proto3" json:"VERSION,omitempty"`
	SEQUENCE  string `protobuf:"bytes,4,opt,name=SEQUENCE,proto3" json:"SEQUENCE,omitempty"`
}

func (x *Fasta) Reset() {
	*x = Fasta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fasta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fasta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fasta) ProtoMessage() {}

func (x *Fasta) ProtoReflect() protoreflect.Message {
	mi := &file_fasta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fasta.ProtoReflect.Descriptor instead.
func (*Fasta) Descriptor() ([]byte, []int) {
	return file_fasta_proto_rawDescGZIP(), []int{0}
}

func (x *Fasta) GetHEADER() string {
	if x != nil {
		return x.HEADER
	}
	return ""
}

func (x *Fasta) GetACCESSION() string {
	if x != nil {
		return x.ACCESSION
	}
	return ""
}

func (x *Fasta) GetVERSION() string {
	if x != nil {
		return x.VERSION
	}
	return ""
}

func (x *Fasta) GetSEQUENCE() string {
	if x != nil {
		return x.SEQUENCE
	}
	return ""
}

var File_fasta_proto protoreflect.FileDescriptor

var file_fasta_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x61, 0x73, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62,
	0x69, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x05, 0x46, 0x61, 0x73, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x3b, 0x62, 0x69, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_fasta_proto_rawDescOnce sync.Once
	file_fasta_proto_rawDescData = file_fasta_proto_rawDesc
)

func file_fasta_proto_rawDescGZIP() []byte {
	file_fasta_proto_rawDescOnce.Do(func() {
		file_fasta_proto_rawDescData = protoimpl.X.CompressGZIP(file_fasta_proto_rawDescData)
	})
	return file_fasta_proto_rawDescData
}

var file_fasta_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fasta_proto_goTypes = []interface{}{
	(*Fasta)(nil), // 0: bioproto.Fasta
}
var file_fasta_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fasta_proto_init() }
func file_fasta_proto_init() {
	if File_fasta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fasta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fasta); i {
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
			RawDescriptor: file_fasta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fasta_proto_goTypes,
		DependencyIndexes: file_fasta_proto_depIdxs,
		MessageInfos:      file_fasta_proto_msgTypes,
	}.Build()
	File_fasta_proto = out.File
	file_fasta_proto_rawDesc = nil
	file_fasta_proto_goTypes = nil
	file_fasta_proto_depIdxs = nil
}
