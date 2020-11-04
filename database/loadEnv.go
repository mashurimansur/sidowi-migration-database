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

var Environment Env

// LoadEnv return env file to struct
func LoadEnv() {
	errLoadEnv := godotenv.Load(".env")
	if errLoadEnv != nil {
		log.Fatal("Error loading .env file")
	}

	Environment.PGDatabase = os.Getenv("PG_DATABASE")
	Environment.PGUser = os.Getenv("PG_USERNAME")
	Environment.PGPassword = os.Getenv("PG_PASSWORD")
	Environment.PGHost = os.Getenv("PG_HOST")
	Environment.PGPort = os.Getenv("PG_PORT")
	Environment.MongoDatabase = os.Getenv("MONGO_DB")
	Environment.MongoHost = os.Getenv("MONGO_HOST")
	Environment.MongoPassword = os.Getenv("MONGO_PASSWORD")
	Environment.MongoUser = os.Getenv("MONGO_USER")
	Environment.MongoPort = os.Getenv("MONGO_PORT")
}
