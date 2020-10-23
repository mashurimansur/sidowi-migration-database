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

	//flag
	migrate := flag.Bool("migrate", false, "to migrate table")
	seed := flag.Bool("seed", false, "to insert data to table")
	drop := flag.Bool("drop", false, "to drop table")
	flag.Parse()

	if *migrate {
		MigrateTable()
	}

	if *seed {
		SeederTable()
	}

	if *drop {
		DropTable()
	}

	fmt.Println("Done!")
}

func SeederTable() {
	postgresData := postgresql.NewPostgresConnection(postgres)
	mongoData := mongodb.NewMongoConnection(mongo)

	mongoKaders, _ := mongoData.GetAllKaders()
	//postgresData.SeederKader(mongoKaders)
	postgresData.WorkerKaders(mongoKaders)

	mongoOpenRegis, _ := mongoData.GetAllOpenRegis()
	postgresData.SeederOpenRegistration(mongoOpenRegis)
}

func MigrateTable() {
	postgres.AutoMigrate(&postgresql.Kaders{}, &postgresql.OpenRegistration{})
}

func DropTable() {
	postgres.DropTableIfExists(&postgresql.Kaders{})
	postgres.DropTableIfExists(&postgresql.OpenRegistration{})
}
