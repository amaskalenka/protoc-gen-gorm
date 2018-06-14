// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/user/user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	example/user/user.proto

It has these top-level messages:
	User
	Email
	Address
	Language
	CreditCard
	Task
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id              int32                       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	CreatedAt       *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt       *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Birthday        *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=birthday" json:"birthday,omitempty"`
	Age             uint32                      `protobuf:"varint,5,opt,name=age" json:"age,omitempty"`
	Num             uint32                      `protobuf:"varint,6,opt,name=num" json:"num,omitempty"`
	CreditCard      *CreditCard                 `protobuf:"bytes,7,opt,name=credit_card,json=creditCard" json:"credit_card,omitempty"`
	Emails          []*Email                    `protobuf:"bytes,8,rep,name=emails" json:"emails,omitempty"`
	Tasks           []*Task                     `protobuf:"bytes,9,rep,name=tasks" json:"tasks,omitempty"`
	BillingAddress  *Address                    `protobuf:"bytes,10,opt,name=billing_address,json=billingAddress" json:"billing_address,omitempty"`
	ShippingAddress *Address                    `protobuf:"bytes,11,opt,name=shipping_address,json=shippingAddress" json:"shipping_address,omitempty"`
	Languages       []*Language                 `protobuf:"bytes,12,rep,name=languages" json:"languages,omitempty"`
	Friends         []*User                     `protobuf:"bytes,13,rep,name=friends" json:"friends,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetBirthday() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *User) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *User) GetCreditCard() *CreditCard {
	if m != nil {
		return m.CreditCard
	}
	return nil
}

func (m *User) GetEmails() []*Email {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *User) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *User) GetBillingAddress() *Address {
	if m != nil {
		return m.BillingAddress
	}
	return nil
}

func (m *User) GetShippingAddress() *Address {
	if m != nil {
		return m.ShippingAddress
	}
	return nil
}

func (m *User) GetLanguages() []*Language {
	if m != nil {
		return m.Languages
	}
	return nil
}

func (m *User) GetFriends() []*User {
	if m != nil {
		return m.Friends
	}
	return nil
}

type Email struct {
	Id         int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Email      string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Subscribed bool   `protobuf:"varint,3,opt,name=subscribed" json:"subscribed,omitempty"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Email) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Email) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Email) GetSubscribed() bool {
	if m != nil {
		return m.Subscribed
	}
	return false
}

type Address struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Address_1 string `protobuf:"bytes,2,opt,name=address_1,json=address1" json:"address_1,omitempty"`
	Address_2 string `protobuf:"bytes,3,opt,name=address_2,json=address2" json:"address_2,omitempty"`
	Post      string `protobuf:"bytes,4,opt,name=post" json:"post,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Address) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Address) GetAddress_1() string {
	if m != nil {
		return m.Address_1
	}
	return ""
}

func (m *Address) GetAddress_2() string {
	if m != nil {
		return m.Address_2
	}
	return ""
}

func (m *Address) GetPost() string {
	if m != nil {
		return m.Post
	}
	return ""
}

type Language struct {
	Id   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Code string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
}

func (m *Language) Reset()                    { *m = Language{} }
func (m *Language) String() string            { return proto.CompactTextString(m) }
func (*Language) ProtoMessage()               {}
func (*Language) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Language) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Language) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Language) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type CreditCard struct {
	Id        int32                       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Number    string                      `protobuf:"bytes,4,opt,name=number" json:"number,omitempty"`
}

func (m *CreditCard) Reset()                    { *m = CreditCard{} }
func (m *CreditCard) String() string            { return proto.CompactTextString(m) }
func (*CreditCard) ProtoMessage()               {}
func (*CreditCard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreditCard) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreditCard) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *CreditCard) GetUpdatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *CreditCard) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type Task struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Priority    int64  `protobuf:"varint,3,opt,name=priority" json:"priority,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Task) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Email)(nil), "user.Email")
	proto.RegisterType((*Address)(nil), "user.Address")
	proto.RegisterType((*Language)(nil), "user.Language")
	proto.RegisterType((*CreditCard)(nil), "user.CreditCard")
	proto.RegisterType((*Task)(nil), "user.Task")
}

