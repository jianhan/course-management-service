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
