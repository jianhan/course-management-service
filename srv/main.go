package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/cli"
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
		micro.Flags(cli.StringFlag{
			Name:  "env",
			Usage: "Development environment: development/staging/production",
		}),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			fmt.Printf("The env flag is: %s\n", c.String("env"))
		}),
	)
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
