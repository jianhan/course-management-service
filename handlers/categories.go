package handlers

// // UpsertCategories upsert multiply categories.
// func (c *CourseManagement) UpsertCategories(ctx context.Context, req *pcourse.UpsertCategoriesRequest, rsp *pcourse.UpsertResult) (err error) {
// 	if err = req.Validate(); err != nil {
// 		return merrors.BadRequest(API+".UpsertCategories", err.Error())
// 	}
// 	result, err := c.CategoryRepository.UpsertCategories(req.Categories)
// 	if err != nil {
// 		return merrors.InternalServerError(API+".GetCategoriesByFilters", err.Error())
// 	}
// 	rsp.RowsAffected, err = result.RowsAffected()
// 	return
// }

//
// // GetCategoriesByFilters retrieves categories by filters.
// func (c *CourseManagement) GetCategoriesByFilters(ctx context.Context, req *pcourse.GetCategoriesByFiltersRequest, rsp *pcourse.GetCategoriesByFiltersResponse) (err error) {
// 	// if err = req.Validate(); err != nil {
// 	// 	return merrors.BadRequest(API+".GetCategoriesByFilters", err.Error())
// 	// }
// 	// if req.FilterSet == nil {
// 	// 	return merrors.BadRequest(API+".GetCategoriesByFilters", "Empty filter set")
// 	// }
// 	// repo := c.GetCategoryRepo()
// 	// defer repo.Close()
// 	// if rsp.Categories, err = repo.GetCategoriesByFilters(req.FilterSet); err != nil {
// 	// 	return merrors.InternalServerError(API+".GetCategoriesByFilters", err.Error())
// 	// }
// 	return
// }
//
// // DeleteCategoriesByIDs remove categories by IDs.
// func (c *CourseManagement) DeleteCategoriesByIDs(ctx context.Context, req *pcourse.DeleteCategoriesByIDsRequest, rsp *pcourse.DeleteCategoriesByIDsResponse) (err error) {
// 	// if err = req.Validate(); err != nil {
// 	// 	return merrors.BadRequest(API+".DeleteCategoriesByIDs", err.Error())
// 	// }
// 	// repo := c.GetCategoryRepo()
// 	// defer repo.Close()
// 	// if rsp.Removed, err = repo.DeleteCategoriesByIDs(req.Ids); err != nil {
// 	// 	return merrors.BadRequest(API+".DeleteCategoriesByIDs", err.Error())
// 	// }
// 	return
// }
