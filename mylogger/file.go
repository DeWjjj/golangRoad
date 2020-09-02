package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//向文件里面写入日志相关信息。

//FileLogger struct >> exported
type FileLogger struct {
	Level       LogLevel
	filePath    string //日子文件保存的路径
	fileName    string //日志文件保存的文件名
	errFileName string //错误日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

//NewFileLogger func >> exported
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("Open log file failed,err: %v.\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("Open log file failed,err: %v.\n", err)
		return err
	}

	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
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

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...) //组合format和a字段。
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s.\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			//如果要记录的日志大于等于ERROR级别，我还要在err日志文件中再记录一遍。
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s.\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)

		}
	}
}

func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level
}

//Debug func >> exported
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

//Trace func >> exported
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

//Info func >> exported
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

//Warning func >> exported
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

//Error func >> exported
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

//Fatal func >> exported
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
