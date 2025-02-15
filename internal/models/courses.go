package models

type Course struct {
	Course_ID   int    `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}