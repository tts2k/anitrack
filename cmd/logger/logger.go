package logger

import (
	"fmt"

	"github.com/tts2k/anitrack/cmd/config"
)

func Println(a ...any) {
	fmt.Println(a...)
}

func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func Debug(a ...any) {
	conf := config.GetConfig()
}
