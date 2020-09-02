package main

import (
	"golangRoad/mylogger"
)

//testing mylogger.
func main() {
	log := mylogger.NewFileLogger("Info", "./", "dew.log", 40*1024*1024)
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
