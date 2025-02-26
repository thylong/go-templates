// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: events/events.proto

package eventsv1

import (
	_type "github.com/thylong/go-templates/06-grpc-sqlc/pkg/proto/google/type"
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

// EventPrivacy represents the different privacty levels of an event
type EventPrivacy int32

const (
	EventPrivacy_PET_PRIVACY_INTERNAL EventPrivacy = 0
	EventPrivacy_PET_PRIVACY_EXTERNAL EventPrivacy = 1
)

// Enum value maps for EventPrivacy.
var (
	EventPrivacy_name = map[int32]string{
		0: "PET_PRIVACY_INTERNAL",
		1: "PET_PRIVACY_EXTERNAL",
	}
	EventPrivacy_value = map[string]int32{
		"PET_PRIVACY_INTERNAL": 0,
		"PET_PRIVACY_EXTERNAL": 1,
	}
)

func (x EventPrivacy) Enum() *EventPrivacy {
	p := new(EventPrivacy)
	*p = x
	return p
}

func (x EventPrivacy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventPrivacy) Descriptor() protoreflect.EnumDescriptor {
	return file_events_events_proto_enumTypes[0].Descriptor()
}

func (EventPrivacy) Type() protoreflect.EnumType {
	return &file_events_events_proto_enumTypes[0]
}

func (x EventPrivacy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventPrivacy.Descriptor instead.
func (EventPrivacy) EnumDescriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{0}
}

// Event represents an event that occured in the company
type Event struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventPrivacy  EventPrivacy           `protobuf:"varint,1,opt,name=event_privacy,json=eventPrivacy,proto3,enum=events.v1.EventPrivacy" json:"event_privacy,omitempty"`
	EventId       string                 `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type          string                 `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Department    string                 `protobuf:"bytes,5,opt,name=department,proto3" json:"department,omitempty"`
	Regions       string                 `protobuf:"bytes,6,opt,name=regions,proto3" json:"regions,omitempty"`
	Tags          string                 `protobuf:"bytes,7,opt,name=tags,proto3" json:"tags,omitempty"`
	StartAt       *_type.DateTime        `protobuf:"bytes,8,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_events_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[0]
	if x != nil {
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
	return file_events_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetEventPrivacy() EventPrivacy {
	if x != nil {
		return x.EventPrivacy
	}
	return EventPrivacy_PET_PRIVACY_INTERNAL
}

func (x *Event) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Event) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *Event) GetRegions() string {
	if x != nil {
		return x.Regions
	}
	return ""
}

func (x *Event) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Event) GetStartAt() *_type.DateTime {
	if x != nil {
		return x.StartAt
	}
	return nil
}

type GetEventsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Pagination parameters (optional)
	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                         // The page number
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"` // Number of events per page
	// Filtering parameters (optional)
	Search        string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"` // Search query for filtering events
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	mi := &file_events_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{1}
}

func (x *GetEventsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetEventsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetEventsRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

// Response message for GetEvents
type GetEventsResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Events []*Event               `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	// Pagination metadata
	TotalCount    int32 `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"` // Total number of events
	Page          int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`                               // Current page
	PageSize      int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`       // Number of events per page
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	mi := &file_events_events_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{2}
}

func (x *GetEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *GetEventsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetEventsResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetEventsResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetEventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventId       string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	mi := &file_events_events_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{3}
}

func (x *GetEventRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type GetEventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Event         *Event                 `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEventResponse) Reset() {
	*x = GetEventResponse{}
	mi := &file_events_events_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventResponse) ProtoMessage() {}

func (x *GetEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventResponse.ProtoReflect.Descriptor instead.
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{4}
}

func (x *GetEventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type PutEventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventPrivacy  EventPrivacy           `protobuf:"varint,1,opt,name=event_privacy,json=eventPrivacy,proto3,enum=events.v1.EventPrivacy" json:"event_privacy,omitempty"`
	EventId       string                 `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type          string                 `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Department    string                 `protobuf:"bytes,5,opt,name=department,proto3" json:"department,omitempty"`
	Regions       string                 `protobuf:"bytes,6,opt,name=regions,proto3" json:"regions,omitempty"`
	Tags          string                 `protobuf:"bytes,7,opt,name=tags,proto3" json:"tags,omitempty"`
	StartAt       *_type.DateTime        `protobuf:"bytes,8,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutEventRequest) Reset() {
	*x = PutEventRequest{}
	mi := &file_events_events_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutEventRequest) ProtoMessage() {}

func (x *PutEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutEventRequest.ProtoReflect.Descriptor instead.
func (*PutEventRequest) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{5}
}

func (x *PutEventRequest) GetEventPrivacy() EventPrivacy {
	if x != nil {
		return x.EventPrivacy
	}
	return EventPrivacy_PET_PRIVACY_INTERNAL
}

func (x *PutEventRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *PutEventRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PutEventRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PutEventRequest) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *PutEventRequest) GetRegions() string {
	if x != nil {
		return x.Regions
	}
	return ""
}

func (x *PutEventRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *PutEventRequest) GetStartAt() *_type.DateTime {
	if x != nil {
		return x.StartAt
	}
	return nil
}

type PutEventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Event         *Event                 `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PutEventResponse) Reset() {
	*x = PutEventResponse{}
	mi := &file_events_events_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PutEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutEventResponse) ProtoMessage() {}

func (x *PutEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutEventResponse.ProtoReflect.Descriptor instead.
func (*PutEventResponse) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{6}
}

func (x *PutEventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type DeleteEventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventID       string                 `protobuf:"bytes,1,opt,name=eventID,proto3" json:"eventID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteEventRequest) Reset() {
	*x = DeleteEventRequest{}
	mi := &file_events_events_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventRequest) ProtoMessage() {}

func (x *DeleteEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventRequest.ProtoReflect.Descriptor instead.
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEventRequest) GetEventID() string {
	if x != nil {
		return x.EventID
	}
	return ""
}

type DeleteEventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteEventResponse) Reset() {
	*x = DeleteEventResponse{}
	mi := &file_events_events_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventResponse) ProtoMessage() {}

func (x *DeleteEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_events_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventResponse.ProtoReflect.Descriptor instead.
func (*DeleteEventResponse) Descriptor() ([]byte, []int) {
	return file_events_events_proto_rawDescGZIP(), []int{8}
}

var File_events_events_proto protoreflect.FileDescriptor

var file_events_events_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x22, 0x5b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x92, 0x02, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x41, 0x74, 0x22, 0x3a, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x42, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x45, 0x54, 0x5f,
	0x50, 0x52, 0x49, 0x56, 0x41, 0x43, 0x59, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x45, 0x54, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x43,
	0x59, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x32, 0xb6, 0x02, 0x0a,
	0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x08, 0x50, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xa9, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x79, 0x6c, 0x6f, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x30, 0x36, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x73, 0x71, 0x6c, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_events_proto_rawDescOnce sync.Once
	file_events_events_proto_rawDescData = file_events_events_proto_rawDesc
)

func file_events_events_proto_rawDescGZIP() []byte {
	file_events_events_proto_rawDescOnce.Do(func() {
		file_events_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_events_proto_rawDescData)
	})
	return file_events_events_proto_rawDescData
}

var file_events_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_events_events_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_events_events_proto_goTypes = []any{
	(EventPrivacy)(0),           // 0: events.v1.EventPrivacy
	(*Event)(nil),               // 1: events.v1.Event
	(*GetEventsRequest)(nil),    // 2: events.v1.GetEventsRequest
	(*GetEventsResponse)(nil),   // 3: events.v1.GetEventsResponse
	(*GetEventRequest)(nil),     // 4: events.v1.GetEventRequest
	(*GetEventResponse)(nil),    // 5: events.v1.GetEventResponse
	(*PutEventRequest)(nil),     // 6: events.v1.PutEventRequest
	(*PutEventResponse)(nil),    // 7: events.v1.PutEventResponse
	(*DeleteEventRequest)(nil),  // 8: events.v1.DeleteEventRequest
	(*DeleteEventResponse)(nil), // 9: events.v1.DeleteEventResponse
	(*_type.DateTime)(nil),      // 10: google.type.DateTime
}
var file_events_events_proto_depIdxs = []int32{
	0,  // 0: events.v1.Event.event_privacy:type_name -> events.v1.EventPrivacy
	10, // 1: events.v1.Event.start_at:type_name -> google.type.DateTime
	1,  // 2: events.v1.GetEventsResponse.events:type_name -> events.v1.Event
	1,  // 3: events.v1.GetEventResponse.event:type_name -> events.v1.Event
	0,  // 4: events.v1.PutEventRequest.event_privacy:type_name -> events.v1.EventPrivacy
	10, // 5: events.v1.PutEventRequest.start_at:type_name -> google.type.DateTime
	1,  // 6: events.v1.PutEventResponse.event:type_name -> events.v1.Event
	2,  // 7: events.v1.EventService.GetEvents:input_type -> events.v1.GetEventsRequest
	4,  // 8: events.v1.EventService.GetEvent:input_type -> events.v1.GetEventRequest
	6,  // 9: events.v1.EventService.PutEvent:input_type -> events.v1.PutEventRequest
	8,  // 10: events.v1.EventService.DeleteEvent:input_type -> events.v1.DeleteEventRequest
	3,  // 11: events.v1.EventService.GetEvents:output_type -> events.v1.GetEventsResponse
	5,  // 12: events.v1.EventService.GetEvent:output_type -> events.v1.GetEventResponse
	7,  // 13: events.v1.EventService.PutEvent:output_type -> events.v1.PutEventResponse
	9,  // 14: events.v1.EventService.DeleteEvent:output_type -> events.v1.DeleteEventResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_events_events_proto_init() }
func file_events_events_proto_init() {
	if File_events_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_events_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_events_events_proto_goTypes,
		DependencyIndexes: file_events_events_proto_depIdxs,
		EnumInfos:         file_events_events_proto_enumTypes,
		MessageInfos:      file_events_events_proto_msgTypes,
	}.Build()
	File_events_events_proto = out.File
	file_events_events_proto_rawDesc = nil
	file_events_events_proto_goTypes = nil
	file_events_events_proto_depIdxs = nil
}
