package courses

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/ptypes"
)

// GenerateConditions generates mysql conditions as string.
func (f *FilterSet) GenerateConditions() (sql string, args []interface{}, err error) {
	var conditions []string
	if len(f.Ids) > 0 {
		var ids []string
		for _, v := range f.Ids {
			ids = append(ids, "?")
			args = append(args, v)
		}
		conditions = append(conditions, fmt.Sprintf("id IN(%s)", strings.Join(ids, ",")))
	}
	if len(f.Names) > 0 {
		var names []string
		for _, v := range f.Names {
			names = append(names, "?")
			args = append(args, v)
		}
		conditions = append(conditions, fmt.Sprintf("name IN(%s)", strings.Join(names, ",")))
	}
	if len(f.Slugs) > 0 {
		var slugs []string
		for _, v := range f.Slugs {
			slugs = append(slugs, "?")
			args = append(args, v)
		}
		conditions = append(conditions, fmt.Sprintf("slug IN(%s)", strings.Join(slugs, ",")))
	}
	if f.Visible != nil && !f.Visible.Ignore {
		conditions = append(conditions, "visible = ?")
		args = append(args, f.Visible.Value)
	}
	if f.Start != nil {
		startDate, tErr := ptypes.Timestamp(f.Start)
		if tErr != nil {
			err = tErr
			return
		}
		conditions = append(conditions, "start <= ?")
		args = append(args, startDate.Format("2006-01-02 15:04:05"))
	}
	if f.End != nil {
		endDate, tErr := ptypes.Timestamp(f.End)
		if tErr != nil {
			err = tErr
			return
		}
		conditions = append(conditions, "end <= ?")
		args = append(args, endDate.Format("2006-01-02 15:04:05"))
	}
	if f.TextSearch != "" {
		conditions = append(conditions, "MATCH(name, description) AGAINST(? IN NATURAL LANGUAGE MODE)")
		args = append(args, f.TextSearch)
	}
	if len(conditions) > 0 {
		sql = fmt.Sprintf(" WHERE %s", strings.Join(conditions, " AND "))
	}
	return
}
