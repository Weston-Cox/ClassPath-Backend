package models

type Node_Positions struct {
	Course_ID int `json:"course_id"`
	Row       int `json:"row"`
	Column    int `json:"column"`
	Degree_ID int `json:"degree_id"`
}