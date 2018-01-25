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
	AddCourses(courses []*pb.Course) error
	UpdateCourses(courses []*pb.Course) error
	DeleteCourses(courses []*pb.Course) error
	GetCourses() error
	Close() *mgo.Session
}

type Course struct {
	session *mgo.Session
}

func (c *Course) AddCourses(courses []*pb.Course) error {
	return nil
}

func (c *Course) UpdateCourses(courses []*pb.Course) error {
	return nil
}

func (c *Course) DeleteCourses(courses []*pb.Course) error {
	return nil
}

func (c *Course) GetCourses() error {
	return nil
}

func (c *Course) Close() {
	c.session.Close()
}

func (c *Course) collection() *mgo.Collection {
	return c.session.DB(dbName).C(vesselCollection)
}
