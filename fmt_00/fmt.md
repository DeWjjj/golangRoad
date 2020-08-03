package main

import "fmt" //引入fmt包，具体应用是Print语句用于输出。

func main() {
	fmt.Println("The life is different,why not use go?")
	//fmt.Println==Printline意味输出一行。
}

tips:go build
go build abc.go 可以将文件转换成服务器可以直接使用的文件。
文件转化方向 .go => .exec (类unix运行文件)。
go build -o <name>.<filetype>