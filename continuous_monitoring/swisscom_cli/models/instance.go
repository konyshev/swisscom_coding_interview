package models

import "gopkg.in/mgo.v2/bson"

type Instance struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Name    string        `bson:"meta_name" json:"meta_name"`
	Host    string        `bson:"host" json:"host"`
	IP      string        `bson:"ip" json:"ip"`
	Contact `bson:"contact" json:"contact"`
}

type Contact struct {
	Email string `bson:"mail" json:"mail"`
	Tel   string `bson:"tel" json:"tel"`
}
