package main

import (
	"fmt"
	"os"
	"time"
)

//切割文件
func fileSplit(fileObj *os.File, now string) {
	fileObj.Close()
	os.Rename(fileObj.Name(), now+".txt")
	time.Sleep(time.Second)
}

//判断大小
func checkSize(fileObj *os.File, maxFileSize int64) bool {
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("Get file info failed,err: %v.\n", err)
	}
	if fileInfo.Size() >= maxFileSize {
		return true
	}
	return false
}

//创建文件
func openFile() *os.File {
	fileObj, err := os.OpenFile("./test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open file failed,err: %v.\n", err)
	}
	return fileObj
}

func main() {
	var maxFileSize int64 = 100
	fileObj := openFile()
cycle1:
	for {
		now := time.Now().Format("20060102150405000")
		fmt.Fprintln(fileObj, "1233211234567")
		if checkSize(fileObj, maxFileSize) {
			fileSplit(fileObj, now)
			fileObj = openFile()
			goto cycle1 //这里尝试使用了一下goto，直接跳转到label。
		}
	}
}
