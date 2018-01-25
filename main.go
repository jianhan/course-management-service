package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jianhan/course-management-service/handlers"
	pb "github.com/jianhan/course-management-service/proto"
	cfgreader "github.com/jianhan/pkg/configs"
	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

func main() {
	serviceConfigs, err := cfgreader.NewReader(os.Getenv("ENV")).Read()
	if err != nil {
		panic(fmt.Sprintf("error while reading configurations: %s", err.Error()))
	}
	srv := micro.NewService(
		micro.Name(serviceConfigs.Name),
		micro.RegisterTTL(time.Duration(serviceConfigs.RegisterTTL)*time.Second),
		micro.RegisterInterval(time.Duration(serviceConfigs.RegisterInterval)*10),
		micro.Version(serviceConfigs.Version),
		micro.Metadata(serviceConfigs.Metadata),
	)
	srv.Init()
	pb.RegisterCourseManagerHandler(srv.Server(), new(handlers.CourseManager))
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
