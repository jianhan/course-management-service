package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"

	pcourses "github.com/jianhan/course-management-service/proto/courses"
	pmysql "github.com/jianhan/pkg/proto/mysql"
)

// CourseRepository contains collection of methods for repository.
type CourseRepository interface {
	UpsertCourses(courses []*pcourses.CourseWithCategories) (result sql.Result, err error)
	GetCoursesByFilters(filterSet *pcourses.FilterSet, sort *pmysql.Sort, pagination *pmysql.Pagination) ([]*pcourses.Course, error)
	DeleteCoursesByFilters(filterSet *pcourses.FilterSet) (deleted int64, err error)
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
func (c *CourseMysql) UpsertCourses(courses []*pcourses.CourseWithCategories) (result sql.Result, err error) {
	// TODO: added upsert categories later
	sql := fmt.Sprintf("INSERT INTO %s (id, name, slug, visible,description, start, end, updated_at) VALUES", c.coursesTable)
	var placeholders []string
	var vals []interface{}
	for _, c := range courses {
		placeholders = append(placeholders, "(?,?,?,?,?,?,?,?)")
		startTime, tErr := ptypes.Timestamp(c.Course.Start)
		if tErr != nil {
			return nil, tErr
		}
		endTime, tErr := ptypes.Timestamp(c.Course.End)
		if tErr != nil {
			return nil, tErr
		}
		updatedAtTime, tErr := ptypes.Timestamp(c.Course.UpdatedAt)
		if tErr != nil {
			return nil, tErr
		}
		vals = append(
			vals,
			c.Course.Id,
			c.Course.Name,
			c.Course.Slug,
			c.Course.Visible,
			c.Course.Description,
			startTime.Format("2006-01-02 15:04:05"),
			endTime.Format("2006-01-02 15:04:05"),
			updatedAtTime.Format("2006-01-02 15:04:05"),
		)
	}
	sql += strings.Join(placeholders, ",")
	sql += `ON DUPLICATE KEY UPDATE name=VALUES(name), slug=VALUES(slug), visible=VALUES(visible), description=VALUES(description), start=VALUES(start), end=VALUES(end), updated_at=VALUES(updated_at)`
	stmt, err := c.db.Prepare(sql)
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

func (c *CourseMysql) rowToCourse(f func(dest ...interface{}) error) (course *pcourses.Course, err error) {
	course = &pcourses.Course{}
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
func (c *CourseMysql) GetCoursesByFilters(filterSet *pcourses.FilterSet, sort *pmysql.Sort, pagination *pmysql.Pagination) (courses []*pcourses.Course, err error) {
	var (
		conditionSQLStr string
		args            []interface{}
	)
	if filterSet != nil {
		// get all filters conditions as SQL string.
		conditionSQLStr, args, err = filterSet.GenerateConditions()
		if err != nil {
			return
		}
	}
	// get pagination part of the SQL.
	var paginationStr = "LIMIT 20 OFFSET 0"
	if pagination != nil {
		paginationStr = pagination.GenerateSQLStr()
	}
	// get sorting part of SQL.
	var sortStr = ""
	if sort != nil {
		sortStr = sort.GenerateSQLStr()
	}
	// start query
	sql := fmt.Sprintf("SELECT * FROM %s %s %s %s", c.coursesTable, conditionSQLStr, sortStr, paginationStr)
	fmt.Println(sql, args)
	stmt, err := c.db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
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

// DeleteCoursesByFilters removes courses according to filters.
func (c *CourseMysql) DeleteCoursesByFilters(filterSet *pcourses.FilterSet) (deleted int64, err error) {
	if filterSet == nil {
		return 0, errors.New("filter set can not be empty while deleting courses")
	}
	var (
		conditionSQLStr string
		args            []interface{}
	)
	conditionSQLStr, args, err = filterSet.GenerateConditions()
	if err != nil {
		return
	}
	sql := fmt.Sprintf("DELETE FROM %s %s", c.coursesTable, conditionSQLStr)
	stmt, err := c.db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	r, err := stmt.Exec(args...)
	if err != nil {
		return
	}
	if deleted, err = r.RowsAffected(); err != nil {
		return
	}
	return
}
