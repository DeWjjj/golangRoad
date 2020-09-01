package logger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//levelType exported
type levelType uint16

const (
	UNKOWN levelType = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLevel(str string) (levelType, error) {
	str = strings.ToLower(str)
	switch str {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "FATAL":
		return FATAL, nil
	case "error":
		return ERROR, nil
	default:
		err := errors.New("err log level")
		return UNKOWN, err
	}

}

//Logger exported
type Logger struct {
	Level levelType
}

//Newlog exported
func Newlog(lv string) Logger {
	level, err := parseLevel(lv)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(lv levelType) bool {
	return lv >= l.Level
}

//Debug exported
func (l Logger) Debug(str string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s][Debug] %s.\n", now.Format("2006-01-02 15-04-05"), str)
	}
}

//Info exported
func (l Logger) Info(str string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s][Info] %s.\n", now.Format("2006-01-02 15-04-05"), str)
	}
}

//Warning exported
func (l Logger) Warning(str string) {
	if l.enable(WARNING) {
		now := time.Now()
		fmt.Printf("[%s][Warning] %s.\n", now.Format("2006-01-02 15-04-05"), str)
	}
}

//Fatal exported
func (l Logger) Fatal(str string) {
	if l.enable(FATAL) {
		now := time.Now()
		fmt.Printf("[%s][Fatal] %s.\n", now.Format("2006-01-02 15-04-05"), str)
	}
}

//Error exported
func (l Logger) Error(str string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s][Error] %s.\n", now.Format("2006-01-02 15-04-05"), str)
	}
}
