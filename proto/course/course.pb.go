// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/course/course.proto

/*
Package course is a generated protocol buffer package.

It is generated from these files:
	proto/course/course.proto

It has these top-level messages:
	Course
	Category
	Courses
	Result
	UpsertCoursesRequest
	UpsertCoursesResponse
	GetCoursesByFiltersResponse
	FilterSet
	BoolField
	GetCoursesByFiltersRequest
	DeleteCoursesRequest
	DeleteCoursesResponse
*/
package course

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Course struct {
	// @inject_tag: bson:"_id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" bson:"_id,omitempty"`
	// @inject_tag: validate:"required"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" validate:"required"`
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty"`
	// @inject_tag: validate:"required"
	Description string                     `protobuf:"bytes,4,opt,name=description" json:"description,omitempty" validate:"required"`
	Visible     bool                       `protobuf:"varint,5,opt,name=visible" json:"visible,omitempty"`
	Start       *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=start" json:"start,omitempty"`
	End         *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=end" json:"end,omitempty"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt  *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at"`
	Order      uint32                     `protobuf:"varint,10,opt,name=order" json:"order,omitempty"`
	Categories []string                   `protobuf:"bytes,11,rep,name=categories" json:"categories,omitempty"`
}

func (m *Course) Reset()                    { *m = Course{} }
func (m *Course) String() string            { return proto.CompactTextString(m) }
func (*Course) ProtoMessage()               {}
func (*Course) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Course) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Course) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Course) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Course) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Course) GetVisible() bool {
	if m != nil {
		return m.Visible
	}
	return false
}

func (m *Course) GetStart() *google_protobuf.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Course) GetEnd() *google_protobuf.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *Course) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Course) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Course) GetOrder() uint32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *Course) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

type Category struct {
	// @inject_tag: bson:"_id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" bson:"_id,omitempty"`
	// @inject_tag: validate:"required"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" validate:"required"`
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty"`
	// @inject_tag: validate:"required"
	Description string   `protobuf:"bytes,4,opt,name=description" json:"description,omitempty" validate:"required"`
	Order       uint32   `protobuf:"varint,5,opt,name=order" json:"order,omitempty"`
	Courses     []string `protobuf:"bytes,6,rep,name=courses" json:"courses,omitempty"`
}

func (m *Category) Reset()                    { *m = Category{} }
func (m *Category) String() string            { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()               {}
func (*Category) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Category) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Category) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Category) GetOrder() uint32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *Category) GetCourses() []string {
	if m != nil {
		return m.Courses
	}
	return nil
}

type Courses struct {
	Courses []*Course `protobuf:"bytes,1,rep,name=courses" json:"courses,omitempty"`
}

func (m *Courses) Reset()                    { *m = Courses{} }
func (m *Courses) String() string            { return proto.CompactTextString(m) }
func (*Courses) ProtoMessage()               {}
func (*Courses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Courses) GetCourses() []*Course {
	if m != nil {
		return m.Courses
	}
	return nil
}

type Result struct {
	RowsEffected uint64 `protobuf:"varint,1,opt,name=rows_effected,json=rowsEffected" json:"rows_effected,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Result) GetRowsEffected() uint64 {
	if m != nil {
		return m.RowsEffected
	}
	return 0
}

type UpsertCoursesRequest struct {
	Courses []*Course `protobuf:"bytes,1,rep,name=courses" json:"courses,omitempty"`
}

func (m *UpsertCoursesRequest) Reset()                    { *m = UpsertCoursesRequest{} }
func (m *UpsertCoursesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpsertCoursesRequest) ProtoMessage()               {}
func (*UpsertCoursesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpsertCoursesRequest) GetCourses() []*Course {
	if m != nil {
		return m.Courses
	}
	return nil
}

type UpsertCoursesResponse struct {
	Updated  uint32 `protobuf:"varint,1,opt,name=updated" json:"updated,omitempty"`
	Inserted uint32 `protobuf:"varint,2,opt,name=inserted" json:"inserted,omitempty"`
}

func (m *UpsertCoursesResponse) Reset()                    { *m = UpsertCoursesResponse{} }
func (m *UpsertCoursesResponse) String() string            { return proto.CompactTextString(m) }
func (*UpsertCoursesResponse) ProtoMessage()               {}
func (*UpsertCoursesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpsertCoursesResponse) GetUpdated() uint32 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *UpsertCoursesResponse) GetInserted() uint32 {
	if m != nil {
		return m.Inserted
	}
	return 0
}

type GetCoursesByFiltersResponse struct {
	Courses []*Course `protobuf:"bytes,1,rep,name=courses" json:"courses,omitempty"`
}

func (m *GetCoursesByFiltersResponse) Reset()                    { *m = GetCoursesByFiltersResponse{} }
func (m *GetCoursesByFiltersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetCoursesByFiltersResponse) ProtoMessage()               {}
func (*GetCoursesByFiltersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetCoursesByFiltersResponse) GetCourses() []*Course {
	if m != nil {
		return m.Courses
	}
	return nil
}

type FilterSet struct {
	Ids        []string                   `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	TextSearch string                     `protobuf:"bytes,2,opt,name=text_search,json=textSearch" json:"text_search,omitempty"`
	Names      []string                   `protobuf:"bytes,3,rep,name=names" json:"names,omitempty"`
	Slugs      []string                   `protobuf:"bytes,4,rep,name=slugs" json:"slugs,omitempty"`
	Start      *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=start" json:"start,omitempty"`
	End        *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=end" json:"end,omitempty"`
	Visible    *BoolField                 `protobuf:"bytes,7,opt,name=visible" json:"visible,omitempty"`
}

