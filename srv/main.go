package main

import (
	"log"
	"time"

	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.course-management"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version("1.0.0"),
		micro.Metadata(map[string]string{
			"course": "sample course management system",
		}),
	)

	// optionally setup command line usage
	service.Init()

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
