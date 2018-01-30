package course

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/gosimple/slug"
	"github.com/jianhan/pkg/validation"
	structvalidator "gopkg.in/go-playground/validator.v9"
)

// Validate checks if any id presented within courses are all valid.
func (c *Courses) Validate() error {
	for _, v := range c.Courses {
		if v.Id != "" && !govalidator.IsUUID(v.Id) {
			return fmt.Errorf("Invalid UUID: %s", v.Id)
		}
	}
	return nil
}

// Validate checks if any invalid slugs or any invalid UUIDs.
func (r *UpsertCoursesRequest) Validate() error {
	sv := structvalidator.New()
	for k, v := range r.Courses {
		if err := sv.Struct(v); err != nil {
			return err
		}
		// if id is not empty, and it is not a valid UUID, throw error.
		if v.Id != "" && !govalidator.IsUUID(v.Id) {
			return fmt.Errorf("Invalid UUID: %s", v.Id)
		}
		// if slug is not empty, regardless if it is insert or update, throw error.
		if v.Slug != "" && !slug.IsSlug(v.Slug) {
			return fmt.Errorf("Invalid slug: %s", v.Slug)
		}
		// if slug is empty then automatically generate one based on name.
		if v.Slug == "" {
			r.Courses[k].Slug = slug.Make(v.Name)
		}
	}
	return nil
}

// Validate performs validation for GetCoursesByFiltersRequest.
func (r *GetCoursesByFiltersRequest) Validate() error {
	if r.FilterSet == nil {
		return errors.New("Filter set is empty")
	}
	if r.FilterSet.Ids != nil && len(r.FilterSet.Ids) > 0 {
		if err := validation.ValidateSliceUUID(r.FilterSet.Ids); err != nil {
			return err
		}
	}
	if r.FilterSet.Slugs != nil && len(r.FilterSet.Slugs) > 0 {
		if err := validation.ValidateSliceSlugs(r.FilterSet.Slugs); err != nil {
			return err
		}
	}
	return nil
}
