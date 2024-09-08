package models

type Show struct {
	Title  string `json:"title"`
	URL    string `json:"url,omitempty"`
	Rating *int   `json:"rating,omitempty"`
}
