package repositories

import (
	pb "github.com/jianhan/course-management-service/proto"
	mgo "gopkg.in/mgo.v2"
)

const (
	dbName           = "course-management-service"
	vesselCollection = "courses"
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
	return c.collection().Insert(courses)
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
	return c.Session.DB(dbName).C(vesselCollection)
}
