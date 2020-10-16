package main

import (
	"flag"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mashurimansur/sidowi-migration-database/database"
	"github.com/mashurimansur/sidowi-migration-database/mongodb"
	"github.com/mashurimansur/sidowi-migration-database/postgresql"
	"gopkg.in/mgo.v2"
)

var (
	mongo    *mgo.Database
	postgres *gorm.DB
)

func init() {
	mongo = database.ConnectMongo()
	postgres = database.ConnectPostgres()
}

func main() {
	// close connection DB
	defer mongo.Session.Close()
	defer postgres.Close()

	migrate := flag.Bool("migrate", false, "to migrate table")
	seed := flag.Bool("seed", false, "to insert data to table")
	flag.Parse()

	if *migrate {
		MigrateTable()
		return
	}

	if *seed {
		mongoData := mongodb.NewMongoConnection(mongo)
		postgresData := postgresql.NewPostgresConnection(postgres)

		mongoKaders, _ := mongoData.GetAllKaders()
		mongoOpenRegis, _ := mongoData.GetAllOpenRegis()

		postgresData.SeederOpenRegistration(mongoOpenRegis)
		postgresData.SeederKader(mongoKaders)
		return
	}

	fmt.Println("Tidak terjadi apa apa")
}

func MigrateTable() {
	postgres.AutoMigrate(&postgresql.Kaders{})
	postgres.AutoMigrate(&postgresql.OpenRegistration{})
}
