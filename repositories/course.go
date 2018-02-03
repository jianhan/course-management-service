package repositories

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	pb "github.com/jianhan/course-management-service/proto/course"
)

// CourseRepository contains collection of methods for repository.
type CourseRepository interface {
	UpsertCourses(courses []*pb.Course, upsertCategories bool) (result sql.Result, err error)
	GetCoursesByFilters(filterSet *pb.FilterSet) ([]*pb.Course, error)
	// DeleteCoursesByIDs(ids []string) (uint32, error)
	// GetCoursesByFilters(filterSet *pb.FilterSet) ([]*pb.Course, error)
}

// CourseMysql is a mysql implementation of CourseRepository.
type CourseMysql struct {
	db           *sql.DB
	coursesTable string
}

// NewCourseRepository returns a interface of CourseRepository.
func NewCourseRepository(db *sql.DB) CourseRepository {
	return &CourseMysql{db: db, coursesTable: "courses"}
}

// UpsertCourses update/insert multiple courses.
func (c *CourseMysql) UpsertCourses(courses []*pb.Course, upsertCategories bool) (result sql.Result, err error) {
	// TODO: added upsert categories later
	sqlStr := fmt.Sprintf("INSERT INTO %s (id, name, slug, visible,description, start, end, updated_at) VALUES", c.coursesTable)
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

// // DeleteCoursesByIDs deletes multiply courses by IDs.
// func (c *Course) DeleteCoursesByIDs(ids []string) (uint32, error) {
// 	changeInfo, err := c.Session.DB(dbName).C(coursesCollection).RemoveAll(bson.M{"_id": bson.M{"$in": ids}})
// 	if err != nil {
// 		return 0, err
// 	}
// 	return uint32(changeInfo.Removed), nil
// }
//

func (c *CourseMysql) rowToCourse(f func(dest ...interface{}) error) (course *pb.Course, err error) {
	course = &pb.Course{}
	var start, end, createdAt, updatedAt time.Time
	err = f(
		&course.Id,
		&course.Name,
		&course.Slug,
		&course.DisplayOrder,
		&course.Description,
		&course.Visible,
		&start,
		&end,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return
	}
	if course.Start, err = ptypes.TimestampProto(start); err != nil {
		return
	}
	if course.End, err = ptypes.TimestampProto(end); err != nil {
		return
	}
	if course.CreatedAt, err = ptypes.TimestampProto(createdAt); err != nil {
		return
	}
	if course.UpdatedAt, err = ptypes.TimestampProto(updatedAt); err != nil {
		return
	}
	return
}

// GetCoursesByFilters retrieves courses.
func (c *CourseMysql) GetCoursesByFilters(filterSet *pb.FilterSet) (courses []*pb.Course, err error) {
	rows, err := c.db.Query(fmt.Sprintf("SELECT * FROM %s", c.coursesTable))
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		course, rErr := c.rowToCourse(rows.Scan)
		if rErr != nil {
			return nil, rErr
		}
		courses = append(courses, course)
	}
	return
}
