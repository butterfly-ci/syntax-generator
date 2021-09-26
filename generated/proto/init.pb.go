// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: init.proto

package proto

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

//
// Init is the config block that runs on every pipeline to define what it is
type Init struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo            *Init_Repo        `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	BaseImage       string            `protobuf:"bytes,2,opt,name=base_image,json=baseImage,proto3" json:"base_image,omitempty"`
	DockerRepo      string            `protobuf:"bytes,3,opt,name=docker_repo,json=dockerRepo,proto3" json:"docker_repo,omitempty"`
	DockerPush      bool              `protobuf:"varint,4,opt,name=docker_push,json=dockerPush,proto3" json:"docker_push,omitempty"`
	DockerPushLocal bool              `protobuf:"varint,5,opt,name=docker_push_local,json=dockerPushLocal,proto3" json:"docker_push_local,omitempty"`
	Env             map[string]string `protobuf:"bytes,6,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Secrets         map[string]string `protobuf:"bytes,7,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Commands        []string          `protobuf:"bytes,9,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *Init) Reset() {
	*x = Init{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Init) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Init) ProtoMessage() {}

func (x *Init) ProtoReflect() protoreflect.Message {
	mi := &file_init_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Init.ProtoReflect.Descriptor instead.
func (*Init) Descriptor() ([]byte, []int) {
	return file_init_proto_rawDescGZIP(), []int{0}
}

func (x *Init) GetRepo() *Init_Repo {
	if x != nil {
		return x.Repo
	}
	return nil
}

func (x *Init) GetBaseImage() string {
	if x != nil {
		return x.BaseImage
	}
	return ""
}

func (x *Init) GetDockerRepo() string {
	if x != nil {
		return x.DockerRepo
	}
	return ""
}

func (x *Init) GetDockerPush() bool {
	if x != nil {
		return x.DockerPush
	}
	return false
}

func (x *Init) GetDockerPushLocal() bool {
	if x != nil {
		return x.DockerPushLocal
	}
	return false
}

func (x *Init) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *Init) GetSecrets() map[string]string {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *Init) GetCommands() []string {
	if x != nil {
		return x.Commands
	}
	return nil
}

type Init_Repo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link         string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	BranchFilter string `protobuf:"bytes,2,opt,name=branch_filter,json=branchFilter,proto3" json:"branch_filter,omitempty"`
}

func (x *Init_Repo) Reset() {
	*x = Init_Repo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Init_Repo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Init_Repo) ProtoMessage() {}

func (x *Init_Repo) ProtoReflect() protoreflect.Message {
	mi := &file_init_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Init_Repo.ProtoReflect.Descriptor instead.
func (*Init_Repo) Descriptor() ([]byte, []int) {
	return file_init_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Init_Repo) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Init_Repo) GetBranchFilter() string {
	if x != nil {
		return x.BranchFilter
	}
	return ""
}

var File_init_proto protoreflect.FileDescriptor

var file_init_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x03, 0x0a,
	0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52,
	0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f,
	0x70, 0x75, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x5f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x03, 0x65, 0x6e, 0x76, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x2e, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x1a, 0x3f,
	0x0a, 0x04, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a,
	0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_init_proto_rawDescOnce sync.Once
	file_init_proto_rawDescData = file_init_proto_rawDesc
)

func file_init_proto_rawDescGZIP() []byte {
	file_init_proto_rawDescOnce.Do(func() {
		file_init_proto_rawDescData = protoimpl.X.CompressGZIP(file_init_proto_rawDescData)
	})
	return file_init_proto_rawDescData
}

var file_init_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_init_proto_goTypes = []interface{}{
	(*Init)(nil),      // 0: Init
	(*Init_Repo)(nil), // 1: Init.Repo
	nil,               // 2: Init.EnvEntry
	nil,               // 3: Init.SecretsEntry
}
var file_init_proto_depIdxs = []int32{
	1, // 0: Init.repo:type_name -> Init.Repo
	2, // 1: Init.env:type_name -> Init.EnvEntry
	3, // 2: Init.secrets:type_name -> Init.SecretsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_init_proto_init() }
func file_init_proto_init() {
	if File_init_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_init_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Init); i {
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
		file_init_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Init_Repo); i {
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
			RawDescriptor: file_init_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_init_proto_goTypes,
		DependencyIndexes: file_init_proto_depIdxs,
		MessageInfos:      file_init_proto_msgTypes,
	}.Build()
	File_init_proto = out.File
	file_init_proto_rawDesc = nil
	file_init_proto_goTypes = nil
	file_init_proto_depIdxs = nil
}