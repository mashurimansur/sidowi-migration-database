package postgresql

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gammazero/workerpool"

	"github.com/cheggaaa/pb"
	"github.com/mashurimansur/sidowi-migration-database/mongodb"

	"github.com/jinzhu/gorm"
)

type PostgresConnection struct {
	DB *gorm.DB
}

func NewPostgresConnection(db *gorm.DB) *PostgresConnection {
	return &PostgresConnection{DB: db}
}

func (postgres *PostgresConnection) InsertProvince() {
	rows := postgres.ReadFileCSV("indonesia_provinces.csv")

	for _, row := range rows[1:] {
		province := IDProvince{
			ID:   row[0],
			Name: row[1],
		}
		postgres.DB.Model(IDProvince{}).Create(&province)
	}
}

func (postgres *PostgresConnection) InsertCity() {
	rows := postgres.ReadFileCSV("indonesia_cities.csv")

	for _, row := range rows[1:] {
		city := IDCities{
			ID:         row[0],
			ProvinceID: row[1],
			Name:       row[2],
		}
		postgres.DB.Model(IDCities{}).Create(&city)
	}
}

func (postgres *PostgresConnection) InsertDistrict() {
	rows := postgres.ReadFileCSV("indonesia_districts.csv")

	for _, row := range rows[1:] {
		district := IDDistricts{
			ID:     row[0],
			CityID: row[1],
			Name:   row[2],
		}
		postgres.DB.Model(IDDistricts{}).Create(&district)
	}
}

func (postgres *PostgresConnection) InsertVillage() {
	rows := postgres.ReadFileCSV("indonesia_villages.csv")

	for _, row := range rows[1:] {
		village := IDVillages{
			ID:         row[0],
			DistrictID: row[1],
			Name:       row[2],
		}
		postgres.DB.Model(IDVillages{}).Create(&village)
	}
}

func (posgres *PostgresConnection) ReadFileCSV(csvFile string) (result [][]string) {
	csvFileProvince, _ := os.Open("./database/csv_indonesia/" + csvFile)
	r := csv.NewReader(csvFileProvince)
	// Iterate through the records
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		data := []string{}

		if len(record) == 3 {
			data = []string{record[0], record[1], record[2]}
		}

		if len(record) == 2 {
			data = []string{record[0], record[1]}
		}

		result = append(result, data)
	}
	return result
}

func (postgres *PostgresConnection) InsertKader(kader *Kaders) (err error) {
	if err = postgres.DB.Model(Kaders{}).Create(&kader).Error; err != nil {
		log.Println(kader.Name)
		return err
	}
	return
}

func (postgres *PostgresConnection) SeederKader(mongoKaders []mongodb.MongoKaders) {
	bar := pb.StartNew(len(mongoKaders))
	for _, value := range mongoKaders {
		idOpenRegis, errFind := postgres.FindOpenRegis(value.RegistrationLog.Title)
		var kader Kaders
		kader.Email = value.Email
		kader.Name = value.Name
		if value.NIK != "" {
			kader.NIK = &value.NIK
		} else {
			kader.NIK = nil
		}
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
			kader.RegistrationID = int(idOpenRegis.ID)
		}
		kader.CreatedAt = value.CreatedAt
		kader.UpdatedAt = value.UpdateAt
		kader.DeletedAt = value.DeletedAt

		errInsert := postgres.InsertKader(&kader)
		if errInsert == nil {
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
	}
	bar.Finish()
	fmt.Println("Succesfully Migrating Kader")
}

func (postgres *PostgresConnection) InsertOpenRegistration(openRegis *OpenRegistration) (err error) {
	if err = postgres.DB.Model(OpenRegistration{}).Create(&openRegis).Error; err != nil {
		return err
	}
	return
}

func (postgres *PostgresConnection) SeederOpenRegistration(mongoOpeRegis []mongodb.MongoOpenRegistration) {
	bar := pb.StartNew(len(mongoOpeRegis))

	for _, value := range mongoOpeRegis {
		var openRegis OpenRegistration
		openRegis.Title = value.Title
		openRegis.Description = value.Description
		openRegis.Code = value.Code
		openRegis.OpenRegistration = value.OpenRegistration
		openRegis.CloseRegistration = value.CloseRegistration
		openRegis.CreatedAt = value.CreatedAt
		openRegis.UpdatedAt = value.UpdateAt
		openRegis.DeletedAt = value.DeletedAt

		errInsert := postgres.InsertOpenRegistration(&openRegis)
		if errInsert == nil {
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
	}
	bar.Finish()
	fmt.Println("Successfully Migrating OpenRegistration")
}

func (postgres *PostgresConnection) FindOpenRegis(title string) (openRegis OpenRegistration, err error) {
	if err = postgres.DB.Model(OpenRegistration{}).Where("title = ?", title).First(&openRegis).Error; err != nil {
		return openRegis, err
	}
	return
}

func (postgres *PostgresConnection) WorkerKaders(kaders []mongodb.MongoKaders) {
	wp := workerpool.New(10)

	for _, kader := range kaders {
		kader := kader
		wp.Submit(func() {
			data := postgres.KadersTransformer(kader)
			postgres.InsertKader(&data)
		})
	}
}

func (postgres *PostgresConnection) KadersTransformer(mongoKader mongodb.MongoKaders) (kader Kaders) {
	kader.Email = mongoKader.Email
	kader.Name = mongoKader.Name
	kader.DateBirth = mongoKader.DateBirth
	kader.PlaceBirth = mongoKader.PlaceBirth
	kader.Avatar = mongoKader.Avatar
	kader.Job = mongoKader.Job
	kader.Skills = mongoKader.Skills
	kader.Office = mongoKader.Office
	kader.Address = mongoKader.Address
	kader.Phone = mongoKader.Phone
	kader.BloodType = mongoKader.BloodType
	kader.Gender = mongoKader.Gender
	kader.ZipCode = mongoKader.ZipCode
	kader.ProvinceID = strconv.Itoa(mongoKader.Province)
	kader.CityID = strconv.Itoa(mongoKader.City)
	kader.DistrictID = strconv.Itoa(mongoKader.District)
	kader.VillageID = strconv.Itoa(mongoKader.Village)
	kader.Password = mongoKader.Password
	kader.Status = mongoKader.Status
	kader.CreatedAt = mongoKader.CreatedAt
	kader.UpdatedAt = mongoKader.UpdateAt
	kader.DeletedAt = mongoKader.DeletedAt

	if mongoKader.NIK != "" {
		kader.NIK = &mongoKader.NIK
	} else {
		kader.NIK = nil
	}

	idOpenRegis, errFind := postgres.FindOpenRegis(mongoKader.RegistrationLog.Title)
	if errFind == nil {
		kader.RegistrationID = int(idOpenRegis.ID)
	}

	return
}
