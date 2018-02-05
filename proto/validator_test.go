package go_micro_srv_cm_test

// func TestCourses_Validate(t *testing.T) {
// 	type fields struct {
// 		Courses []*pb.Course
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		{
// 			name: "valid uuids success",
// 			fields: fields{
// 				Courses: []*pb.Course{
// 					&pb.Course{
// 						Id: "cf5f6f48-0359-11e8-ba89-0ed5f89f718b",
// 					},
// 					&pb.Course{
// 						Id: "cf5f743e-0359-11e8-ba89-0ed5f89f718b",
// 					},
// 					&pb.Course{
// 						Id: "b4a2345a-39f2-4d3a-9139-78c5864d6a51",
// 					},
// 					&pb.Course{
// 						Id: "84f28c80-7cea-462d-a912-0a6de4a89769",
// 					},
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "invalid uuids fail",
// 			fields: fields{
// 				Courses: []*pb.Course{
// 					&pb.Course{
// 						Id: "invalid UUID",
// 					},
// 				},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "allow empty uuid success",
// 			fields: fields{
// 				Courses: []*pb.Course{
// 					&pb.Course{
// 						Id: "",
// 					},
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := &pb.Courses{
// 				Courses: tt.fields.Courses,
// 			}
// 			if err := c.Validate(); (err != nil) != tt.wantErr {
// 				t.Errorf("Courses.Validate() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
//
// func TestUpsertCoursesRequest_Validate(t *testing.T) {
// 	validCourses := generateValidCourses(1, false, true)
// 	type fields struct {
// 		Courses []*pb.Course
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		{
// 			name: "all fields exists success",
// 			fields: fields{
// 				Courses: validCourses,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "missing name will fail",
// 			fields: fields{
// 				Courses: []*pb.Course{
// 					&pb.Course{
// 						Id:          uuid.Must(uuid.NewV4(), nil).String(),
// 						Slug:        "test-slug",
// 						Description: "test description",
// 					},
// 				},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "missing description will fail",
// 			fields: fields{
// 				Courses: []*pb.Course{
// 					&pb.Course{
// 						Id:   uuid.Must(uuid.NewV4(), nil).String(),
// 						Name: "test name",
// 						Slug: "test-slug",
// 					},
// 				},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "invalid uuid will fail",
// 			fields: fields{
// 				Courses: []*pb.Course{
// 					&pb.Course{
// 						Id:          "invalid uuid",
// 						Name:        "test name",
// 						Slug:        "test-slug",
// 						Description: "test description",
// 					},
// 				},
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := &pb.UpsertCoursesRequest{
// 				Courses: tt.fields.Courses,
// 			}
// 			if err := r.Validate(); (err != nil) != tt.wantErr {
// 				t.Errorf("UpsertCoursesRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
//
// func TestUpsertCoursesRequest_ValidateAutoFillSlug(t *testing.T) {
// 	upsertCoursesRequest := &pb.UpsertCoursesRequest{
// 		Courses: generateValidCourses(1, true, true),
// 	}
// 	upsertCoursesRequest.Validate()
// 	for _, v := range upsertCoursesRequest.Courses {
// 		assert.Equal(t, slug.Make(v.Name), v.Slug)
// 	}
// }
//
// func generateValidCourses(num int, withEmptySlug bool, withEmptyID bool) []*pb.Course {
// 	validCourses := []*pb.Course{}
// 	for i := 0; i < num; i++ {
// 		c := pb.Course{}
// 		faker.FakeData(&c)
// 		c.Id = uuid.Must(uuid.NewV4(), nil).String()
// 		if withEmptySlug {
// 			c.Slug = ""
// 		} else {
// 			c.Slug = slug.Make(c.Name)
// 		}
// 		if withEmptyID {
// 			c.Id = ""
// 		}
// 		validCourses = append(validCourses, &c)
// 	}
// 	return validCourses
// }
