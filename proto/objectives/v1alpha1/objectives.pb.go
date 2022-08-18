// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: objectives/v1alpha1/objectives.proto

package objectivesv1alpha1

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

type ListObjectivesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expr string `protobuf:"bytes,1,opt,name=expr,proto3" json:"expr,omitempty"`
}

func (x *ListObjectivesRequest) Reset() {
	*x = ListObjectivesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectives_v1alpha1_objectives_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListObjectivesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListObjectivesRequest) ProtoMessage() {}

func (x *ListObjectivesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_objectives_v1alpha1_objectives_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListObjectivesRequest.ProtoReflect.Descriptor instead.
func (*ListObjectivesRequest) Descriptor() ([]byte, []int) {
	return file_objectives_v1alpha1_objectives_proto_rawDescGZIP(), []int{0}
}

func (x *ListObjectivesRequest) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

type ListObjectivesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Objectives []*Objective `protobuf:"bytes,1,rep,name=objectives,proto3" json:"objectives,omitempty"`
}

func (x *ListObjectivesResponse) Reset() {
	*x = ListObjectivesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectives_v1alpha1_objectives_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListObjectivesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListObjectivesResponse) ProtoMessage() {}

func (x *ListObjectivesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_objectives_v1alpha1_objectives_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListObjectivesResponse.ProtoReflect.Descriptor instead.
func (*ListObjectivesResponse) Descriptor() ([]byte, []int) {
	return file_objectives_v1alpha1_objectives_proto_rawDescGZIP(), []int{1}
}

func (x *ListObjectivesResponse) GetObjectives() []*Objective {
	if x != nil {
		return x.Objectives
	}
	return nil
}

type Objective struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels      map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Description string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Target      float64           `protobuf:"fixed64,3,opt,name=target,proto3" json:"target,omitempty"`
	Window      int64             `protobuf:"varint,4,opt,name=window,proto3" json:"window,omitempty"`
}

func (x *Objective) Reset() {
	*x = Objective{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objectives_v1alpha1_objectives_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Objective) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Objective) ProtoMessage() {}

func (x *Objective) ProtoReflect() protoreflect.Message {
	mi := &file_objectives_v1alpha1_objectives_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Objective.ProtoReflect.Descriptor instead.
func (*Objective) Descriptor() ([]byte, []int) {
	return file_objectives_v1alpha1_objectives_proto_rawDescGZIP(), []int{2}
}

func (x *Objective) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Objective) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Objective) GetTarget() float64 {
	if x != nil {
		return x.Target
	}
	return 0
}

func (x *Objective) GetWindow() int64 {
	if x != nil {
		return x.Window
	}
	return 0
}

var File_objectives_v1alpha1_objectives_proto protoreflect.FileDescriptor

var file_objectives_v1alpha1_objectives_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x2b, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x70, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x70, 0x72, 0x22, 0x58, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x73, 0x22, 0xdc, 0x01, 0x0a, 0x09, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x42, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x32, 0x7f, 0x0a, 0x10, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x32, 0x86, 0x01, 0x0a, 0x17, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73,
	0x12, 0x2a, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x49, 0x5a, 0x47, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x79, 0x72, 0x72, 0x61, 0x2d,
	0x64, 0x65, 0x76, 0x2f, 0x70, 0x79, 0x72, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x3b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_objectives_v1alpha1_objectives_proto_rawDescOnce sync.Once
	file_objectives_v1alpha1_objectives_proto_rawDescData = file_objectives_v1alpha1_objectives_proto_rawDesc
)

func file_objectives_v1alpha1_objectives_proto_rawDescGZIP() []byte {
	file_objectives_v1alpha1_objectives_proto_rawDescOnce.Do(func() {
		file_objectives_v1alpha1_objectives_proto_rawDescData = protoimpl.X.CompressGZIP(file_objectives_v1alpha1_objectives_proto_rawDescData)
	})
	return file_objectives_v1alpha1_objectives_proto_rawDescData
}

var file_objectives_v1alpha1_objectives_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_objectives_v1alpha1_objectives_proto_goTypes = []interface{}{
	(*ListObjectivesRequest)(nil),  // 0: objectives.v1alpha1.ListObjectivesRequest
	(*ListObjectivesResponse)(nil), // 1: objectives.v1alpha1.ListObjectivesResponse
	(*Objective)(nil),              // 2: objectives.v1alpha1.Objective
	nil,                            // 3: objectives.v1alpha1.Objective.LabelsEntry
}
var file_objectives_v1alpha1_objectives_proto_depIdxs = []int32{
	2, // 0: objectives.v1alpha1.ListObjectivesResponse.objectives:type_name -> objectives.v1alpha1.Objective
	3, // 1: objectives.v1alpha1.Objective.labels:type_name -> objectives.v1alpha1.Objective.LabelsEntry
	0, // 2: objectives.v1alpha1.ObjectiveService.ListObjectives:input_type -> objectives.v1alpha1.ListObjectivesRequest
	0, // 3: objectives.v1alpha1.ObjectiveBackendService.ListObjectives:input_type -> objectives.v1alpha1.ListObjectivesRequest
	1, // 4: objectives.v1alpha1.ObjectiveService.ListObjectives:output_type -> objectives.v1alpha1.ListObjectivesResponse
	1, // 5: objectives.v1alpha1.ObjectiveBackendService.ListObjectives:output_type -> objectives.v1alpha1.ListObjectivesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_objectives_v1alpha1_objectives_proto_init() }
func file_objectives_v1alpha1_objectives_proto_init() {
	if File_objectives_v1alpha1_objectives_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_objectives_v1alpha1_objectives_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListObjectivesRequest); i {
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
		file_objectives_v1alpha1_objectives_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListObjectivesResponse); i {
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
		file_objectives_v1alpha1_objectives_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Objective); i {
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
			RawDescriptor: file_objectives_v1alpha1_objectives_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_objectives_v1alpha1_objectives_proto_goTypes,
		DependencyIndexes: file_objectives_v1alpha1_objectives_proto_depIdxs,
		MessageInfos:      file_objectives_v1alpha1_objectives_proto_msgTypes,
	}.Build()
	File_objectives_v1alpha1_objectives_proto = out.File
	file_objectives_v1alpha1_objectives_proto_rawDesc = nil
	file_objectives_v1alpha1_objectives_proto_goTypes = nil
	file_objectives_v1alpha1_objectives_proto_depIdxs = nil
}
