// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: google/cloud/discoveryengine/v1beta/user_event_service.proto

package discoveryenginepb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for WriteUserEvent method.
type WriteUserEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent DataStore resource name, such as
	// `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. User event to write.
	UserEvent *UserEvent `protobuf:"bytes,2,opt,name=user_event,json=userEvent,proto3,oneof" json:"user_event,omitempty"`
}

func (x *WriteUserEventRequest) Reset() {
	*x = WriteUserEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_user_event_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteUserEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteUserEventRequest) ProtoMessage() {}

func (x *WriteUserEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_user_event_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteUserEventRequest.ProtoReflect.Descriptor instead.
func (*WriteUserEventRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDescGZIP(), []int{0}
}

func (x *WriteUserEventRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *WriteUserEventRequest) GetUserEvent() *UserEvent {
	if x != nil {
		return x.UserEvent
	}
	return nil
}

// Request message for CollectUserEvent method.
type CollectUserEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent DataStore resource name, such as
	// `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. URL encoded UserEvent proto with a length limit of 2,000,000
	// characters.
	UserEvent string `protobuf:"bytes,2,opt,name=user_event,json=userEvent,proto3" json:"user_event,omitempty"`
	// The URL including cgi-parameters but excluding the hash fragment with a
	// length limit of 5,000 characters. This is often more useful than the
	// referer URL, because many browsers only send the domain for third-party
	// requests.
	Uri *string `protobuf:"bytes,3,opt,name=uri,proto3,oneof" json:"uri,omitempty"`
	// The event timestamp in milliseconds. This prevents browser caching of
	// otherwise identical get requests. The name is abbreviated to reduce the
	// payload bytes.
	Ets *int64 `protobuf:"varint,4,opt,name=ets,proto3,oneof" json:"ets,omitempty"`
}

func (x *CollectUserEventRequest) Reset() {
	*x = CollectUserEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_user_event_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectUserEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectUserEventRequest) ProtoMessage() {}

