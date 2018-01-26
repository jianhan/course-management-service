package repositories

import (
	"github.com/davecgh/go-spew/spew"
	pb "github.com/jianhan/course-management-service/proto"
	mgo "gopkg.in/mgo.v2"
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
	Session *mgo.Session
}

func (c *Course) CreateCourses(courses []*pb.Course) error {
	for _, v := range courses {
		err := c.collection().Insert(v)
		if err != nil {
			spew.Dump(err)
		}
	}
	return nil
}

func (c *Course) UpdateCourses(courses []*pb.Course) error {
	for _, v := range courses {
		c.collection().UpdateId(v.Id, v)
	}
	return nil
}

func (c *Course) DeleteCourses(courses []*pb.Course) error {
	return nil
}

func (c *Course) GetCourses() error {
	return nil
}

func (c *Course) Close() {
	c.Session.Close()
}

func (c *Course) collection() *mgo.Collection {
	return c.Session.DB(dbName).C(coursesCollection)
}
