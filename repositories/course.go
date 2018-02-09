package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"

	pcourses "github.com/jianhan/course-management-service/proto/courses"
	pmysql "github.com/jianhan/pkg/proto/mysql"
)

// CourseRepository contains collection of methods for repository.
type CourseRepository interface {
	// course
	UpsertCourses(ctx context.Context, records []*pcourses.UpsertCourseRecord) (insertedCount, updatedCount int64, err error)
	InsertCourses(ctx context.Context, records []*pcourses.UpsertCourseRecord) (result sql.Result, err error)
	UpdateCourses(ctx context.Context, records []*pcourses.UpsertCourseRecord) (updated int64, err error)
	GetCoursesByFilters(filterSet *pcourses.FilterSet, sort *pmysql.Sort, pagination *pmysql.Pagination) ([]*pcourses.Course, error)
	DeleteCoursesByFilters(filterSet *pcourses.FilterSet) (deleted int64, err error)
	// course & category relationship
	AddCategories(ctx context.Context, courseID string, categoryIDs []string) (result sql.Result, err error)
	RemoveCategories(ctx context.Context, courseID string, categoryIDs []string) (result sql.Result, err error)
	SyncCategories(ctx context.Context, courseID string, categoryIDs []string) (err error)
}

// CourseMysql is a mysql implementation of CourseRepository.
type CourseMysql struct {
	db                  *sql.DB
	coursesTable        string
	courseCategoryTable string
}

// NewCourseRepository returns a interface of CourseRepository.
func NewCourseRepository(db *sql.DB) CourseRepository {
	return &CourseMysql{db: db, coursesTable: "courses", courseCategoryTable: "course_category"}
}

// UpsertCourses update/insert multiple courses.
func (c *CourseMysql) UpsertCourses(ctx context.Context, records []*pcourses.UpsertCourseRecord) (insertedCount, updatedCount int64, err error) {
	inserts, updates := []*pcourses.UpsertCourseRecord{}, []*pcourses.UpsertCourseRecord{}
	for _, v := range records {
		if v.Course.Id == "" {
			inserts = append(inserts, v)
		} else {
			updates = append(updates, v)
		}
	}
	insertedCountChan, updatedCountChan, errChan := make(chan int64), make(chan int64), make(chan error)
	go func(insertedCountChan chan int64, errChan chan error) {
		ir, iErr := c.InsertCourses(ctx, inserts)
		if iErr != nil {
			errChan <- iErr
			return
		}
		count, raErr := ir.RowsAffected()
		if raErr != nil {
			errChan <- raErr
			return
		}
		insertedCountChan <- count
	}(insertedCountChan, errChan)
	go func(updatedCountChan chan int64, errChan chan error) {
		c, uErr := c.UpdateCourses(ctx, updates)
		if uErr != nil {
			errChan <- uErr
			return
		}
		updatedCountChan <- c
	}(updatedCountChan, errChan)
	for i := 0; i < 2; i++ {
		select {
		case err = <-errChan:
		case insertedCount = <-insertedCountChan:
		case updatedCount = <-updatedCountChan:
		}
	}
	return
}

func (c *CourseMysql) AddCategories(ctx context.Context, courseID string, categoryIDs []string) (result sql.Result, err error) {
	return
}

func (c *CourseMysql) RemoveCategories(ctx context.Context, courseID string, categoryIDs []string) (result sql.Result, err error) {
	return
}

func (c *CourseMysql) InsertCourses(ctx context.Context, records []*pcourses.UpsertCourseRecord) (result sql.Result, err error) {
	if len(records) == 0 {
		return
	}
	sql := fmt.Sprintf("INSERT INTO %s (id, name, slug, visible,description, start, end, updated_at) VALUES", c.coursesTable)
	var placeholders []string
	var vals []interface{}
	for _, v := range records {
		newID, _ := uuid.NewUUID()
		v.Course.Id = newID.String()
		placeholders = append(placeholders, "(?,?,?,?,?,?,?,?)")
		startTime, tErr := ptypes.Timestamp(v.Course.Start)
		if tErr != nil {
			return nil, tErr
		}
		endTime, tErr := ptypes.Timestamp(v.Course.End)
		if tErr != nil {
			return nil, tErr
		}
		updatedAtTime, tErr := ptypes.Timestamp(v.Course.UpdatedAt)
		if tErr != nil {
			return nil, tErr
		}
		vals = append(
			vals,
			v.Course.Id,
			v.Course.Name,
			v.Course.Slug,
			v.Course.Visible,
			v.Course.Description,
			startTime.Format("2006-01-02 15:04:05"),
			endTime.Format("2006-01-02 15:04:05"),
			updatedAtTime.Format("2006-01-02 15:04:05"),
		)
	}
	sql += strings.Join(placeholders, ",")
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

func (c *CourseMysql) UpdateCourses(ctx context.Context, records []*pcourses.UpsertCourseRecord) (updated int64, err error) {
	if len(records) == 0 {
		return
	}
	sql := fmt.Sprintf(`UPDATE %s
											SET (name = ?, slug = ?, visible = ?, description = ?, start = ?, end = ?, updated_at = ?)
											WHERE id = ?`, c.coursesTable)
	stmt, err := c.db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	for _, v := range records {
		startTime, tErr := ptypes.Timestamp(v.Course.Start)
		if tErr != nil {
			return 0, tErr
		}
		endTime, tErr := ptypes.Timestamp(v.Course.End)
		if tErr != nil {
			return 0, tErr
		}
		updatedAtTime := time.Now()
		_, err = stmt.Exec(
			v.Course.Name,
			v.Course.Slug,
			v.Course.Visible,
			v.Course.Description,
			startTime.Format("2006-01-02 15:04:05"),
			endTime.Format("2006-01-02 15:04:05"),
			updatedAtTime.Format("2006-01-02 15:04:05"),
			v.Course.Id,
		)
		if err != nil {
			return
		}
		updated++
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

// SyncCategories syncs categories for an course by course ID.
func (c *CourseMysql) SyncCategories(ctx context.Context, courseID string, categoryIDs []string) (err error) {
	var (
		placeholders []string
		vals         []interface{}
	)
	for _, cID := range categoryIDs {
		placeholders = append(placeholders, "?")
		vals = append(vals, cID)
	}
	// start to update categories related to course
	stmt, err := c.db.Prepare(fmt.Sprintf("DELETE FROM %s WHERE course_id IN (%s)", c.courseCategoryTable, strings.Join(placeholders, ",")))
	spew.Dump(fmt.Sprintf("DELETE FROM %s WHERE course_id IN (%s)", c.courseCategoryTable, strings.Join(placeholders, ",")), vals)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(vals...)
	if err != nil {
		return
	}
	return
}
