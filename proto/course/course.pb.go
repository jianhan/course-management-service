// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/course/course.proto

/*
Package course is a generated protocol buffer package.

It is generated from these files:
	proto/course/course.proto

It has these top-level messages:
	Course
	Category
	UpsertResult
	BoolField
	UpsertCoursesRequest
	FilterSet
	Pagination
	Sort
	GetCoursesByFiltersRequest
	GetCoursesByFiltersResponse
	DeleteCoursesByIDsRequest
	DeleteCoursesByIDsResponse
	CategoryFilterSet
	UpsertCategoriesRequest
	UpsertCategoriesResponse
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

type SortDirection int32

const (
	SortDirection_ASC  SortDirection = 0
	SortDirection_DESC SortDirection = 1
)

var SortDirection_name = map[int32]string{
	0: "ASC",
	1: "DESC",
}
var SortDirection_value = map[string]int32{
	"ASC":  0,
	"DESC": 1,
}

func (x SortDirection) String() string {
	return proto.EnumName(SortDirection_name, int32(x))
}
func (SortDirection) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Course struct {
	// @inject_tag: bson:"_id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty"`
	// @inject_tag: validate:"required"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required"`
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty"`
	// @inject_tag: bson:"display_order,omitempty"
	DisplayOrder uint64 `protobuf:"varint,4,opt,name=display_order,json=displayOrder" json:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty" bson:"display_order,omitempty"`
	// @inject_tag: validate:"required"
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required"`
	Visible     bool   `protobuf:"varint,6,opt,name=visible" json:"visible,omitempty"`
	// @inject_tag: validate:"required"
	Start *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=start" json:"start,omitempty" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required"`
	// @inject_tag: validate:"required"
	End *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=end" json:"end,omitempty" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt  *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at"`
	Categories map[string]string          `protobuf:"bytes,11,rep,name=categories" json:"categories,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
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

func (m *Course) GetDisplayOrder() uint64 {
	if m != nil {
		return m.DisplayOrder
	}
	return 0
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

func (m *Course) GetCategories() map[string]string {
	if m != nil {
		return m.Categories
	}
	return nil
}

type Category struct {
	// @inject_tag: bson:"_id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty" bson:"_id,omitempty"`
	// @inject_tag: validate:"required"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required"`
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty"`
	// @inject_tag: validate:"required"
	Description  string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required" validate:"required"`
	DisplayOrder uint32 `protobuf:"varint,5,opt,name=display_order,json=displayOrder" json:"display_order,omitempty"`
	Visible      bool   `protobuf:"varint,6,opt,name=visible" json:"visible,omitempty"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at" bson:"updated_at"`
	Courses   map[string]string          `protobuf:"bytes,9,rep,name=courses" json:"courses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
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

func (m *Category) GetDisplayOrder() uint32 {
	if m != nil {
		return m.DisplayOrder
	}
	return 0
}

func (m *Category) GetVisible() bool {
	if m != nil {
		return m.Visible
	}
	return false
}

func (m *Category) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Category) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Category) GetCourses() map[string]string {
	if m != nil {
		return m.Courses
	}
	return nil
}

type UpsertResult struct {
	RowsAffected int64 `protobuf:"varint,1,opt,name=rows_affected,json=rowsAffected" json:"rows_affected,omitempty"`
	Updated      int64 `protobuf:"varint,2,opt,name=updated" json:"updated,omitempty"`
	Inserted     int64 `protobuf:"varint,3,opt,name=inserted" json:"inserted,omitempty"`
}

func (m *UpsertResult) Reset()                    { *m = UpsertResult{} }
func (m *UpsertResult) String() string            { return proto.CompactTextString(m) }
func (*UpsertResult) ProtoMessage()               {}
func (*UpsertResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpsertResult) GetRowsAffected() int64 {
	if m != nil {
		return m.RowsAffected
	}
	return 0
}

func (m *UpsertResult) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *UpsertResult) GetInserted() int64 {
	if m != nil {
		return m.Inserted
	}
	return 0
}

type BoolField struct {
	Value  bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Ignore bool `protobuf:"varint,2,opt,name=ignore" json:"ignore,omitempty"`
}

func (m *BoolField) Reset()                    { *m = BoolField{} }
func (m *BoolField) String() string            { return proto.CompactTextString(m) }
func (*BoolField) ProtoMessage()               {}
func (*BoolField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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

type UpsertCoursesRequest struct {
	Courses          []*Course `protobuf:"bytes,1,rep,name=courses" json:"courses,omitempty"`
	UpsertCategories bool      `protobuf:"varint,2,opt,name=upsert_categories,json=upsertCategories" json:"upsert_categories,omitempty"`
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

func (m *UpsertCoursesRequest) GetUpsertCategories() bool {
	if m != nil {
		return m.UpsertCategories
	}
	return false
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
func (*FilterSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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

type Pagination struct {
	Page    uint64 `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	PerPage uint64 `protobuf:"varint,2,opt,name=per_page,json=perPage" json:"per_page,omitempty"`
}

func (m *Pagination) Reset()                    { *m = Pagination{} }
func (m *Pagination) String() string            { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()               {}
func (*Pagination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Pagination) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pagination) GetPerPage() uint64 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

type Sort struct {
	Field     string        `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	Direction SortDirection `protobuf:"varint,2,opt,name=direction,enum=go_micro_srv_course.SortDirection" json:"direction,omitempty"`
}

func (m *Sort) Reset()                    { *m = Sort{} }
func (m *Sort) String() string            { return proto.CompactTextString(m) }
func (*Sort) ProtoMessage()               {}
func (*Sort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Sort) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *Sort) GetDirection() SortDirection {
	if m != nil {
		return m.Direction
	}
	return SortDirection_ASC
}

type GetCoursesByFiltersRequest struct {
	FilterSet  *FilterSet  `protobuf:"bytes,1,opt,name=filter_set,json=filterSet" json:"filter_set,omitempty"`
	Pagination *Pagination `protobuf:"bytes,8,opt,name=pagination" json:"pagination,omitempty"`
	Sort       *Sort       `protobuf:"bytes,9,opt,name=sort" json:"sort,omitempty"`
}

func (m *GetCoursesByFiltersRequest) Reset()                    { *m = GetCoursesByFiltersRequest{} }
func (m *GetCoursesByFiltersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCoursesByFiltersRequest) ProtoMessage()               {}
func (*GetCoursesByFiltersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetCoursesByFiltersRequest) GetFilterSet() *FilterSet {
	if m != nil {
		return m.FilterSet
	}
	return nil
}

func (m *GetCoursesByFiltersRequest) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *GetCoursesByFiltersRequest) GetSort() *Sort {
	if m != nil {
		return m.Sort
	}
	return nil
}

type GetCoursesByFiltersResponse struct {
	Courses []*Course `protobuf:"bytes,1,rep,name=courses" json:"courses,omitempty"`
}

func (m *GetCoursesByFiltersResponse) Reset()                    { *m = GetCoursesByFiltersResponse{} }
func (m *GetCoursesByFiltersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetCoursesByFiltersResponse) ProtoMessage()               {}
func (*GetCoursesByFiltersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetCoursesByFiltersResponse) GetCourses() []*Course {
	if m != nil {
		return m.Courses
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

type CategoryFilterSet struct {
	Ids        []string   `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	TextSearch string     `protobuf:"bytes,2,opt,name=text_search,json=textSearch" json:"text_search,omitempty"`
	Names      []string   `protobuf:"bytes,3,rep,name=names" json:"names,omitempty"`
	Slugs      []string   `protobuf:"bytes,4,rep,name=slugs" json:"slugs,omitempty"`
	Visible    *BoolField `protobuf:"bytes,5,opt,name=visible" json:"visible,omitempty"`
}

func (m *CategoryFilterSet) Reset()                    { *m = CategoryFilterSet{} }
func (m *CategoryFilterSet) String() string            { return proto.CompactTextString(m) }
func (*CategoryFilterSet) ProtoMessage()               {}
func (*CategoryFilterSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CategoryFilterSet) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *CategoryFilterSet) GetTextSearch() string {
	if m != nil {
		return m.TextSearch
	}
	return ""
}

func (m *CategoryFilterSet) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *CategoryFilterSet) GetSlugs() []string {
	if m != nil {
		return m.Slugs
	}
	return nil
}

func (m *CategoryFilterSet) GetVisible() *BoolField {
	if m != nil {
		return m.Visible
	}
	return nil
}

type UpsertCategoriesRequest struct {
	Categories []*Category `protobuf:"bytes,1,rep,name=categories" json:"categories,omitempty"`
}

func (m *UpsertCategoriesRequest) Reset()                    { *m = UpsertCategoriesRequest{} }
func (m *UpsertCategoriesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpsertCategoriesRequest) ProtoMessage()               {}
func (*UpsertCategoriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UpsertCategoriesRequest) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type UpsertCategoriesResponse struct {
	Updated  uint32 `protobuf:"varint,1,opt,name=updated" json:"updated,omitempty"`
	Inserted uint32 `protobuf:"varint,2,opt,name=inserted" json:"inserted,omitempty"`
}

func (m *UpsertCategoriesResponse) Reset()                    { *m = UpsertCategoriesResponse{} }
func (m *UpsertCategoriesResponse) String() string            { return proto.CompactTextString(m) }
func (*UpsertCategoriesResponse) ProtoMessage()               {}
func (*UpsertCategoriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *UpsertCategoriesResponse) GetUpdated() uint32 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *UpsertCategoriesResponse) GetInserted() uint32 {
	if m != nil {
		return m.Inserted
	}
	return 0
}

func init() {
	proto.RegisterType((*Course)(nil), "go_micro_srv_course.Course")
	proto.RegisterType((*Category)(nil), "go_micro_srv_course.Category")
	proto.RegisterType((*UpsertResult)(nil), "go_micro_srv_course.UpsertResult")
	proto.RegisterType((*BoolField)(nil), "go_micro_srv_course.BoolField")
	proto.RegisterType((*UpsertCoursesRequest)(nil), "go_micro_srv_course.UpsertCoursesRequest")
	proto.RegisterType((*FilterSet)(nil), "go_micro_srv_course.FilterSet")
	proto.RegisterType((*Pagination)(nil), "go_micro_srv_course.Pagination")
	proto.RegisterType((*Sort)(nil), "go_micro_srv_course.Sort")
	proto.RegisterType((*GetCoursesByFiltersRequest)(nil), "go_micro_srv_course.GetCoursesByFiltersRequest")
	proto.RegisterType((*GetCoursesByFiltersResponse)(nil), "go_micro_srv_course.GetCoursesByFiltersResponse")
	proto.RegisterType((*DeleteCoursesByIDsRequest)(nil), "go_micro_srv_course.DeleteCoursesByIDsRequest")
	proto.RegisterType((*DeleteCoursesByIDsResponse)(nil), "go_micro_srv_course.DeleteCoursesByIDsResponse")
	proto.RegisterType((*CategoryFilterSet)(nil), "go_micro_srv_course.CategoryFilterSet")
	proto.RegisterType((*UpsertCategoriesRequest)(nil), "go_micro_srv_course.UpsertCategoriesRequest")
	proto.RegisterType((*UpsertCategoriesResponse)(nil), "go_micro_srv_course.UpsertCategoriesResponse")
	proto.RegisterEnum("go_micro_srv_course.SortDirection", SortDirection_name, SortDirection_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CourseManager service

type CourseManagerClient interface {
	// Courses
	UpsertCourses(ctx context.Context, in *UpsertCoursesRequest, opts ...client.CallOption) (*UpsertResult, error)
	GetCoursesByFilters(ctx context.Context, in *GetCoursesByFiltersRequest, opts ...client.CallOption) (*GetCoursesByFiltersResponse, error)
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

func (c *courseManagerClient) UpsertCourses(ctx context.Context, in *UpsertCoursesRequest, opts ...client.CallOption) (*UpsertResult, error) {
	req := c.c.NewRequest(c.serviceName, "CourseManager.UpsertCourses", in)
	out := new(UpsertResult)
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

// Server API for CourseManager service

type CourseManagerHandler interface {
	// Courses
	UpsertCourses(context.Context, *UpsertCoursesRequest, *UpsertResult) error
	GetCoursesByFilters(context.Context, *GetCoursesByFiltersRequest, *GetCoursesByFiltersResponse) error
}

func RegisterCourseManagerHandler(s server.Server, hdlr CourseManagerHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&CourseManager{hdlr}, opts...))
}

type CourseManager struct {
	CourseManagerHandler
}

func (h *CourseManager) UpsertCourses(ctx context.Context, in *UpsertCoursesRequest, out *UpsertResult) error {
	return h.CourseManagerHandler.UpsertCourses(ctx, in, out)
}

func (h *CourseManager) GetCoursesByFilters(ctx context.Context, in *GetCoursesByFiltersRequest, out *GetCoursesByFiltersResponse) error {
	return h.CourseManagerHandler.GetCoursesByFilters(ctx, in, out)
}

func init() { proto.RegisterFile("proto/course/course.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 946 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x6d, 0x6f, 0x1b, 0x45,
	0x10, 0xce, 0xf9, 0xce, 0x6f, 0xe3, 0xb8, 0xa4, 0xdb, 0x0a, 0x2e, 0xae, 0x20, 0xe6, 0xf8, 0x62,
	0x5a, 0xea, 0x54, 0x41, 0xa0, 0xb6, 0xa8, 0x82, 0x24, 0x6e, 0x11, 0x42, 0x88, 0xe8, 0xdc, 0x4a,
	0x88, 0x0f, 0x9c, 0x2e, 0xbe, 0xf1, 0xb1, 0xe2, 0xec, 0x3d, 0x76, 0xd7, 0x01, 0xf7, 0x2b, 0xfc,
	0x0b, 0xfe, 0x02, 0xff, 0x84, 0xbf, 0xc2, 0x8f, 0x40, 0xfb, 0x72, 0x67, 0x27, 0xb9, 0xbc, 0x7e,
	0xe0, 0x93, 0x77, 0x66, 0xe7, 0x99, 0xdb, 0x99, 0xe7, 0x99, 0x31, 0x6c, 0xe7, 0x9c, 0x49, 0xb6,
	0x3b, 0x61, 0x0b, 0x2e, 0xd0, 0xfe, 0x0c, 0xb5, 0x8f, 0xdc, 0x4b, 0x59, 0x34, 0xa3, 0x13, 0xce,
	0x22, 0xc1, 0x4f, 0x22, 0x73, 0xd5, 0xdb, 0x49, 0x19, 0x4b, 0x33, 0xdc, 0xd5, 0x21, 0xc7, 0x8b,
	0xe9, 0xae, 0xa4, 0x33, 0x14, 0x32, 0x9e, 0xe5, 0x06, 0x15, 0xfc, 0xe1, 0x41, 0xe3, 0x50, 0xc7,
	0x92, 0x3b, 0x50, 0xa3, 0x89, 0xef, 0xf4, 0x9d, 0x41, 0x3b, 0xac, 0xd1, 0x84, 0x10, 0xf0, 0xe6,
	0xf1, 0x0c, 0xfd, 0x9a, 0xf6, 0xe8, 0xb3, 0xf2, 0x89, 0x6c, 0x91, 0xfa, 0xae, 0xf1, 0xa9, 0x33,
	0xf9, 0x08, 0xba, 0x09, 0x15, 0x79, 0x16, 0x2f, 0x23, 0xc6, 0x13, 0xe4, 0xbe, 0xd7, 0x77, 0x06,
	0x5e, 0xb8, 0x69, 0x9d, 0xdf, 0x2b, 0x1f, 0xe9, 0x43, 0x27, 0x41, 0x31, 0xe1, 0x34, 0x97, 0x94,
	0xcd, 0xfd, 0xba, 0xc6, 0xaf, 0xbb, 0x88, 0x0f, 0xcd, 0x13, 0x2a, 0xe8, 0x71, 0x86, 0x7e, 0xa3,
	0xef, 0x0c, 0x5a, 0x61, 0x61, 0x92, 0x27, 0x50, 0x17, 0x32, 0xe6, 0xd2, 0x6f, 0xf6, 0x9d, 0x41,
	0x67, 0xaf, 0x37, 0x34, 0x45, 0x0d, 0x8b, 0xa2, 0x86, 0xaf, 0x8b, 0xa2, 0x42, 0x13, 0x48, 0x3e,
	0x01, 0x17, 0xe7, 0x89, 0xdf, 0xba, 0x32, 0x5e, 0x85, 0x91, 0x67, 0x00, 0x13, 0x8e, 0xb1, 0xc4,
	0x24, 0x8a, 0xa5, 0xdf, 0xbe, 0x12, 0xd4, 0xb6, 0xd1, 0xfb, 0x52, 0x41, 0x17, 0x79, 0x52, 0x40,
	0xe1, 0x6a, 0xa8, 0x8d, 0xde, 0x97, 0xe4, 0x5b, 0x80, 0x49, 0x2c, 0x31, 0x65, 0x9c, 0xa2, 0xf0,
	0x3b, 0x7d, 0x77, 0xd0, 0xd9, 0x7b, 0x34, 0xac, 0x20, 0x71, 0x78, 0x68, 0x7f, 0xca, 0xe8, 0x97,
	0x73, 0xc9, 0x97, 0xe1, 0x1a, 0xbc, 0xf7, 0x02, 0xde, 0x39, 0x73, 0x4d, 0xb6, 0xc0, 0xfd, 0x05,
	0x97, 0x96, 0x4f, 0x75, 0x24, 0xf7, 0xa1, 0x7e, 0x12, 0x67, 0x8b, 0x82, 0x51, 0x63, 0x3c, 0xaf,
	0x3d, 0x75, 0x82, 0xbf, 0x5c, 0x68, 0x59, 0xfc, 0xf2, 0xd6, 0x3a, 0x38, 0x43, 0xb1, 0x77, 0x9e,
	0xe2, 0x73, 0x4a, 0x51, 0x32, 0xe8, 0x9e, 0x51, 0xca, 0xc5, 0x3a, 0x38, 0xcd, 0x53, 0xf3, 0xf6,
	0x3c, 0xb5, 0x6e, 0xc2, 0xd3, 0x08, 0x9a, 0x86, 0x07, 0xe1, 0xb7, 0x35, 0x49, 0x0f, 0xab, 0x49,
	0xb2, 0xed, 0xb3, 0x6c, 0x59, 0x8e, 0x0a, 0x68, 0xef, 0x39, 0x6c, 0xae, 0x5f, 0xdc, 0x88, 0x1d,
	0x0a, 0x9b, 0x6f, 0x72, 0x81, 0x5c, 0x86, 0x28, 0x16, 0x99, 0x54, 0x6d, 0xe4, 0xec, 0x37, 0x11,
	0xc5, 0xd3, 0x29, 0x4e, 0x24, 0x1a, 0xae, 0xdc, 0x70, 0x53, 0x39, 0xf7, 0xad, 0x4f, 0xb5, 0xd1,
	0xd6, 0xa0, 0x13, 0xba, 0x61, 0x61, 0x92, 0x1e, 0xb4, 0xe8, 0x5c, 0xa5, 0xc3, 0x44, 0xf3, 0xe7,
	0x86, 0xa5, 0x1d, 0x3c, 0x83, 0xf6, 0x01, 0x63, 0xd9, 0x2b, 0x8a, 0x59, 0xb2, 0x7a, 0x91, 0xa3,
	0x79, 0x30, 0x06, 0x79, 0x17, 0x1a, 0x34, 0x9d, 0x33, 0x6e, 0x1e, 0xda, 0x0a, 0xad, 0x15, 0xbc,
	0x85, 0xfb, 0xe6, 0x95, 0xb6, 0xce, 0x10, 0x7f, 0x5d, 0xa0, 0x90, 0xe4, 0xb3, 0x55, 0xff, 0x1c,
	0xdd, 0xbf, 0x07, 0x97, 0x88, 0xbc, 0x6c, 0x18, 0x79, 0x04, 0x77, 0x17, 0x3a, 0x5d, 0xb4, 0x36,
	0x25, 0xe6, 0x8b, 0x5b, 0xe6, 0x62, 0x25, 0xf8, 0xe0, 0xcf, 0x1a, 0xb4, 0x5f, 0xd1, 0x4c, 0x22,
	0x1f, 0xa3, 0x54, 0xbd, 0xa5, 0x89, 0xf9, 0x5a, 0x3b, 0x54, 0x47, 0xb2, 0x03, 0x1d, 0x89, 0xbf,
	0xcb, 0x48, 0x60, 0xcc, 0x27, 0x3f, 0xdb, 0x0e, 0x83, 0x72, 0x8d, 0xb5, 0x47, 0x95, 0xaa, 0x74,
	0x2d, 0x7c, 0x57, 0x83, 0x8c, 0xa1, 0xbc, 0x4a, 0xd9, 0xc2, 0xf7, 0x8c, 0x57, 0x1b, 0xab, 0x75,
	0x54, 0xbf, 0xe1, 0x3a, 0x6a, 0x5c, 0x6f, 0x1d, 0x3d, 0x5d, 0x0d, 0x80, 0xd1, 0xf8, 0x07, 0x95,
	0x0d, 0x2b, 0x79, 0x2a, 0x07, 0x24, 0xf8, 0x02, 0xe0, 0x28, 0x4e, 0xe9, 0x3c, 0xd6, 0xd3, 0x46,
	0xc0, 0xcb, 0xe3, 0xd4, 0xb0, 0xe7, 0x85, 0xfa, 0x4c, 0xb6, 0xa1, 0x95, 0x23, 0x8f, 0xb4, 0xbf,
	0xa6, 0xfd, 0xcd, 0x1c, 0xf9, 0x51, 0x9c, 0x62, 0xf0, 0x13, 0x78, 0x63, 0xc6, 0xa5, 0x2a, 0x7a,
	0xaa, 0xd2, 0x5a, 0x6d, 0x1a, 0x83, 0x7c, 0x05, 0xed, 0x84, 0x72, 0x9c, 0xe8, 0xd1, 0x56, 0xc8,
	0x3b, 0x7b, 0x41, 0xe5, 0xb3, 0x54, 0x8e, 0x51, 0x11, 0x19, 0xae, 0x40, 0xc1, 0x3f, 0x0e, 0xf4,
	0xbe, 0xc6, 0x42, 0x1d, 0x07, 0x4b, 0xc3, 0x57, 0x29, 0x93, 0x17, 0x00, 0x53, 0xed, 0x89, 0x04,
	0x4a, 0xfd, 0xed, 0x8b, 0x0a, 0x2f, 0x89, 0x0e, 0xdb, 0xd3, 0x92, 0xf3, 0x2f, 0x01, 0xf2, 0xb2,
	0x74, 0x3b, 0xe0, 0x3b, 0x95, 0xf0, 0x55, 0x87, 0xc2, 0x35, 0x08, 0x79, 0x0c, 0x9e, 0x60, 0xbc,
	0x58, 0xff, 0xdb, 0x17, 0xd6, 0x16, 0xea, 0xb0, 0xe0, 0x35, 0x3c, 0xa8, 0x2c, 0x46, 0xe4, 0x6c,
	0x2e, 0xf0, 0x96, 0xa2, 0x0f, 0x1e, 0xc3, 0xf6, 0x08, 0x33, 0x94, 0x58, 0x26, 0xfe, 0x66, 0x54,
	0x76, 0xe8, 0x9c, 0xac, 0x83, 0xcf, 0xa1, 0x57, 0x15, 0x6e, 0xdf, 0xe0, 0x43, 0x93, 0xe3, 0x8c,
	0x9d, 0xd8, 0x05, 0xd1, 0x0d, 0x0b, 0x33, 0xf8, 0xdb, 0x81, 0xbb, 0xc5, 0xbe, 0xfa, 0xdf, 0xc6,
	0x66, 0x4d, 0xd6, 0xf5, 0x9b, 0xc9, 0xfa, 0x07, 0x78, 0xef, 0xcd, 0x99, 0x89, 0x5f, 0x53, 0xcd,
	0xda, 0x7a, 0x30, 0xad, 0x7e, 0xff, 0xd2, 0xfd, 0xbc, 0xfe, 0xb7, 0x19, 0x1c, 0x81, 0x7f, 0x3e,
	0xf3, 0xaa, 0x7d, 0xc5, 0x02, 0xb5, 0xed, 0xab, 0x5a, 0xa0, 0x35, 0x7d, 0x55, 0xda, 0x0f, 0x03,
	0xe8, 0x9e, 0x9a, 0x00, 0xd2, 0x04, 0x77, 0x7f, 0x7c, 0xb8, 0xb5, 0x41, 0x5a, 0xe0, 0x8d, 0x5e,
	0x8e, 0x0f, 0xb7, 0x9c, 0xbd, 0x7f, 0x1d, 0xe8, 0x1a, 0xc6, 0xbe, 0x8b, 0xe7, 0x71, 0x8a, 0x9c,
	0x44, 0xd0, 0x3d, 0xb5, 0x3b, 0xc9, 0xc7, 0x95, 0x35, 0x54, 0xed, 0xd7, 0xde, 0x87, 0x97, 0x84,
	0x9a, 0x3f, 0x8c, 0x60, 0x83, 0xbc, 0x85, 0x7b, 0x15, 0x72, 0x25, 0xbb, 0x95, 0xd8, 0x8b, 0xa7,
	0xb4, 0xf7, 0xe4, 0xfa, 0x00, 0xd3, 0xc6, 0x60, 0xe3, 0xa0, 0xf5, 0x63, 0xc3, 0xc4, 0x1d, 0x37,
	0xf4, 0xca, 0xfb, 0xf4, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x30, 0xb2, 0xfe, 0xc6, 0x0a,
	0x00, 0x00,
}
