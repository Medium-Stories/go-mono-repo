// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: discount/proto/discount.proto

package proto

import (
	proto "github.com/medium-stories/go-mono-repo/product/proto"
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

type DiscountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*proto.ProductMessage `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *DiscountsRequest) Reset() {
	*x = DiscountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_discount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountsRequest) ProtoMessage() {}

func (x *DiscountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_discount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountsRequest.ProtoReflect.Descriptor instead.
func (*DiscountsRequest) Descriptor() ([]byte, []int) {
	return file_discount_proto_discount_proto_rawDescGZIP(), []int{0}
}

func (x *DiscountsRequest) GetProducts() []*proto.ProductMessage {
	if x != nil {
		return x.Products
	}
	return nil
}

type DiscountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*proto.ProductMessage `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *DiscountsResponse) Reset() {
	*x = DiscountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_discount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountsResponse) ProtoMessage() {}

func (x *DiscountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_discount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountsResponse.ProtoReflect.Descriptor instead.
func (*DiscountsResponse) Descriptor() ([]byte, []int) {
	return file_discount_proto_discount_proto_rawDescGZIP(), []int{1}
}

func (x *DiscountsResponse) GetProducts() []*proto.ProductMessage {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_discount_proto_discount_proto protoreflect.FileDescriptor

var file_discount_proto_discount_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22,
	0x48, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0x5e, 0x0a, 0x10, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x4a, 0x0a,
	0x0d, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x2d, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x2d, 0x72,
	0x65, 0x70, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discount_proto_discount_proto_rawDescOnce sync.Once
	file_discount_proto_discount_proto_rawDescData = file_discount_proto_discount_proto_rawDesc
)

func file_discount_proto_discount_proto_rawDescGZIP() []byte {
	file_discount_proto_discount_proto_rawDescOnce.Do(func() {
		file_discount_proto_discount_proto_rawDescData = protoimpl.X.CompressGZIP(file_discount_proto_discount_proto_rawDescData)
	})
	return file_discount_proto_discount_proto_rawDescData
}

var file_discount_proto_discount_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_discount_proto_discount_proto_goTypes = []interface{}{
	(*DiscountsRequest)(nil),     // 0: discount.DiscountsRequest
	(*DiscountsResponse)(nil),    // 1: discount.DiscountsResponse
	(*proto.ProductMessage)(nil), // 2: product.ProductMessage
}
var file_discount_proto_discount_proto_depIdxs = []int32{
	2, // 0: discount.DiscountsRequest.products:type_name -> product.ProductMessage
	2, // 1: discount.DiscountsResponse.products:type_name -> product.ProductMessage
	0, // 2: discount.DiscountProvider.ApplyDiscount:input_type -> discount.DiscountsRequest
	1, // 3: discount.DiscountProvider.ApplyDiscount:output_type -> discount.DiscountsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_discount_proto_discount_proto_init() }
func file_discount_proto_discount_proto_init() {
	if File_discount_proto_discount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discount_proto_discount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountsRequest); i {
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
		file_discount_proto_discount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountsResponse); i {
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
			RawDescriptor: file_discount_proto_discount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discount_proto_discount_proto_goTypes,
		DependencyIndexes: file_discount_proto_discount_proto_depIdxs,
		MessageInfos:      file_discount_proto_discount_proto_msgTypes,
	}.Build()
	File_discount_proto_discount_proto = out.File
	file_discount_proto_discount_proto_rawDesc = nil
	file_discount_proto_discount_proto_goTypes = nil
	file_discount_proto_discount_proto_depIdxs = nil
}
