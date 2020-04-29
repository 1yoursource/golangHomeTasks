package main

import (
	"fmt"

	log "github.com/vigo5190/goimports-example/b"
)

type person struct {
	name    string
	surname string
	age     int
}

type employee struct {
	salary float32
	person
}

type teacher struct {
	specialization string
	formMaster     bool
	employee
}

type driver struct {
	category   string
	experience uint8
	employee
}

type hairdresser struct {
	certificate bool
	employee
}

type psychologist struct {
	license        bool
	academicDegree string
	employee
}
type electrician struct {
	category uint
	employee
}
type lawyer struct {
	direction string
	license   bool
	employee
}

func main() {
	var testTeacher = teacher{
		specialization: "math",
		formMaster:     false,
		employee: employee{
			salary: 9200.1,
			person: person{
				name:    "testTeacherName",
				surname: "testTeacherSurname",
				age:     23,
			},
		},
	}
	var testDriver = driver{
		category:   "B",
		experience: 5,
		employee: employee{
			salary: 9000.9,
			person: person{
				name: "testDriverName",
				age:  30,
			},
		},
	}
	var testHairdress = hairdresser{
		certificate: true,
		employee: employee{
			salary: 14232.1,
			person: person{
				name: "testHairdressName",
			},
		},
	}
	var testPsyhologist = psychologist{
		license: true,
		employee: employee{
			salary: 11000.1,
			person: person{
				age:     33,
				name:    "testPsyhName",
				surname: "testPsyhSurname",
			},
		},
	}
	var testElectrician = electrician{
		category: 4,
		employee: employee{
			salary: 11000.35,
			person: person{
				name: "testNameEl",
			},
		},
	}
	var testLawyer = lawyer{
		direction: "Criminal Law",
		employee: employee{
			salary: 33000.9,
			person: person{
				age:  26,
				name: "testLawyerName",
			},
		},
	}

	fmt.Println("teacher name: ", testTeacher.name)
	fmt.Println("Driver Age: ", testDriver.age)
	fmt.Println("electrician salary: ", testElectrician.salary)
	fmt.Println("hairdress certificate: ", testHairdress.certificate)
	fmt.Println("Psyhologist surname: ", testPsyhologist.surname)
	fmt.Println("lawyer direction ", testLawyer.direction)

	fmt.Println("result of import package", log.Foo)
}
