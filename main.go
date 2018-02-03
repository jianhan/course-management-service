package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jianhan/course-management-service/handlers"
	cfgreader "github.com/jianhan/pkg/configs"
	micro "github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// serviceConfigs, err := cfgreader.NewReader(os.Getenv("ENV")).Read()
	serviceConfigs, err := cfgreader.NewReader("development").Read()
	if err != nil {
		panic(fmt.Sprintf("error while reading configurations: %s", err.Error()))
	}
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			viper.GetString("mysql.username"),
			viper.GetString("mysql.password"),
			viper.GetString("mysql.host"),
			viper.GetString("mysql.port"),
			viper.GetString("mysql.database"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	srv := micro.NewService(
		micro.Name(serviceConfigs.Name),
		micro.RegisterTTL(time.Duration(serviceConfigs.RegisterTTL)*time.Second),
		micro.RegisterInterval(time.Duration(serviceConfigs.RegisterInterval)*10),
		micro.Version(serviceConfigs.Version),
		micro.Metadata(serviceConfigs.Metadata),
	)
	srv.Init()
	pb.RegisterCourseManagerHandler(srv.Server(), &handlers.CourseManagement{DB: db})
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
