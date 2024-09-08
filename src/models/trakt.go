package models

type TraktMovie struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	IDs   IDs    `json:"ids"`
}

type IDs struct {
	Trakt  int     `json:"trakt"`
	Slug   string  `json:"slug"`
	IMDB   string  `json:"imdb"`
	TMDB   int     `json:"tmdb"`
	TVDB   int     `json:"tvdb"`
	TVRage *string `json:"tvrage"`
}

type RatingItem struct {
	RatedAt string     `json:"rated_at"`
	Rating  int        `json:"rating"`
	Show    *TraktShow `json:"show,omitempty"`
	Type    string     `json:"type"`
}

type WatchedItem struct {
	Plays         int            `json:"plays"`
	LastWatchedAt string         `json:"last_watched_at"`
	LastUpdatedAt string         `json:"last_updated_at"`
	ResetAt       *string        `json:"reset_at"`
	Movie         *TraktMovie    `json:"movie,omitempty"`
	Show          *TraktShow     `json:"show,omitempty"`
	Seasons       *[]TraktSeason `json:"seasons,omitempty"`
}

type TraktShow struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	IDs   IDs    `json:"ids"`
}

type TraktSeason struct {
	Episodes []TraktEpisode `json:"episodes"`
	Number   int            `json:"number"`
}

type TraktEpisode struct {
	LastWatchedAt string `json:"last_watched_at"`
	Number        int    `json:"number"`
	Plays         int    `json:"plays"`
}
