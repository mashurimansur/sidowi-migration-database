package mongodb

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoConnection struct {
	Mongo *mgo.Database
}

type MongoKaders struct {
	ID              bson.ObjectId   `bson:"_id,omitempty"`
	Name            string          `bson:"name"`
	Email           string          `bson:"email"`
	PlaceBirth      string          `bson:"place_birth"`
	DateBirth       time.Time       `bson:"date_birth"`
	Gender          string          `bson:"gender"`
	Job             string          `bson:"job"`
	Office          string          `bson:"office"`
	Skills          string          `bson:"skills"`
	Address         string          `bson:"address"`
	Phone           string          `bson:"phone"`
	City            int             `bson:"city"`
	Province        int             `bson:"province"`
	District        int             `bson:"district"`
	Village         int             `bson:"village"`
	BloodType       string          `bson:"blood_type"`
	Password        string          `bson:"password"`
	ZipCode         string          `bson:"zip_code"`
	Status          string          `bson:"status"`
	Avatar          string          `bson:"avatar"`
	UpdateAt        time.Time       `bson:"update_at"`
	CreatedAt       time.Time       `bson:"created_at"`
	DeletedAt       *time.Time      `bson:"deleted_at"`
	NIK             string          `bson:"nik"`
	RegistrationLog RegistrationLog `bson:"registration_log"`
}

type RegistrationLog struct {
	IDRegistration string `bson:"id_registration"`
	Title          string `bson:"title"`
}

type MongoOpenRegistration struct {
	ID                bson.ObjectId `bson:"_id,omitempty"`
	UpdateAt          time.Time     `bson:"update_at"`
	CreatedAt         time.Time     `bson:"created_at"`
	DeletedAt         *time.Time    `bson:"deleted_at"`
	Title             string        `bson:"title"`
	Description       string        `bson:"description"`
	Code              string        `bson:"code"`
	OpenRegistration  time.Time     `bson:"open_registration"`
	CloseRegistration time.Time     `bson:"close_registration"`
}

func NewMongoConnection(mongo *mgo.Database) *MongoConnection {
	return &MongoConnection{Mongo: mongo}
}

func (mongo *MongoConnection) GetAllKaders() (kaders []MongoKaders, err error) {
	if err = mongo.Mongo.C("kaders").Find(nil).All(&kaders); err != nil {
		return nil, err
	}
	return
}

func (mongo *MongoConnection) GetAllOpenRegis() (openRegis []MongoOpenRegistration, err error) {
	if err = mongo.Mongo.C("open_registration").Find(nil).All(&openRegis); err != nil {
		return nil, err
	}
	return
}
