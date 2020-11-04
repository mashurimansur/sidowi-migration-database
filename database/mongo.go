package database

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

func ConnectMongo() *mgo.Database {
	mongoAdds := fmt.Sprintf("%s:%s", Environment.MongoHost, Environment.MongoPort)
	infoDial := &mgo.DialInfo{
		Addrs:    []string{mongoAdds},
		Timeout:  120 * time.Second,
		Database: Environment.MongoDatabase,
		Username: Environment.MongoUser,
		Password: Environment.MongoPassword,
	}

	session, _ := mgo.DialWithInfo(infoDial)
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(Environment.MongoDatabase)
	return collection
}
