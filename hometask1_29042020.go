package main

import (
	"fmt"
	log "github.com/vigo5190/goimports-example/b"
	"math/rand"
	"time"
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
	Post   string  `json:"post"`
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

func main() {
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

	/*
		_map := map[uId]*employee{0: testTeacher.employee, 1: testDriver.employee, 2: testHairdress.employee,
			3: testPsyhologist.employee, 4: testElectrician.employee, 5: testLawyer.employee}
	*/
	_map := map[uId]interface{}{0: testTeacher.employee, 1: testDriver.employee, 2: testHairdress.employee,
		3: testPsyhologist.employee, 4: testElectrician.employee, 5: testLawyer.employee}
	/*
		fmt.Println("Start")

		fmt.Println("teacher name: ", (*testTeacher).employee.person.Name)
		fmt.Println("Driver Age: ", (*testDriver).employee.person.Age)
		fmt.Println("electrician salary: ", (*testElectrician).employee.Salary)
		fmt.Println("hairdress certificate: ", (*testHairdress).Certificate)
		fmt.Println("Psyhologist surname: ", (*testPsyhologist).employee.person.Surname)
		fmt.Println("lawyer direction: ", (*testLawyer).Direction)
	*/

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
	/*panic
	if (*person).GetAge((*testLawyer).employee.person) != 26 {
		f_panic()
	}
	*/
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

/*
func Poll(_map map[uId]*employee) {
	for _, v := range _map {
		fmt.Println(v.person.GetName())
		fmt.Println(v.GetPost())
	}
}
*/
func Poll(_map map[uId]interface{}) {
	for _, v := range _map {
		fmt.Println(v.(*employee).person.GetName())
		fmt.Println(v.(*employee).GetPost())
	}
}
func GetType(temp interface{}) {
	switch v := temp.(type) {
	case int:
		fmt.Printf("Your type is: %T\n", v)
	case string:
		fmt.Printf("Your type is: %T\n", v)
	case float64:
		fmt.Printf("Your type is: %T\n", v)
	case *teacher:
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
/*
func f_panic() {
	panic("It`s panic!")
}
*/
