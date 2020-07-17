package controllers

import (
	"awesomeProject/forms"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"regexp"
	"strconv"
)

var personModel = new(models.PersonModel)

type PersonCotroller struct{}

//Add template
func (p *PersonCotroller) GetAllPersons(c *gin.Context) {
	persons, err := personModel.GetAllPersons()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "PAGE", persons)
}
func (p *PersonCotroller) AddNewPerson(c *gin.Context) {
	var data forms.AddPersonCommand

	data.Name = c.PostForm("name")
	data.Surname = c.PostForm("surname")
	data.Sex = c.PostForm("sex")
	age, _ := strconv.ParseInt(c.PostForm("age"), 10, 32)
	data.Age = uint8(age)

	if CheckFieldNoEmpty(data.Name, data.Surname, data.Sex, strconv.FormatInt(age, 10)) {
		c.JSON(401, gin.H{"message": "Provide relevant fields!"})
		//fmt.Println(401,"Provide relevant fields")
		return
	}

	boolTemp, index := CheckHasNoDigits(data.Name, data.Surname, data.Sex)
	if boolTemp {
		c.JSON(402, gin.H{"message": "Field" + strconv.Itoa(index) + "can`t contain any digits!"})
		//fmt.Println(406,"Field",index,"can`t contain any digits!")
		return
	}

	if CheckAgeHasNoAlphabetic(age) {
		c.JSON(402, gin.H{"message": "Age can`t contain any alphabetic symbols!"})
		//fmt.Println(406,"Age can`t contain any alphabetic symbols!")
		return
	}

	result, _ := personModel.GetPersonByName(data.Name, data.Surname)
	if result.Name != "" {
		c.JSON(405, gin.H{"message": "This person is already exist!"})
		return
	}

	err := personModel.AddPerson(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Problem adding a person! Try again later"})
		return
	}
}
func (p *PersonCotroller) ChangePersonInfo(c *gin.Context) {
	var data models.Person

	data.Id = bson.ObjectIdHex(c.Param("id"))
	data.Name = c.PostForm("name")
	data.Surname = c.PostForm("surname")
	data.Sex = c.PostForm("sex")
	age, _ := strconv.ParseInt(c.PostForm("age"), 10, 32)
	data.Age = uint8(age)

	if CheckFieldNoEmpty(data.Name, data.Surname, data.Sex, strconv.FormatInt(age, 10)) {
		c.JSON(401, gin.H{"message": "Provide relevant fields!"})
		//fmt.Println(401,"Provide relevant fields")
		return
	}

	boolTemp, index := CheckHasNoDigits(data.Name, data.Surname, data.Sex)
	if boolTemp {
		c.JSON(402, gin.H{"message": "Field" + strconv.Itoa(index) + "can`t contain any digits!"})
		//fmt.Println(406,"Field",index,"can`t contain any digits!")
		return
	}

	if CheckAgeHasNoAlphabetic(age) {
		c.JSON(402, gin.H{"message": "Age can`t contain any alphabetic symbols!"})
		//fmt.Println(406,"Age can`t contain any alphabetic symbols!")
		return
	}

	result, _ := personModel.GetPersonByName(data.Name, data.Surname)
	//result,_:= personModel.GetPersonByID(data.Id)
	if result.Name != "" {
		c.JSON(405, gin.H{"message": "This person is already exist!"})
		return
	}

	err := personModel.ChangePersonInfo(data)
	if err != nil {
		c.JSON(407, gin.H{"message": "Problem saving data! Try again later"})
		return
	}
}
func (p *PersonCotroller) DeletePerson(c *gin.Context) {
	id := bson.ObjectIdHex(c.Param("id"))

	err := personModel.DeletePerson(id)
	if err != nil {
		c.JSON(408, gin.H{"message": "Problem with deleting!"})
		return
	}
	c.Status(http.StatusOK)
}

