// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: genbank.proto

package gbparse

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

type Genbank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LOCUS      string       `protobuf:"bytes,1,opt,name=LOCUS,proto3" json:"LOCUS,omitempty"`
	ACCESSION  []string     `protobuf:"bytes,2,rep,name=ACCESSION,proto3" json:"ACCESSION,omitempty"`
	DEFINITION string       `protobuf:"bytes,3,opt,name=DEFINITION,proto3" json:"DEFINITION,omitempty"`
	VERSION    string       `protobuf:"bytes,4,opt,name=VERSION,proto3" json:"VERSION,omitempty"`
	DBLINK     []string     `protobuf:"bytes,5,rep,name=DBLINK,proto3" json:"DBLINK,omitempty"`
	KEYWORDS   string       `protobuf:"bytes,6,opt,name=KEYWORDS,proto3" json:"KEYWORDS,omitempty"`
	SOURCE     string       `protobuf:"bytes,7,opt,name=SOURCE,proto3" json:"SOURCE,omitempty"`
	ORGANISM   []string     `protobuf:"bytes,8,rep,name=ORGANISM,proto3" json:"ORGANISM,omitempty"`
	COMMENT    string       `protobuf:"bytes,9,opt,name=COMMENT,proto3" json:"COMMENT,omitempty"`
	SEQUENCE   string       `protobuf:"bytes,10,opt,name=SEQUENCE,proto3" json:"SEQUENCE,omitempty"`
	CONTIG     string       `protobuf:"bytes,11,opt,name=CONTIG,proto3" json:"CONTIG,omitempty"`
	REFERENCES []*Reference `protobuf:"bytes,12,rep,name=REFERENCES,proto3" json:"REFERENCES,omitempty"`
	FEATURES   []*Feature   `protobuf:"bytes,13,rep,name=FEATURES,proto3" json:"FEATURES,omitempty"`
}

func (x *Genbank) Reset() {
	*x = Genbank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genbank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genbank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genbank) ProtoMessage() {}

