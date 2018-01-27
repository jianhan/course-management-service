package repositories

import (
	"github.com/google/uuid"
	pb "github.com/jianhan/course-management-service/proto/course"
	jmongod "github.com/jianhan/pkg/mongod"
)

const (
	dbName            = "course-management-service"
	coursesCollection = "courses"
)

type CourseRepository interface {
	UpsertCourses(courses []*pb.Course) (uint32, uint32, error)
	DeleteCourses(courses []*pb.Course) error
	GetCourses() ([]*pb.Course, error)
	Close()
}

type Course struct {
	jmongod.MRepository
}

func (c *Course) UpsertCourses(courses []*pb.Course) (uint32, uint32, error) {
	var updated, inserted uint32
	for _, v := range courses {
		if v.Id == "" {
			tmpId, err := uuid.NewUUID()
			if err != nil {
				return 0, 0, nil
			}
			v.Id = tmpId.String()
		}
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
