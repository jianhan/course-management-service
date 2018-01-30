package handlers

import (
	"context"

	pcourse "github.com/jianhan/course-management-service/proto/course"
	rp "github.com/jianhan/course-management-service/repositories"
	merrors "github.com/micro/go-micro/errors"
	mgo "gopkg.in/mgo.v2"
)

const API = "go_micro_srv_course"

// CourseManagement handles all incomming request related to course.
type CourseManagement struct {
	Session *mgo.Session
}

// GetRepo retrieve an mongo db session instance.
func (c *CourseManagement) GetRepo() rp.CourseRepository {
	t := &rp.Course{}
	t.Session = c.Session.Clone()
	return t
}

// UpsertCourses upsert multiply courses.
func (c *CourseManagement) UpsertCourses(ctx context.Context, req *pcourse.UpsertCoursesRequest, rsp *pcourse.UpsertCoursesResponse) (err error) {
	repo := c.GetRepo()
	defer repo.Close()
	if err = req.Validate(); err != nil {
		return merrors.BadRequest(API+".GetCoursesByFilters", err.Error())
	}
	if rsp.Updated, rsp.Inserted, err = repo.UpsertCourses(req.Courses); err != nil {
		return merrors.InternalServerError(API+".GetCoursesByFilters", err.Error())
	}
	return
}

// GetCoursesByFilters retrieves courses by filters.
func (c *CourseManagement) GetCoursesByFilters(ctx context.Context, req *pcourse.GetCoursesByFiltersRequest, rsp *pcourse.GetCoursesByFiltersResponse) (err error) {
	if err = req.Validate(); err != nil {
		return merrors.BadRequest(API+".GetCoursesByFilters", err.Error())
	}
	repo := c.GetRepo()
	defer repo.Close()
	if req.FilterSet == nil {
		return merrors.BadRequest(API+".GetCoursesByFilters", "Empty filter set")
	}
	if rsp.Courses, err = repo.GetCoursesByFilters(req.FilterSet); err != nil {
		return merrors.InternalServerError(API+".GetCoursesByFilters", err.Error())
	}
	return
}
