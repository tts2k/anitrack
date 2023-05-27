package lib

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type authData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (k *Kitsu) Login() (string, string, error) {
	const EndPoint = "oauth/token"
	var email, password string
	reader := bufio.NewReader(os.Stdin)

	// Get username
	fmt.Println("Please enter your kitsu credential")
	fmt.Print("Email: ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// Get password
	fmt.Print("Password: ")
	b, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}
	password = string(b)
	fmt.Println()

	body := make(map[string]string)
	body["grant_type"] = "password"
	body["username"] = email
	body["password"] = password

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return "", "", err
	}

	joinedURL, err := url.JoinPath(k.baseURL, EndPoint)
	if err != nil {
		return "", "", err
	}

	resp, err := http.Post(
		joinedURL,
		"application/json",
		bytes.NewBuffer(bodyJSON),
	)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", nil
	}

	if resp.StatusCode != 200 {
		var errResp errRes
		err = json.Unmarshal(bodyBytes, &errResp)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Request failed with unknown error")
			return "", "", err
		}

		fmt.Println("Error:", errResp.Name)
		fmt.Println(errResp.Description)
		return "", "", errors.New("Login failed")
	}

	var respBody authData
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		fmt.Println("Malformed response body")
		return "", "", errors.New("Login failed")
	}

	return respBody.AccessToken, respBody.RefreshToken, nil
}
