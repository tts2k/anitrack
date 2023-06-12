package lib

import (
	"encoding/json"
	"github.com/tts2k/anitrack/lib"
	"net/url"
)

type kitsuAnime struct {
	Title struct {
		Eng   string `json:"en"`
		Roman string `json:"en_jp"`
		Jap   string `json:"ja_jp"`
	} `json:"titles"`
	Rating       string `json:"averageRating"`
	Status       string `json:"status"`
	EpisodeCount int    `json:"episodeCount"`
	SubType      string `json:"subType"`
}

type kitsuAnimeRepsonse struct {
	Data []struct {
		Attributes kitsuAnime `json:"attributes"`
	} `json:"data"`
}

func (k *Kitsu) Trending() ([]lib.Anime, error) {
	const EndPoint = "edge/anime"

	joinedURL, err := url.JoinPath(k.baseURL, EndPoint)
	if err != nil {
		return nil, err
	}

	// Do request
	resp, err := k.client.Get(
		joinedURL,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse json
	var respBody kitsuAnimeRepsonse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, err
	}

	// Map response to data struct
	var result []lib.Anime
	for _, d := range respBody.Data {
		anime := lib.Anime{
			Title:        d.Attributes.Title.Eng, // TODO: support mutliple titile types
			Rating:       d.Attributes.Rating,
			Status:       d.Attributes.Status,
			EpisodeCount: d.Attributes.EpisodeCount,
			SubType:      d.Attributes.SubType,
		}

		result = append(result, anime)
	}

	return result, nil
}
