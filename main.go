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
	database.LoadEnv()
	mongo = database.ConnectMongo()
	postgres = database.ConnectPostgres()
}

func main() {
	// close connection DB
	defer func() {
		mongo.Session.Close()
		postgres.Close()
	}()

	//flag
	migrate := flag.Bool("migrate", false, "to migrate table")
	seed := flag.Bool("seed", false, "to insert data to table")
	drop := flag.Bool("drop", false, "to drop table")
	flag.Parse()

	if *migrate {
		MigrateTable()
		return
	}

	if *seed {
		SeederTable()
		return
	}

	if *drop {
		DropTable()
		return
	}

	fmt.Println("Done!")
}

func SeederTable() {
	postgresData := postgresql.NewPostgresConnection(postgres)
	mongoData := mongodb.NewMongoConnection(mongo)
	//postgresData.WorkerKaders(mongoKaders)

	//get data from mongodb
	mongoOpenRegis, _ := mongoData.GetAllOpenRegis()
	mongoKaders, _ := mongoData.GetAllKaders()

	//insert data
	postgresData.WorkerProvince()
	postgresData.WorkerVillage()
	postgresData.WorkerDistrict()
	postgresData.WorkerCity()
	postgresData.SeederOpenRegistration(mongoOpenRegis)
	postgresData.SeederKader(mongoKaders)
}

func MigrateTable() {
	postgres.AutoMigrate(
		&postgresql.Kaders{},
		&postgresql.OpenRegistration{},
		&postgresql.IDProvince{},
		&postgresql.IDCities{},
		&postgresql.IDDistricts{},
		&postgresql.IDVillages{},
	)

	postgres.Model(&postgresql.Kaders{}).AddForeignKey("registration_id", "open_registrations(id)", "RESTRICT", "CASCADE").
		AddForeignKey("province_id", "id_provinces(id)", "RESTRICT", "CASCADE").
		AddForeignKey("city_id", "id_cities(id)", "RESTRICT", "CASCADE").
		AddForeignKey("district_id", "id_districts(id)", "RESTRICT", "CASCADE").
		AddForeignKey("village_id", "id_villages(id)", "RESTRICT", "CASCADE")
}

func DropTable() {
	postgres.DropTableIfExists(&postgresql.Kaders{})
	postgres.DropTableIfExists(&postgresql.OpenRegistration{})
	postgres.DropTableIfExists(&postgresql.IDProvince{})
	postgres.DropTableIfExists(&postgresql.IDCities{})
	postgres.DropTableIfExists(&postgresql.IDDistricts{})
	postgres.DropTableIfExists(&postgresql.IDVillages{})
}
