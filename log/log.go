package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//log
func main() {
	logFile, err := os.OpenFile("./test.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("err: %v.\n", err)
	}
	log.SetOutput(logFile)
	for i := 0; i <= 5; i++ {
		log.Println("This is a log test.")
		time.Sleep(5 * time.Second)
	}
	defer logFile.Close()
}
