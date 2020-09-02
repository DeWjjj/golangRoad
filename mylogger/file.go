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
	if err != nil { //如果拿到了err值，则panic停止运行，panic中填入err的信息
		panic(err)
	}
	fl := &FileLogger{
		Level:       LogLevel, //将上面通过字符串解析到的等级传入
		filePath:    fp,       //以下三条都是默认传入的值
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {     //如果拿到了err值，则panic停止运行，panic中填入err的信息
		panic(err)
	}
	return fl //函数结束返回一个结构体
}

func (f *FileLogger) initFile() error { //传入一个结构体作为方法
	fullFileName := path.Join(f.filePath, f.fileName)                                    //拼接结构体内的地址和名字
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //以拼接的地址和名字作为路径和文件名打开文件
	if err != nil {                                                                      //默认报错信息填写
		fmt.Printf("Open log file failed,err: %v.\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //以拼接的地址和名字加上.err作为路径和文件名打开文件
	if err != nil {                                                                                //默认报错信息填写
		fmt.Printf("Open log file failed,err: %v.\n", err)
		return err
	}

	f.fileObj = fileObj //将创建的文件值赋值给结构体中的字段
	f.errFileObj = errFileObj
	return nil
}

//以*FileLogger为方法传入一个LogLevel等级，如果传入等级大于方法中的等级返回一个true
func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level
}

//确认大小小于设置范围
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat() //将传入文件的信息赋值给fileInfo
	if err != nil {              //默认报错信息填写
		fmt.Printf("Get file into failed,err: %v.\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize //如果文件的大小大于设置的值返回一个true
}

//分割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102150405000") //格式化日期
	fileInfo, err := file.Stat()                     //拿到传入文件的信息
	if err != nil {                                  //默认报错信息填写
		fmt.Printf("Get file info failed,err: %v.\n", err)
		return nil, err //这里写两个的原因是因为函数设置了返回值
	}
	logName := path.Join(f.filePath, fileInfo.Name())                               //拼接拿到当前日志文件路径和日志文件名
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)                          //拼接一个备份日志文件的名字
	file.Close()                                                                    //将文件关闭
	os.Rename(logName, newLogName)                                                  //将旧的名字改成新的名字
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755) //打开的一个新的文件，名字以83行的代码为基础
	if err != nil {                                                                 //默认报错信息填写
		fmt.Printf("Open new log file failed,err: %v.\n", err)
		return nil, err
	}
	return fileObj, nil //若是没有报错，则返回打开的fileObj对象，和空的error值
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...) //组合format和a字段。
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s.\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			//如果要记录的日志大于等于ERROR级别，我还要在err日志文件中再记录一遍。
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s.\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)

		}
	}
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

//Close func >> exported
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
