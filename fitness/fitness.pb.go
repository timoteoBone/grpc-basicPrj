// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: fitness.proto

package grpc_basicPrj

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

type UserData_Gender int32

const (
	UserData_MALE  UserData_Gender = 0
	UserData_WOMAN UserData_Gender = 1
)

// Enum value maps for UserData_Gender.
var (
	UserData_Gender_name = map[int32]string{
		0: "MALE",
		1: "WOMAN",
	}
	UserData_Gender_value = map[string]int32{
		"MALE":  0,
		"WOMAN": 1,
	}
)

func (x UserData_Gender) Enum() *UserData_Gender {
	p := new(UserData_Gender)
	*p = x
	return p
}

func (x UserData_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserData_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_fitness_proto_enumTypes[0].Descriptor()
}

func (UserData_Gender) Type() protoreflect.EnumType {
	return &file_fitness_proto_enumTypes[0]
}

func (x UserData_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserData_Gender.Descriptor instead.
func (UserData_Gender) EnumDescriptor() ([]byte, []int) {
	return file_fitness_proto_rawDescGZIP(), []int{0, 0}
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weight float32 `protobuf:"fixed32,1,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Heigt  float32 `protobuf:"fixed32,2,opt,name=Heigt,proto3" json:"Heigt,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fitness_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_fitness_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_fitness_proto_rawDescGZIP(), []int{0}
}

func (x *UserData) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *UserData) GetHeigt() float32 {
	if x != nil {
		return x.Heigt
	}
	return 0
}

type DietResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Basis    string `protobuf:"bytes,1,opt,name=Basis,proto3" json:"Basis,omitempty"`
	Calories int64  `protobuf:"varint,2,opt,name=Calories,proto3" json:"Calories,omitempty"`
}

func (x *DietResponse) Reset() {
	*x = DietResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fitness_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DietResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DietResponse) ProtoMessage() {}

func (x *DietResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fitness_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DietResponse.ProtoReflect.Descriptor instead.
func (*DietResponse) Descriptor() ([]byte, []int) {
	return file_fitness_proto_rawDescGZIP(), []int{1}
}

func (x *DietResponse) GetBasis() string {
	if x != nil {
		return x.Basis
	}
	return ""
}

func (x *DietResponse) GetCalories() int64 {
	if x != nil {
		return x.Calories
	}
	return 0
}

var File_fitness_proto protoreflect.FileDescriptor

var file_fitness_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x48, 0x65,
	0x69, 0x67, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x48, 0x65, 0x69, 0x67, 0x74,
	0x22, 0x1d, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41,
	0x4c, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x4f, 0x4d, 0x41, 0x4e, 0x10, 0x01, 0x22,
	0x40, 0x0a, 0x0c, 0x44, 0x69, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x42, 0x61, 0x73, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x42, 0x61, 0x73, 0x69, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x32, 0x43, 0x0a, 0x0c, 0x4e, 0x75, 0x74, 0x72, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x69, 0x73,
	0x74, 0x12, 0x33, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6d, 0x6f, 0x74, 0x65, 0x6f, 0x42, 0x6f, 0x6e, 0x65,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x61, 0x73, 0x69, 0x63, 0x50, 0x72, 0x6a, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fitness_proto_rawDescOnce sync.Once
	file_fitness_proto_rawDescData = file_fitness_proto_rawDesc
)

func file_fitness_proto_rawDescGZIP() []byte {
	file_fitness_proto_rawDescOnce.Do(func() {
		file_fitness_proto_rawDescData = protoimpl.X.CompressGZIP(file_fitness_proto_rawDescData)
	})
	return file_fitness_proto_rawDescData
}

var file_fitness_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fitness_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fitness_proto_goTypes = []interface{}{
	(UserData_Gender)(0), // 0: proto.UserData.Gender
	(*UserData)(nil),     // 1: proto.UserData
	(*DietResponse)(nil), // 2: proto.DietResponse
}
var file_fitness_proto_depIdxs = []int32{
	1, // 0: proto.Nutricionist.SendRoutine:input_type -> proto.UserData
	2, // 1: proto.Nutricionist.SendRoutine:output_type -> proto.DietResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fitness_proto_init() }
func file_fitness_proto_init() {
	if File_fitness_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fitness_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_fitness_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DietResponse); i {
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
			RawDescriptor: file_fitness_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fitness_proto_goTypes,
		DependencyIndexes: file_fitness_proto_depIdxs,
		EnumInfos:         file_fitness_proto_enumTypes,
		MessageInfos:      file_fitness_proto_msgTypes,
	}.Build()
	File_fitness_proto = out.File
	file_fitness_proto_rawDesc = nil
	file_fitness_proto_goTypes = nil
	file_fitness_proto_depIdxs = nil
}
