package repositories

import (
	mgo "gopkg.in/mgo.v2"
)

// RepoInitFunc represents a common initialization function.
type RepoInitFunc func(session *mgo.Session)

// Initialize takes all initialization as functions and execute them.
func Initialize(session *mgo.Session, funcs ...RepoInitFunc) {
	for _, v := range funcs {
		v(session)
	}
}

// InitCourse performs initializations for courses collection such as index, etc..
var InitCourse RepoInitFunc = func(session *mgo.Session) {
	s := session.Clone()
	defer s.Close()
	c := s.DB(dbName).C(coursesCollection)
	// create unique index on slug
	slugIndex := mgo.Index{
		Key:        []string{"slug"},
		Unique:     true,
		DropDups:   false,
		Background: false,
		Sparse:     true,
	}
	err := c.EnsureIndex(slugIndex)
	if err != nil {
		panic(err)
	}
	textIndex := mgo.Index{
		Key: []string{"$text:name", "$text:description"},
	}
	err = c.EnsureIndex(textIndex)
	if err != nil {
		panic(err)
	}
}

// InitCategories performs initializations for categories collection such as index, etc..
var InitCategories RepoInitFunc = func(session *mgo.Session) {
	s := session.Clone()
	defer s.Close()
	c := s.DB(dbName).C(categoriesCollection)
	// create unique index on slug
	slugIndex := mgo.Index{
		Key:        []string{"slug"},
		Unique:     true,
		DropDups:   false,
		Background: false,
		Sparse:     true,
	}
	err := c.EnsureIndex(slugIndex)
	if err != nil {
		panic(err)
	}
	textIndex := mgo.Index{
		Key: []string{"$text:name", "$text:description"},
	}
	err = c.EnsureIndex(textIndex)
	if err != nil {
		panic(err)
	}
}
