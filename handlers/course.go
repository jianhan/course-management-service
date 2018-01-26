package handlers

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	gpb "github.com/golang/protobuf/ptypes/empty"
	pcourse "github.com/jianhan/course-management-service/proto/course"
	rp "github.com/jianhan/course-management-service/repositories"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func (c *CourseManagement) Courses(ctx context.Context, _ *gpb.Empty, rsp *pcourse.CoursesResponse) error {
	repo := c.GetRepo()
	defer repo.Close()
	courses, err := repo.GetCourses()
	if err != nil {
		return err
	}
	for _, v := range courses {
		spew.Dump(bson.ObjectId(v.Id).Hex())
	}
	rsp.Courses = courses
	return nil
}
