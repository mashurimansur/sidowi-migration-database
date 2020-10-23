package database

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

func ConnectMongo() *mgo.Database {
	env := LoadEnv()
	mongoAdds := fmt.Sprintf("%s:%s", env.MongoHost, env.MongoPort)
	infoDial := &mgo.DialInfo{
		Addrs:    []string{mongoAdds},
		Timeout:  120 * time.Second,
		Database: env.MongoDatabase,
		Username: env.MongoUser,
		Password: env.MongoPassword,
	}

	session, _ := mgo.DialWithInfo(infoDial)
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(env.MongoDatabase)
	return collection
}