func (x *Genbank) ProtoReflect() protoreflect.Message {
	mi := &file_genbank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genbank.ProtoReflect.Descriptor instead.
func (*Genbank) Descriptor() ([]byte, []int) {
	return file_genbank_proto_rawDescGZIP(), []int{0}
}

func (x *Genbank) GetLOCUS() string {
	if x != nil {
		return x.LOCUS
	}
	return ""
}

func (x *Genbank) GetACCESSION() []string {
	if x != nil {
		return x.ACCESSION
	}
	return nil
}

func (x *Genbank) GetDEFINITION() string {
	if x != nil {
		return x.DEFINITION
	}
	return ""
}

func (x *Genbank) GetVERSION() string {
	if x != nil {
		return x.VERSION
	}
	return ""
}

func (x *Genbank) GetDBLINK() []string {
	if x != nil {
		return x.DBLINK
	}
	return nil
}

func (x *Genbank) GetKEYWORDS() string {
	if x != nil {
		return x.KEYWORDS
	}
	return ""
}

func (x *Genbank) GetSOURCE() string {
	if x != nil {
		return x.SOURCE
	}
	return ""
}

func (x *Genbank) GetORGANISM() []string {
	if x != nil {
		return x.ORGANISM
	}
	return nil
}

func (x *Genbank) GetCOMMENT() string {
	if x != nil {
		return x.COMMENT
	}
	return ""
}

func (x *Genbank) GetSEQUENCE() string {
	if x != nil {
		return x.SEQUENCE
	}
	return ""
}

func (x *Genbank) GetCONTIG() string {
	if x != nil {
		return x.CONTIG
	}
	return ""
}

func (x *Genbank) GetREFERENCES() []*Reference {
	if x != nil {
		return x.REFERENCES
	}
	return nil
}

func (x *Genbank) GetFEATURES() []*Feature {
	if x != nil {
		return x.FEATURES
	}
	return nil
}

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number  int32  `protobuf:"varint,1,opt,name=Number,proto3" json:"Number,omitempty"`
	ORIGIN  string `protobuf:"bytes,2,opt,name=ORIGIN,proto3" json:"ORIGIN,omitempty"`
	AUTHORS string `protobuf:"bytes,3,opt,name=AUTHORS,proto3" json:"AUTHORS,omitempty"`
	CONSRTM string `protobuf:"bytes,4,opt,name=CONSRTM,proto3" json:"CONSRTM,omitempty"`
	TITLE   string `protobuf:"bytes,5,opt,name=TITLE,proto3" json:"TITLE,omitempty"`
	JOURNAL string `protobuf:"bytes,6,opt,name=JOURNAL,proto3" json:"JOURNAL,omitempty"`
	PUBMED  string `protobuf:"bytes,7,opt,name=PUBMED,proto3" json:"PUBMED,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genbank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_genbank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_genbank_proto_rawDescGZIP(), []int{1}
}

func (x *Reference) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Reference) GetORIGIN() string {
	if x != nil {
		return x.ORIGIN
	}
	return ""
}

func (x *Reference) GetAUTHORS() string {
	if x != nil {
		return x.AUTHORS
	}
	return ""
}

func (x *Reference) GetCONSRTM() string {
	if x != nil {
		return x.CONSRTM
	}
	return ""
}

func (x *Reference) GetTITLE() string {
	if x != nil {
		return x.TITLE
	}
	return ""
}

func (x *Reference) GetJOURNAL() string {
	if x != nil {
		return x.JOURNAL
	}
	return ""
}

func (x *Reference) GetPUBMED() string {
	if x != nil {
		return x.PUBMED
	}
	return ""
}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TYPE         string       `protobuf:"bytes,1,opt,name=TYPE,proto3" json:"TYPE,omitempty"`
	IsCompliment bool         `protobuf:"varint,2,opt,name=isCompliment,proto3" json:"isCompliment,omitempty"`
	IsJoined     bool         `protobuf:"varint,3,opt,name=isJoined,proto3" json:"isJoined,omitempty"`
	START        string       `protobuf:"bytes,4,opt,name=START,proto3" json:"START,omitempty"`
	STOP         string       `protobuf:"bytes,5,opt,name=STOP,proto3" json:"STOP,omitempty"`
	QUALIFIERS   []*Qualifier `protobuf:"bytes,6,rep,name=QUALIFIERS,proto3" json:"QUALIFIERS,omitempty"`
	ID           string       `protobuf:"bytes,7,opt,name=ID,proto3" json:"ID,omitempty"`
	ACCESSION    string       `protobuf:"bytes,8,opt,name=ACCESSION,proto3" json:"ACCESSION,omitempty"`
	ACCESSION_ID string       `protobuf:"bytes,9,opt,name=ACCESSION_ID,json=ACCESSIONID,proto3" json:"ACCESSION_ID,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genbank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_genbank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_genbank_proto_rawDescGZIP(), []int{2}
}

func (x *Feature) GetTYPE() string {
	if x != nil {
		return x.TYPE
	}
	return ""
}

func (x *Feature) GetIsCompliment() bool {
	if x != nil {
		return x.IsCompliment
	}
	return false
}

func (x *Feature) GetIsJoined() bool {
	if x != nil {
		return x.IsJoined
	}
	return false
}

func (x *Feature) GetSTART() string {
	if x != nil {
		return x.START
	}
	return ""
}

func (x *Feature) GetSTOP() string {
	if x != nil {
		return x.STOP
	}
	return ""
}

func (x *Feature) GetQUALIFIERS() []*Qualifier {
	if x != nil {
		return x.QUALIFIERS
	}
	return nil
}

func (x *Feature) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Feature) GetACCESSION() string {
	if x != nil {
		return x.ACCESSION
	}
	return ""
}

func (x *Feature) GetACCESSION_ID() string {
	if x != nil {
		return x.ACCESSION_ID
	}
	return ""
}

type Qualifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Qualifier) Reset() {
	*x = Qualifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genbank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Qualifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Qualifier) ProtoMessage() {}

func (x *Qualifier) ProtoReflect() protoreflect.Message {
	mi := &file_genbank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Qualifier.ProtoReflect.Descriptor instead.
func (*Qualifier) Descriptor() ([]byte, []int) {
	return file_genbank_proto_rawDescGZIP(), []int{3}
}

func (x *Qualifier) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Qualifier) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_genbank_proto protoreflect.FileDescriptor

