package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jianhan/course-management-service/handlers"
	pb "github.com/jianhan/course-management-service/proto/course"
	"github.com/jianhan/course-management-service/repositories"
	cfgreader "github.com/jianhan/pkg/configs"
	jmongod "github.com/jianhan/pkg/mongod"
	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

func main() {
	session, err := jmongod.CreateSession("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	repositories.Initialize(session, repositories.InitCourse)
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
	pb.RegisterCourseManagerHandler(srv.Server(), &handlers.CourseManagement{Session: session})
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
