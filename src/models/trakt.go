package models

type TraktMovie struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	IDs   IDs    `json:"ids"`
}

type IDs struct {
	Trakt int    `json:"trakt"`
	Slug  string `json:"slug"`
	IMDB  string `json:"imdb"`
	TMDB  int    `json:"tmdb"`
}

type WatchedItem struct {
	WatchedAt string      `json:"watched_at"`
	Type      string      `json:"type"`
	Movie     *TraktMovie `json:"movie,omitempty"`
	Show      *TraktShow  `json:"show,omitempty"`
}

type TraktShow struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	IDs   IDs    `json:"ids"`
}
