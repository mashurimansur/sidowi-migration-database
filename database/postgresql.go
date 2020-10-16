package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

// ConnectPostgres to connect database postgresql
func ConnectPostgres() *gorm.DB {
	env := LoadEnv()
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env.PGHost, env.PGPort, env.PGUser, env.PGPassword, env.PGDatabase)
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		log.Panic(err.Error())
	}
	return db
}
