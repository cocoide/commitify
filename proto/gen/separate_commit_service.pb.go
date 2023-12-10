// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: separate_commit_service.proto

package src

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

// ChangeType specifies the type of
type ChangeType int32

const (
	ChangeType_UNKNOWN_CHANGE ChangeType = 0
	ChangeType_CREATE         ChangeType = 1
	ChangeType_UPDATE         ChangeType = 2
	ChangeType_DELETE         ChangeType = 3
)

// Enum value maps for ChangeType.
var (
	ChangeType_name = map[int32]string{
		0: "UNKNOWN_CHANGE",
		1: "CREATE",
		2: "UPDATE",
		3: "DELETE",
	}
	ChangeType_value = map[string]int32{
		"UNKNOWN_CHANGE": 0,
		"CREATE":         1,
		"UPDATE":         2,
		"DELETE":         3,
	}
)

func (x ChangeType) Enum() *ChangeType {
	p := new(ChangeType)
	*p = x
	return p
}

func (x ChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_separate_commit_service_proto_enumTypes[0].Descriptor()
}

func (ChangeType) Type() protoreflect.EnumType {
	return &file_separate_commit_service_proto_enumTypes[0]
}

func (x ChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeType.Descriptor instead.
func (ChangeType) EnumDescriptor() ([]byte, []int) {
	return file_separate_commit_service_proto_rawDescGZIP(), []int{0}
}

// SeparateCommitRequest is the request format for generating messages
type SeparateCommitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChanges []*FileChange  `protobuf:"bytes,1,rep,name=fileChanges,proto3" json:"fileChanges,omitempty"`
	CodeFormat  CodeFormatType `protobuf:"varint,2,opt,name=codeFormat,proto3,enum=code_type.CodeFormatType" json:"codeFormat,omitempty"`
	Language    LanguageType   `protobuf:"varint,3,opt,name=language,proto3,enum=code_type.LanguageType" json:"language,omitempty"`
}

func (x *SeparateCommitRequest) Reset() {
	*x = SeparateCommitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_separate_commit_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeparateCommitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeparateCommitRequest) ProtoMessage() {}

func (x *SeparateCommitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_separate_commit_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeparateCommitRequest.ProtoReflect.Descriptor instead.
func (*SeparateCommitRequest) Descriptor() ([]byte, []int) {
	return file_separate_commit_service_proto_rawDescGZIP(), []int{0}
}

func (x *SeparateCommitRequest) GetFileChanges() []*FileChange {
	if x != nil {
		return x.FileChanges
	}
	return nil
}

func (x *SeparateCommitRequest) GetCodeFormat() CodeFormatType {
	if x != nil {
		return x.CodeFormat
	}
	return CodeFormatType_UNKNOWN_FORMAT
}

func (x *SeparateCommitRequest) GetLanguage() LanguageType {
	if x != nil {
		return x.Language
	}
	return LanguageType_UNKNOWN_LANGUAGE
}

type LineDiff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Line  string `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
}

func (x *LineDiff) Reset() {
	*x = LineDiff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_separate_commit_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineDiff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineDiff) ProtoMessage() {}

func (x *LineDiff) ProtoReflect() protoreflect.Message {
	mi := &file_separate_commit_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineDiff.ProtoReflect.Descriptor instead.
func (*LineDiff) Descriptor() ([]byte, []int) {
	return file_separate_commit_service_proto_rawDescGZIP(), []int{1}
}

func (x *LineDiff) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *LineDiff) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

type CodeDiff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Added   []*LineDiff `protobuf:"bytes,1,rep,name=added,proto3" json:"added,omitempty"`
	Deleted []*LineDiff `protobuf:"bytes,2,rep,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *CodeDiff) Reset() {
	*x = CodeDiff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_separate_commit_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeDiff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeDiff) ProtoMessage() {}

