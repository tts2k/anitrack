package logger

import (
	"log"
	"os"
	"runtime"
)

var logger *log.Logger
var verbose bool = false

type loggerMode int

const (
	none loggerMode = iota
	info
	debug
	warn
	error
)

var colors = []string{
	"",
	"\033[36m", // Cyan
	"\033[32m", // Green
	"\033[33m", // Yellow
	"\033[32m", // Red
}

func InitLogger(enableVerbose bool) {
	logger = log.New(os.Stdout, "", 0)
	verbose = enableVerbose

	// Reset color on windows
	if runtime.GOOS == "windows" {
		for i := range colors {
			colors[i] = ""
		}
	}
}

func setLoggerOptions(mode loggerMode) {
	logger.SetFlags(0)
	switch mode {
	case none, debug, info:
		logger.SetOutput(os.Stdout)
	case warn, error:
		logger.SetOutput(os.Stderr)
	default:
		panic("No such logger mode")
	}
	// Set color
	logger.Print(string(colors[mode]))
}

func Println(a ...any) {
	setLoggerOptions(none)
	log.Println(a...)
}

func Print(a ...any) {
	setLoggerOptions(none)
	log.Print(a...)
}

func Printf(format string, a ...any) {
	setLoggerOptions(none)
	log.Printf(format, a...)
}

func Infoln(a ...any) {
	setLoggerOptions(info)
	log.Println(a...)
}

func Info(a ...any) {
	setLoggerOptions(info)
	log.Print(a...)
}

func Infof(format string, a ...any) {
	setLoggerOptions(info)
	logger.Printf(format, a...)
}

func Debugln(a ...any) {
	if verbose {
		setLoggerOptions(debug)
		logger.Println(a...)
	}
}

func Debugf(format string, a ...any) {
	if verbose {
		setLoggerOptions(debug)
		logger.Printf(format, a...)
	}
}

func WarnLn(a ...any) {
	setLoggerOptions(warn)
	logger.Println(a...)
}

func Warnf(format string, a ...any) {
	setLoggerOptions(warn)
	logger.Printf(format, a...)
}

func Errorln(a ...any) {
	setLoggerOptions(error)
	logger.Println(a...)
}

func Errorf(format string, a ...any) {
	setLoggerOptions(error)
	logger.Printf(format, a...)
}
