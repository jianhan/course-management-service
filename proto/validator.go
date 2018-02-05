package go_micro_srv_cm

// import (
// 	"errors"
// 	"fmt"
// 	"time"
//
// 	"github.com/asaskevich/govalidator"
// 	"github.com/gosimple/slug"
// 	"github.com/jianhan/pkg/validation"
// 	"github.com/micro/protobuf/ptypes"
// )
//
//
//
// // Validate checks if any invalid slugs or any invalid UUIDs.
// func (r *UpsertCoursesRequest) Validate() error {
// 	// Struct validation
// 	sv := structvalidator.New()
// 	for k, v := range r.Courses {
// 		if err := sv.Struct(v); err != nil {
// 			return err
// 		}
// 		// if id is not empty, and it is not a valid UUID, throw error.
// 		if v.Id != "" && !govalidator.IsUUID(v.Id) {
// 			return fmt.Errorf("Invalid UUID: %s", v.Id)
// 		}
// 		// if slug is not empty, regardless if it is insert or update, throw error.
// 		if v.Slug != "" && !slug.IsSlug(v.Slug) {
// 			return fmt.Errorf("Invalid slug: %s", v.Slug)
// 		}
// 		// if slug is empty then automatically generate one based on name.
// 		if v.Slug == "" {
// 			r.Courses[k].Slug = slug.Make(v.Name)
// 		}
// 		if v.UpdatedAt == nil {
// 			t, err := ptypes.TimestampProto(time.Now())
// 			if err != nil {
// 				return err
// 			}
// 			r.Courses[k].UpdatedAt = t
// 		}
// 	}
// 	return nil
// }
//
// // Validate performs validation up on DeleteCoursesByFiltersRequest.
// func (r *DeleteCoursesByFiltersRequest) Validate() error {
// 	// when deleting courses never allow empty filter set to wipe out entire table.
// 	if r.FilterSet == nil {
// 		return errors.New("filter set can not be empty")
// 	}
// 	if err := r.FilterSet.Validate(); err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// // Validate performs validation on filterSet.
// func (f *FilterSet) Validate() error {
// 	if len(f.Ids) > 0 {
// 		if err := validation.ValidateSliceUUID(f.Ids); err != nil {
// 			return err
// 		}
// 	}
// 	if f.Slugs != nil && len(f.Slugs) > 0 {
// 		if err := validation.ValidateSliceSlugs(f.Slugs); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
//
// // Validate validates request for upsert.
// func (r *UpsertCategoriesRequest) Validate() error {
// 	sv := structvalidator.New()
// 	for k, v := range r.Categories {
// 		if err := sv.Struct(v); err != nil {
// 			return err
// 		}
// 		// if id is not empty, and it is not a valid UUID, throw error.
// 		if v.Id != "" && !govalidator.IsUUID(v.Id) {
// 			return fmt.Errorf("Invalid UUID: %s", v.Id)
// 		}
// 		// if slug is not empty, regardless if it is insert or update, throw error.
// 		if v.Slug != "" && !slug.IsSlug(v.Slug) {
// 			return fmt.Errorf("Invalid slug: %s", v.Slug)
// 		}
// 		// if slug is empty then automatically generate one based on name.
// 		if v.Slug == "" {
// 			r.Categories[k].Slug = slug.Make(v.Name)
// 		}
// 		// automatically set updated at
// 		if v.UpdatedAt == nil {
// 			t, err := ptypes.TimestampProto(time.Now())
// 			if err != nil {
// 				return err
// 			}
// 			r.Categories[k].UpdatedAt = t
// 		}
// 	}
// 	return nil
// }

// // Validate performs validation on GetCategoriesByFiltersRequest.
// func (r *GetCategoriesByFiltersRequest) Validate() error {
// 	if r.FilterSet == nil {
// 		return errors.New("Filter set is empty while fetching categories")
// 	}
// 	if len(r.FilterSet.Ids) > 0 {
// 		if err := validation.ValidateSliceUUID(r.FilterSet.Ids); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
//
// // Validate performs validation on DeleteCategoriesByIDsRequest.
// func (r *DeleteCategoriesByIDsRequest) Validate() error {
// 	if len(r.Ids) == 0 {
// 		return errors.New("Empty IDs while trying to delete categories")
// 	}
// 	if err := validation.ValidateSliceUUID(r.Ids); err != nil {
// 		return err
// 	}
// 	return nil
// }
