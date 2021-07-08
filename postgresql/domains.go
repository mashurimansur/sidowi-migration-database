package postgresql

import (
	"time"

	guuid "github.com/google/uuid"
)

// Kaders adfasd
type Kaders struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	Name           string     `gorm:"column:name;type:varchar(100);not null;" json:"name"`
	Email          string     `gorm:"column:email;type:varchar(255);unique;not null;" json:"email"`
	NIK            *string    `gorm:"column:nik;type:varchar(16);unique" json:"nik"`
	DateBirth      time.Time  `gorm:"column:date_birth;type:date;not null;" json:"date_birth"`
	PlaceBirth     string     `gorm:"column:place_birth;type:varchar(50);not null;" json:"place_birth"`
	Avatar         string     `gorm:"column:avatar;type:varchar(255);not null;" json:"avatar"`
	Job            string     `gorm:"column:job;type:varchar(100);not null;" json:"job"`
	Office         string     `gorm:"column:office;type:varchar(100);not null;" json:"office"`
	Skills         string     `gorm:"column:skills;type:varchar(255);not null;" json:"skills"`
	Address        string     `gorm:"column:address;type:varchar(255);not null;" json:"address"`
	Phone          string     `gorm:"column:phone;type:varchar(15);unique;not null;" json:"phone"`
	BloodType      string     `gorm:"column:blood_type;type:varchar(2);not null;" json:"blood_type"`
	Gender         string     `gorm:"column:gender;type:varchar(1);not null;" json:"gender"`
	ZipCode        string     `gorm:"column:zip_code;type:varchar(7);not null;" json:"zip_code"`
	ProvinceID     string     `gorm:"column:province_id;type:char(2);not null;" json:"province_id"`
	CityID         string     `gorm:"column:city_id;type:char(4);not null;" json:"city_id"`
	DistrictID     string     `gorm:"column:district_id;type:char(7);not null;" json:"district_id"`
	VillageID      string     `gorm:"column:village_id;type:char(10);not null;" json:"village_id"`
	RegistrationID *uint      `gorm:"column:registration_id;" json:"registration_id"`
	Password       string     `gorm:"column:password;type:varchar(255);not null;" json:"password"`
	Status         string     `gorm:"column:status;type:varchar(30)" json:"status"`
	CampusName     string     `gorm:"column:campus_name;type:varchar(30)" json:"campus_name"`
	CampusMajor    string     `gorm:"column:campus_major;type:varchar(30)" json:"campus_major"`
	CampusBatch    string     `gorm:"column:campus_batch;type:varchar(4)" json:"campus_batch"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at"`

	//Roles []Roles `gorm:"many2many:kaders_roles" json:",omitempty"`
	//Province IDProvince `gorm:"foreignkey:ID;references:province_id"`
	Province         IDProvinces
	OpenRegistration OpenRegistration
}

type OpenRegistration struct {
	ID                uint       `gorm:"primary_key" json:"id"`
	Title             string     `gorm:"column:title;type:varchar(100)" json:"title"`
	Description       string     `gorm:"column:description;type:varchar(100)" json:"description"`
	OpenRegistration  time.Time  `gorm:"column:open_registration;type:date" json:"open_registration"`
	CloseRegistration time.Time  `gorm:"column:close_registration;type:date" json:"close_registration"`
	Code              string     `gorm:"column:code;type:varchar(5)" json:"code"`
	UUID              guuid.UUID `json:"uuid"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `sql:"index" json:"deleted_at"`
}

type Marhalahs struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type IDProvinces struct {
	ID   string `gorm:"type:char(2)" json:"id"`
	Name string `json:"name"`

	Kaders []Kaders `gorm:"ForeignKey:ProvinceID"`
}

// IDCities struct for city
type IDCities struct {
	ID         string `gorm:"type:char(4)" json:"id"`
	ProvinceID string `gorm:"type:char(2)" json:"province_id"`
	Name       string `json:"name"`
}

// IDDistricts struct for district
type IDDistricts struct {
	ID     string `gorm:"type:char(7)" json:"id"`
	CityID string `gorm:"type:char(4)" json:"city_id"`
	Name   string `json:"name"`
}

// IDVillages struct for village
type IDVillages struct {
	ID         string `gorm:"type:char(10)" json:"id"`
	DistrictID string `gorm:"type:char(7)" json:"district_id"`
	Name       string `json:"name"`
}

type Roles struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `gorm:"column:name;type:varchar(20);unique;not null;" json:"name"`
	Description string    `gorm:"column:description;type:varchar(100);" json:"description"`

	//Kaders []Kaders `gorm:"many2many:kaders_roles" json:",omitempty"`
}

type KadersRoles struct {
	KadersID uint
	RolesID  uint
}
