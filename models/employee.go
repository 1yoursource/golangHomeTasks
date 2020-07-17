package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Employee struct {
	Id     bson.ObjectId `json:"_id" bson:"_id"`
	Salary float32       `json:"salary" bson:"salary"`
	Post   string        `json:"post" bson:"post"`
	person *Person
}
type EmployeeModel struct {
	Collection *mgo.Collection
}

var EmployeeModelConnectionTemp = EmployeeModel{Collection: dbConnect.Use(databaseName, "Employees")}

func (e *EmployeeModel) GetAllEmployees() (employees []*Employee, err error) {
	err = EmployeeModelConnectionTemp.Collection.Find(bson.M{}).All(&employees)
	return employees, err
}