func init() { proto.RegisterFile("example/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xae, 0x13, 0x27, 0x75, 0x4e, 0xfa, 0x13, 0xc4, 0xd8, 0xdc, 0x0c, 0xb6, 0x90, 0xdd, 0x84,
	0xb1, 0x26, 0x34, 0x1b, 0x83, 0xb6, 0x30, 0xd6, 0x96, 0xdd, 0xf5, 0x4a, 0x74, 0x0c, 0x76, 0x13,
	0x64, 0x4b, 0x75, 0x45, 0x6d, 0xcb, 0x93, 0x64, 0x68, 0xdf, 0x6b, 0x37, 0xcd, 0x33, 0xed, 0x21,
	0x86, 0x25, 0xd9, 0x31, 0x2b, 0xac, 0xbb, 0xdc, 0x8d, 0x39, 0xfa, 0xce, 0x77, 0xbe, 0xf3, 0xe3,
	0x23, 0xc1, 0x0b, 0x76, 0x47, 0xb2, 0x22, 0x65, 0x8b, 0x52, 0x31, 0x69, 0x3e, 0xf3, 0x42, 0x0a,
	0x2d, 0x90, 0x5f, 0xd9, 0xe3, 0x93, 0x84, 0xeb, 0x9b, 0x32, 0x9a, 0xc7, 0x22, 0x5b, 0xf0, 0xfc,
	0x5a, 0x44, 0xa9, 0xb8, 0x13, 0x05, 0xcb, 0x17, 0x86, 0x14, 0x1f, 0x26, 0x2c, 0x3f, 0x4c, 0x84,
	0xcc, 0x16, 0xa2, 0xd0, 0x5c, 0xe4, 0x6a, 0x51, 0x1d, 0xac, 0xc2, 0xf8, 0x75, 0x22, 0x44, 0x92,
	0x32, 0x4b, 0x8d, 0xca, 0xeb, 0x85, 0xe6, 0x19, 0x53, 0x9a, 0x64, 0x85, 0x25, 0x4c, 0x7f, 0xf9,
	0xe0, 0x7f, 0x55, 0x4c, 0xa2, 0x3d, 0xe8, 0x70, 0x1a, 0x7a, 0x13, 0x6f, 0xd6, 0xc3, 0x1d, 0x4e,
	0xd1, 0x31, 0x40, 0x2c, 0x19, 0xd1, 0x8c, 0xae, 0x88, 0x0e, 0x3b, 0x13, 0x6f, 0x36, 0x5c, 0x8e,
	0xe7, 0x56, 0x6e, 0x5e, 0xcb, 0xcd, 0xaf, 0x6a, 0x39, 0x3c, 0x70, 0xec, 0x33, 0x5d, 0x85, 0x96,
	0x05, 0xad, 0x43, 0xbb, 0x4f, 0x87, 0x3a, 0xf6, 0x99, 0x46, 0x1f, 0x21, 0x88, 0xb8, 0xd4, 0x37,
	0x94, 0xdc, 0x87, 0xfe, 0x93, 0x81, 0x0d, 0x17, 0x85, 0xd0, 0x25, 0x09, 0x0b, 0x7b, 0x13, 0x6f,
	0xb6, 0x7b, 0xde, 0x5f, 0x3f, 0x1c, 0x74, 0x46, 0x1e, 0xae, 0x20, 0x34, 0x82, 0x6e, 0x5e, 0x66,
	0x61, 0xbf, 0xf2, 0xe0, 0xca, 0x44, 0x47, 0x30, 0x8c, 0x25, 0xa3, 0x5c, 0xaf, 0x62, 0x22, 0x69,
	0xb8, 0x6d, 0xd2, 0x8c, 0xe6, 0x66, 0xee, 0x17, 0xc6, 0x71, 0x41, 0x24, 0xc5, 0x10, 0x37, 0x36,
	0x7a, 0x03, 0x7d, 0x96, 0x11, 0x9e, 0xaa, 0x30, 0x98, 0x74, 0x67, 0xc3, 0xe5, 0xd0, 0xb2, 0xbf,
	0x54, 0x18, 0x76, 0x2e, 0xb4, 0x84, 0x9e, 0x26, 0xea, 0x56, 0x85, 0x03, 0xc3, 0x01, 0xcb, 0xb9,
	0x22, 0xea, 0xf6, 0x7c, 0xb4, 0x7e, 0x38, 0xd8, 0x79, 0x0b, 0xd3, 0xa0, 0x90, 0x5c, 0x48, 0xae,
	0xef, 0xb1, 0xa5, 0xa2, 0x4f, 0xb0, 0x1f, 0xf1, 0x34, 0xe5, 0x79, 0xb2, 0x22, 0x94, 0x4a, 0xa6,
	0x54, 0x08, 0xa6, 0x9e, 0x5d, 0x1b, 0x7d, 0x66, 0x41, 0xdb, 0xd2, 0x74, 0x0b, 0xef, 0x39, 0xb6,
	0xc3, 0xd1, 0x67, 0x18, 0xa9, 0x1b, 0x5e, 0x14, 0x6d, 0x81, 0xe1, 0xdf, 0x04, 0xf6, 0x6b, 0x7a,
	0xad, 0xf0, 0x01, 0x06, 0x29, 0xc9, 0x93, 0x92, 0x24, 0x4c, 0x85, 0x3b, 0xa6, 0xf2, 0x3d, 0x1b,
	0x7a, 0xe9, 0x60, 0x1b, 0xbb, 0xdc, 0xc2, 0x1b, 0x22, 0x7a, 0x07, 0xdb, 0xd7, 0x92, 0xb3, 0x9c,
	0xaa, 0x70, 0xb7, 0xdd, 0x6d, 0xb5, 0x4a, 0x0d, 0xbf, 0xa6, 0x9c, 0x04, 0xeb, 0x87, 0x03, 0x3f,
	0xf0, 0x26, 0xde, 0xf4, 0x1b, 0xf4, 0xcc, 0xd0, 0x1e, 0xad, 0xdb, 0x33, 0xe8, 0x99, 0x31, 0x9a,
	0x4d, 0x1b, 0x60, 0x7b, 0x40, 0xaf, 0x00, 0x54, 0x19, 0xa9, 0x58, 0xf2, 0x88, 0x51, 0xb3, 0x49,
	0x01, 0x6e, 0x21, 0x2d, 0xe1, 0x1f, 0xb0, 0x5d, 0x77, 0xf4, 0xa7, 0xf4, 0x4b, 0x18, 0xb8, 0xd1,
	0xac, 0x8e, 0x9c, 0x7c, 0xe0, 0x80, 0xa3, 0xb6, 0x73, 0x69, 0x12, 0x6c, 0x9c, 0x4b, 0x84, 0xc0,
	0x2f, 0x84, 0xd2, 0x66, 0x13, 0x07, 0xd8, 0xd8, 0xad, 0x94, 0x97, 0x10, 0xd4, 0x23, 0x7a, 0x94,
	0x13, 0x81, 0x9f, 0x93, 0x8c, 0xb9, 0x74, 0xc6, 0xae, 0xb0, 0x58, 0x50, 0xe6, 0xb2, 0x18, 0xbb,
	0xa5, 0xf6, 0xd3, 0x03, 0xd8, 0x6c, 0xdf, 0x7f, 0x72, 0x1d, 0x9f, 0x43, 0x3f, 0x2f, 0xb3, 0x88,
	0x49, 0x37, 0x02, 0x77, 0x6a, 0x95, 0x1d, 0x81, 0x5f, 0x6d, 0x78, 0xd3, 0xb0, 0xd7, 0x6a, 0x78,
	0x02, 0x43, 0xca, 0xaa, 0x5f, 0x65, 0xde, 0x25, 0x37, 0x8b, 0x36, 0x84, 0xc6, 0xd0, 0xdc, 0x08,
	0x53, 0x58, 0x17, 0x37, 0xe7, 0x4d, 0x8e, 0xf3, 0xd3, 0xef, 0xc7, 0xff, 0xfa, 0x04, 0xb6, 0x5f,
	0xd2, 0xd3, 0xea, 0x13, 0xf5, 0x0d, 0xe5, 0xfd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x5c,
	0xa1, 0x65, 0x65, 0x05, 0x00, 0x00,
}