func (x *CollectUserEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_user_event_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectUserEventRequest.ProtoReflect.Descriptor instead.
func (*CollectUserEventRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDescGZIP(), []int{1}
}

func (x *CollectUserEventRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CollectUserEventRequest) GetUserEvent() string {
	if x != nil {
		return x.UserEvent
	}
	return ""
}

func (x *CollectUserEventRequest) GetUri() string {
	if x != nil && x.Uri != nil {
		return *x.Uri
	}
	return ""
}

func (x *CollectUserEventRequest) GetEts() int64 {
	if x != nil && x.Ets != nil {
		return *x.Ets
	}
	return 0
}

var File_google_cloud_discoveryengine_v1beta_user_event_service_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x15, 0x57, 0x72, 0x69, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x48, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x30, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x57, 0x0a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x48, 0x00, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x88,
	0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0xc5, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0,
	0x41, 0x02, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x75,
	0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x69, 0x88,
	0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x01, 0x52, 0x03, 0x65, 0x74, 0x73, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72,
	0x69, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x65, 0x74, 0x73, 0x32, 0xdb, 0x08, 0x0a, 0x10, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbc,
	0x02, 0x0a, 0x0e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xbd, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0xb6, 0x01, 0x3a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5a, 0x61, 0x3a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x53, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0x45, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f,
	0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x92, 0x02,
	0x0a, 0x10, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0xa9, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xa2, 0x01,
	0x5a, 0x57, 0x12, 0x55, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x3a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x47, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x12, 0x9e, 0x03, 0x0a, 0x10, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c,
	0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xac, 0x02, 0xca, 0x41, 0x7c, 0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xa6, 0x01, 0x3a, 0x01, 0x2a,
	0x5a, 0x59, 0x3a, 0x01, 0x2a, 0x22, 0x54, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x3a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x46, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a,
	0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x1a, 0x52, 0xca, 0x41, 0x1e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x9c, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x42, 0x15, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xa2,
	0x02, 0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49, 0x4e,
	0x45, 0xaa, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x26,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDescData = file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1beta_user_event_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_discoveryengine_v1beta_user_event_service_proto_goTypes = []interface{}{
	(*WriteUserEventRequest)(nil),   // 0: google.cloud.discoveryengine.v1beta.WriteUserEventRequest
	(*CollectUserEventRequest)(nil), // 1: google.cloud.discoveryengine.v1beta.CollectUserEventRequest
	(*UserEvent)(nil),               // 2: google.cloud.discoveryengine.v1beta.UserEvent
	(*ImportUserEventsRequest)(nil), // 3: google.cloud.discoveryengine.v1beta.ImportUserEventsRequest
	(*httpbody.HttpBody)(nil),       // 4: google.api.HttpBody
	(*longrunningpb.Operation)(nil), // 5: google.longrunning.Operation
}
var file_google_cloud_discoveryengine_v1beta_user_event_service_proto_depIdxs = []int32{
	2, // 0: google.cloud.discoveryengine.v1beta.WriteUserEventRequest.user_event:type_name -> google.cloud.discoveryengine.v1beta.UserEvent
	0, // 1: google.cloud.discoveryengine.v1beta.UserEventService.WriteUserEvent:input_type -> google.cloud.discoveryengine.v1beta.WriteUserEventRequest
	1, // 2: google.cloud.discoveryengine.v1beta.UserEventService.CollectUserEvent:input_type -> google.cloud.discoveryengine.v1beta.CollectUserEventRequest
	3, // 3: google.cloud.discoveryengine.v1beta.UserEventService.ImportUserEvents:input_type -> google.cloud.discoveryengine.v1beta.ImportUserEventsRequest
	2, // 4: google.cloud.discoveryengine.v1beta.UserEventService.WriteUserEvent:output_type -> google.cloud.discoveryengine.v1beta.UserEvent
	4, // 5: google.cloud.discoveryengine.v1beta.UserEventService.CollectUserEvent:output_type -> google.api.HttpBody
	5, // 6: google.cloud.discoveryengine.v1beta.UserEventService.ImportUserEvents:output_type -> google.longrunning.Operation
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1beta_user_event_service_proto_init() }
func file_google_cloud_discoveryengine_v1beta_user_event_service_proto_init() {
	if File_google_cloud_discoveryengine_v1beta_user_event_service_proto != nil {
		return
	}
	file_google_cloud_discoveryengine_v1beta_import_config_proto_init()
	file_google_cloud_discoveryengine_v1beta_user_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1beta_user_event_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteUserEventRequest); i {
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
		file_google_cloud_discoveryengine_v1beta_user_event_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectUserEventRequest); i {
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
	file_google_cloud_discoveryengine_v1beta_user_event_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_google_cloud_discoveryengine_v1beta_user_event_service_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1beta_user_event_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1beta_user_event_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1beta_user_event_service_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1beta_user_event_service_proto = out.File
	file_google_cloud_discoveryengine_v1beta_user_event_service_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1beta_user_event_service_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1beta_user_event_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserEventServiceClient is the client API for UserEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserEventServiceClient interface {
	// Writes a single user event.
	WriteUserEvent(ctx context.Context, in *WriteUserEventRequest, opts ...grpc.CallOption) (*UserEvent, error)
	// Writes a single user event from the browser. This uses a GET request to
	// due to browser restriction of POST-ing to a third-party domain.
	//
	// This method is used only by the Discovery Engine API JavaScript pixel and
	// Google Tag Manager. Users should not call this method directly.
	CollectUserEvent(ctx context.Context, in *CollectUserEventRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	// Bulk import of User events. Request processing might be
	// synchronous. Events that already exist are skipped.
	// Use this method for backfilling historical user events.
	//
	// Operation.response is of type ImportResponse. Note that it is
	// possible for a subset of the items to be successfully inserted.
	// Operation.metadata is of type ImportMetadata.
	ImportUserEvents(ctx context.Context, in *ImportUserEventsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type userEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserEventServiceClient(cc grpc.ClientConnInterface) UserEventServiceClient {
	return &userEventServiceClient{cc}
}

func (c *userEventServiceClient) WriteUserEvent(ctx context.Context, in *WriteUserEventRequest, opts ...grpc.CallOption) (*UserEvent, error) {
	out := new(UserEvent)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1beta.UserEventService/WriteUserEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEventServiceClient) CollectUserEvent(ctx context.Context, in *CollectUserEventRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1beta.UserEventService/CollectUserEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEventServiceClient) ImportUserEvents(ctx context.Context, in *ImportUserEventsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1beta.UserEventService/ImportUserEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserEventServiceServer is the server API for UserEventService service.
type UserEventServiceServer interface {
	// Writes a single user event.
	WriteUserEvent(context.Context, *WriteUserEventRequest) (*UserEvent, error)
	// Writes a single user event from the browser. This uses a GET request to
	// due to browser restriction of POST-ing to a third-party domain.
	//
	// This method is used only by the Discovery Engine API JavaScript pixel and
	// Google Tag Manager. Users should not call this method directly.
	CollectUserEvent(context.Context, *CollectUserEventRequest) (*httpbody.HttpBody, error)
	// Bulk import of User events. Request processing might be
	// synchronous. Events that already exist are skipped.
	// Use this method for backfilling historical user events.
	//
	// Operation.response is of type ImportResponse. Note that it is
	// possible for a subset of the items to be successfully inserted.
	// Operation.metadata is of type ImportMetadata.
	ImportUserEvents(context.Context, *ImportUserEventsRequest) (*longrunningpb.Operation, error)
}

// UnimplementedUserEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserEventServiceServer struct {
}

func (*UnimplementedUserEventServiceServer) WriteUserEvent(context.Context, *WriteUserEventRequest) (*UserEvent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteUserEvent not implemented")
}
func (*UnimplementedUserEventServiceServer) CollectUserEvent(context.Context, *CollectUserEventRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectUserEvent not implemented")
}
func (*UnimplementedUserEventServiceServer) ImportUserEvents(context.Context, *ImportUserEventsRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportUserEvents not implemented")
}

func RegisterUserEventServiceServer(s *grpc.Server, srv UserEventServiceServer) {
	s.RegisterService(&_UserEventService_serviceDesc, srv)
}

func _UserEventService_WriteUserEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteUserEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServiceServer).WriteUserEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1beta.UserEventService/WriteUserEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServiceServer).WriteUserEvent(ctx, req.(*WriteUserEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEventService_CollectUserEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectUserEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServiceServer).CollectUserEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1beta.UserEventService/CollectUserEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServiceServer).CollectUserEvent(ctx, req.(*CollectUserEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEventService_ImportUserEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportUserEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEventServiceServer).ImportUserEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1beta.UserEventService/ImportUserEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEventServiceServer).ImportUserEvents(ctx, req.(*ImportUserEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserEventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1beta.UserEventService",
	HandlerType: (*UserEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteUserEvent",
			Handler:    _UserEventService_WriteUserEvent_Handler,
		},
		{
			MethodName: "CollectUserEvent",
			Handler:    _UserEventService_CollectUserEvent_Handler,
		},
		{
			MethodName: "ImportUserEvents",
			Handler:    _UserEventService_ImportUserEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1beta/user_event_service.proto",
}
