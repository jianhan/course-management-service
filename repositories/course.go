package repositories

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/jianhan/course-management-service/proto/course"
	jmongod "github.com/jianhan/pkg/mongod"
	"github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName            = "course-management-service"
	coursesCollection = "courses"
)

// CourseRepository contains collection of methods for repository.
type CourseRepository interface {
	UpsertCourses(courses []*pb.Course) (uint32, uint32, error)
	DeleteCourses(courses []*pb.Course) error
	GetCoursesByFilters(filterSet *pb.FilterSet) ([]*pb.Course, error)
	Close()
}

// Course is a struct which will implement CourseRepository interface.
type Course struct {
	jmongod.MRepository
}

// UpsertCourses update/insert multiple courses.
func (c *Course) UpsertCourses(courses []*pb.Course) (uint32, uint32, error) {
	var updated, inserted uint32
	for _, v := range courses {
		now, err := ptypes.TimestampProto(time.Now())
		if err != nil {
			return 0, 0, nil
		}
		if v.Id == "" {
			v.Id = uuid.Must(uuid.NewV4(), nil).String()
			v.CreatedAt = now
		}
		if v.CreatedAt == nil {
			v.CreatedAt = now
		}
		v.UpdatedAt = now
		info, err := c.Session.DB(dbName).C(coursesCollection).UpsertId(v.Id, v)
		if err != nil {
			return 0, 0, nil
		}
		if info.Updated > 0 {
			updated++
		} else {
			inserted++
		}
	}
	return updated, inserted, nil
}

// DeleteCourses deletes multiply courses.
func (c *Course) DeleteCourses(courses []*pb.Course) error {
	return nil
}

// GetCoursesByFilters retrieves courses.
func (c *Course) GetCoursesByFilters(filterSet *pb.FilterSet) ([]*pb.Course, error) {
	var queries []map[string]interface{}
	queries = append(queries, bson.M{"visiable": filterSet.Visible})
	if len(filterSet.Ids) > 0 {
		queries = append(queries, bson.M{"_id": bson.M{"$in": filterSet.Ids}})
	}
	if filterSet.Start != nil {
		queries = append(queries, bson.M{"start": bson.M{"$lte": filterSet.Start}})
	}
	if filterSet.End != nil {
		queries = append(queries, bson.M{"start": bson.M{"$gte": filterSet.End}})
	}
	var courses []*pb.Course
	if err := c.Session.DB(dbName).C(coursesCollection).Find(bson.M{"$and": queries}).All(&courses); err != nil {
		return nil, err
	}
	return courses, nil
}

// InitCourse performs initializations for courses collection such as index, etc..
var InitCourse RepoInitFunc = func(session *mgo.Session) {
	s := session.Clone()
	defer s.Close()
	c := s.DB(dbName).C(coursesCollection)
	index := mgo.Index{
		Key:        []string{"slug"},
		Unique:     true,
		DropDups:   false,
		Background: false,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
