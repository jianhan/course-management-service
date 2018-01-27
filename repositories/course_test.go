package repositories

import (
	"reflect"
	"testing"

	pb "github.com/jianhan/course-management-service/proto/course"
)

func TestCourse_UpsertCourses(t *testing.T) {
	type args struct {
		courses []*pb.Course
	}
	tests := []struct {
		name    string
		c       *Course
		args    args
		want    uint32
		want1   uint32
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.c.UpsertCourses(tt.args.courses)
			if (err != nil) != tt.wantErr {
				t.Errorf("Course.UpsertCourses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Course.UpsertCourses() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Course.UpsertCourses() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCourse_DeleteCourses(t *testing.T) {
	type args struct {
		courses []*pb.Course
	}
	tests := []struct {
		name    string
		c       *Course
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.DeleteCourses(tt.args.courses); (err != nil) != tt.wantErr {
				t.Errorf("Course.DeleteCourses() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCourse_GetCourses(t *testing.T) {
	tests := []struct {
		name    string
		c       *Course
		want    []*pb.Course
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetCourses()
			if (err != nil) != tt.wantErr {
				t.Errorf("Course.GetCourses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Course.GetCourses() = %v, want %v", got, tt.want)
			}
		})
	}
}
