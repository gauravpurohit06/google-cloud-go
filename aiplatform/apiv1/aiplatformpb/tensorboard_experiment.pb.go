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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: google/cloud/aiplatform/v1/tensorboard_experiment.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A TensorboardExperiment is a group of TensorboardRuns, that are typically the
// results of a training job run, in a Tensorboard.
type TensorboardExperiment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Name of the TensorboardExperiment.
	// Format:
	// `projects/{project}/locations/{location}/tensorboards/{tensorboard}/experiments/{experiment}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// User provided name of this TensorboardExperiment.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Description of this TensorboardExperiment.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Timestamp when this TensorboardExperiment was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this TensorboardExperiment was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The labels with user-defined metadata to organize your Datasets.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// No more than 64 user labels can be associated with one Dataset (System
	// labels are excluded).
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	// System reserved label keys are prefixed with "aiplatform.googleapis.com/"
	// and are immutable. Following system labels exist for each Dataset:
	// * "aiplatform.googleapis.com/dataset_metadata_schema":
	//   - output only, its value is the
	//     [metadata_schema's][metadata_schema_uri] title.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag string `protobuf:"bytes,7,opt,name=etag,proto3" json:"etag,omitempty"`
	// Immutable. Source of the TensorboardExperiment. Example: a custom training job.
	Source string `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *TensorboardExperiment) Reset() {
	*x = TensorboardExperiment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TensorboardExperiment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TensorboardExperiment) ProtoMessage() {}

func (x *TensorboardExperiment) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TensorboardExperiment.ProtoReflect.Descriptor instead.
func (*TensorboardExperiment) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDescGZIP(), []int{0}
}

func (x *TensorboardExperiment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TensorboardExperiment) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *TensorboardExperiment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TensorboardExperiment) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *TensorboardExperiment) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *TensorboardExperiment) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TensorboardExperiment) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *TensorboardExperiment) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_google_cloud_aiplatform_v1_tensorboard_experiment_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd1, 0x04, 0x0a, 0x15, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x55,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x3a, 0x92, 0x01, 0xea, 0x41, 0x8e, 0x01, 0x0a, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f,
	0x7b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x7d, 0x2f, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x42, 0xde, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x1a, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0xaa, 0x02, 0x1a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDescData = file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_goTypes = []interface{}{
	(*TensorboardExperiment)(nil), // 0: google.cloud.aiplatform.v1.TensorboardExperiment
	nil,                           // 1: google.cloud.aiplatform.v1.TensorboardExperiment.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_depIdxs = []int32{
	2, // 0: google.cloud.aiplatform.v1.TensorboardExperiment.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.cloud.aiplatform.v1.TensorboardExperiment.update_time:type_name -> google.protobuf.Timestamp
	1, // 2: google.cloud.aiplatform.v1.TensorboardExperiment.labels:type_name -> google.cloud.aiplatform.v1.TensorboardExperiment.LabelsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_init() }
func file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_init() {
	if File_google_cloud_aiplatform_v1_tensorboard_experiment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TensorboardExperiment); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_tensorboard_experiment_proto = out.File
	file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_tensorboard_experiment_proto_depIdxs = nil
}
