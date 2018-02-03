package repositories

import (
	"database/sql"

	pb "github.com/jianhan/course-management-service/proto/course"
)

// TODO: move follwing to Course struct.
const (
	dbName            = "course-management-service"
	coursesCollection = "courses"
)

// CourseRepository contains collection of methods for repository.
type CourseRepository interface {
	UpsertCourses(courses []*pb.Course, upsertCategories bool) (rowsAffected uint32, updated uint32, inserted uint32, err error)

	// DeleteCoursesByIDs(ids []string) (uint32, error)
	// GetCoursesByFilters(filterSet *pb.FilterSet) ([]*pb.Course, error)
}

// CourseMysql is a mysql implementation of CourseRepository.
type CourseMysql struct {
	db *sql.DB
}

// NewCourseRepository returns a interface of CourseRepository.
func NewCourseRepository(db *sql.DB) CourseRepository {
	return &CourseMysql{db: db}
}

// UpsertCourses update/insert multiple courses.
func (c *CourseMysql) UpsertCourses(courses []*pb.Course, upsertCategories bool) (rowsAffected uint32, updated uint32, inserted uint32, err error) {
	return
}

// // DeleteCoursesByIDs deletes multiply courses by IDs.
// func (c *Course) DeleteCoursesByIDs(ids []string) (uint32, error) {
// 	changeInfo, err := c.Session.DB(dbName).C(coursesCollection).RemoveAll(bson.M{"_id": bson.M{"$in": ids}})
// 	if err != nil {
// 		return 0, err
// 	}
// 	return uint32(changeInfo.Removed), nil
// }
//
// // GetCoursesByFilters retrieves courses.
// func (c *Course) GetCoursesByFilters(filterSet *pb.FilterSet) ([]*pb.Course, error) {
// 	var conditions []map[string]interface{}
// 	if len(filterSet.Ids) > 0 {
// 		conditions = append(conditions, bson.M{"_id": bson.M{"$in": filterSet.Ids}})
// 	}
// 	if filterSet.Start != nil {
// 		conditions = append(conditions, bson.M{"start": bson.M{"$lte": filterSet.Start}})
// 	}
// 	if filterSet.End != nil {
// 		conditions = append(conditions, bson.M{"end": bson.M{"$gte": filterSet.End}})
// 	}
// 	if len(filterSet.Names) > 0 {
// 		conditions = append(conditions, bson.M{"name": bson.M{"$in": filterSet.Names}})
// 	}
// 	if strings.TrimSpace(filterSet.TextSearch) != "" {
// 		conditions = append(conditions, bson.M{"$text": bson.M{"$search": strings.TrimSpace(filterSet.TextSearch)}})
// 	}
// 	if filterSet.Visible != nil && !filterSet.Visible.Ignore {
// 		conditions = append(conditions, bson.M{"visible": bson.M{"$eq": filterSet.Visible.Value}})
// 	}
// 	var query interface{}
// 	if len(conditions) > 0 {
// 		query = bson.M{"$and": conditions}
// 	}
// 	var courses []*pb.Course
// 	if err := c.Session.DB(dbName).C(coursesCollection).Find(query).All(&courses); err != nil {
// 		return nil, err
// 	}
// 	return courses, nil
// }
