package go_micro_srv_course

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

func (c *Courses) Validate() error {
	for _, v := range c.Courses {
		if v.Id != "" && !govalidator.IsUUID(v.Id) {
			return fmt.Errorf("Invalid UUID: %s", v.Id)
		}
	}
	return nil
}
