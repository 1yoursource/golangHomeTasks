package main

import (
	"fmt"
	log "github.com/vigo5190/goimports-example/b"
)

type person struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Nationality string `json:"nationality"`
	Sex         string `json:"sex"`
	Age         uint8  `json:"age"`
	Height      uint8  `json:"height"`
	Weight      uint8  `json:"weight"`
}
type employee struct {
	Salary float32 `json:"salary"`
	person *person
}
type teacher struct {
	Specialization string `json:"specialization"`
	FormMaster     bool   `json:"form_master"`
	employee       *employee
}
type driver struct {
	Category   string `json:"category"`
	Experience uint8  `json:"experience"`
	employee   *employee
}
type hairdresser struct {
	Certificate bool `json:"certificate"`
	employee    *employee
}
type psychologist struct {
	License        bool   `json:"license"`
	AcademicDegree string `json:"academic_degree"`
	employee       *employee
}
type electrician struct {
	Category uint `json:"category"`
	employee *employee
}
type lawyer struct {
	Direction string `json:"direction"`
	License   bool   `json:"license"`
	employee  *employee
}

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

func main() {
	var testTeacher *teacher = &teacher{
		Specialization: "Math",
		FormMaster:     false,
		employee: &employee{
			Salary: 9200.1,
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

	fmt.Println("Start")

	fmt.Println("teacher name: ", (*testTeacher).employee.person.Name)
	fmt.Println("Driver Age: ", (*testDriver).employee.person.Age)
	fmt.Println("electrician salary: ", (*testElectrician).employee.Salary)
	fmt.Println("hairdress certificate: ", (*testHairdress).Certificate)
	fmt.Println("Psyhologist surname: ", (*testPsyhologist).employee.person.Surname)
	fmt.Println("lawyer direction: ", (*testLawyer).Direction)

	fmt.Println("result of import package: ", log.Foo)

	fmt.Println("Start 2")
	fmt.Println((*person).GetName((*testLawyer).employee.person))
	fmt.Println((*person).GetSurname((*testPsyhologist).employee.person))
	fmt.Println((*person).GetNationality((*testElectrician).employee.person))
	fmt.Println((*person).GetSex((*testDriver).employee.person))
	fmt.Println((*person).GetAge((*testTeacher).employee.person))
	fmt.Println((*person).GetHeight((*testHairdress).employee.person))
	fmt.Println((*person).GetWeight((*testHairdress).employee.person))

	/*panic*/
	if (*person).GetAge((*testLawyer).employee.person) == 26 {
		f_panic()
	}
}
func f_panic() {
	panic("It`s panic!")
}
