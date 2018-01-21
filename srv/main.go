package main

import (
	"fmt"
	"os"
	"time"

	creader "github.com/jianhan/course-management-service/configs_reader"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

func main() {
	env, ok := os.LookupEnv("ENV")
	if !ok {
		panic("missing environment variable 'ENV', failed to start service")
	}
	serviceConfigs, err := creader.NewConfigsReader(env).Read()
	if err != nil {
		panic(fmt.Sprintf("error while reading configurations: %s", err.Error()))
	}
	service := micro.NewService(
		micro.Name(serviceConfigs.Name),
		micro.RegisterTTL(time.Duration(serviceConfigs.RegisterTTL)*time.Second),
		micro.RegisterInterval(time.Duration(serviceConfigs.RegisterInterval)*10),
		micro.Version(serviceConfigs.Version),
		micro.Metadata(serviceConfigs.Metadata),
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
