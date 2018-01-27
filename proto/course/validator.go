package go_micro_srv_course

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/gosimple/slug"
	structvalidator "gopkg.in/go-playground/validator.v9"
)

func (c *Courses) Validate() error {
	for _, v := range c.Courses {
		if v.Id != "" && !govalidator.IsUUID(v.Id) {
			return fmt.Errorf("Invalid UUID: %s", v.Id)
		}
	}
	return nil
}

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
		// if id is empty, means it is an insert, and slug is empty too then
		// automatically generate one based on name.
		if v.Id == "" && v.Slug == "" {
			r.Courses[k].Slug = slug.Make(v.Name)
		}
	}
	return nil
}
