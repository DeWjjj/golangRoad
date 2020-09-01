package mylogger

//向文件里面写入日志相关信息。

//FileLogger struct >> exported
type FileLogger struct {
	Level       LogLevel
	filePath    string //日子文件保存的路径
	fileName    string //日志文件保存的文件名
	errFileName string //错误日志文件保存的文件名
}
