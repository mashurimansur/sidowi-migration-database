package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/cheggaaa/pb"
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
		InsertDataOpenRegistration()
		InsertDataKaders()
		return
	}

	fmt.Println("Tidak terjadi apa apa")
}

func MigrateTable() {
	postgres.AutoMigrate(&postgresql.Kaders{})
	postgres.AutoMigrate(&postgresql.OpenRegistration{})
}

func InsertDataKaders() {
	mongoData := mongodb.NewMongoConnection(mongo)
	mongoKaders, _ := mongoData.GetAllKaders()

	postgresData := postgresql.NewPostgresConnection(postgres)

	bar := pb.StartNew(len(mongoKaders))

	for _, value := range mongoKaders {

		idOpenRegis, errFind := postgresData.FindOpenRegis(value.RegistrationLog.Title)

		var kader postgresql.Kaders
		kader.Email = value.Email
		kader.Name = value.Name
		kader.NIK = value.NIK
		kader.DateBirth = value.DateBirth
		kader.PlaceBirth = value.PlaceBirth
		kader.Avatar = value.Avatar
		kader.Job = value.Job
		kader.Skills = value.Skills
		kader.Office = value.Office
		kader.Address = value.Address
		kader.Phone = value.Phone
		kader.BloodType = value.BloodType
		kader.Gender = value.Gender
		kader.ZipCode = value.ZipCode
		kader.ProvinceID = strconv.Itoa(value.Province)
		kader.CityID = strconv.Itoa(value.City)
		kader.DistrictID = strconv.Itoa(value.District)
		kader.VillageID = strconv.Itoa(value.Village)
		kader.Password = value.Password
		kader.Status = value.Status
		if errFind == nil {
			kader.RegistrationID = fmt.Sprint(idOpenRegis.ID)
		}
		kader.CreatedAt = value.CreatedAt
		kader.UpdatedAt = value.UpdateAt
		kader.DeletedAt = value.DeletedAt
		errInsert := postgresData.InsertKader(&kader)
		if errInsert == nil {
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
	}
	bar.Finish()
	fmt.Println("Succesfully Migrating Kader")
}

func InsertDataOpenRegistration() {
	mongoData := mongodb.NewMongoConnection(mongo)
	mongoOpeRegis, _ := mongoData.GetAllOpenRegis()

	postgresData := postgresql.NewPostgresConnection(postgres)

	bar := pb.StartNew(len(mongoOpeRegis))

	for _, value := range mongoOpeRegis {
		var openRegis postgresql.OpenRegistration
		openRegis.Title = value.Title
		openRegis.Description = value.Description
		openRegis.Code = value.Code
		openRegis.OpenRegistration = value.OpenRegistration
		openRegis.CloseRegistration = value.CloseRegistration
		openRegis.CreatedAt = value.CreatedAt
		openRegis.UpdatedAt = value.UpdateAt
		openRegis.DeletedAt = value.DeletedAt

		errInsert := postgresData.InsertOpenRegistration(&openRegis)
		if errInsert == nil {
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
	}
	bar.Finish()
	fmt.Println("Successfully Migrating OpenRegistration")
}
