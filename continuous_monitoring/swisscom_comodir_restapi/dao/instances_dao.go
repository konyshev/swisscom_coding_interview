package dao

import (
	"log"

	. "github.com/konyshev/swisscom_comodir_restapi/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type InstanceDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "instances"
)

func (m *InstanceDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *InstanceDAO) FindAll() ([]Instance, error) {
	var instances []Instance
	err := db.C(COLLECTION).Find(bson.M{}).All(&instances)
	return instances, err
}

func (m *InstanceDAO) FindById(id string) (Instance, error) {
	var instance Instance
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&instance)
	return instance, err
}

func (m *InstanceDAO) Insert(instance Instance) error {
	err := db.C(COLLECTION).Insert(&instance)
	return err
}

func (m *InstanceDAO) Delete(instance Instance) error {
	err := db.C(COLLECTION).Remove(&instance)
	return err
}

func (m *InstanceDAO) Update(instance Instance) error {
	err := db.C(COLLECTION).UpdateId(instance.ID, &instance)
	return err
}
