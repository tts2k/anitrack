package lib

type Anime struct {
	Title        string
	Rating       string `json:"averageRating"`
	Status       string `json:"status"`
	EpisodeCount int    `json:"episodeCount"`
	SubType      string `json:"subType"`
}

type User struct {
	ID              string
	ProfileLink     string
	Name            string
	FavouritesCount int
	ReviewsCount    int
}

type Site interface {
	Login() (string, string, error)
	Trending() ([]Anime, error)
	UserAnime(page uint, limit uint) ([]Anime, error)
	User() (User, error)
}
