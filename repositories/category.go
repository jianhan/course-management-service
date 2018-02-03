package repositories

const categoriesCollection = "categories"

// // CategoryRepository contains collection of methods for repository.
// type CategoryRepository interface {
// 	UpsertCategories(categories []*pb.Category) (uint32, uint32, error)
// 	DeleteCategoriesByIDs(ids []string) (uint32, error)
// 	GetCategoriesByFilters(filterSet *pb.CategoryFilterSet) ([]*pb.Category, error)
// }
//
// // Category is a struct which will implement CategoryRepository interface.
// type Category struct {
// }
//
// // UpsertCategories update/insert multiple categories.
// func (c *Category) UpsertCategories(categories []*pb.Category) (uint32, uint32, error) {
//
// }
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
