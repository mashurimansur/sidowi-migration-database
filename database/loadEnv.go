package database

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Env struct for env file
type Env struct {
	PGUser     string
	PGPassword string
	PGDatabase string
	PGHost     string
	PGPort     string

	MongoUser     string
	MongoPort     string
	MongoDatabase string
	MongoHost     string
	MongoPassword string
}

// LoadEnv return env file to struct
func LoadEnv() (c Env) {
	errLoadEnv := godotenv.Load(".env")
	if errLoadEnv != nil {
		log.Fatal("Error loading .env file")
	}

	c.PGDatabase = os.Getenv("PG_DATABASE")
	c.PGUser = os.Getenv("PG_USERNAME")
	c.PGPassword = os.Getenv("PG_PASSWORD")
	c.PGHost = os.Getenv("PG_HOST")
	c.PGPort = os.Getenv("PG_PORT")
	c.MongoDatabase = os.Getenv("MONGO_DB")
	c.MongoHost = os.Getenv("MONGO_HOST")
	c.MongoPassword = os.Getenv("MONGO_PASSWORD")
	c.MongoUser = os.Getenv("MONGO_USER")
	c.MongoPort = os.Getenv("MONGO_PORT")

	return c
}
