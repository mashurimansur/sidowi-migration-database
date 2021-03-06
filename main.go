package main

import (
	"encoding/json"
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
	dataseed := flag.Bool("dataseed", false, "to drop migrate and seed table")
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

	// drop -> migrate -> seed data
	if *dataseed {
		DropTable()
		fmt.Printf("\033[1;31m%s\033[0m", "Succesfully Drop Table.\n")
		MigrateTable()
		fmt.Printf("\033[1;33m%s\033[0m", "Succesfully Migrate Table.\n")
		SeederTable()
		return
	}

	// test issue

	//var kaders postgresql.Kaders
	//postgres.Preload("IDProvinces").First(&kaders)
	//postgres.Model(postgresql.Kaders{}).Preload("Roles").Where("id = ?", 2).First(&kaders)
	////postgres.Model(postgresql.Kaders{}).Preload("IDProvince").Where("id = ?", 440).First(&kaders)
	//
	//var province postgresql.IDProvinces
	//postgres.Preload("Kaders").Where("id = ?", 73).First(&province)
	//s, _ := json.MarshalIndent(province, "", "\t")
	//fmt.Println(string(s))
	//
	//fmt.Println("Done!")

	var halaqah []postgresql.Halaqahs
	postgres.Model(postgresql.Halaqahs{}).Preload("Kaders").Find(&halaqah)
	s, _ := json.MarshalIndent(halaqah, "", "\t")
	fmt.Println(string(s))

}

// SeederTable : function for seed data to table database
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

// MigrateTable : function for migrate table to database
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
		&postgresql.Halaqahs{},
		&postgresql.HalaqahsKaders{},
		&postgresql.HistoryHalaqah{},
		&postgresql.HistoryKader{},
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

	postgres.Model(postgresql.HalaqahsKaders{}).
		AddForeignKey("kaders_id", "kaders(id)", "CASCADE", "CASCADE").
		AddForeignKey("halaqahs_id", "halaqahs(id)", "CASCADE", "CASCADE")

	postgres.Model(postgresql.Halaqahs{}).AddForeignKey("marhalah_id", "marhalahs(id)", "RESTRICT", "CASCADE")

	postgres.Model(postgresql.HistoryKader{}).AddForeignKey("kader_id", "kaders(id)", "CASCADE", "CASCADE")
	postgres.Model(postgresql.HistoryHalaqah{}).AddForeignKey("halaqah_id", "halaqahs(id)", "CASCADE", "CASCADE")
}

// DropTable : function for drop table database
func DropTable() {
	postgres.DropTableIfExists(
		&postgresql.HistoryHalaqah{},
		&postgresql.HistoryKader{},
		&postgresql.HalaqahsKaders{},
		&postgresql.KadersRoles{},
		&postgresql.Roles{},
		&postgresql.Halaqahs{},
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
