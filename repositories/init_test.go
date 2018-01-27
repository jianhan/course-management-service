package repositories

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func TestInitialize(t *testing.T) {
	type args struct {
		session *mgo.Session
		funcs   []RepoInitFunc
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Initialize(tt.args.session, tt.args.funcs...)
		})
	}
}
