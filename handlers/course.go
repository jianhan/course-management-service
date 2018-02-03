package handlers

import (
	"context"

	pcourse "github.com/jianhan/course-management-service/proto/course"
	"github.com/jianhan/course-management-service/repositories"
	merrors "github.com/micro/go-micro/errors"
)

// API is a constant to define the name of API.
const API = "go_micro_srv_course"

// CourseManagement handles all incomming request related to course.
type CourseManagement struct {
	CourseRepository repositories.CourseRepository
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

// // GetCoursesByFilters retrieves courses by filters.
// func (c *CourseManagement) GetCoursesByFilters(ctx context.Context, req *pcourse.GetCoursesByFiltersRequest, rsp *pcourse.GetCoursesByFiltersResponse) (err error) {
// 	// if err = req.Validate(); err != nil {
// 	// 	return merrors.BadRequest(API+".GetCoursesByFilters", err.Error())
// 	// }
// 	// if req.FilterSet == nil {
// 	// 	return merrors.BadRequest(API+".GetCoursesByFilters", "Empty filter set")
// 	// }
// 	// repo := c.GetCourseRepo()
// 	// defer repo.Close()
// 	// if rsp.Courses, err = repo.GetCoursesByFilters(req.FilterSet); err != nil {
// 	// 	return merrors.InternalServerError(API+".GetCoursesByFilters", err.Error())
// 	// }
// 	return
// }
//
// // DeleteCoursesByIDs remove courses by IDs.
// func (c *CourseManagement) DeleteCoursesByIDs(ctx context.Context, req *pcourse.DeleteCoursesByIDsRequest, rsp *pcourse.DeleteCoursesByIDsResponse) (err error) {
// 	// if err = req.Validate(); err != nil {
// 	// 	return merrors.BadRequest(API+".DeleteCoursesByIDs", err.Error())
// 	// }
// 	// repo := c.GetCourseRepo()
// 	// defer repo.Close()
// 	// if rsp.Removed, err = repo.DeleteCoursesByIDs(req.Ids); err != nil {
// 	// 	return merrors.BadRequest(API+".DeleteCoursesByIDs", err.Error())
// 	// }
// 	return
// }
