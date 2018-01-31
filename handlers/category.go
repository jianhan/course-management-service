package handlers

import (
	"context"

	pcourse "github.com/jianhan/course-management-service/proto/course"
	rp "github.com/jianhan/course-management-service/repositories"
	merrors "github.com/micro/go-micro/errors"
)

// GetCategoryRepo retrieve an mongo db session instance.
func (c *CourseManagement) GetCategoryRepo() rp.CategoryRepository {
	t := &rp.Category{}
	t.Session = c.Session.Clone()
	return t
}

// UpsertCategories upsert multiply categories.
func (c *CourseManagement) UpsertCategories(ctx context.Context, req *pcourse.UpsertCategoriesRequest, rsp *pcourse.UpsertCategoriesResponse) (err error) {
	if err = req.Validate(); err != nil {
		return merrors.BadRequest(API+".GetCategoriesByFilters", err.Error())
	}
	repo := c.GetCategoryRepo()
	defer repo.Close()
	if rsp.Updated, rsp.Inserted, err = repo.UpsertCategories(req.Categories); err != nil {
		return merrors.InternalServerError(API+".GetCategoriesByFilters", err.Error())
	}
	return
}

// GetCategoriesByFilters retrieves categories by filters.
func (c *CourseManagement) GetCategoriesByFilters(ctx context.Context, req *pcourse.GetCategoriesByFiltersRequest, rsp *pcourse.GetCategoriesByFiltersResponse) (err error) {
	if err = req.Validate(); err != nil {
		return merrors.BadRequest(API+".GetCategoriesByFilters", err.Error())
	}
	if req.FilterSet == nil {
		return merrors.BadRequest(API+".GetCategoriesByFilters", "Empty filter set")
	}
	repo := c.GetCategoryRepo()
	defer repo.Close()
	if rsp.Categories, err = repo.GetCategoriesByFilters(req.FilterSet); err != nil {
		return merrors.InternalServerError(API+".GetCategoriesByFilters", err.Error())
	}
	return
}

// DeleteCategoriesByIDs remove categories by IDs.
func (c *CourseManagement) DeleteCategoriesByIDs(ctx context.Context, req *pcourse.DeleteCategoriesByIDsRequest, rsp *pcourse.DeleteCategoriesByIDsResponse) (err error) {
	if err = req.Validate(); err != nil {
		return merrors.BadRequest(API+".DeleteCategoriesByIDs", err.Error())
	}
	repo := c.GetRepo()
	defer repo.Close()
	if rsp.Removed, err = repo.DeleteCategoriesByIDs(req.Ids); err != nil {
		return merrors.BadRequest(API+".DeleteCategoriesByIDs", err.Error())
	}
	return
}
