package handlers

import (
	"context"

	pcourse "github.com/jianhan/course-management-service/proto"
	"github.com/jianhan/course-management-service/repositories"
	merrors "github.com/micro/go-micro/errors"
)

// API is a constant to define the name of API.
const API = "go_micro_srv_course"

// Courses handles all incomming request related to course.
type CourseManagement struct {
	CourseRepository   repositories.CourseRepository
	CategoryRepository repositories.CategoryRepository
}

// UpsertCourses upsert multiply courses.
func (c *CourseManagement) UpsertCourses(ctx context.Context, req *pcourse.UpsertCoursesRequest, rsp *pcourse.UpsertResult) (err error) {
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
func (c *CourseManagement) GetCoursesByFilters(ctx context.Context, req *pcourse.GetCoursesByFiltersRequest, rsp *pcourse.GetCoursesByFiltersResponse) (err error) {
	if rsp.Courses, err = c.CourseRepository.GetCoursesByFilters(req.FilterSet, req.Sort, req.Pagination); err != nil {
		return err
	}
	return
}

// DeleteCoursesByFilters remove courses according to filter set.
func (c *CourseManagement) DeleteCoursesByFilters(ctx context.Context, req *pcourse.DeleteCoursesByFiltersRequest, rsp *pcourse.DeleteCoursesByFiltersResponse) (err error) {
	if err = req.Validate(); err != nil {
		return merrors.BadRequest(API+".DeleteCoursesByFilters", err.Error())
	}
	if rsp.Deleted, err = c.CourseRepository.DeleteCoursesByFilters(req.FilterSet); err != nil {
		return err
	}
	return
}
