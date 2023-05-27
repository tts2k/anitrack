package lib

type Anime struct{
	Title string
	Rating string `json:"averageRating"`
	Status string `json:"status"`
	EpisodeCount int `json:"episodeCount"`
	SubType string `json:"subType"`
}

type Site interface {
	Login() (string, string, error)
	Trending() ([]Anime, error)
	UserAnime(page int, limit int) ([]Anime, error)
}
