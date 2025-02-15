package models

type Degree_Elective struct {
	Degree_ID                int `json:"degree_id"`
	General_Electives        int `json:"general_electives"`
	Specific_Electives       int `json:"specific_electives"`
	History_Elective         int `json:"history_elective"`
	Fine_Arts_Elective       int `json:"fine_arts_elective"`
	Social_Sciences_Elective int `json:"social_sciences_elective"`
	Science_Elective         int `json:"science_elective"`
}