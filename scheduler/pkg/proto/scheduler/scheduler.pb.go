// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: scheduler/pkg/proto/scheduler/scheduler.proto

package scheduler

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{0}
}

type EventID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventID int64 `protobuf:"varint,1,opt,name=EventID,proto3" json:"EventID,omitempty"`
}

func (x *EventID) Reset() {
	*x = EventID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventID) ProtoMessage() {}

func (x *EventID) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventID.ProtoReflect.Descriptor instead.
func (*EventID) Descriptor() ([]byte, []int) {
	return file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *EventID) GetEventID() int64 {
	if x != nil {
		return x.EventID
	}
	return 0
}

type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int32 `protobuf:"varint,1,opt,name=Year,proto3" json:"Year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=Month,proto3" json:"Month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=Day,proto3" json:"Day,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{2}
}

func (x *Date) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Date) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

type CustomerID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CustomerID) Reset() {
	*x = CustomerID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerID) ProtoMessage() {}

func (x *CustomerID) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerID.ProtoReflect.Descriptor instead.
func (*CustomerID) Descriptor() ([]byte, []int) {
	return file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{3}
}

func (x *CustomerID) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type CustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerID int64 `protobuf:"varint,1,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Since      int64 `protobuf:"varint,2,opt,name=Since,proto3" json:"Since,omitempty"`
	Days       int32 `protobuf:"varint,3,opt,name=Days,proto3" json:"Days,omitempty"`
}

func (x *CustomerRequest) Reset() {
	*x = CustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerRequest) ProtoMessage() {}

func (x *CustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerRequest.ProtoReflect.Descriptor instead.
func (*CustomerRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{4}
}

func (x *CustomerRequest) GetCustomerID() int64 {
	if x != nil {
		return x.CustomerID
	}
	return 0
}

func (x *CustomerRequest) GetSince() int64 {
	if x != nil {
		return x.Since
	}
	return 0
}

func (x *CustomerRequest) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventID  int64 `protobuf:"varint,1,opt,name=EventID,proto3" json:"EventID,omitempty"`
	Date     int64 `protobuf:"varint,2,opt,name=Date,proto3" json:"Date,omitempty"`
	WithNext bool  `protobuf:"varint,3,opt,name=WithNext,proto3" json:"WithNext,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRequest) GetEventID() int64 {
	if x != nil {
		return x.EventID
	}
	return 0
}

func (x *DeleteRequest) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *DeleteRequest) GetWithNext() bool {
	if x != nil {
		return x.WithNext
	}
	return false
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date     int64  `protobuf:"varint,1,opt,name=Date,proto3" json:"Date,omitempty"`
	WithNext bool   `protobuf:"varint,2,opt,name=WithNext,proto3" json:"WithNext,omitempty"`
	Event    *Event `protobuf:"bytes,3,opt,name=Event,proto3" json:"Event,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRequest) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *UpdateRequest) GetWithNext() bool {
	if x != nil {
		return x.WithNext
	}
	return false
}

func (x *UpdateRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CustomerID       int64  `protobuf:"varint,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	StartTime        string `protobuf:"bytes,3,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime          string `protobuf:"bytes,4,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	Description      string `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	IsRecurring      bool   `protobuf:"varint,6,opt,name=IsRecurring,proto3" json:"IsRecurring,omitempty"`
	Period           string `protobuf:"bytes,7,opt,name=Period,proto3" json:"Period,omitempty"`
	RecurringEndTime string `protobuf:"bytes,8,opt,name=RecurringEndTime,proto3" json:"RecurringEndTime,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{7}
}

func (x *Event) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Event) GetCustomerID() int64 {
	if x != nil {
		return x.CustomerID
	}
	return 0
}

func (x *Event) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *Event) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetIsRecurring() bool {
	if x != nil {
		return x.IsRecurring
	}
	return false
}

func (x *Event) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Event) GetRecurringEndTime() string {
	if x != nil {
		return x.RecurringEndTime
	}
	return ""
}

type Events struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty"`
}

func (x *Events) Reset() {
	*x = Events{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Events) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events) ProtoMessage() {}

func (x *Events) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events.ProtoReflect.Descriptor instead.
func (*Events) Descriptor() ([]byte, []int) {
	return file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{8}
}

func (x *Events) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_scheduler_pkg_proto_scheduler_scheduler_proto protoreflect.FileDescriptor

