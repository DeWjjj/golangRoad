package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//LogLevel type >> exported
type LogLevel uint16

const (
	//UNKOWN const >> exported
	UNKOWN LogLevel = iota
	//DEBUG const >> exported
	DEBUG
	//TRACE const >> exported
	TRACE
	//INFO const >> exported
	INFO
	//WARNING const >> exported
	WARNING
	//ERROR const >> exported
	ERROR
	//FATAL const >> exported
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKOWN, err
	}

}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed.")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1] //得到的funcName是main.mian函数，因为是main所以，我们只需要一个就够了。
	//所以将其用strings里面的Split()以 "." 为基础分割，取得切片其中的第一个。
	return
}
