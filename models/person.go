package models

import (
	"awesomeProject/forms"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Id      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string        `json:"name" bson:"name"`
	Surname string        `json:"surname" bson:"surname"`
	Sex     string        `json:"sex" bson:"sex"`
	Age     uint8         `json:"age" bson:"age"`
	//Nationality string `json:"nationality" bson:"nationality"`
}
type PersonModel struct {
	Collection *mgo.Collection
}

var PersonModelConnectionTemp = PersonModel{Collection: dbConnect.Use(databaseName, "Persons")}

func (p *PersonModel) GetAllPersons() (persons []*Person, err error) {
	err = PersonModelConnectionTemp.Collection.Find(bson.M{}).All(&persons)
	return persons, err
}

/*func (p *PersonModel) GetPersonByID(id bson.ObjectId) (person *Person, err error) {
	err = PersonModelConnectionTemp.Collection.Find(bson.M{"_id":id}).One(&person)
	return person, err
}
*/
func (p *PersonModel) GetPersonByName(name, surname string) (person *Person, err error) {
	err = PersonModelConnectionTemp.Collection.Find(bson.M{"name": name, "surname": surname}).One(&person)
	return person, err
}
func (p *PersonModel) AddPerson(data forms.AddPersonCommand) error {
	err := PersonModelConnectionTemp.Collection.Insert(data)
	return err
}
func (p *PersonModel) ChangePersonInfo(data Person) error {
	err := PersonModelConnectionTemp.Collection.Update(bson.M{"_id": data.Id}, data)
	return err
}
func (p *PersonModel) DeletePerson(id bson.ObjectId) error {
	err := PersonModelConnectionTemp.Collection.Remove(bson.M{"_id": id})
	return err
}
