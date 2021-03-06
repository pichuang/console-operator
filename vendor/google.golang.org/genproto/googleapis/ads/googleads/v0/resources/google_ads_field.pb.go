// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/google_ads_field.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A field or resource (artifact) used by GoogleAdsService.
type GoogleAdsField struct {
	// The resource name of the artifact.
	// Artifact resource names have the form:
	//
	// `googleAdsFields/{name}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The name of the artifact.
	Name *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The category of the artifact.
	Category enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory `protobuf:"varint,3,opt,name=category,proto3,enum=google.ads.googleads.v0.enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory" json:"category,omitempty"`
	// Whether the artifact can be used in a SELECT clause in search
	// queries.
	Selectable *wrappers.BoolValue `protobuf:"bytes,4,opt,name=selectable,proto3" json:"selectable,omitempty"`
	// Whether the artifact can be used in a WHERE clause in search
	// queries.
	Filterable *wrappers.BoolValue `protobuf:"bytes,5,opt,name=filterable,proto3" json:"filterable,omitempty"`
	// Whether the artifact can be used in a ORDER BY clause in search
	// queries.
	Sortable *wrappers.BoolValue `protobuf:"bytes,6,opt,name=sortable,proto3" json:"sortable,omitempty"`
	// The names of all resources, segments, and metrics that are selectable with
	// the described artifact.
	SelectableWith []*wrappers.StringValue `protobuf:"bytes,7,rep,name=selectable_with,json=selectableWith,proto3" json:"selectable_with,omitempty"`
	// The names of all resources that are selectable with the described
	// artifact. Fields from these resources do not segment metrics when included
	// in search queries.
	//
	// This field is only set for artifacts whose category is RESOURCE.
	AttributeResources []*wrappers.StringValue `protobuf:"bytes,8,rep,name=attribute_resources,json=attributeResources,proto3" json:"attribute_resources,omitempty"`
	// The names of all metrics that are selectable with the described artifact.
	//
	// This field is only set for artifacts whose category is either RESOURCE or
	// SEGMENT.
	Metrics []*wrappers.StringValue `protobuf:"bytes,9,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// The names of all artifacts, whether a segment or another resource, that
	// segment metrics when included in search queries.
	//
	// This field is only set for artifacts of category RESOURCE, SEGMENT or
	// METRIC.
	Segments []*wrappers.StringValue `protobuf:"bytes,10,rep,name=segments,proto3" json:"segments,omitempty"`
	// Values the artifact can assume if it is a field of type ENUM.
	//
	// This field is only set for artifacts of category SEGMENT or ATTRIBUTE.
	EnumValues []*wrappers.StringValue `protobuf:"bytes,11,rep,name=enum_values,json=enumValues,proto3" json:"enum_values,omitempty"`
	// This field determines the operators that can be used with the artifact
	// in WHERE clauses.
	DataType enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType `protobuf:"varint,12,opt,name=data_type,json=dataType,proto3,enum=google.ads.googleads.v0.enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType" json:"data_type,omitempty"`
	// The URL of proto describing the artifact's data type.
	TypeUrl *wrappers.StringValue `protobuf:"bytes,13,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Whether the field artifact is repeated.
	IsRepeated           *wrappers.BoolValue `protobuf:"bytes,14,opt,name=is_repeated,json=isRepeated,proto3" json:"is_repeated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GoogleAdsField) Reset()         { *m = GoogleAdsField{} }
func (m *GoogleAdsField) String() string { return proto.CompactTextString(m) }
func (*GoogleAdsField) ProtoMessage()    {}
func (*GoogleAdsField) Descriptor() ([]byte, []int) {
	return fileDescriptor_google_ads_field_c22565e2dfed71e8, []int{0}
}
func (m *GoogleAdsField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleAdsField.Unmarshal(m, b)
}
func (m *GoogleAdsField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleAdsField.Marshal(b, m, deterministic)
}
func (dst *GoogleAdsField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleAdsField.Merge(dst, src)
}
func (m *GoogleAdsField) XXX_Size() int {
	return xxx_messageInfo_GoogleAdsField.Size(m)
}
func (m *GoogleAdsField) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleAdsField.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleAdsField proto.InternalMessageInfo

func (m *GoogleAdsField) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *GoogleAdsField) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GoogleAdsField) GetCategory() enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory {
	if m != nil {
		return m.Category
	}
	return enums.GoogleAdsFieldCategoryEnum_UNSPECIFIED
}

func (m *GoogleAdsField) GetSelectable() *wrappers.BoolValue {
	if m != nil {
		return m.Selectable
	}
	return nil
}

func (m *GoogleAdsField) GetFilterable() *wrappers.BoolValue {
	if m != nil {
		return m.Filterable
	}
	return nil
}

func (m *GoogleAdsField) GetSortable() *wrappers.BoolValue {
	if m != nil {
		return m.Sortable
	}
	return nil
}

func (m *GoogleAdsField) GetSelectableWith() []*wrappers.StringValue {
	if m != nil {
		return m.SelectableWith
	}
	return nil
}

func (m *GoogleAdsField) GetAttributeResources() []*wrappers.StringValue {
	if m != nil {
		return m.AttributeResources
	}
	return nil
}

func (m *GoogleAdsField) GetMetrics() []*wrappers.StringValue {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *GoogleAdsField) GetSegments() []*wrappers.StringValue {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *GoogleAdsField) GetEnumValues() []*wrappers.StringValue {
	if m != nil {
		return m.EnumValues
	}
	return nil
}

func (m *GoogleAdsField) GetDataType() enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType {
	if m != nil {
		return m.DataType
	}
	return enums.GoogleAdsFieldDataTypeEnum_UNSPECIFIED
}

func (m *GoogleAdsField) GetTypeUrl() *wrappers.StringValue {
	if m != nil {
		return m.TypeUrl
	}
	return nil
}

func (m *GoogleAdsField) GetIsRepeated() *wrappers.BoolValue {
	if m != nil {
		return m.IsRepeated
	}
	return nil
}

func init() {
	proto.RegisterType((*GoogleAdsField)(nil), "google.ads.googleads.v0.resources.GoogleAdsField")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/google_ads_field.proto", fileDescriptor_google_ads_field_c22565e2dfed71e8)
}

var fileDescriptor_google_ads_field_c22565e2dfed71e8 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6a, 0xdb, 0x3e,
	0x14, 0xc7, 0x49, 0xda, 0x5f, 0xe3, 0x2a, 0x6d, 0x7e, 0xa0, 0xde, 0x98, 0x30, 0x46, 0xba, 0x51,
	0xc8, 0x95, 0x62, 0x3a, 0xe8, 0x8a, 0x4b, 0x07, 0xce, 0xd6, 0x15, 0x06, 0x1b, 0xc1, 0xdb, 0x32,
	0x18, 0x01, 0xa3, 0x44, 0x27, 0xae, 0x41, 0xb6, 0x8c, 0x24, 0xa7, 0xe4, 0x6e, 0xcf, 0xb2, 0xcb,
	0xdd, 0xed, 0x35, 0xf6, 0x28, 0x7b, 0x8a, 0xe1, 0x7f, 0x4a, 0x4b, 0x97, 0x25, 0xdb, 0xdd, 0xf1,
	0x39, 0xdf, 0xcf, 0xf9, 0x23, 0x1f, 0x09, 0x9d, 0x87, 0x42, 0x84, 0x1c, 0x06, 0x94, 0xa9, 0x41,
	0x69, 0xe6, 0xd6, 0xc2, 0x19, 0x48, 0x50, 0x22, 0x93, 0x33, 0xa8, 0xdd, 0x01, 0x65, 0x2a, 0x98,
	0x47, 0xc0, 0x19, 0x49, 0xa5, 0xd0, 0x02, 0x1f, 0x97, 0x7e, 0x42, 0x99, 0x22, 0x86, 0x24, 0x0b,
	0x87, 0x18, 0xb2, 0x7b, 0xb9, 0x2e, 0x39, 0x24, 0x59, 0xfc, 0x30, 0x71, 0x30, 0xa3, 0x1a, 0x42,
	0x21, 0x97, 0x65, 0x85, 0xee, 0x8b, 0xbf, 0xc4, 0x19, 0xd5, 0x34, 0xd0, 0xcb, 0x14, 0x2a, 0xfe,
	0x71, 0xc5, 0x17, 0x5f, 0xd3, 0x6c, 0x3e, 0xb8, 0x95, 0x34, 0x4d, 0x41, 0xaa, 0x32, 0xfe, 0xe4,
	0x7b, 0x0b, 0x75, 0xae, 0x0b, 0x89, 0xc7, 0xd4, 0xeb, 0x3c, 0x05, 0x7e, 0x8a, 0x0e, 0xeb, 0xf6,
	0x83, 0x84, 0xc6, 0x60, 0x37, 0x7a, 0x8d, 0xfe, 0xbe, 0x7f, 0x50, 0x3b, 0xdf, 0xd1, 0x18, 0xb0,
	0x83, 0x76, 0x8b, 0x58, 0xb3, 0xd7, 0xe8, 0xb7, 0x4f, 0x1f, 0x55, 0xd3, 0x93, 0xba, 0x0c, 0x79,
	0xaf, 0x65, 0x94, 0x84, 0x63, 0xca, 0x33, 0xf0, 0x0b, 0x25, 0xe6, 0xc8, 0xaa, 0x67, 0xb3, 0x77,
	0x7a, 0x8d, 0x7e, 0xe7, 0x74, 0x44, 0xd6, 0x1d, 0x5f, 0x31, 0x1c, 0xb9, 0xdf, 0xd7, 0xcb, 0x0a,
	0xbe, 0x4a, 0xb2, 0x78, 0x4d, 0xc8, 0x37, 0x15, 0xb0, 0x8b, 0x90, 0x02, 0x0e, 0x33, 0x4d, 0xa7,
	0x1c, 0xec, 0xdd, 0xa2, 0xcb, 0xee, 0x83, 0x2e, 0x87, 0x42, 0xf0, 0xb2, 0xc7, 0x3b, 0xea, 0x9c,
	0x9d, 0x47, 0x5c, 0x83, 0x2c, 0xd8, 0xff, 0x36, 0xb3, 0x2b, 0x35, 0x3e, 0x43, 0x96, 0x12, 0xb2,
	0xac, 0xba, 0xb7, 0x91, 0x34, 0x5a, 0x7c, 0x85, 0xfe, 0x5f, 0x75, 0x10, 0xdc, 0x46, 0xfa, 0xc6,
	0x6e, 0xf5, 0x76, 0x36, 0x1e, 0x6d, 0x67, 0x05, 0x7d, 0x8a, 0xf4, 0x0d, 0x7e, 0x8b, 0x8e, 0xa8,
	0xd6, 0x32, 0x9a, 0x66, 0x1a, 0x02, 0xb3, 0x84, 0xb6, 0xb5, 0x45, 0x2a, 0x6c, 0x40, 0xbf, 0xe6,
	0xf0, 0x19, 0x6a, 0xc5, 0xa0, 0x65, 0x34, 0x53, 0xf6, 0xfe, 0x16, 0x29, 0x6a, 0x31, 0x3e, 0x47,
	0x96, 0x82, 0x30, 0x86, 0x44, 0x2b, 0x1b, 0x6d, 0x01, 0x1a, 0x35, 0xbe, 0x44, 0xed, 0xfc, 0xe7,
	0x07, 0x8b, 0xdc, 0xaf, 0xec, 0xf6, 0x16, 0x30, 0xca, 0x81, 0xc2, 0x54, 0x38, 0x46, 0xfb, 0xe6,
	0x06, 0xd8, 0x07, 0xff, 0xb0, 0x65, 0xaf, 0xa8, 0xa6, 0x1f, 0x96, 0x29, 0xfc, 0x66, 0xcb, 0xea,
	0x90, 0x6f, 0xb1, 0xca, 0xc2, 0xcf, 0x91, 0x95, 0x57, 0x0a, 0x32, 0xc9, 0xed, 0xc3, 0x2d, 0x6e,
	0x42, 0x2b, 0x57, 0x7f, 0x94, 0x1c, 0x5f, 0xa0, 0x76, 0xa4, 0x02, 0x09, 0x29, 0x50, 0x0d, 0xcc,
	0xee, 0x6c, 0xde, 0xb1, 0x48, 0xf9, 0x95, 0x7a, 0xf8, 0xa5, 0x89, 0x4e, 0x66, 0x22, 0x26, 0x1b,
	0x1f, 0x9f, 0xe1, 0xd1, 0xfd, 0x09, 0x46, 0x79, 0xde, 0x51, 0xe3, 0xf3, 0x9b, 0x8a, 0x0c, 0x05,
	0xa7, 0x49, 0x48, 0x84, 0x0c, 0x07, 0x21, 0x24, 0x45, 0xd5, 0xfa, 0x91, 0x49, 0x23, 0xf5, 0x87,
	0xf7, 0xf0, 0xc2, 0x58, 0x5f, 0x9b, 0x3b, 0xd7, 0x9e, 0xf7, 0xad, 0x79, 0x5c, 0x56, 0x22, 0x1e,
	0xbb, 0x73, 0xa2, 0x64, 0xec, 0x10, 0xb3, 0x4c, 0x3f, 0x6a, 0xcd, 0xc4, 0x63, 0x6a, 0x62, 0x34,
	0x93, 0xb1, 0x33, 0x31, 0x9a, 0x9f, 0xcd, 0x93, 0x32, 0xe0, 0xba, 0x1e, 0x53, 0xae, 0x6b, 0x54,
	0xae, 0x3b, 0x76, 0x5c, 0xd7, 0xe8, 0xa6, 0x7b, 0x45, 0xb3, 0xcf, 0x7e, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xb6, 0x0d, 0xf5, 0xb9, 0xbb, 0x05, 0x00, 0x00,
}