func (m *FilterSet) Reset()                    { *m = FilterSet{} }
func (m *FilterSet) String() string            { return proto.CompactTextString(m) }
func (*FilterSet) ProtoMessage()               {}
func (*FilterSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FilterSet) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *FilterSet) GetTextSearch() string {
	if m != nil {
		return m.TextSearch
	}
	return ""
}

func (m *FilterSet) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *FilterSet) GetSlugs() []string {
	if m != nil {
		return m.Slugs
	}
	return nil
}

func (m *FilterSet) GetStart() *google_protobuf.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *FilterSet) GetEnd() *google_protobuf.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *FilterSet) GetVisible() *BoolField {
	if m != nil {
		return m.Visible
	}
	return nil
}

type BoolField struct {
	Value  bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Ignore bool `protobuf:"varint,2,opt,name=ignore" json:"ignore,omitempty"`
}

func (m *BoolField) Reset()                    { *m = BoolField{} }
func (m *BoolField) String() string            { return proto.CompactTextString(m) }
func (*BoolField) ProtoMessage()               {}
func (*BoolField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BoolField) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func (m *BoolField) GetIgnore() bool {
	if m != nil {
		return m.Ignore
	}
	return false
}

type GetCoursesByFiltersRequest struct {
	FilterSet *FilterSet `protobuf:"bytes,1,opt,name=filter_set,json=filterSet" json:"filter_set,omitempty"`
}

func (m *GetCoursesByFiltersRequest) Reset()                    { *m = GetCoursesByFiltersRequest{} }
func (m *GetCoursesByFiltersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCoursesByFiltersRequest) ProtoMessage()               {}
func (*GetCoursesByFiltersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetCoursesByFiltersRequest) GetFilterSet() *FilterSet {
	if m != nil {
		return m.FilterSet
	}
	return nil
}

type DeleteCoursesRequest struct {
	FilterSet *FilterSet `protobuf:"bytes,1,opt,name=filter_set,json=filterSet" json:"filter_set,omitempty"`
}

func (m *DeleteCoursesRequest) Reset()                    { *m = DeleteCoursesRequest{} }
func (m *DeleteCoursesRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteCoursesRequest) ProtoMessage()               {}
func (*DeleteCoursesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteCoursesRequest) GetFilterSet() *FilterSet {
	if m != nil {
		return m.FilterSet
	}
	return nil
}

type DeleteCoursesResponse struct {
	Removed uint32 `protobuf:"varint,1,opt,name=removed" json:"removed,omitempty"`
}

func (m *DeleteCoursesResponse) Reset()                    { *m = DeleteCoursesResponse{} }
func (m *DeleteCoursesResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteCoursesResponse) ProtoMessage()               {}
func (*DeleteCoursesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteCoursesResponse) GetRemoved() uint32 {
	if m != nil {
		return m.Removed
	}
	return 0
}

