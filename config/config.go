package dmas

import (
	"gopkg.in/mgo.v2"
)

// Config stores global configuration - todo load it from configuration file
const (
	Host   = "localhost"
	DbName = "quiz_db"
)

//just start a mongo session and share it via variable to other files
var (
	MgoSession *mgo.Session
)

//init function is required if we want to initialize something
func init() {
	session, err := mgo.Dial(Host)
	if err != nil {
		panic(err)
	}

	MgoSession = session
}