var file_genbank_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x62, 0x70, 0x61, 0x72, 0x73, 0x65, 0x22, 0x8f, 0x03, 0x0a, 0x07, 0x47, 0x65, 0x6e,
	0x62, 0x61, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x55, 0x53, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x4f, 0x43, 0x55, 0x53, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x41,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x45, 0x46, 0x49,
	0x4e, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x45,
	0x46, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x42, 0x4c, 0x49, 0x4e, 0x4b, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x44, 0x42, 0x4c, 0x49, 0x4e, 0x4b, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x45,
	0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x45,
	0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x12, 0x1a,
	0x0a, 0x08, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x53, 0x4d, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x53, 0x4d, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x4f,
	0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x4f, 0x4d,
	0x4d, 0x45, 0x4e, 0x54, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45,
	0x12, 0x16, 0x0a, 0x06, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x47, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x47, 0x12, 0x32, 0x0a, 0x0a, 0x52, 0x45, 0x46, 0x45,
	0x52, 0x45, 0x4e, 0x43, 0x45, 0x53, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x62, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x0a, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x53, 0x12, 0x2c, 0x0a, 0x08,
	0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x53, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x67, 0x62, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x08, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x53, 0x22, 0xb7, 0x01, 0x0a, 0x09, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x55, 0x54, 0x48,
	0x4f, 0x52, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x55, 0x54, 0x48, 0x4f,
	0x52, 0x53, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x53, 0x52, 0x54, 0x4d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x4f, 0x4e, 0x53, 0x52, 0x54, 0x4d, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x49, 0x54, 0x4c, 0x45, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x49, 0x54,
	0x4c, 0x45, 0x12, 0x18, 0x0a, 0x07, 0x4a, 0x4f, 0x55, 0x52, 0x4e, 0x41, 0x4c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4a, 0x4f, 0x55, 0x52, 0x4e, 0x41, 0x4c, 0x12, 0x16, 0x0a, 0x06,
	0x50, 0x55, 0x42, 0x4d, 0x45, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x55,
	0x42, 0x4d, 0x45, 0x44, 0x22, 0x8c, 0x02, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x59, 0x50, 0x45, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x54, 0x59, 0x50, 0x45, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4a, 0x6f,
	0x69, 0x6e, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4a, 0x6f,
	0x69, 0x6e, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x52, 0x54, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x54, 0x41, 0x52, 0x54, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x54,
	0x4f, 0x50, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x54, 0x4f, 0x50, 0x12, 0x32,
	0x0a, 0x0a, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x46, 0x49, 0x45, 0x52, 0x53, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x62, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x46, 0x49, 0x45,
	0x52, 0x53, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x12, 0x21, 0x0a, 0x0c, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x49, 0x44, 0x22, 0x33, 0x0a, 0x09, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x67, 0x62,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_genbank_proto_rawDescOnce sync.Once
	file_genbank_proto_rawDescData = file_genbank_proto_rawDesc
)

func file_genbank_proto_rawDescGZIP() []byte {
	file_genbank_proto_rawDescOnce.Do(func() {
		file_genbank_proto_rawDescData = protoimpl.X.CompressGZIP(file_genbank_proto_rawDescData)
	})
	return file_genbank_proto_rawDescData
}

var file_genbank_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_genbank_proto_goTypes = []interface{}{
	(*Genbank)(nil),   // 0: gbparse.Genbank
	(*Reference)(nil), // 1: gbparse.Reference
	(*Feature)(nil),   // 2: gbparse.Feature
	(*Qualifier)(nil), // 3: gbparse.Qualifier
}
var file_genbank_proto_depIdxs = []int32{
	1, // 0: gbparse.Genbank.REFERENCES:type_name -> gbparse.Reference
	2, // 1: gbparse.Genbank.FEATURES:type_name -> gbparse.Feature
	3, // 2: gbparse.Feature.QUALIFIERS:type_name -> gbparse.Qualifier
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_genbank_proto_init() }
func file_genbank_proto_init() {
	if File_genbank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_genbank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genbank); i {
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
		file_genbank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
		file_genbank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
		file_genbank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Qualifier); i {
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
			RawDescriptor: file_genbank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_genbank_proto_goTypes,
		DependencyIndexes: file_genbank_proto_depIdxs,
		MessageInfos:      file_genbank_proto_msgTypes,
	}.Build()
	File_genbank_proto = out.File
	file_genbank_proto_rawDesc = nil
	file_genbank_proto_goTypes = nil
	file_genbank_proto_depIdxs = nil
}
