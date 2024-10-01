package models

type Book struct {
	Id     int      `json:"key"`
	Title  string   `json:"title"`
	Author []string `json:"author_name"`
	Image  string   `json:"cover_i"`
}
