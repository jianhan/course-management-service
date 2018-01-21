package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ServiceConfigs interface {
	GetServiceName() string
	GetServiceRegisterTTL() int
	GetServiceRegisterInterval() int
	GetServiceVersion() string
	GetMetaData() map[string]string
}

type scfgs struct {
	serviceName             string
	serviceRegisterTTL      int
	serviceRegisterInterval int
	serviceVersion          string
	metaData                map[string]string
}

func NewServiceConfigs(r io.Reader) ServiceConfigs {

	r.Read([]byte("../configs/base.yml"))
	return nil
}

type ServiceConfigsReader struct {
}

func (s *ServiceConfigsReader) Read(p []byte) (n int, err error) {
	filePath := string(p)
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.WithFields(log.Fields{
			"config_file_path": filePath,
		}).Warn("file path not exists")
	}
	fileName, fileExt := fileInfo.Name(), filepath.Ext(fileInfo.Name())
	fileBaseName := fileName[0 : len(fileName)-len(fileExt)]
	fileDir := filepath.Dir(filePath)
	viper.SetConfigName(fileBaseName)
	viper.AddConfigPath(fileDir) // path to look for the config file in
	err = viper.MergeInConfig()  // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	spew.Dump(viper.GetString("service.name"))
	// spew.Dump(fileBaseName)
	return 0, nil
}

func NewServiceConfigsReader() io.Reader {
	return &ServiceConfigsReader{}
}

func main() {
	spew.Dump(NewServiceConfigs(NewServiceConfigsReader()))
	// rttl, err := strconv.Atoi(os.Getenv("SERVICE_REGISTERTTL"))
	// if err != nil {
	// 	panic(fmt.Sprintf("Error while booting service: %s", err.Error()))
	// }
	// service := micro.NewService(
	// 	micro.Name(os.Getenv("SERVICE_NAME")),
	// 	micro.RegisterTTL(time.Duration(rttl)*time.Second),
	// 	micro.RegisterInterval(time.Second*10),
	// 	micro.Version("1.0.0"),
	// 	micro.Metadata(map[string]string{
	// 		"course": "sample course management system",
	// 	}),
	// 	micro.Flags(cli.StringFlag{
	// 		Name:  "env",
	// 		Usage: "Development environment: development/staging/production",
	// 	}),
	// )
	// service.Init(
	// 	micro.Action(func(c *cli.Context) {
	// 		fmt.Printf("The env flag is: %s\n", c.String("env"))
	// 	}),
	// )
	// // Run server
	// if err := service.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
