package handlers

import (
	"context"

	proto "github.com/jianhan/course-management-service/proto"
	mgo "gopkg.in/mgo.v2"
)

type CourseManager struct {
	session *mgo.Session
}

func (c *CourseManager) Add(ctx context.Context, req *proto.AddCourseRequest, rsp *proto.AddCourseResponse) error {
	rsp.Msg = "test"
	return nil
}
