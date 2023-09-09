// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: commit_message_service.proto

package pb

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

// CodeFormatType specifies the type of commit message to generate
type CodeFormatType int32

const (
	CodeFormatType_EMOJI  CodeFormatType = 0
	CodeFormatType_PREFIX CodeFormatType = 1
	CodeFormatType_NORMAL CodeFormatType = 2
)

// Enum value maps for CodeFormatType.
var (
	CodeFormatType_name = map[int32]string{
		0: "EMOJI",
		1: "PREFIX",
		2: "NORMAL",
	}
	CodeFormatType_value = map[string]int32{
		"EMOJI":  0,
		"PREFIX": 1,
		"NORMAL": 2,
	}
)

func (x CodeFormatType) Enum() *CodeFormatType {
	p := new(CodeFormatType)
	*p = x
	return p
}

func (x CodeFormatType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodeFormatType) Descriptor() protoreflect.EnumDescriptor {
	return file_commit_message_service_proto_enumTypes[0].Descriptor()
}

func (CodeFormatType) Type() protoreflect.EnumType {
	return &file_commit_message_service_proto_enumTypes[0]
}

func (x CodeFormatType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodeFormatType.Descriptor instead.
func (CodeFormatType) EnumDescriptor() ([]byte, []int) {
	return file_commit_message_service_proto_rawDescGZIP(), []int{0}
}

// LanguageType specifies the language of the commit message to generate
type LanguageType int32

const (
	LanguageType_ENGLISH  LanguageType = 0
	LanguageType_JAPANESE LanguageType = 1
)

// Enum value maps for LanguageType.
var (
	LanguageType_name = map[int32]string{
		0: "ENGLISH",
		1: "JAPANESE",
	}
	LanguageType_value = map[string]int32{
		"ENGLISH":  0,
		"JAPANESE": 1,
	}
)

func (x LanguageType) Enum() *LanguageType {
	p := new(LanguageType)
	*p = x
	return p
}

func (x LanguageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LanguageType) Descriptor() protoreflect.EnumDescriptor {
	return file_commit_message_service_proto_enumTypes[1].Descriptor()
}

func (LanguageType) Type() protoreflect.EnumType {
	return &file_commit_message_service_proto_enumTypes[1]
}

func (x LanguageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LanguageType.Descriptor instead.
func (LanguageType) EnumDescriptor() ([]byte, []int) {
	return file_commit_message_service_proto_rawDescGZIP(), []int{1}
}

// CommitMessageRequest is the request format for generating messages
type CommitMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputCode  string         `protobuf:"bytes,1,opt,name=inputCode,proto3" json:"inputCode,omitempty"`
	CodeFormat CodeFormatType `protobuf:"varint,2,opt,name=codeFormat,proto3,enum=commit_message.CodeFormatType" json:"codeFormat,omitempty"`
	Language   LanguageType   `protobuf:"varint,3,opt,name=language,proto3,enum=commit_message.LanguageType" json:"language,omitempty"`
}

func (x *CommitMessageRequest) Reset() {
	*x = CommitMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commit_message_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitMessageRequest) ProtoMessage() {}

func (x *CommitMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commit_message_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitMessageRequest.ProtoReflect.Descriptor instead.
func (*CommitMessageRequest) Descriptor() ([]byte, []int) {
	return file_commit_message_service_proto_rawDescGZIP(), []int{0}
}

func (x *CommitMessageRequest) GetInputCode() string {
	if x != nil {
		return x.InputCode
	}
	return ""
}

func (x *CommitMessageRequest) GetCodeFormat() CodeFormatType {
	if x != nil {
		return x.CodeFormat
	}
	return CodeFormatType_EMOJI
}

func (x *CommitMessageRequest) GetLanguage() LanguageType {
	if x != nil {
		return x.Language
	}
	return LanguageType_ENGLISH
}

// CommitMessageResponse returns generated commit messages
type CommitMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []string `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *CommitMessageResponse) Reset() {
	*x = CommitMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commit_message_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitMessageResponse) ProtoMessage() {}

func (x *CommitMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commit_message_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitMessageResponse.ProtoReflect.Descriptor instead.
func (*CommitMessageResponse) Descriptor() ([]byte, []int) {
	return file_commit_message_service_proto_rawDescGZIP(), []int{1}
}

func (x *CommitMessageResponse) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_commit_message_service_proto protoreflect.FileDescriptor

var file_commit_message_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xae,
	0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22,
	0x33, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2a, 0x33, 0x0a, 0x0e, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x4f, 0x4a, 0x49, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x02, 0x2a, 0x29, 0x0a, 0x0c, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x47,
	0x4c, 0x49, 0x53, 0x48, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4a, 0x41, 0x50, 0x41, 0x4e, 0x45,
	0x53, 0x45, 0x10, 0x01, 0x32, 0x7c, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x15,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commit_message_service_proto_rawDescOnce sync.Once
	file_commit_message_service_proto_rawDescData = file_commit_message_service_proto_rawDesc
)

func file_commit_message_service_proto_rawDescGZIP() []byte {
	file_commit_message_service_proto_rawDescOnce.Do(func() {
		file_commit_message_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_commit_message_service_proto_rawDescData)
	})
	return file_commit_message_service_proto_rawDescData
}

var file_commit_message_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_commit_message_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commit_message_service_proto_goTypes = []interface{}{
	(CodeFormatType)(0),           // 0: commit_message.CodeFormatType
	(LanguageType)(0),             // 1: commit_message.LanguageType
	(*CommitMessageRequest)(nil),  // 2: commit_message.CommitMessageRequest
	(*CommitMessageResponse)(nil), // 3: commit_message.CommitMessageResponse
}
var file_commit_message_service_proto_depIdxs = []int32{
	0, // 0: commit_message.CommitMessageRequest.codeFormat:type_name -> commit_message.CodeFormatType
	1, // 1: commit_message.CommitMessageRequest.language:type_name -> commit_message.LanguageType
	2, // 2: commit_message.CommitMessageService.GenerateCommitMessage:input_type -> commit_message.CommitMessageRequest
	3, // 3: commit_message.CommitMessageService.GenerateCommitMessage:output_type -> commit_message.CommitMessageResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_commit_message_service_proto_init() }
func file_commit_message_service_proto_init() {
	if File_commit_message_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commit_message_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitMessageRequest); i {
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
		file_commit_message_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitMessageResponse); i {
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
			RawDescriptor: file_commit_message_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commit_message_service_proto_goTypes,
		DependencyIndexes: file_commit_message_service_proto_depIdxs,
		EnumInfos:         file_commit_message_service_proto_enumTypes,
		MessageInfos:      file_commit_message_service_proto_msgTypes,
	}.Build()
	File_commit_message_service_proto = out.File
	file_commit_message_service_proto_rawDesc = nil
	file_commit_message_service_proto_goTypes = nil
	file_commit_message_service_proto_depIdxs = nil
}
