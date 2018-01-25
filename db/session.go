package db

import (
	mgo "gopkg.in/mgo.v2"
)

func NewMongodSession(url string) (*mgo.Session, error) {
	s, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	s.SetMode(mgo.Monotonic, true)
	return s, nil
}
