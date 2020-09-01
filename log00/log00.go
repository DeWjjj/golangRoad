package main

import (
	"golangRoad/logger"
	"time"
)

func main() {
	log := logger.Newlog("info")
	for {
		log.Debug("Debug log.")
		log.Info("Info log.")
		log.Warning("Warning log.")
		log.Fatal("Fatal log.")
		log.Error("Error log.")
		time.Sleep(time.Second * 5)
	}
}