var file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x23, 0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x42, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x59, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x61,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x44, 0x61, 0x79, 0x22, 0x1c, 0x0a, 0x0a,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x5b, 0x0a, 0x0f, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x69,
	0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x44, 0x61, 0x79, 0x73, 0x22, 0x59, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x65,
	0x78, 0x74, 0x22, 0x67, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x4e,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x57, 0x69, 0x74, 0x68, 0x4e,
	0x65, 0x78, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xf7, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x49, 0x73, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x73, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x63,
	0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x52, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x28, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xb0, 0x03, 0x0a, 0x09, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x12, 0x0f, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x11, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x00, 0x12, 0x47, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f,
	0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x15, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x43, 0x5a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x2d,
	0x62, 0x79, 0x2f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x2d, 0x62, 0x75, 0x73, 0x79, 0x2d, 0x62,
	0x61, 0x63, 0x6b, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescOnce sync.Once
	file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescData = file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDesc
)

func file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescGZIP() []byte {
	file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescOnce.Do(func() {
		file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescData)
	})
	return file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDescData
}

var file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_scheduler_pkg_proto_scheduler_scheduler_proto_goTypes = []interface{}{
	(*Empty)(nil),           // 0: scheduler.Empty
	(*EventID)(nil),         // 1: scheduler.EventID
	(*Date)(nil),            // 2: scheduler.Date
	(*CustomerID)(nil),      // 3: scheduler.CustomerID
	(*CustomerRequest)(nil), // 4: scheduler.CustomerRequest
	(*DeleteRequest)(nil),   // 5: scheduler.DeleteRequest
	(*UpdateRequest)(nil),   // 6: scheduler.UpdateRequest
	(*Event)(nil),           // 7: scheduler.Event
	(*Events)(nil),          // 8: scheduler.Events
}
var file_scheduler_pkg_proto_scheduler_scheduler_proto_depIdxs = []int32{
	7, // 0: scheduler.UpdateRequest.Event:type_name -> scheduler.Event
	7, // 1: scheduler.Events.Events:type_name -> scheduler.Event
	1, // 2: scheduler.Scheduler.GetEvent:input_type -> scheduler.EventID
	2, // 3: scheduler.Scheduler.GetEventsFor:input_type -> scheduler.Date
	4, // 4: scheduler.Scheduler.GetEventsForCustomer:input_type -> scheduler.CustomerRequest
	7, // 5: scheduler.Scheduler.CreateEvent:input_type -> scheduler.Event
	6, // 6: scheduler.Scheduler.UpdateEvent:input_type -> scheduler.UpdateRequest
	5, // 7: scheduler.Scheduler.DeleteEvent:input_type -> scheduler.DeleteRequest
	3, // 8: scheduler.Scheduler.DeleteAllForCustomer:input_type -> scheduler.CustomerID
	7, // 9: scheduler.Scheduler.GetEvent:output_type -> scheduler.Event
	8, // 10: scheduler.Scheduler.GetEventsFor:output_type -> scheduler.Events
	8, // 11: scheduler.Scheduler.GetEventsForCustomer:output_type -> scheduler.Events
	7, // 12: scheduler.Scheduler.CreateEvent:output_type -> scheduler.Event
	0, // 13: scheduler.Scheduler.UpdateEvent:output_type -> scheduler.Empty
	0, // 14: scheduler.Scheduler.DeleteEvent:output_type -> scheduler.Empty
	0, // 15: scheduler.Scheduler.DeleteAllForCustomer:output_type -> scheduler.Empty
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_scheduler_pkg_proto_scheduler_scheduler_proto_init() }
func file_scheduler_pkg_proto_scheduler_scheduler_proto_init() {
	if File_scheduler_pkg_proto_scheduler_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventID); i {
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
		file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
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
		file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerID); i {
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
		file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerRequest); i {
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
		file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Events); i {
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
			RawDescriptor: file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scheduler_pkg_proto_scheduler_scheduler_proto_goTypes,
		DependencyIndexes: file_scheduler_pkg_proto_scheduler_scheduler_proto_depIdxs,
		MessageInfos:      file_scheduler_pkg_proto_scheduler_scheduler_proto_msgTypes,
	}.Build()
	File_scheduler_pkg_proto_scheduler_scheduler_proto = out.File
	file_scheduler_pkg_proto_scheduler_scheduler_proto_rawDesc = nil
	file_scheduler_pkg_proto_scheduler_scheduler_proto_goTypes = nil
	file_scheduler_pkg_proto_scheduler_scheduler_proto_depIdxs = nil
}