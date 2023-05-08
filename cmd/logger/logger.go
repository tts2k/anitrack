package cmd

import (
	"io"
	"log"
	"os"
	"runtime"
)

type level int
type levelConfig struct {
	color  string
	output io.Writer
}

var logger *log.Logger

const (
	none level = iota
	info
	debug
	warn
	err
)

var levels = []levelConfig{
	{"\033[0m", os.Stdout},  // No color
	{"\033[36m", os.Stdout}, // Cyan
	{"\033[32m", os.Stdout}, // Green
	{"\033[33m", os.Stderr}, // Yellow
	{"\033[31m", os.Stderr}, // Red
}


func init() {
	// Reset color on windows
	if runtime.GOOS == "windows" {
		for i := range levels {
			levels[i].color = ""
		}
	}

	// Initialize logger
	logger = log.New(os.Stdout, "", 0)
}

func printColor(m level, a any) {
	color := levels[m].color
	noColor := levels[none].color
	logger.Print(string(color), a, string(noColor))
}

// func printfColor(m level, f string, a ...any) {
// 	color := levels[m].color
// 	noColor := levels[none].color
// 	format := "%s" + f + "%s"
// 	a = append(a, noColor)
// 	a = append([]any{color}, a)
//
// 	logger.Printf(format, a...)
// }

func Info(s string) {
	logger.SetOutput(levels[info].output)
	printColor(info, s)
}

func Warn(s string) {
	logger.SetOutput(levels[warn].output)
	printColor(warn, s)
}

func Error(a any) {
	logger.SetOutput(levels[err].output)
	printColor(err, a)
}
