package lib

import (
	"net/http"
	"time"
)

type Kitsu struct {
	client       *http.Client
	baseURL      string
	accessToken  string
	refreshToken string
}

type kitsuErrRes struct {
	Name        string `json:"error"`
	Description string `json:"error_description"`
}

func New(accessToken string, refreshToken string) *Kitsu {
	client := &http.Client{
		Transport: http.DefaultClient.Transport,
		Timeout:   time.Duration(10) * time.Second,
	}

	return &Kitsu{
		client:       client,
		baseURL:      "https://kitsu.io/api",
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}
