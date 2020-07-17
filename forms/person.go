package forms

type AddPersonCommand struct {
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
	Sex     string `json:"sex" bson:"sex"`
	Age     uint8  `json:"age" bson:"age"`
}
