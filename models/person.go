package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Nationality string `json:"nationality"`
	Sex         string `json:"sex"`
	Age         uint8  `json:"age"`
	Height      uint8  `json:"height"`
	Weight      uint8  `json:"weight"`
}
type Employee struct {
	Salary float32 `json:"salary"`
	Post   string  `json:"post"`
	person *Person
}

type PersonModel struct {
	Collection *mgo.Collection
}

var PersonModelConnectionTemp = PersonModel{Collection: dbConnect.Use(databaseName, "Persons")}

func (p *PersonModel) GetAllPersons() (persons []*Person, err error) {
	err = PersonModelConnectionTemp.Collection.Find(bson.M{}).All(&persons)
	return persons, err
}
