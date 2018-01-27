package go_micro_srv_course

import (
	"fmt"

	"github.com/asaskevich/govalidator"
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
	for _, v := range r.Courses {
		if err := sv.Struct(v); err != nil {
			return err
		}
	}
	return nil
}
