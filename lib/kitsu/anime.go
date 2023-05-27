package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tts2k/anitrack/lib"
	"io"
	"net/http"
	"net/url"
)

type Kitsu struct {
	baseURL      string
	accessToken  string
	refreshToken string
}

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

type kitsuRepsonse struct {
	Data []struct {
		Attributes kitsuAnime `json:"attributes"`
	} `json:"data"`
}

type errRes struct {
	Name        string `json:"error"`
	Description string `json:"error_description"`
}

func New(accessToken string, refreshToken string) *Kitsu {
	return &Kitsu{
		baseURL:      "https://kitsu.io/api",
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

func (k *Kitsu) Trending() ([]lib.Anime, error) {
	const EndPoint = "edge/anime"

	joinedURL, err := url.JoinPath(k.baseURL, EndPoint)
	if err != nil {
		return nil, err
	}

	// Do request
	resp, err := http.Get(
		joinedURL,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBuffer := bytes.NewBuffer([]byte{})

	// Copy response body to buffer
	_, err = io.Copy(bodyBuffer, resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse json
	var respBody kitsuRepsonse
	err = json.Unmarshal(bodyBuffer.Bytes(), &respBody)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("malformed json body")
	}

	// Map response to data struct
	var result []lib.Anime
	for _, d := range respBody.Data {
		anime := lib.Anime{
			Title:        d.Attributes.Title.Eng,
			Rating:       d.Attributes.Rating,
			Status:       d.Attributes.Status,
			EpisodeCount: d.Attributes.EpisodeCount,
			SubType:      d.Attributes.SubType,
		}

		result = append(result, anime)
	}

	return result, nil
}

func (k *Kitsu) UserAnime(page int, limit int) ([]lib.Anime, error) {
	const EndPoint = "edge/library-entries"

	joinedURL, err := url.JoinPath(k.baseURL, EndPoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodGet,
		joinedURL,
		nil,
	)
	req.Header.Add("Authorization", "Bearer "+k.accessToken)
	if err != nil {
		return nil, err
	}

	// Do request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBuffer := bytes.NewBuffer([]byte{})
	// Copy response body to buffer
	_, err = io.Copy(bodyBuffer, resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(bodyBuffer.String())

	return nil, nil
}
