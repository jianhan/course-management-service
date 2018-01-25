package db

import (
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

func NewMongodSession(url string) (*mgo.Session, error) {
	s, err := mgo.Dial(viper.GetString("mongo.url"))
	if err != nil {
		return nil, err
	}
	return s, nil
}
