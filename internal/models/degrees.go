package models

type Degree struct {
	Degree_ID       int    `json:"degree_id"`
	Program         string `json:"program"`
	Title           string `json:"title"`
	Shorthand_Title string `json:"shorthand_title"`
	Rows            int    `json:"rows"`
	Columns         int    `json:"columns"`
}
