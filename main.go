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

	//var kaders postgresql.Kaders
	//postgres.Model(postgresql.Kaders{}).Preload("Roles").Where("id = ?", 2).First(&kaders)
	////postgres.Model(postgresql.Kaders{}).Preload("IDProvince").Where("id = ?", 440).First(&kaders)
	//
	//s, _ := json.MarshalIndent(kaders, "", "\t")
	//fmt.Println(string(s))

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
	postgresData.RunningWorkerIndonesia()
	postgresData.InsertRoles()
	postgresData.SeederOpenRegistration(mongoOpenRegis)
	postgresData.SeederKader(mongoKaders)
}

func MigrateTable() {
	postgres.AutoMigrate(
		&postgresql.Kaders{},
		&postgresql.Marhalahs{},
		&postgresql.OpenRegistration{},
		&postgresql.IDProvinces{},
		&postgresql.IDCities{},
		&postgresql.IDDistricts{},
		&postgresql.IDVillages{},
		&postgresql.Roles{},
		&postgresql.KadersRoles{},
	)

	postgres.Model(&postgresql.Kaders{}).
		AddForeignKey("registration_id", "open_registrations(id)", "RESTRICT", "CASCADE").
		AddForeignKey("province_id", "id_provinces(id)", "RESTRICT", "CASCADE").
		AddForeignKey("city_id", "id_cities(id)", "RESTRICT", "CASCADE").
		AddForeignKey("district_id", "id_districts(id)", "RESTRICT", "CASCADE").
		AddForeignKey("village_id", "id_villages(id)", "RESTRICT", "CASCADE")

	postgres.Model(postgresql.KadersRoles{}).
		AddForeignKey("kaders_id", "kaders(id)", "CASCADE", "CASCADE").
		AddForeignKey("roles_id", "roles(id)", "CASCADE", "CASCADE")
}

func DropTable() {
	postgres.DropTableIfExists(
		&postgresql.KadersRoles{},
		&postgresql.Roles{},
		&postgresql.Marhalahs{},
		&postgresql.Kaders{},
		&postgresql.OpenRegistration{},
		&postgresql.IDProvinces{},
		&postgresql.IDCities{},
		&postgresql.IDCities{},
		&postgresql.IDDistricts{},
		&postgresql.IDVillages{},
	)

}
