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
	t := &rp.Course{}
	t.Session = c.Session.Clone()
	return t
}

// func (c *CourseManagement) Create(ctx context.Context, course *pcourse.Course, rsp *pcourse.CreateCourseResponse) error {
// 	repo := c.GetRepo()
// 	defer repo.Close()
// 	return repo.CreateCourses([]*pcourse.Course{course})
// }

// func (c *CourseManagement) Courses(ctx context.Context, _ *gpb.Empty, rsp *pcourse.CoursesResponse) error {
// 	repo := c.GetRepo()
// 	defer repo.Close()
// 	courses, err := repo.GetCourses()
// 	if err != nil {
// 		return err
// 	}
// 	for _, v := range courses {
// 		spew.Dump(bson.ObjectId(v.Id).Hex())
// 	}
// 	rsp.Courses = courses
// 	return nil
// }

func (c *CourseManagement) UpsertCourses(ctx context.Context, req *pcourse.UpsertCoursesRequest, rsp *pcourse.UpsertCoursesResponse) (err error) {
	repo := c.GetRepo()
	defer repo.Close()
	if rsp.Updated, rsp.Inserted, err = repo.UpsertCourses(req.Courses); err != nil {
		return err
	}
	return
}
