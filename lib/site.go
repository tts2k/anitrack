package lib

import (
	"fmt"
	"strings"

	kitsu "github.com/tts2k/anitrack/lib/kitsu"
)

type Site interface {
	Login() (string, string, error)
}

func New(site string) Site {
	switch strings.ToLower(site) {
	case "kitsu":
		return kitsu.New()
	case "mal":
	case "anilist":
	default:
		fmt.Println("Not implemented", site)
	}
	return nil
}
