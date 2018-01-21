package handlers

import (
	"context"

	proto "github.com/jianhan/course-management-service/proto"
)

type CourseManager struct {
}

func (c *CourseManager) Add(ctx context.Context, req *proto.AddCourseRequest, rsp *proto.AddCourseResponse) error {
	rsp.Msg = "test"
	return nil
}
