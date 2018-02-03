package repositories

import (
	"database/sql"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	pb "github.com/jianhan/course-management-service/proto/course"
)

// TODO: move follwing to Course struct.
const (
	dbName            = "course-management-service"
	coursesCollection = "courses"
)

// CourseRepository contains collection of methods for repository.
type CourseRepository interface {
	UpsertCourses(courses []*pb.Course, upsertCategories bool) (result sql.Result, err error)

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
func (c *CourseMysql) UpsertCourses(courses []*pb.Course, upsertCategories bool) (result sql.Result, err error) {
	sqlStr := `INSERT INTO courses (id, name, slug, visible,description, start, end, updated_at) VALUES`
	var placeholders []string
	var vals []interface{}
	for _, c := range courses {
		if c.Id == "" {
			c.Id = uuid.New().String()
		}
		placeholders = append(placeholders, "(?,?,?,?,?,?,?,?)")
		startTime, tErr := ptypes.Timestamp(c.Start)
		if tErr != nil {
			return nil, tErr
		}
		endTime, tErr := ptypes.Timestamp(c.End)
		if tErr != nil {
			return nil, tErr
		}
		updatedAtTime, tErr := ptypes.Timestamp(c.UpdatedAt)
		if tErr != nil {
			return nil, tErr
		}
		vals = append(
			vals,
			c.Id,
			c.Name,
			c.Slug,
			c.Visible,
			c.Description,
			startTime.Format("2006-01-02 15:04:05"),
			endTime.Format("2006-01-02 15:04:05"),
			updatedAtTime.Format("2006-01-02 15:04:05"),
		)
	}
	sqlStr += strings.Join(placeholders, ",")
	sqlStr += `ON DUPLICATE KEY UPDATE name=VALUES(name), slug=VALUES(slug), visible=VALUES(visible), description=VALUES(description), start=VALUES(start), end=VALUES(end), updated_at=VALUES(updated_at)`
	stmt, err := c.db.Prepare(sqlStr)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err = stmt.Exec(vals...)
	if err != nil {
		return
	}
	return
}

// // UpsertCourses update/insert multiple courses.
// func (c *CourseMysql) UpsertCourses(courses []*pb.Course, upsertCategories bool) (result sql.Result, err error) {
// 	sqlStr := `INSERT INTO courses
// 						(id, name, slug, display_order, description, visible, start, end, updated_at)
// 						VALUES`
// 	var placeholders []string
// 	var vals []interface{}
// 	for _, c := range courses {
// 		placeholders = append(placeholders, "(?,?,?,?,?,?,?,?,?)")
// 		vals = append(vals, c.Id, c.Name, c.Slug, c.DisplayOrder, c.Description, c.Visible, "2018-12-12 11:11:11", "2018-12-12 11:11:11", "2018-12-12 11:11:11")
// 	}
// 	sqlStr += strings.Join(placeholders, ",")
// 	sqlStr += `ON DUPLICATE KEY UPDATE
// 														 	name=VALUES(name),
// 															slug=VALUES(slug),
// 															display_order=VALUES(display_order),
// 															description=VALUES(description),
// 															visible=VALUES(visible)
// 															start=VALUES(start)
// 															end=VALUES(end)
// 															updated_at=VALUES(updated_at)`
// 	stmt, err := c.db.Prepare(sqlStr)
// 	spew.Dump(sqlStr, vals)
// 	if err != nil {
// 		return
// 	}
// 	defer stmt.Close()
// 	result, err = stmt.Exec(vals...)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

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
