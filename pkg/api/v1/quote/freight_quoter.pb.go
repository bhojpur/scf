// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: pkg/api/v1/quote/freight_quoter.proto

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package quote

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quantity uint64  `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price    float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Width    uint64  `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height   uint64  `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Length   uint64  `protobuf:"varint,5,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_quote_freight_quoter_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Item) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Item) GetWidth() uint64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Item) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Item) GetLength() uint64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ToPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ToPackage) Reset() {
	*x = ToPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToPackage) ProtoMessage() {}

func (x *ToPackage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToPackage.ProtoReflect.Descriptor instead.
func (*ToPackage) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_quote_freight_quoter_proto_rawDescGZIP(), []int{1}
}

func (x *ToPackage) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quantity uint64  `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price    float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Volume   float32 `protobuf:"fixed32,3,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_quote_freight_quoter_proto_rawDescGZIP(), []int{2}
}

func (x *Package) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Package) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Package) GetVolume() float32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

type Quotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	ZipCode string   `protobuf:"bytes,2,opt,name=zipCode,proto3" json:"zipCode,omitempty"`
}

func (x *Quotation) Reset() {
	*x = Quotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quotation) ProtoMessage() {}

func (x *Quotation) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quotation.ProtoReflect.Descriptor instead.
func (*Quotation) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_quote_freight_quoter_proto_rawDescGZIP(), []int{3}
}

func (x *Quotation) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

func (x *Quotation) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

type Delivery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  uint64  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Price float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Delivery) Reset() {
	*x = Delivery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delivery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delivery) ProtoMessage() {}

func (x *Delivery) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delivery.ProtoReflect.Descriptor instead.
func (*Delivery) Descriptor() ([]byte, []int) {
	return file_pkg_api_v1_quote_freight_quoter_proto_rawDescGZIP(), []int{4}
}

func (x *Delivery) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Delivery) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_pkg_api_v1_quote_freight_quoter_proto protoreflect.FileDescriptor

var file_pkg_api_v1_quote_freight_quoter_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x2f, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x22, 0x7e, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x22, 0x31, 0x0a, 0x09, 0x54, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x24,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x53, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x09, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x34, 0x0a,
	0x08, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x32, 0x7d, 0x0a, 0x0d, 0x46, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0c, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e,
	0x54, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x76,
	0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x68, 0x6f, 0x6a, 0x70, 0x75, 0x72, 0x2f, 0x73, 0x63, 0x66, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x3b, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_v1_quote_freight_quoter_proto_rawDescOnce sync.Once
	file_pkg_api_v1_quote_freight_quoter_proto_rawDescData = file_pkg_api_v1_quote_freight_quoter_proto_rawDesc
)

func file_pkg_api_v1_quote_freight_quoter_proto_rawDescGZIP() []byte {
	file_pkg_api_v1_quote_freight_quoter_proto_rawDescOnce.Do(func() {
		file_pkg_api_v1_quote_freight_quoter_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_v1_quote_freight_quoter_proto_rawDescData)
	})
	return file_pkg_api_v1_quote_freight_quoter_proto_rawDescData
}

var file_pkg_api_v1_quote_freight_quoter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_api_v1_quote_freight_quoter_proto_goTypes = []interface{}{
	(*Item)(nil),      // 0: v1.quote.Item
	(*ToPackage)(nil), // 1: v1.quote.ToPackage
	(*Package)(nil),   // 2: v1.quote.Package
	(*Quotation)(nil), // 3: v1.quote.Quotation
	(*Delivery)(nil),  // 4: v1.quote.Delivery
}
var file_pkg_api_v1_quote_freight_quoter_proto_depIdxs = []int32{
	0, // 0: v1.quote.ToPackage.items:type_name -> v1.quote.Item
	2, // 1: v1.quote.Quotation.package:type_name -> v1.quote.Package
	1, // 2: v1.quote.FreightQuoter.MountPackage:input_type -> v1.quote.ToPackage
	3, // 3: v1.quote.FreightQuoter.Quote:input_type -> v1.quote.Quotation
	2, // 4: v1.quote.FreightQuoter.MountPackage:output_type -> v1.quote.Package
	4, // 5: v1.quote.FreightQuoter.Quote:output_type -> v1.quote.Delivery
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_api_v1_quote_freight_quoter_proto_init() }
func file_pkg_api_v1_quote_freight_quoter_proto_init() {
	if File_pkg_api_v1_quote_freight_quoter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToPackage); i {
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
		file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quotation); i {
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
		file_pkg_api_v1_quote_freight_quoter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delivery); i {
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
			RawDescriptor: file_pkg_api_v1_quote_freight_quoter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_v1_quote_freight_quoter_proto_goTypes,
		DependencyIndexes: file_pkg_api_v1_quote_freight_quoter_proto_depIdxs,
		MessageInfos:      file_pkg_api_v1_quote_freight_quoter_proto_msgTypes,
	}.Build()
	File_pkg_api_v1_quote_freight_quoter_proto = out.File
	file_pkg_api_v1_quote_freight_quoter_proto_rawDesc = nil
	file_pkg_api_v1_quote_freight_quoter_proto_goTypes = nil
	file_pkg_api_v1_quote_freight_quoter_proto_depIdxs = nil
}
