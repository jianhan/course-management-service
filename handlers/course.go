package handlers

import (
	"context"

	pcourse "github.com/jianhan/course-management-service/proto/course"
	rp "github.com/jianhan/course-management-service/repositories"
	mgo "gopkg.in/mgo.v2"
)

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
	if err := req.Validate(); err != nil {
		return err
	}
	if rsp.Updated, rsp.Inserted, err = repo.UpsertCourses(req.Courses); err != nil {
		return err
	}
	return
}
