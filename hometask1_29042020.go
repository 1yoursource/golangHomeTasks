package main

import (
	"fmt"

	log "github.com/vigo5190/goimports-example/b"
)

type person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

type employee struct {
	Salary float32 `json:"salary"`
	person
}

type teacher struct {
	Specialization string `json:"specialization"`
	FormMaster     bool   `json:"form_master"`
	employee
}

type driver struct {
	Category   string `json:"category"`
	Experience uint8  `json:"experience"`
	employee
}

type hairdresser struct {
	Certificate bool `json:"certificate"`
	employee
}

type psychologist struct {
	License        bool   `json:"license"`
	AcademicDegree string `json:"academic_degree"`
	employee
}
type electrician struct {
	Category uint `json:"category"`
	employee
}
type lawyer struct {
	Direction string `json:"direction"`
	License   bool   `json:"license"`
	employee
}

func main() {
	var testTeacher = teacher{
		Specialization: "math",
		FormMaster:     false,
		employee: employee{
			Salary: 9200.1,
			person: person{
				Name:    "testTeacherName",
				Surname: "testTeacherSurname",
				Age:     23,
			},
		},
	}
	var testDriver = driver{
		Category:   "B",
		Experience: 5,
		employee: employee{
			Salary: 9000.9,
			person: person{
				Name: "testDriverName",
				Age:  30,
			},
		},
	}
	var testHairdress = hairdresser{
		Certificate: true,
		employee: employee{
			Salary: 14232.1,
			person: person{
				Name: "testHairdressName",
			},
		},
	}
	var testPsyhologist = psychologist{
		License: true,
		employee: employee{
			Salary: 11000.1,
			person: person{
				Age:     33,
				Name:    "testPsyhName",
				Surname: "testPsyhSurname",
			},
		},
	}
	var testElectrician = electrician{
		Category: 4,
		employee: employee{
			Salary: 11000.35,
			person: person{
				Name: "testNameEl",
			},
		},
	}
	var testLawyer = lawyer{
		Direction: "Criminal Law",
		employee: employee{
			Salary: 33000.9,
			person: person{
				Age:  26,
				Name: "testLawyerName",
			},
		},
	}
	fmt.Println("Start")

	fmt.Println("teacher name: ", testTeacher.Name)
	fmt.Println("Driver Age: ", testDriver.Age)
	fmt.Println("electrician salary: ", testElectrician.Salary)
	fmt.Println("hairdress certificate: ", testHairdress.Certificate)
	fmt.Println("Psyhologist surname: ", testPsyhologist.Surname)
	fmt.Println("lawyer direction: ", testLawyer.Direction)

	fmt.Println("result of import package", log.Foo)
}
