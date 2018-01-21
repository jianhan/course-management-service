package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jianhan/course-management-service/handlers"
	proto "github.com/jianhan/course-management-service/proto"
	cfgreader "github.com/jianhan/pkg/configs"
	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

func main() {
	serviceConfigs, err := cfgreader.NewReader(os.Getenv("ENV")).Read()
	if err != nil {
		panic(fmt.Sprintf("error while reading configurations: %s", err.Error()))
	}
	service := micro.NewService(
		micro.Name(serviceConfigs.Name),
		micro.RegisterTTL(time.Duration(serviceConfigs.RegisterTTL)*time.Second),
		micro.RegisterInterval(time.Duration(serviceConfigs.RegisterInterval)*10),
		micro.Version(serviceConfigs.Version),
		micro.Metadata(serviceConfigs.Metadata),
	)
	service.Init()
	proto.RegisterCourseManagerHandler(service.Server(), new(handlers.CourseManager))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
