package handlers

import (
	"context"

	pcourse "github.com/jianhan/course-management-service/proto/courses"
	"github.com/jianhan/course-management-service/repositories"
	pmysql "github.com/jianhan/pkg/proto/mysql"
	merrors "github.com/micro/go-micro/errors"
)

// API is a constant to define the name of API.
const API = "go_micro_srv_course"

// Courses handles all incomming request related to course.
type Courses struct {
	CourseRepository repositories.CourseRepository
}

// UpsertCourses upsert multiply courses.
func (c *Courses) UpsertCourses(ctx context.Context, req *pcourse.UpsertCoursesRequest, rsp *pmysql.UpsertResult) (err error) {
	if err = req.Validate(); err != nil {
		return merrors.BadRequest(API+".GetCoursesByFilters", err.Error())
	}
	result, err := c.CourseRepository.UpsertCourses(req.Courses, req.UpsertCategories)
	if err != nil {
		return
	}
	rsp.RowsAffected, err = result.RowsAffected()
	return
}

// GetCoursesByFilters retrieves courses by filters.
func (c *Courses) GetCoursesByFilters(ctx context.Context, req *pcourse.GetCoursesByFiltersRequest, rsp *pcourse.GetCoursesByFiltersResponse) (err error) {
	if err = req.Validate(); err != nil {
		return err
	}
	if rsp.Courses, err = c.CourseRepository.GetCoursesByFilters(req.FilterSet, req.Sort, req.Pagination); err != nil {
		return err
	}
	return
}

// DeleteCoursesByFilters remove courses according to filter set.
func (c *Courses) DeleteCoursesByFilters(ctx context.Context, req *pcourse.DeleteCoursesByFiltersRequest, rsp *pcourse.DeleteCoursesByFiltersResponse) (err error) {
	if err = req.Validate(); err != nil {
		return merrors.BadRequest(API+".DeleteCoursesByFilters", err.Error())
	}
	if rsp.Deleted, err = c.CourseRepository.DeleteCoursesByFilters(req.FilterSet); err != nil {
		return err
	}
	return
}
