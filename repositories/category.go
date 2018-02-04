package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	pb "github.com/jianhan/course-management-service/proto/course"
)

// CategoryRepository contains collection of methods for repository.
type CategoryRepository interface {
	UpsertCategories(categories []*pb.Category) (result sql.Result, err error)
	// DeleteCategoriesByIDs(ids []string) (uint32, error)
	// GetCategoriesByFilters(filterSet *pb.CategoryFilterSet) ([]*pb.Category, error)
}

// CategoryMysql is a struct which will implement CategoryRepository interface.
type CategoryMysql struct {
	db              *sql.DB
	categoriesTable string
}

// NewCategoryRepository returns a interface of CategoryRepository
func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryMysql{db: db, categoriesTable: "categories"}
}

// UpsertCategories update/insert multiple categories.
func (c *CategoryMysql) UpsertCategories(categories []*pb.Category) (result sql.Result, err error) {
	// TODO: added upsert courses later
	sql := fmt.Sprintf("INSERT INTO %s (id, name, slug, description, display_order, visible, updated_at) VALUES", c.categoriesTable)
	var placeholders []string
	var vals []interface{}
	for _, c := range categories {
		if c.Id == "" {
			c.Id = uuid.New().String()
		}
		placeholders = append(placeholders, "(?,?,?,?,?,?,?,?)")
		updatedAtTime, tErr := ptypes.Timestamp(c.UpdatedAt)
		if tErr != nil {
			return nil, tErr
		}
		vals = append(
			vals,
			c.Id,
			c.Name,
			c.Slug,
			c.Description,
			c.DisplayOrder,
			c.Visible,
			updatedAtTime.Format("2006-01-02 15:04:05"),
		)
	}
	sql += strings.Join(placeholders, ",")
	sql += `ON DUPLICATE KEY UPDATE name=VALUES(name), slug=VALUES(slug), description=VALUES(description), display_order=VALUES(display_order), visible=VALUES(visible), updated_at=VALUES(updated_at)`
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

//
// // DeleteCategoriesByIDs deletes multiply categories by IDs.
// func (c *Category) DeleteCategoriesByIDs(ids []string) (uint32, error) {
// 	changeInfo, err := c.Session.DB(dbName).C(categoriesCollection).RemoveAll(bson.M{"_id": bson.M{"$in": ids}})
// 	if err != nil {
// 		return 0, err
// 	}
// 	// if err := c.Session.DB(dbName).C(coursesCollection).Find(bson.M{"name": bson.M{"$in": ids}}).All(&courses); err != nil {
// 	// 	return nil, err
// 	// }
// 	return uint32(changeInfo.Removed), nil
// }
//
// // GetCategoriesByFilters retrieves categories.
// func (c *Category) GetCategoriesByFilters(filterSet *pb.CategoryFilterSet) ([]*pb.Category, error) {
// 	var conditions []map[string]interface{}
// 	if len(filterSet.Ids) > 0 {
// 		conditions = append(conditions, bson.M{"_id": bson.M{"$in": filterSet.Ids}})
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
// 	var categories []*pb.Category
// 	if err := c.Session.DB(dbName).C(categoriesCollection).Find(query).All(&categories); err != nil {
// 		return nil, err
// 	}
// 	return categories, nil
// }
