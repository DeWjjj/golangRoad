package main

import (
	"golangRoad/mylogger"
)

var log mylogger.Logger //声明一个全局的接口变量

//testing mylogger.
func main() {
	log = mylogger.NewConsoleLog("Info")                                //终端日志实例
	log = mylogger.NewFileLogger("Info", "./", "dew.log", 40*1024*1024) //文件日志实例
	for {
		log.Debug("Debug log")
		log.Trace("Trace log")
		log.Info("Info log")
		log.Warning("Warning log")
		id := 10010
		name := "dew"
		log.Error("Error log.id: %d,name: %s", id, name)
		log.Fatal("Fatal log")
		// time.Sleep(time.Second * 2)
	}
}
