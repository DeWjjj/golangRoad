package mylogger

import (
	"fmt"
	"time"
)

//ConsoleLogger struct >> exported
type ConsoleLogger struct {
	Level LogLevel
}

// //Newlog func >> exported
// func Newlog(levelStr string) ConsoleLogger {
// 	level, err := parseLogLevel(levelStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return ConsoleLogger{
// 		Level: level,
// 	}
// }

func (c ConsoleLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...) //组合format和a字段。
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s.\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

//Debug func >> exported
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

//Trace func >> exported
func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log(TRACE, format, a...)
}

//Info func >> exported
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

//Warning func >> exported
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

//Error func >> exported
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

//Fatal func >> exported
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
