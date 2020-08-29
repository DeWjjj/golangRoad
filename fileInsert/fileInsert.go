package main

import (
	"fmt"
	"io"
	"os"
)

func fileInsert(fileName string, sentence string, position int) {
	//目标文件
	tmpFile1, err := os.OpenFile(fileName, os.O_RDWR, 0755)
	if err != nil {
		fmt.Printf("err:%v.\n", err)
	}
	defer tmpFile1.Close()
	//临时文件
	tmpFile0, err := os.OpenFile("./tmpFile00.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("err:%v.\n", err)
	}
	defer tmpFile0.Close()
	//前部复制内容
	tmp := make([]byte, position)
	n, err := tmpFile1.Read(tmp[:]) //12345
	if err != nil {
		fmt.Printf("err:%v.\n", err)
	}
	tmpFile0.Write(tmp[:n])
	//中间插入内容
	tmpSentence := []byte(sentence) //123
	tmpFile0.Write(tmpSentence)
	//后部复制内容
	//获取全部的内容，然后取出第一个n。
	tmp0 := make([]byte, 1024) //创建一个切片tmp0，其空间为1024。
	for {
		n, err = tmpFile1.Read(tmp0[:]) //67 go语言中很多函数，其所指如果被读取过了，那么实际上其并不会再被读出，像是被消耗掉了。
		if err == io.EOF {
			tmpFile0.Write(tmp[:n])
			break
		}
		if err != nil {
			fmt.Printf("err:%v.\n", err)
		}
		tmpFile0.Write(tmp0[:n])
	}
	os.Rename("./tmpFile00.txt", fileName)
}

func main() {
	fileInsert("./test.txt", "123", 5)

}
