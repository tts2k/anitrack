package lib

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tts2k/anitrack/lib"
)

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
