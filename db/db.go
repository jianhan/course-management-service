package db

import (
	proto "github.com/jianhan/course-management-service/proto"
	"github.com/jianhan/pkg/mongod"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

type DB interface {
	AddCourses(courses []*proto.Course) error
	UpdateCourses(courses []*proto.Course) error
	DeleteCourses(courses []*proto.Course) error
	GetCourses() error
	getSession() *mgo.Session
}

type Database struct {
	session *mgo.Session
}

func (d *Database) getSession() *mgo.Session {
	s := d.getSession()
	defer s.Close()
	return d.session.Copy()
}

func (d *Database) AddCourses(courses []*proto.Course) error {
	s := d.getSession()
	defer s.Close()
	return nil
}

func (d *Database) UpdateCourses(courses []*proto.Course) error {
	s := d.getSession()
	defer s.Close()
	return nil
}

func (d *Database) DeleteCourses(courses []*proto.Course) error {
	s := d.getSession()
	defer s.Close()
	return nil
}

func (d *Database) GetCourses() error {
	s := d.getSession()
	defer s.Close()
	return nil
}

func NewDatabase() (DB, error) {
	mongodSession, f, err := mongod.NewSession(viper.GetString("mongo.url"))
	defer f()
	if err != nil {
		return nil, err
	}
	return &Database{
		session: mongodSession,
	}, nil
}
