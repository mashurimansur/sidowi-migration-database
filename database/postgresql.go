package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// ConnectPostgres to connect database postgresql
func ConnectPostgres() *gorm.DB {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Environment.PGHost, Environment.PGPort, Environment.PGUser, Environment.PGPassword, Environment.PGDatabase)
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		log.Panic(err.Error())
	}
	return db
}
