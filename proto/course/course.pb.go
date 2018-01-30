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
	DeleteCoursesByIDsRequest
	DeleteCoursesByIDsResponse
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

type DeleteCoursesByIDsRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *DeleteCoursesByIDsRequest) Reset()                    { *m = DeleteCoursesByIDsRequest{} }
func (m *DeleteCoursesByIDsRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteCoursesByIDsRequest) ProtoMessage()               {}
func (*DeleteCoursesByIDsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteCoursesByIDsRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type DeleteCoursesByIDsResponse struct {
	Removed uint32 `protobuf:"varint,1,opt,name=removed" json:"removed,omitempty"`
}

func (m *DeleteCoursesByIDsResponse) Reset()                    { *m = DeleteCoursesByIDsResponse{} }
func (m *DeleteCoursesByIDsResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteCoursesByIDsResponse) ProtoMessage()               {}
func (*DeleteCoursesByIDsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteCoursesByIDsResponse) GetRemoved() uint32 {
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
	proto.RegisterType((*DeleteCoursesByIDsRequest)(nil), "go_micro_srv_course.DeleteCoursesByIDsRequest")
	proto.RegisterType((*DeleteCoursesByIDsResponse)(nil), "go_micro_srv_course.DeleteCoursesByIDsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CourseManager service

type CourseManagerClient interface {
	UpsertCourses(ctx context.Context, in *UpsertCoursesRequest, opts ...client.CallOption) (*UpsertCoursesResponse, error)
	GetCoursesByFilters(ctx context.Context, in *GetCoursesByFiltersRequest, opts ...client.CallOption) (*GetCoursesByFiltersResponse, error)
	DeleteCoursesByIDs(ctx context.Context, in *DeleteCoursesByIDsRequest, opts ...client.CallOption) (*DeleteCoursesByIDsResponse, error)
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

func (c *courseManagerClient) DeleteCoursesByIDs(ctx context.Context, in *DeleteCoursesByIDsRequest, opts ...client.CallOption) (*DeleteCoursesByIDsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CourseManager.DeleteCoursesByIDs", in)
	out := new(DeleteCoursesByIDsResponse)
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
	DeleteCoursesByIDs(context.Context, *DeleteCoursesByIDsRequest, *DeleteCoursesByIDsResponse) error
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

func (h *CourseManager) DeleteCoursesByIDs(ctx context.Context, in *DeleteCoursesByIDsRequest, out *DeleteCoursesByIDsResponse) error {
	return h.CourseManagerHandler.DeleteCoursesByIDs(ctx, in, out)
}

func init() { proto.RegisterFile("proto/course/course.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x25, 0x5f, 0x9b, 0xec, 0x84, 0x20, 0xe4, 0x16, 0xe4, 0x6e, 0xa5, 0x36, 0x5a, 0x2e, 0x01,
	0xd1, 0xa4, 0x2a, 0x02, 0xd1, 0x03, 0x12, 0xfd, 0xa0, 0x88, 0x43, 0x2f, 0x6e, 0xb9, 0xc0, 0x21,
	0xda, 0x66, 0x27, 0xa9, 0xa5, 0x4d, 0x1c, 0x6c, 0x27, 0xa5, 0x9c, 0x39, 0xf0, 0x03, 0xf8, 0x9d,
	0xfc, 0x06, 0x64, 0x7b, 0x77, 0x53, 0xe8, 0x56, 0x0d, 0x95, 0x38, 0xd5, 0xf3, 0xf2, 0xe6, 0xf9,
	0xad, 0xe7, 0x4d, 0x61, 0x6d, 0x2a, 0x85, 0x16, 0xbd, 0x81, 0x98, 0x49, 0x85, 0xe9, 0x9f, 0xae,
	0xc5, 0xc8, 0xca, 0x48, 0xf4, 0xc7, 0x7c, 0x20, 0x45, 0x5f, 0xc9, 0x79, 0xdf, 0xfd, 0x14, 0x6c,
	0x8e, 0x84, 0x18, 0x25, 0xd8, 0xb3, 0x94, 0xb3, 0xd9, 0xb0, 0xa7, 0xf9, 0x18, 0x95, 0x8e, 0xc6,
	0x53, 0xd7, 0x15, 0xfe, 0xa8, 0x80, 0x77, 0x60, 0xb9, 0xe4, 0x01, 0x94, 0x79, 0x4c, 0x4b, 0xed,
	0x52, 0xc7, 0x67, 0x65, 0x1e, 0x13, 0x02, 0xd5, 0x49, 0x34, 0x46, 0x5a, 0xb6, 0x88, 0x3d, 0x1b,
	0x4c, 0x25, 0xb3, 0x11, 0xad, 0x38, 0xcc, 0x9c, 0x49, 0x1b, 0x9a, 0x31, 0xaa, 0x81, 0xe4, 0x53,
	0xcd, 0xc5, 0x84, 0x56, 0xed, 0x4f, 0x57, 0x21, 0x42, 0xa1, 0x3e, 0xe7, 0x8a, 0x9f, 0x25, 0x48,
	0x6b, 0xed, 0x52, 0xa7, 0xc1, 0xb2, 0x92, 0x6c, 0x43, 0x4d, 0xe9, 0x48, 0x6a, 0xea, 0xb5, 0x4b,
	0x9d, 0xe6, 0x4e, 0xd0, 0x75, 0x7e, 0xbb, 0x99, 0xdf, 0xee, 0x69, 0xe6, 0x97, 0x39, 0x22, 0x79,
	0x0e, 0x15, 0x9c, 0xc4, 0xb4, 0x7e, 0x2b, 0xdf, 0xd0, 0xc8, 0x2e, 0xc0, 0x40, 0x62, 0xa4, 0x31,
	0xee, 0x47, 0x9a, 0x36, 0x6e, 0x6d, 0xf2, 0x53, 0xf6, 0x9e, 0x36, 0xad, 0xb3, 0x69, 0x9c, 0xb5,
	0xfa, 0xb7, 0xb7, 0xa6, 0xec, 0x3d, 0x4d, 0x56, 0xa1, 0x26, 0x64, 0x8c, 0x92, 0x42, 0xbb, 0xd4,
	0x69, 0x31, 0x57, 0x90, 0x0d, 0x80, 0x41, 0xa4, 0x71, 0x24, 0x24, 0x47, 0x45, 0x9b, 0xed, 0x4a,
	0xc7, 0x67, 0x57, 0x90, 0xf0, 0x67, 0x09, 0x1a, 0x07, 0xae, 0xbc, 0xfc, 0x8f, 0xc3, 0xc8, 0xcd,
	0xd5, 0xae, 0x9a, 0xa3, 0x50, 0x77, 0x91, 0x51, 0xd4, 0xb3, 0xce, 0xb2, 0x32, 0x7c, 0x0b, 0x75,
	0x17, 0x10, 0x45, 0x5e, 0x2e, 0x48, 0xa5, 0x76, 0xa5, 0xd3, 0xdc, 0x59, 0xef, 0x16, 0x84, 0xae,
	0xeb, 0xe8, 0x0b, 0x85, 0x2d, 0xf0, 0x18, 0xaa, 0x59, 0xa2, 0xc9, 0x13, 0x68, 0x49, 0x71, 0xa1,
	0xfa, 0x38, 0x1c, 0xe2, 0x40, 0xa3, 0xfb, 0xc0, 0x2a, 0xbb, 0x6f, 0xc0, 0x77, 0x29, 0x16, 0x1e,
	0xc3, 0xea, 0xc7, 0xa9, 0x42, 0xa9, 0xd3, 0x6b, 0x19, 0x7e, 0x99, 0xa1, 0xd2, 0x77, 0xbd, 0xfd,
	0x18, 0x1e, 0xfd, 0x25, 0xa7, 0xa6, 0x62, 0xa2, 0xd0, 0x7c, 0x72, 0x3a, 0x32, 0x6b, 0xa3, 0xc5,
	0xb2, 0x92, 0x04, 0xd0, 0xe0, 0x13, 0xd3, 0x82, 0xb1, 0x7d, 0xf0, 0x16, 0xcb, 0xeb, 0xf0, 0x14,
	0xd6, 0xdf, 0x63, 0xa6, 0xb5, 0x7f, 0x79, 0xc4, 0x13, 0x8d, 0x72, 0x21, 0x7a, 0x47, 0x93, 0xdf,
	0xcb, 0xe0, 0x3b, 0xa9, 0x13, 0xd4, 0xe4, 0x21, 0x54, 0x78, 0xec, 0x04, 0x7c, 0x66, 0x8e, 0x64,
	0x13, 0x9a, 0x1a, 0xbf, 0xea, 0xbe, 0xc2, 0x48, 0x0e, 0xce, 0xd3, 0x14, 0x80, 0x81, 0x4e, 0x2c,
	0x62, 0xa6, 0x6a, 0x32, 0xa1, 0x68, 0xc5, 0x36, 0xb9, 0xc2, 0xa0, 0x26, 0x15, 0x8a, 0x56, 0x1d,
	0x6a, 0x8b, 0xc5, 0xd2, 0xd5, 0xfe, 0x71, 0xe9, 0xbc, 0xe5, 0x96, 0xee, 0xf5, 0x62, 0xdd, 0xdd,
	0x9a, 0x6e, 0x14, 0xbe, 0xc1, 0xbe, 0x10, 0xc9, 0x11, 0xc7, 0x24, 0xce, 0xff, 0x1d, 0x84, 0xbb,
	0xe0, 0xe7, 0xa8, 0x31, 0x3f, 0x8f, 0x92, 0x19, 0xda, 0xe9, 0x34, 0x98, 0x2b, 0xc8, 0x63, 0xf0,
	0xf8, 0x68, 0x22, 0xa4, 0x5b, 0x85, 0x06, 0x4b, 0xab, 0xf0, 0x33, 0x04, 0x85, 0x73, 0x71, 0xd9,
	0x79, 0x03, 0x30, 0xb4, 0x48, 0x5f, 0xa1, 0xb6, 0x82, 0x37, 0xb9, 0xca, 0xa7, 0xc0, 0xfc, 0x61,
	0x76, 0x0c, 0xb7, 0x60, 0xed, 0x10, 0x13, 0xd4, 0x98, 0xeb, 0x7f, 0x38, 0xcc, 0xb5, 0xaf, 0x4d,
	0x2b, 0x7c, 0x05, 0x41, 0x11, 0x7d, 0x91, 0x3b, 0x89, 0x63, 0x31, 0x5f, 0xe4, 0x2e, 0x2d, 0x77,
	0x7e, 0x95, 0xa1, 0xe5, 0x5a, 0x8e, 0xa3, 0x49, 0x34, 0x42, 0x49, 0xce, 0xa1, 0xf5, 0x47, 0x78,
	0xc9, 0xd3, 0x42, 0xd3, 0x45, 0xfb, 0x12, 0x3c, 0x5b, 0x86, 0xea, 0x3c, 0x85, 0xf7, 0xc8, 0x37,
	0x58, 0x29, 0x78, 0x3f, 0xd2, 0x2b, 0x14, 0xb9, 0xf9, 0xa5, 0x83, 0xed, 0xe5, 0x1b, 0xf2, 0xbb,
	0x2f, 0x80, 0x5c, 0x7f, 0x2f, 0xd2, 0x2d, 0x54, 0xba, 0x71, 0x0e, 0x41, 0x6f, 0x69, 0x7e, 0x76,
	0xf1, 0x7e, 0xe3, 0x93, 0xe7, 0x68, 0x67, 0x9e, 0x0d, 0xf3, 0x8b, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xbf, 0x08, 0x1d, 0xc5, 0x61, 0x07, 0x00, 0x00,
}