func CheckFieldNoEmpty(fields ...string) bool {
	var boolTemp bool
	for _, v := range fields {
		if v == "" {
			boolTemp = false
		} else {
			boolTemp = true
		}
	}
	return boolTemp
}
func CheckHasNoDigits(fields ...string) (matched bool, index int) {
	pattern := "[[:digit:]]"
	//pattern := "[\d]"
	for k, v := range fields {
		matched, _ := regexp.MatchString(pattern, v)
		if matched {
			index = k
			return matched, index
		}
	}
	return matched, index
}
func CheckAgeHasNoAlphabetic(age int64) bool {
	pattern := "[[:alpha:]]"
	//pattern := "[\D]" // or [\^d]
	matched, _ := regexp.MatchString(pattern, strconv.FormatInt(age, 10))
	return matched
}

/*
package main

import (
	"fmt"
	log "github.com/vigo5190/goimports-example/b"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"time"
)

func Dial(url string)(*mgo.Session, error)



type uId uint8

func (receiver *person) GetName() string {
	fmt.Println("\nPerson name: ")
	return receiver.Name
	//return "\nPerson name: " + receiver.Name
}
func (receiver *person) GetSurname() string {
	fmt.Println("\nPerson surname: ")
	return receiver.Surname
}
func (receiver *person) GetNationality() string {
	fmt.Println("\nPerson nationality: ")
	return receiver.Nationality
}
func (receiver *person) GetSex() string {
	fmt.Println("\nPerson sex: ")
	return receiver.Sex
}
func (receiver *person) GetAge() uint8 {
	fmt.Println("\nPerson age: ")
	return receiver.Age
}
func (receiver *person) GetHeight() uint8 {
	fmt.Println("\nPerson height: ")
	return receiver.Height
}
func (receiver *person) GetWeight() uint8 {
	fmt.Println("\nPerson weight: ")
	return receiver.Weight
}
func (receiver *employee) GetPost() string {
	fmt.Println("Person post: ")
	return receiver.Post
}


var testTeacher *teacher = &teacher{
	Specialization: "Math",
	FormMaster:     false,
	employee: &employee{
		Salary: 9200.1,
		Post:   "Headmaster",
		person: &person{
			Name:        "testTeacherName",
			Surname:     "testTeacherSurname",
			Nationality: "German",
			Sex:         "Female",
			Age:         23,
			Weight:      65,
			Height:      173,
		},
	},
}
var testDriver *driver = &driver{
	Category:   "B",
	Experience: 5,
	employee: &employee{
		Salary: 9000.9,
		Post:   "Trucker",
		person: &person{
			Name:        "testDriverName",
			Surname:     "testDriverSurame",
			Nationality: "Brazilian",
			Sex:         "Male",
			Age:         30,
			Weight:      95,
			Height:      180,
		},
	},
}
var testHairdress *hairdresser = &hairdresser{
	Certificate: true,
	employee: &employee{
		Salary: 14232.1,
		Post:   "Hairdresser",
		person: &person{
			Name:        "testHairdressName",
			Surname:     "testHairdressSurname",
			Nationality: "Ukrainian",
			Sex:         "Male",
			Age:         22,
			Weight:      75,
			Height:      175,
		},
	},
}
var testPsyhologist *psychologist = &psychologist{
	License: true,
	employee: &employee{
		Salary: 11000.1,
		Post:   "Head doctor",
		person: &person{
			Name:        "testPsyhologistName",
			Surname:     "testPsyhologistSurname",
			Nationality: "Poland",
			Sex:         "Male",
			Age:         33,
			Weight:      84,
			Height:      179,
		},
	},
}
var testElectrician *electrician = &electrician{
	Category: 4,
	employee: &employee{
		Salary: 11000.35,
		Post:   "Senior electrician",
		person: &person{
			Name:        "testNameElectrician",
			Surname:     "testSurnameameElectrician",
			Nationality: "Belgium",
			Sex:         "Male",
			Age:         41,
			Weight:      75,
			Height:      180,
		},
	},
}
var testLawyer *lawyer = &lawyer{
	Direction: "Criminal Law",
	employee: &employee{
		Salary: 33000.9,
		Post:   "Judge",
		person: &person{
			Name:        "testLawyerName",
			Surname:     "testLawyerSurname",
			Nationality: "Ukrainian",
			Sex:         "Female",
			Age:         26,
			Weight:      62,
			Height:      174,
		},
	},
}

_map := map[uId]interface{}{0: testTeacher.employee, 1: testDriver.employee, 2: testHairdress.employee,
3: testPsyhologist.employee, 4: testElectrician.employee, 5: testLawyer.employee}


fmt.Println("result of import package: ", log.Foo)

fmt.Println("Start 2")
fmt.Println((*person).GetName((*testLawyer).employee.person))
fmt.Println((*person).GetSurname((*testPsyhologist).employee.person))
fmt.Println((*person).GetNationality((*testElectrician).employee.person))
fmt.Println((*person).GetSex((*testDriver).employee.person))
fmt.Println((*person).GetAge((*testTeacher).employee.person))
fmt.Println((*person).GetHeight((*testHairdress).employee.person))
fmt.Println((*person).GetWeight((*testHairdress).employee.person))

fmt.Println("Start 3")
//fmt.Println(MapOfTeachers)
//fmt.Println(_map)
//Poll(_map)
//fmt.Println(map1)
Poll(_map)

fmt.Println("Start 4")
GetType(testDriver)
GetType(testElectrician)
fmt.Println("Start 5")
go randomNumber(10, 20)
randomNumber(1, 10)

fmt.Println("Start 6")
DNA_HammingDistance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT")

//fmt.Println("Start 7")
//luhnAlgorithm("9875132414827548")
}
func Poll(_map map[uId]interface{}) {
	for _, v := range _map {
		fmt.Println(v.(*employee).person.GetName())
		fmt.Println(v.(*employee).GetPost())
	}
}
func CurrentTypeToPersonType(temp interface{}) {
	switch v := temp.(type) {
	case *teacher:
		var temp1 person = person(*v);
		fmt.Printf("Your type is: %T\n", v)
	case *driver:
		fmt.Printf("Your type is: %T\n", v)
	case *hairdresser:
		fmt.Printf("Your type is: %T\n", v)
	case *psychologist:
		fmt.Printf("Your type is: %T\n", v)
	case *electrician:
		fmt.Printf("Your type is: %T\n", v)
	case *lawyer:
		fmt.Printf("Your type is: %T\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func randomNumber(min, max int) {
	for i := 0; i < 5; i++ {
		time.Sleep(400 * time.Millisecond)
		if min > 9 {
			fmt.Println("IT`S GOrutine i =", i, "->", rand.Intn(max-min)+min)
		} else {
			fmt.Println("i =", i, "->", rand.Intn(max-min)+min)
		}
	}
}
func DNA_HammingDistance(str1, str2 string) {
	var count int
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			count++
		}
	}
	fmt.Println("Hamming distance = ", count)
}

/*
func luhnAlgorithm(pPurported string){
	sum:=0
	num, _ := strconv.ParseFloat(pPurported,64)
	var digitCount int = int(math.Log10(num)+1)
	var array []int =[] int{digitCount}
	for i:=0;i<digitCount;i++{
		array[i]=int(num) % 10
		num/=10
	}
	var p int
	if digitCount%2==0{
		for i:=0;i<digitCount-1;i++{
			if i%2==0{
				p=array[i]*2
				if p>9{
					p-=9
				}
				array[i]=p
			} else {
				array[i]=array[i]
			}
			sum+=array[i]
		}
		if sum%10==0{
			fmt.Println("Number is valid")
		} else {
			fmt.Println("Number is not valid")
		}
	}
}
/*
func luhnAlgorithm(pPurported string)bool{
	var sum int
	for i:=0;i<len(pPurported)-1;i++{
		 digit := int(pPurported[i])
		 if i%2==len(pPurported)%2{
		 	digit:=digit*2
		 	if digit>9{
		 		digit:=digit-9
			}
		 }
		 sum:=sum+digit
	}
	return sum%10==0
}
*/
