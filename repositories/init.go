package repositories

import (
	mgo "gopkg.in/mgo.v2"
)

type RepoInitFunc func(session *mgo.Session)

func Initialize(session *mgo.Session, funcs ...RepoInitFunc) {
	for _, v := range funcs {
		v(session)
	}
}
