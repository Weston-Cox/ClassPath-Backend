package models

type Degree_Course struct {
	Degree_ID        int      `json:"degree_id"`
	Course_ID        int      `json:"course_id"`
	Semester_Offered []string `json:"semester_offered"`
}