func (x *CodeDiff) ProtoReflect() protoreflect.Message {
	mi := &file_separate_commit_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeDiff.ProtoReflect.Descriptor instead.
func (*CodeDiff) Descriptor() ([]byte, []int) {
	return file_separate_commit_service_proto_rawDescGZIP(), []int{2}
}

func (x *CodeDiff) GetAdded() []*LineDiff {
	if x != nil {
		return x.Added
	}
	return nil
}

func (x *CodeDiff) GetDeleted() []*LineDiff {
	if x != nil {
		return x.Deleted
	}
	return nil
}

type FileChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeDiff   *CodeDiff  `protobuf:"bytes,1,opt,name=codeDiff,proto3" json:"codeDiff,omitempty"`
	Filename   string     `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	ChangeType ChangeType `protobuf:"varint,3,opt,name=changeType,proto3,enum=separate_commit.ChangeType" json:"changeType,omitempty"`
}

func (x *FileChange) Reset() {
	*x = FileChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_separate_commit_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChange) ProtoMessage() {}

func (x *FileChange) ProtoReflect() protoreflect.Message {
	mi := &file_separate_commit_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChange.ProtoReflect.Descriptor instead.
func (*FileChange) Descriptor() ([]byte, []int) {
	return file_separate_commit_service_proto_rawDescGZIP(), []int{3}
}

func (x *FileChange) GetCodeDiff() *CodeDiff {
	if x != nil {
		return x.CodeDiff
	}
	return nil
}

func (x *FileChange) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileChange) GetChangeType() ChangeType {
	if x != nil {
		return x.ChangeType
	}
	return ChangeType_UNKNOWN_CHANGE
}

// SeparateCommitResponse returns generated and separated commit messages
type SeparateCommitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeparatedCommits []*SeparatedCommitMessages `protobuf:"bytes,1,rep,name=separatedCommits,proto3" json:"separatedCommits,omitempty"`
}

func (x *SeparateCommitResponse) Reset() {
	*x = SeparateCommitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_separate_commit_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeparateCommitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeparateCommitResponse) ProtoMessage() {}

func (x *SeparateCommitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_separate_commit_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeparateCommitResponse.ProtoReflect.Descriptor instead.
func (*SeparateCommitResponse) Descriptor() ([]byte, []int) {
	return file_separate_commit_service_proto_rawDescGZIP(), []int{4}
}

func (x *SeparateCommitResponse) GetSeparatedCommits() []*SeparatedCommitMessages {
	if x != nil {
		return x.SeparatedCommits
	}
	return nil
}

type SeparatedCommitMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages   []string   `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	Filename   string     `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	ChangeType ChangeType `protobuf:"varint,3,opt,name=changeType,proto3,enum=separate_commit.ChangeType" json:"changeType,omitempty"`
}

func (x *SeparatedCommitMessages) Reset() {
	*x = SeparatedCommitMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_separate_commit_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeparatedCommitMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeparatedCommitMessages) ProtoMessage() {}

func (x *SeparatedCommitMessages) ProtoReflect() protoreflect.Message {
	mi := &file_separate_commit_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeparatedCommitMessages.ProtoReflect.Descriptor instead.
func (*SeparatedCommitMessages) Descriptor() ([]byte, []int) {
	return file_separate_commit_service_proto_rawDescGZIP(), []int{5}
}

func (x *SeparatedCommitMessages) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *SeparatedCommitMessages) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *SeparatedCommitMessages) GetChangeType() ChangeType {
	if x != nil {
		return x.ChangeType
	}
	return ChangeType_UNKNOWN_CHANGE
}

var File_separate_commit_service_proto protoreflect.FileDescriptor

var file_separate_commit_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x1a, 0x0f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0b, 0x66,
	0x69, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0b, 0x66,
	0x69, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x6f,
	0x64, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x34, 0x0a, 0x08, 0x4c, 0x69,
	0x6e, 0x65, 0x44, 0x69, 0x66, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65,
	0x22, 0x70, 0x0a, 0x08, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x66, 0x66, 0x12, 0x2f, 0x0a, 0x05,
	0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69,
	0x6e, 0x65, 0x44, 0x69, 0x66, 0x66, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x33, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x44, 0x69, 0x66, 0x66, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x66, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x66, 0x66, 0x52, 0x08,
	0x63, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x66, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x70, 0x61, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x6e, 0x0a, 0x16, 0x53, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x73,
	0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x10, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x73, 0x22, 0x8e, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x70, 0x61,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x2a, 0x44, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x32, 0x89, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x70,
	0x61, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x70, 0x0a, 0x1d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x65,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x53, 0x65,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_separate_commit_service_proto_rawDescOnce sync.Once
	file_separate_commit_service_proto_rawDescData = file_separate_commit_service_proto_rawDesc
)

func file_separate_commit_service_proto_rawDescGZIP() []byte {
	file_separate_commit_service_proto_rawDescOnce.Do(func() {
		file_separate_commit_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_separate_commit_service_proto_rawDescData)
	})
	return file_separate_commit_service_proto_rawDescData
}

var file_separate_commit_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_separate_commit_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_separate_commit_service_proto_goTypes = []interface{}{
	(ChangeType)(0),                 // 0: separate_commit.ChangeType
	(*SeparateCommitRequest)(nil),   // 1: separate_commit.SeparateCommitRequest
	(*LineDiff)(nil),                // 2: separate_commit.LineDiff
	(*CodeDiff)(nil),                // 3: separate_commit.CodeDiff
	(*FileChange)(nil),              // 4: separate_commit.FileChange
	(*SeparateCommitResponse)(nil),  // 5: separate_commit.SeparateCommitResponse
	(*SeparatedCommitMessages)(nil), // 6: separate_commit.SeparatedCommitMessages
	(CodeFormatType)(0),             // 7: code_type.CodeFormatType
	(LanguageType)(0),               // 8: code_type.LanguageType
}
var file_separate_commit_service_proto_depIdxs = []int32{
	4,  // 0: separate_commit.SeparateCommitRequest.fileChanges:type_name -> separate_commit.FileChange
	7,  // 1: separate_commit.SeparateCommitRequest.codeFormat:type_name -> code_type.CodeFormatType
	8,  // 2: separate_commit.SeparateCommitRequest.language:type_name -> code_type.LanguageType
	2,  // 3: separate_commit.CodeDiff.added:type_name -> separate_commit.LineDiff
	2,  // 4: separate_commit.CodeDiff.deleted:type_name -> separate_commit.LineDiff
	3,  // 5: separate_commit.FileChange.codeDiff:type_name -> separate_commit.CodeDiff
	0,  // 6: separate_commit.FileChange.changeType:type_name -> separate_commit.ChangeType
	6,  // 7: separate_commit.SeparateCommitResponse.separatedCommits:type_name -> separate_commit.SeparatedCommitMessages
	0,  // 8: separate_commit.SeparatedCommitMessages.changeType:type_name -> separate_commit.ChangeType
	1,  // 9: separate_commit.SeparateCommitService.GenerateMultipleCommitMessage:input_type -> separate_commit.SeparateCommitRequest
	5,  // 10: separate_commit.SeparateCommitService.GenerateMultipleCommitMessage:output_type -> separate_commit.SeparateCommitResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_separate_commit_service_proto_init() }
func file_separate_commit_service_proto_init() {
	if File_separate_commit_service_proto != nil {
		return
	}
	file_code_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_separate_commit_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeparateCommitRequest); i {
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
		file_separate_commit_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineDiff); i {
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
		file_separate_commit_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeDiff); i {
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
		file_separate_commit_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileChange); i {
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
		file_separate_commit_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeparateCommitResponse); i {
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
		file_separate_commit_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeparatedCommitMessages); i {
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
			RawDescriptor: file_separate_commit_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_separate_commit_service_proto_goTypes,
		DependencyIndexes: file_separate_commit_service_proto_depIdxs,
		EnumInfos:         file_separate_commit_service_proto_enumTypes,
		MessageInfos:      file_separate_commit_service_proto_msgTypes,
	}.Build()
	File_separate_commit_service_proto = out.File
	file_separate_commit_service_proto_rawDesc = nil
	file_separate_commit_service_proto_goTypes = nil
	file_separate_commit_service_proto_depIdxs = nil
}
