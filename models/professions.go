package models

type Teacher struct {
	Specialization string `json:"specialization"`
	FormMaster     bool   `json:"form_master"`
	employee       *Employee
}
type Driver struct {
	Category   string `json:"category"`
	Experience uint8  `json:"experience"`
	employee   *Employee
}
type Hairdresser struct {
	Certificate bool `json:"certificate"`
	employee    *Employee
}
type Psychologist struct {
	License        bool   `json:"license"`
	AcademicDegree string `json:"academic_degree"`
	employee       *Employee
}
type Electrician struct {
	Category uint `json:"category"`
	employee *Employee
}
type Lawyer struct {
	Direction string `json:"direction"`
	License   bool   `json:"license"`
	employee  *Employee
}
