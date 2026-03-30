package main

import "fmt"

type LogLevel int

const (
	LevelTrace = 0
	LevelDebug = 1
	LevelInfo  = 2
	LevelWarn  = 3
	LevelError = 4
)

var levelNames = []string{"Trace", "Debug", "Info", "Warn", "Error"}

func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "Unknown"
	}
	return levelNames[l]
}

func printLog(level LogLevel) {
	fmt.Println("Log level:", level.String())
}

func main() {

	printLog(LevelTrace)
	printLog(LevelDebug)
	printLog(LevelInfo)
	printLog(LevelWarn)
	printLog(LevelError)

}
