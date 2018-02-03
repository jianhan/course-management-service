package db

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

// GetDB returns an sql.DB.
func GetDB() (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			viper.GetString("mysql.username"),
			viper.GetString("mysql.password"),
			viper.GetString("mysql.host"),
			viper.GetString("mysql.port"),
			viper.GetString("mysql.database"),
		),
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
