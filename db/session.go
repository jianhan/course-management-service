package db

import (
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

func NewMongodSession() (*mgo.Session, func(), error) {
	closeFunc := func() {}
	db, err := mgo.Dial(viper.GetString("mongo.url"))
	if err != nil {
		return nil, nil, err
	}
	closeFunc = func() {
		db.Close()
	}
	return db, closeFunc, nil
}
