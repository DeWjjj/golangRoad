package mylogger

import (
	"fmt"
	"time"
)

//ConsoleLogger struct >> exported
type ConsoleLogger struct {
	Level LogLevel
}

//Newlog func >> exported
func Newlog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (l ConsoleLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= l.Level
}

func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...) //组合format和a字段。
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s.\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}

//Debug func >> exported
func (l ConsoleLogger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		log(DEBUG, format, a...)
	}
}

//Trace func >> exported
func (l ConsoleLogger) Trace(format string, a ...interface{}) {
	if l.enable(TRACE) {
		log(TRACE, format, a...)

	}
}

//Info func >> exported
func (l ConsoleLogger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		log(INFO, format, a...)

	}
}

//Warning func >> exported
func (l ConsoleLogger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING, format, a...)

	}
}

//Error func >> exported
func (l ConsoleLogger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR, format, a...)

	}
}

//Fatal func >> exported
func (l ConsoleLogger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL, format, a...)

	}
}
