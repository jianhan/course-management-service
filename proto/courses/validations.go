package courses

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gosimple/slug"
	pkgvalidation "github.com/jianhan/pkg/validation"
	"github.com/micro/protobuf/ptypes"
)

// Validate performs validation for request.
func (r *GetCoursesByFiltersRequest) Validate() error {
	if r.FilterSet != nil {
		if err := r.FilterSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate checks if any invalid slugs or any invalid UUIDs.
func (r *UpsertCoursesRequest) Validate() error {
	if r.Records == nil {
		return errors.New("Courses is empty")
	}
	if len(r.Records) == 0 {
		return errors.New("Empty courses supplied")
	}
	// validate course with categories if wish to update categories at the same time
	for k, v := range r.Records {
		if err := v.Course.Validate(); err != nil {
			return err
		}
		if err := v.Course.Validate(); err != nil {
			return err
		}
		// if slug is empty then automatically generate one based on name.
		if v.Course.Slug == "" {
			r.Records[k].Course.Slug = slug.Make(v.Course.Name)
		} else if !slug.IsSlug(v.Course.Slug) {
			return fmt.Errorf("course ID: %s is not a valid slug", v.Course.Slug)
		}
		if v.Course.UpdatedAt == nil {
			t, err := ptypes.TimestampProto(time.Now())
			if err != nil {
				return err
			}
			r.Records[k].Course.UpdatedAt = t
		}
	}
	return nil
}

// Validate performs validation up on DeleteCoursesByFiltersRequest.
func (r *DeleteCoursesByFiltersRequest) Validate() error {
	// when deleting courses never allow empty filter set to wipe out entire table.
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
		if err := pkgvalidation.ValidateSliceUUID(f.Ids); err != nil {
			return err
		}
	}
	if f.Slugs != nil && len(f.Slugs) > 0 {
		if err := pkgvalidation.ValidateSliceSlugs(f.Slugs); err != nil {
			return err
		}
	}
	return nil
}

// Validate performs validation on course struct.
func (c Course) Validate() error {
	return validation.ValidateStruct(&c,
		// ID can be empty but it must be UUID if it is not empty.
		validation.Field(&c.Id, is.UUID.Error("invalid UUID")),
		// Name is required.
		validation.Field(&c.Name, validation.Required.Error("course name is required")),
		// Slug can be empty but must be a valid slug if it is not.
		validation.Field(&c.Slug, validation.Match(regexp.MustCompile("^[a-z0-9-]+$")).Error("slug must be in a valid format")),
		// Description is required.
		validation.Field(&c.Description, validation.Required.Error("course description is required")),
		// Start is required.
		validation.Field(&c.Start, validation.Required.Error("course start is required")),
		// End is required.
		validation.Field(&c.End, validation.Required.Error("course end is required")),
	)
}
