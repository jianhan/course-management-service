package courses

import (
	"errors"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gosimple/slug"
	"github.com/jianhan/pkg/validation"
	"github.com/micro/protobuf/ptypes"
)

// Validate checks if any invalid slugs or any invalid UUIDs.
func (r *UpsertCoursesRequest) Validate() error {
	// Struct validation
	for k, v := range r.Courses {
		if _, err := govalidator.ValidateStruct(v); err != nil {
			return err
		}
		// if slug is not empty, regardless if it is insert or update, throw error.
		if v.Slug != "" && !slug.IsSlug(v.Slug) {
			return fmt.Errorf("Invalid slug: %s", v.Slug)
		}
		// if slug is empty then automatically generate one based on name.
		if v.Slug == "" {
			r.Courses[k].Slug = slug.Make(v.Name)
		}
		if v.UpdatedAt == nil {
			t, err := ptypes.TimestampProto(time.Now())
			if err != nil {
				return err
			}
			r.Courses[k].UpdatedAt = t
		}
	}
	return nil
}

// Validate performs validation up on DeleteCoursesByFiltersRequest.
func (r *DeleteCoursesByFiltersRequest) Validate() error {
	// when deleting courses never allow empty filter set to wipe out entire table.
	if r.FilterSet == nil {
		return errors.New("filter set can not be empty")
	}
	if _, err := govalidator.ValidateStruct(r.FilterSet); err != nil {
		return err
	}
	return nil
}

// Validate performs validation for request.
func (r *GetCoursesByFiltersRequest) Validate() error {
	if r.FilterSet != nil {
		if err := r.FilterSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate performs validation on filter set.
func (f *FilterSet) Validate() error {
	if len(f.Ids) > 0 {
		if err := validation.ValidateSliceUUID(f.Ids); err != nil {
			return err
		}
	}
	if f.Slugs != nil && len(f.Slugs) > 0 {
		if err := validation.ValidateSliceSlugs(f.Slugs); err != nil {
			return err
		}
	}
	return nil
}
