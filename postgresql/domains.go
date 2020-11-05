package postgresql

import "time"

// Kaders adfasd
type Kaders struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	Name           string     `gorm:"column:name;type:varchar(100);not null;" json:"name"`
	Email          string     `gorm:"column:email;type:varchar(255);unique;not null;" json:"email"`
	NIK            *string    `gorm:"column:nik;type:char(16);unique" json:"nik"`
	DateBirth      time.Time  `gorm:"column:date_birth;type:date;not null;" json:"date_birth"`
	PlaceBirth     string     `gorm:"column:place_birth;type:varchar(50);not null;" json:"place_birth"`
	Avatar         string     `gorm:"column:avatar;type:varchar(255);not null;" json:"avatar"`
	Job            string     `gorm:"column:job;type:varchar(100);not null;" json:"job"`
	Office         string     `gorm:"column:office;type:varchar(100);not null;" json:"office"`
	Skills         string     `gorm:"column:skills;type:varchar(255);not null;" json:"skills"`
	Address        string     `gorm:"column:address;type:varchar(255);not null;" json:"address"`
	Phone          string     `gorm:"column:phone;type:varchar(15);unique;not null;" json:"phone"`
	BloodType      string     `gorm:"column:blood_type;type:char(2);not null;" json:"blood_type"`
	Gender         string     `gorm:"column:gender;type:char(2);not null;" json:"gender"`
	ZipCode        string     `gorm:"column:zip_code;type:varchar(7);not null;" json:"zip_code"`
	ProvinceID     string     `gorm:"column:province_id;type:varchar(2);not null;" json:"province_id"`
	CityID         string     `gorm:"column:city_id;type:varchar(4);not null;" json:"city_id"`
	DistrictID     string     `gorm:"column:district_id;type:varchar(7);not null;" json:"district_id"`
	VillageID      string     `gorm:"column:village_id;type:varchar(10);not null;" json:"village_id"`
	Password       string     `gorm:"column:password;type:varchar(255);not null;" json:"password"`
	Status         string     `gorm:"column:status;type:varchar(30)" json:"status"`
	RegistrationID int     `gorm:"column:registration_id;" json:"registration_id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at"`
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


type IDProvince struct {
	ID   string `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

// IDCities struct for city
type IDCities struct {
	ID         string `gorm:"primary_key" json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

// IDDistricts struct for district
type IDDistricts struct {
	ID     string `gorm:"primary_key" json:"id"`
	CityID string `json:"city_id"`
	Name   string `json:"name"`
}

// IDVillages struct for village
type IDVillages struct {
	ID         string `gorm:"primary_key" json:"id"`
	DistrictID string `json:"district_id"`
	Name       string `json:"name"`
}