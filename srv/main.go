package main

import (
	"fmt"
	"os"
	"time"

	creader "github.com/jianhan/pkg/configs"
	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

func main() {
	serviceConfigs, err := creader.NewReader(os.Getenv("ENV")).Read()
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
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
