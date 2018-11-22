package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// InitMongo starts a mongo session
func InitMongo(arg string) *mgo.Session {
	session, err := mgo.Dial(arg)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	return session
}
