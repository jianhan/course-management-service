package repositories

import (
	"github.com/davecgh/go-spew/spew"
	pb "github.com/jianhan/course-management-service/proto/course"
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
	GetCourses() ([]*pb.Course, error)
	Close()
}

type Course struct {
	Session *mgo.Session
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
		c.Collection(dbName, coursesCollection).UpdateId(v.Id, v)
	}
	return nil
}

func (c *Course) DeleteCourses(courses []*pb.Course) error {
	return nil
}

func (c *Course) GetCourses() ([]*pb.Course, error) {
	var courses []*pb.Course
	if err := c.Session.DB(dbName).C(coursesCollection).Find(nil).All(&courses); err != nil {
		return nil, err
	}
	return courses, nil
}

func (c *Course) Close() {
	c.Session.Close()
}

func (c *Course) Collection(dbName, collection string) *mgo.Collection {
	return c.Session.DB(dbName).C(collection)
}
