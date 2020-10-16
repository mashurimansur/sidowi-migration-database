package postgresql

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PostgresConnection struct {
	DB *gorm.DB
}

func NewPostgresConnection(db *gorm.DB) *PostgresConnection {
	return &PostgresConnection{DB: db}
}

type Kaders struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at"`
	Name           string     `gorm:"column:name;type:varchar(100)" json:"name"`
	Email          string     `gorm:"column:email;type:varchar(255)" json:"email"`
	NIK            string     `gorm:"column:nik;type:char(16)"" json:"nik"`
	DateBirth      time.Time  `gorm:"column:date_birth;type:date" json:"date_birth"`
	PlaceBirth     string     `gorm:"column:place_birth;type:varchar(50)" json:"place_birth"`
	Avatar         string     `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
	Job            string     `gorm:"column:job;type:varchar(100)" json:"job"`
	Office         string     `gorm:"column:office;type:varchar(100)" json:"office"`
	Skills         string     `gorm:"column:skills;type:varchar(255)" json:"skills"`
	Address        string     `gorm:"column:address;type:varchar(255)" json:"address"`
	Phone          string     `gorm:"column:phone;type:varchar(15)" json:"phone"`
	BloodType      string     `gorm:"column:blood_type;type:char(2)" json:"blood_type"`
	Gender         string     `gorm:"column:gender;type:char(2)" json:"gender"`
	ZipCode        string     `gorm:"column:zip_code;type:varchar(7)" json:"zip_code"`
	ProvinceID     string     `gorm:"column:province_id;type:varchar(2)" json:"province_id"`
	CityID         string     `gorm:"column:city_id;type:varchar(4)" json:"city_id"`
	DistrictID     string     `gorm:"column:district_id;type:varchar(7)" json:"district_id"`
	VillageID      string     `gorm:"column:village_id;type:varchar(10)" json:"village_id"`
	Password       string     `gorm:"column:password;type:varchar(255)" json:"password"`
	Status         string     `gorm:"column:status;type:varchar(30)" json:"status"`
	RegistrationID string     `gorm:"column:registration_id;type:varchar(5)" json:"registration_id"`
}

type OpenRegistration struct {
	ID                uint       `gorm:"primary_key" json:"id"`
	Title             string     `gorm:"column:title;type:varchar(100)" json:"title"`
	Description       string     `gorm:"column:description;type:varchar(100)" json:"description"`
	OpenRegistration  time.Time  `gorm:"column:open_registration;type:date" json:"open_registration"`
	CloseRegistration time.Time  `gorm:"column:close_registration;type:date" json:"close_registration"`
	Code              string     `gorm:"column:code;type:varchar(5)" json:"code"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `sql:"index" json:"deleted_at"`
}

func (postgres *PostgresConnection) InsertKader(kader *Kaders) (err error) {
	if err = postgres.DB.Model(Kaders{}).Create(&kader).Error; err != nil {
		return err
	}
	return
}

func (postgres *PostgresConnection) InsertOpenRegistration(openRegis *OpenRegistration) (err error) {
	if err = postgres.DB.Model(OpenRegistration{}).Create(&openRegis).Error; err != nil {
		return err
	}
	return
}

func (postgres *PostgresConnection) FindOpenRegis(title string) (openRegis OpenRegistration, err error) {
	if err = postgres.DB.Model(OpenRegistration{}).Where("title = ?", title).First(&openRegis).Error; err != nil {
		return openRegis, err
	}
	return
}
