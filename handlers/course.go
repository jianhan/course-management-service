package handlers

import (
	"context"

	pb "github.com/jianhan/course-management-service/proto"
	rp "github.com/jianhan/course-management-service/repositories"
	mgo "gopkg.in/mgo.v2"
)

type CourseManagement struct {
	Session *mgo.Session
}

func (c *CourseManagement) GetRepo() rp.CourseRepository {
	return &rp.Course{Session: c.Session.Clone()}
}

func (c *CourseManagement) Create(ctx context.Context, course *pb.Course, rsp *pb.CreateCourseResponse) error {
	repo := c.GetRepo()
	defer repo.Close()
	repo.CreateCourses([]*pb.Course{course})
	return nil
}
