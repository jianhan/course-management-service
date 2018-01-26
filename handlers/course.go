package handlers

import (
	"context"

	pcourse "github.com/jianhan/course-management-service/proto/course"
	rp "github.com/jianhan/course-management-service/repositories"
	mgo "gopkg.in/mgo.v2"
)

type CourseManagement struct {
	Session *mgo.Session
}

func (c *CourseManagement) GetRepo() rp.CourseRepository {
	return &rp.Course{Session: c.Session.Clone()}
}

func (c *CourseManagement) Create(ctx context.Context, course *pcourse.Course, rsp *pcourse.CreateCourseResponse) error {
	repo := c.GetRepo()
	defer repo.Close()
	return repo.CreateCourses([]*pcourse.Course{course})
}
