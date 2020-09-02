package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
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
	s = strings.ToLower(s) //将传入的字符串转换成消协
	switch s {             //将小写的字符串传入switch，返回一个变量（上面const中以iota为方法赋予的uint16）
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
	default: //考虑到可能会输入匹配不上的的字符串
		err := errors.New("无效的日志级别") //新建一种err，内容为“无效的日志级别”
		return UNKOWN, err
	}

}

//解析获得的变量（上面const中以iota为方法赋予的uint16）,返回一个字符串
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

//console.go中用来打印的代码。
func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...) //组合format和a字段。
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s.\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}
