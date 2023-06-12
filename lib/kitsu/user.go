package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tts2k/anitrack/lib"
)

type kitsuUser struct {
	Name           string `json:"name"`
	FavoritesCount int    `json:"favioritesCount"`
	ReviewsCount   int    `json:"reviewsCount"`
}

type kitsuUserResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Attributes kitsuUser `json:"attributes"`
	} `json:"data"`
}

func (k *Kitsu) User() (lib.User, error) {
	const EndPoint = "edge/users"

	joinedURL, err := url.JoinPath(k.baseURL, EndPoint)
	if err != nil {
		return lib.User{}, err
	}

	req, err := http.NewRequest(
		http.MethodGet,
		joinedURL,
		nil,
	)
	if err != nil {
		return lib.User{}, err
	}

	// Query
	query := req.URL.Query()
	query.Add("filter[self]", "true")
	req.URL.RawQuery = query.Encode()

	// Header
	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Content-Type", "application/vnd.api+json")
	req.Header.Add("Authorization", "Bearer "+k.accessToken)

	// Do request
	resp, err := k.client.Do(req)
	if err != nil {
		return lib.User{}, err
	}
	defer resp.Body.Close()

	bodyBuffer := bytes.NewBuffer([]byte{})
	// Copy response body to buffer
	_, err = io.Copy(bodyBuffer, resp.Body)
	if err != nil {
		return lib.User{}, err
	}

	// Parse json
	var respBody kitsuUserResponse
	err = json.Unmarshal(bodyBuffer.Bytes(), &respBody)
	if err != nil {
		return lib.User{}, err
	}

	// Check empty
	if len(respBody.Data) == 0 {
		return lib.User{}, errors.New("User not found. Try re-logging in")
	}

	// Map response to data struct
	result := lib.User{}
	result.ID = respBody.Data[0].ID
	result.ProfileLink = respBody.Data[0].Links.Self
	result.Name = respBody.Data[0].Attributes.Name
	result.FavouritesCount = respBody.Data[0].Attributes.FavoritesCount
	result.ReviewsCount = respBody.Data[0].Attributes.ReviewsCount

	return result, nil
}

func (k *Kitsu) UserAnime(page int, limit int) ([]lib.Anime, error) {
	const EndPoint = "edge/library-entries/"

	// Get user id
	user, err := k.User()
	if err != nil {
		return nil, err
	}

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

	// Query
	query := req.URL.Query()
	query.Add("filter[userId]", user.ID)
	req.URL.RawQuery = query.Encode()

	// Do request
	resp, err := k.client.Do(req)
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
