// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/user.proto

package user

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateUserInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string               `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName string               `protobuf:"bytes,2,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName   string               `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Nickname   string               `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email      string               `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password   string               `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	Birthdate  *timestamp.Timestamp `protobuf:"bytes,6,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
}

func (x *CreateUserInput) Reset() {
	*x = CreateUserInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserInput) ProtoMessage() {}

func (x *CreateUserInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserInput.ProtoReflect.Descriptor instead.
func (*CreateUserInput) Descriptor() ([]byte, []int) {
	return file_protos_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserInput) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateUserInput) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *CreateUserInput) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateUserInput) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *CreateUserInput) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserInput) GetBirthdate() *timestamp.Timestamp {
	if x != nil {
		return x.Birthdate
	}
	return nil
}

type UpdateUserInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string           `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserData *CreateUserInput `protobuf:"bytes,2,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
}

func (x *UpdateUserInput) Reset() {
	*x = UpdateUserInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInput) ProtoMessage() {}

func (x *UpdateUserInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInput.ProtoReflect.Descriptor instead.
func (*UpdateUserInput) Descriptor() ([]byte, []int) {
	return file_protos_user_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateUserInput) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateUserInput) GetUserData() *CreateUserInput {
	if x != nil {
		return x.UserData
	}
	return nil
}

type DeleteUserInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DeleteUserInput) Reset() {
	*x = DeleteUserInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserInput) ProtoMessage() {}

func (x *DeleteUserInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserInput.ProtoReflect.Descriptor instead.
func (*DeleteUserInput) Descriptor() ([]byte, []int) {
	return file_protos_user_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteUserInput) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_protos_user_proto protoreflect.FileDescriptor

var file_protos_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x38, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x22, 0x56, 0x0a, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x2d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x27, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xaf, 0x01, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x10, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x11, 0x5a,
	0x0f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_user_proto_rawDescOnce sync.Once
	file_protos_user_proto_rawDescData = file_protos_user_proto_rawDesc
)

func file_protos_user_proto_rawDescGZIP() []byte {
	file_protos_user_proto_rawDescOnce.Do(func() {
		file_protos_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_user_proto_rawDescData)
	})
	return file_protos_user_proto_rawDescData
}

var file_protos_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_user_proto_goTypes = []interface{}{
	(*CreateUserInput)(nil),     // 0: CreateUserInput
	(*UpdateUserInput)(nil),     // 1: UpdateUserInput
	(*DeleteUserInput)(nil),     // 2: DeleteUserInput
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_protos_user_proto_depIdxs = []int32{
	3, // 0: CreateUserInput.birthdate:type_name -> google.protobuf.Timestamp
	0, // 1: UpdateUserInput.user_data:type_name -> CreateUserInput
	0, // 2: UserService.Create:input_type -> CreateUserInput
	1, // 3: UserService.Update:input_type -> UpdateUserInput
	2, // 4: UserService.Delete:input_type -> DeleteUserInput
	4, // 5: UserService.Create:output_type -> google.protobuf.Empty
	4, // 6: UserService.Update:output_type -> google.protobuf.Empty
	4, // 7: UserService.Delete:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_user_proto_init() }
func file_protos_user_proto_init() {
	if File_protos_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserInput); i {
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
		file_protos_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInput); i {
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
		file_protos_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserInput); i {
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
			RawDescriptor: file_protos_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_user_proto_goTypes,
		DependencyIndexes: file_protos_user_proto_depIdxs,
		MessageInfos:      file_protos_user_proto_msgTypes,
	}.Build()
	File_protos_user_proto = out.File
	file_protos_user_proto_rawDesc = nil
	file_protos_user_proto_goTypes = nil
	file_protos_user_proto_depIdxs = nil
}