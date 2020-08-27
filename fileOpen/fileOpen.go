package main

import (
	"fmt"
	"os"
)

func main() {
	val, err := os.Open("./test.txt") //val >> *File err >> error
	if err != nil {
		fmt.Printf("Open file failed.err:%v\n", err)
		return
	}
	fmt.Println(val)  //&{0xc000056180}
	defer val.Close() //最后自动退出。
	//读文件内容
	var tmp [1]byte //定义一个长度128byte的数组。
	n, err := val.Read(tmp[:])
	if err != nil {
		fmt.Printf("Read file err:%v.\n", err)
		return
	}
	fmt.Printf("Read %d byte.\n", n)
	fmt.Println(string(tmp[:])) //读取数组内所有的元素，转换成string型输出。
}