func init() {
	proto.RegisterType((*Course)(nil), "go_micro_srv_course.Course")
	proto.RegisterType((*Category)(nil), "go_micro_srv_course.Category")
	proto.RegisterType((*Courses)(nil), "go_micro_srv_course.Courses")
	proto.RegisterType((*Result)(nil), "go_micro_srv_course.Result")
	proto.RegisterType((*UpsertCoursesRequest)(nil), "go_micro_srv_course.UpsertCoursesRequest")
	proto.RegisterType((*UpsertCoursesResponse)(nil), "go_micro_srv_course.UpsertCoursesResponse")
	proto.RegisterType((*GetCoursesByFiltersResponse)(nil), "go_micro_srv_course.GetCoursesByFiltersResponse")
	proto.RegisterType((*FilterSet)(nil), "go_micro_srv_course.FilterSet")
	proto.RegisterType((*BoolField)(nil), "go_micro_srv_course.BoolField")
	proto.RegisterType((*GetCoursesByFiltersRequest)(nil), "go_micro_srv_course.GetCoursesByFiltersRequest")
	proto.RegisterType((*DeleteCoursesRequest)(nil), "go_micro_srv_course.DeleteCoursesRequest")
	proto.RegisterType((*DeleteCoursesResponse)(nil), "go_micro_srv_course.DeleteCoursesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CourseManager service

type CourseManagerClient interface {
	UpsertCourses(ctx context.Context, in *UpsertCoursesRequest, opts ...client.CallOption) (*UpsertCoursesResponse, error)
	GetCoursesByFilters(ctx context.Context, in *GetCoursesByFiltersRequest, opts ...client.CallOption) (*GetCoursesByFiltersResponse, error)
	DeleteCourses(ctx context.Context, in *DeleteCoursesRequest, opts ...client.CallOption) (*DeleteCoursesResponse, error)
}

type courseManagerClient struct {
	c           client.Client
	serviceName string
}

func NewCourseManagerClient(serviceName string, c client.Client) CourseManagerClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go_micro_srv_course"
	}
	return &courseManagerClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *courseManagerClient) UpsertCourses(ctx context.Context, in *UpsertCoursesRequest, opts ...client.CallOption) (*UpsertCoursesResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CourseManager.UpsertCourses", in)
	out := new(UpsertCoursesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseManagerClient) GetCoursesByFilters(ctx context.Context, in *GetCoursesByFiltersRequest, opts ...client.CallOption) (*GetCoursesByFiltersResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CourseManager.GetCoursesByFilters", in)
	out := new(GetCoursesByFiltersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseManagerClient) DeleteCourses(ctx context.Context, in *DeleteCoursesRequest, opts ...client.CallOption) (*DeleteCoursesResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CourseManager.DeleteCourses", in)
	out := new(DeleteCoursesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CourseManager service

type CourseManagerHandler interface {
	UpsertCourses(context.Context, *UpsertCoursesRequest, *UpsertCoursesResponse) error
	GetCoursesByFilters(context.Context, *GetCoursesByFiltersRequest, *GetCoursesByFiltersResponse) error
	DeleteCourses(context.Context, *DeleteCoursesRequest, *DeleteCoursesResponse) error
}

func RegisterCourseManagerHandler(s server.Server, hdlr CourseManagerHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&CourseManager{hdlr}, opts...))
}

type CourseManager struct {
	CourseManagerHandler
}

func (h *CourseManager) UpsertCourses(ctx context.Context, in *UpsertCoursesRequest, out *UpsertCoursesResponse) error {
	return h.CourseManagerHandler.UpsertCourses(ctx, in, out)
}

func (h *CourseManager) GetCoursesByFilters(ctx context.Context, in *GetCoursesByFiltersRequest, out *GetCoursesByFiltersResponse) error {
	return h.CourseManagerHandler.GetCoursesByFilters(ctx, in, out)
}

func (h *CourseManager) DeleteCourses(ctx context.Context, in *DeleteCoursesRequest, out *DeleteCoursesResponse) error {
	return h.CourseManagerHandler.DeleteCourses(ctx, in, out)
}

func init() { proto.RegisterFile("proto/course/course.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x6f, 0x13, 0x3b,
	0x10, 0x6e, 0x7e, 0x6d, 0xb2, 0x93, 0x97, 0xa7, 0x27, 0xb7, 0x7d, 0xf2, 0x4b, 0xa5, 0x36, 0xda,
	0x77, 0x09, 0x08, 0x92, 0x52, 0x84, 0x44, 0x0f, 0x48, 0xb4, 0x85, 0x72, 0xea, 0xc5, 0x6d, 0x2f,
	0x70, 0x88, 0xb6, 0xd9, 0x49, 0x6a, 0x69, 0xb3, 0x0e, 0xb6, 0x37, 0x50, 0xce, 0x1c, 0xf8, 0x03,
	0xf8, 0xff, 0xf8, 0x57, 0x90, 0xed, 0xdd, 0x6c, 0x5b, 0x6d, 0xd4, 0x52, 0xc4, 0x29, 0x9e, 0x6f,
	0xbf, 0x19, 0x7f, 0x9e, 0xf9, 0x26, 0xf0, 0xdf, 0x5c, 0x0a, 0x2d, 0x86, 0x63, 0x91, 0x4a, 0x85,
	0xd9, 0xcf, 0xc0, 0x62, 0x64, 0x7d, 0x2a, 0x46, 0x33, 0x3e, 0x96, 0x62, 0xa4, 0xe4, 0x62, 0xe4,
	0x3e, 0x75, 0x77, 0xa6, 0x42, 0x4c, 0x63, 0x1c, 0x5a, 0xca, 0x45, 0x3a, 0x19, 0x6a, 0x3e, 0x43,
	0xa5, 0xc3, 0xd9, 0xdc, 0x65, 0x05, 0xdf, 0x6a, 0xe0, 0x1d, 0x59, 0x2e, 0xf9, 0x1b, 0xaa, 0x3c,
	0xa2, 0x95, 0x5e, 0xa5, 0xef, 0xb3, 0x2a, 0x8f, 0x08, 0x81, 0x7a, 0x12, 0xce, 0x90, 0x56, 0x2d,
	0x62, 0xcf, 0x06, 0x53, 0x71, 0x3a, 0xa5, 0x35, 0x87, 0x99, 0x33, 0xe9, 0x41, 0x3b, 0x42, 0x35,
	0x96, 0x7c, 0xae, 0xb9, 0x48, 0x68, 0xdd, 0x7e, 0xba, 0x0e, 0x11, 0x0a, 0xcd, 0x05, 0x57, 0xfc,
	0x22, 0x46, 0xda, 0xe8, 0x55, 0xfa, 0x2d, 0x96, 0x87, 0x64, 0x17, 0x1a, 0x4a, 0x87, 0x52, 0x53,
	0xaf, 0x57, 0xe9, 0xb7, 0xf7, 0xba, 0x03, 0xa7, 0x77, 0x90, 0xeb, 0x1d, 0x9c, 0xe5, 0x7a, 0x99,
	0x23, 0x92, 0x27, 0x50, 0xc3, 0x24, 0xa2, 0xcd, 0x3b, 0xf9, 0x86, 0x46, 0xf6, 0x01, 0xc6, 0x12,
	0x43, 0x8d, 0xd1, 0x28, 0xd4, 0xb4, 0x75, 0x67, 0x92, 0x9f, 0xb1, 0x0f, 0xb4, 0x49, 0x4d, 0xe7,
	0x51, 0x9e, 0xea, 0xdf, 0x9d, 0x9a, 0xb1, 0x0f, 0x34, 0xd9, 0x80, 0x86, 0x90, 0x11, 0x4a, 0x0a,
	0xbd, 0x4a, 0xbf, 0xc3, 0x5c, 0x40, 0xb6, 0x01, 0xc6, 0xa1, 0xc6, 0xa9, 0x90, 0x1c, 0x15, 0x6d,
	0xf7, 0x6a, 0x7d, 0x9f, 0x5d, 0x43, 0x82, 0xef, 0x15, 0x68, 0x1d, 0xb9, 0xf0, 0xea, 0x0f, 0x0e,
	0x63, 0x29, 0xae, 0x71, 0x5d, 0x1c, 0x85, 0xa6, 0xb3, 0x8c, 0xa2, 0x9e, 0x55, 0x96, 0x87, 0xc1,
	0x6b, 0x68, 0x3a, 0x83, 0x28, 0xf2, 0xa2, 0x20, 0x55, 0x7a, 0xb5, 0x7e, 0x7b, 0x6f, 0x6b, 0x50,
	0x62, 0xba, 0x81, 0xa3, 0x17, 0x15, 0x9e, 0x82, 0xc7, 0x50, 0xa5, 0xb1, 0x26, 0xff, 0x43, 0x47,
	0x8a, 0x4f, 0x6a, 0x84, 0x93, 0x09, 0x8e, 0x35, 0xba, 0x07, 0xd6, 0xd9, 0x5f, 0x06, 0x7c, 0x9b,
	0x61, 0xc1, 0x09, 0x6c, 0x9c, 0xcf, 0x15, 0x4a, 0x9d, 0x5d, 0xcb, 0xf0, 0x63, 0x8a, 0x4a, 0x3f,
	0xf4, 0xf6, 0x13, 0xd8, 0xbc, 0x55, 0x4e, 0xcd, 0x45, 0xa2, 0xd0, 0x3c, 0x39, 0x1b, 0x99, 0x95,
	0xd1, 0x61, 0x79, 0x48, 0xba, 0xd0, 0xe2, 0x89, 0x49, 0xc1, 0xc8, 0x36, 0xbc, 0xc3, 0x96, 0x71,
	0x70, 0x06, 0x5b, 0xef, 0x30, 0xaf, 0x75, 0x78, 0x75, 0xcc, 0x63, 0x8d, 0xb2, 0x28, 0xfa, 0x40,
	0x91, 0x5f, 0xab, 0xe0, 0xbb, 0x52, 0xa7, 0xa8, 0xc9, 0x3f, 0x50, 0xe3, 0x91, 0x2b, 0xe0, 0x33,
	0x73, 0x24, 0x3b, 0xd0, 0xd6, 0xf8, 0x59, 0x8f, 0x14, 0x86, 0x72, 0x7c, 0x99, 0xb9, 0x00, 0x0c,
	0x74, 0x6a, 0x11, 0x33, 0x55, 0xe3, 0x09, 0x45, 0x6b, 0x36, 0xc9, 0x05, 0x06, 0x35, 0xae, 0x50,
	0xb4, 0xee, 0x50, 0x1b, 0x14, 0x4b, 0xd7, 0xf8, 0xc5, 0xa5, 0xf3, 0xee, 0xb7, 0x74, 0x2f, 0x8b,
	0x75, 0x77, 0x6b, 0xba, 0x5d, 0xda, 0x83, 0x43, 0x21, 0xe2, 0x63, 0x8e, 0x71, 0xb4, 0xfc, 0x3b,
	0x08, 0xf6, 0xc1, 0x5f, 0xa2, 0x46, 0xfc, 0x22, 0x8c, 0x53, 0xb4, 0xd3, 0x69, 0x31, 0x17, 0x90,
	0x7f, 0xc1, 0xe3, 0xd3, 0x44, 0x48, 0xb7, 0x0a, 0x2d, 0x96, 0x45, 0xc1, 0x07, 0xe8, 0x96, 0xce,
	0xc5, 0x79, 0xe7, 0x15, 0xc0, 0xc4, 0x22, 0x23, 0x85, 0xda, 0x16, 0x5c, 0xa5, 0x6a, 0x39, 0x05,
	0xe6, 0x4f, 0xf2, 0x63, 0x70, 0x0e, 0x1b, 0x6f, 0x30, 0x46, 0x8d, 0xb7, 0x2c, 0xf9, 0x9b, 0x65,
	0x9f, 0xc1, 0xe6, 0xad, 0xb2, 0x85, 0x35, 0x25, 0xce, 0xc4, 0xa2, 0xb0, 0x66, 0x16, 0xee, 0xfd,
	0xa8, 0x42, 0xc7, 0xb1, 0x4f, 0xc2, 0x24, 0x9c, 0xa2, 0x24, 0x97, 0xd0, 0xb9, 0xe1, 0x6f, 0xf2,
	0xa8, 0x54, 0x40, 0xd9, 0x4a, 0x75, 0x1f, 0xdf, 0x87, 0xea, 0x34, 0x05, 0x6b, 0xe4, 0x0b, 0xac,
	0x97, 0xb4, 0x98, 0x0c, 0x4b, 0x8b, 0xac, 0x1e, 0x46, 0x77, 0xf7, 0xfe, 0x09, 0xcb, 0xbb, 0x2f,
	0xa1, 0x73, 0xa3, 0x55, 0x2b, 0x5e, 0x59, 0x36, 0xa5, 0x15, 0xaf, 0x2c, 0xed, 0x7c, 0xb0, 0x76,
	0xd8, 0x7a, 0xef, 0x39, 0xc6, 0x85, 0x67, 0x0d, 0xfe, 0xfc, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd3, 0x72, 0xfb, 0x44, 0x75, 0x07, 0x00, 0x00,
}
