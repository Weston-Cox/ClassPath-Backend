package models

type Course_Requisite struct {
	Source      int  `json:"source"`
	Target      int  `json:"target"`
	Corequisite bool `json:"corequisite"`
}