package repositories

import (
	"github.com/davecgh/go-spew/spew"
	pb "github.com/jianhan/course-management-service/proto/course"
	jmongod "github.com/jianhan/pkg/mongod"
)

const (
	dbName            = "course-management-service"
	coursesCollection = "courses"
)

type CourseRepository interface {
	CreateCourses(courses []*pb.Course) error
	UpdateCourses(courses []*pb.Course) error
	DeleteCourses(courses []*pb.Course) error
	GetCourses() error
	Close()
}

type Course struct {
	jmongod.MRepository
}

func (c *Course) CreateCourses(courses []*pb.Course) error {
	for _, v := range courses {
		err := c.Collection(dbName, coursesCollection).Insert(v)
		if err != nil {
			spew.Dump(err)
		}
	}
	return nil
}

func (c *Course) UpdateCourses(courses []*pb.Course) error {
	for _, v := range courses {
		c.Collection(dbName, coursesCollection).UpdateId(v.XId, v)
	}
	return nil
}

func (c *Course) DeleteCourses(courses []*pb.Course) error {
	return nil
}

func (c *Course) GetCourses() error {
	return nil
}
