// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: test.proto

package test

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

type Person_PhoneType int32

const (
	Person_mobile Person_PhoneType = 0
	Person_home   Person_PhoneType = 1
	Person_work   Person_PhoneType = 2
)

// Enum value maps for Person_PhoneType.
var (
	Person_PhoneType_name = map[int32]string{
		0: "mobile",
		1: "home",
		2: "work",
	}
	Person_PhoneType_value = map[string]int32{
		"mobile": 0,
		"home":   1,
		"work":   2,
	}
)

func (x Person_PhoneType) Enum() *Person_PhoneType {
	p := new(Person_PhoneType)
	*p = x
	return p
}

func (x Person_PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Person_PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[0].Descriptor()
}

func (Person_PhoneType) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[0]
}

func (x Person_PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Person_PhoneType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Person_PhoneType(num)
	return nil
}

// Deprecated: Use Person_PhoneType.Descriptor instead.
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0, 0}
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   *string               `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Id     *int32                `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Email  *string               `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Person) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Person) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *Person) GetPhones() []*Person_PhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

type AddressBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (x *AddressBook) Reset() {
	*x = AddressBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBook) ProtoMessage() {}

func (x *AddressBook) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressBook.ProtoReflect.Descriptor instead.
func (*AddressBook) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *AddressBook) GetPeople() []*Person {
	if x != nil {
		return x.People
	}
	return nil
}

type Person_PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number *string           `protobuf:"bytes,1,req,name=number" json:"number,omitempty"`
	Type   *Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=Person_PhoneType,def=1" json:"type,omitempty"`
}

// Default values for Person_PhoneNumber fields.
const (
	Default_Person_PhoneNumber_Type = Person_home
)

func (x *Person_PhoneNumber) Reset() {
	*x = Person_PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_PhoneNumber) ProtoMessage() {}

func (x *Person_PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_PhoneNumber.ProtoReflect.Descriptor instead.
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Person_PhoneNumber) GetNumber() string {
	if x != nil && x.Number != nil {
		return *x.Number
	}
	return ""
}

func (x *Person_PhoneNumber) GetType() Person_PhoneType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_Person_PhoneNumber_Type
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a,
	0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x1a, 0x52,
	0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x2b, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x68,
	0x6f, 0x6d, 0x65, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x6b, 0x10, 0x02, 0x22,
	0x2e, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1f,
	0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x74, 0x65, 0x73, 0x74,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_test_proto_goTypes = []interface{}{
	(Person_PhoneType)(0),      // 0: Person.PhoneType
	(*Person)(nil),             // 1: Person
	(*AddressBook)(nil),        // 2: AddressBook
	(*Person_PhoneNumber)(nil), // 3: Person.PhoneNumber
}
var file_test_proto_depIdxs = []int32{
	3, // 0: Person.phones:type_name -> Person.PhoneNumber
	1, // 1: AddressBook.people:type_name -> Person
	0, // 2: Person.PhoneNumber.type:type_name -> Person.PhoneType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressBook); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_PhoneNumber); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		EnumInfos:         file_test_proto_enumTypes,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
