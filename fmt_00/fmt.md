package main

import "fmt" //引入fmt包，具体应用是Print语句用于输出。

func main() {
	fmt.Println("The life is different,why not use go?")
	//fmt.Println==Printline意味输出一行。
}

tips:go build
go build abc.go 可以将文件转换成服务器可以直接使用的文件。
文件转化方向 .go => .exec (类unix运行文件)。
go build -o <name>.<filetype>。
//上诉代码的意思是，go build 生成文件由自己设置文件名。

此外，golang可以通过go build <fileaddress>/<name>.<filetype> <fileaddress>来改变文件生成位置但是并不实用。
tips:go run <name>.<filetype>
//直接执行go文件。

此外，golang的安装我们并不赘述，安装的问题大多是因为服务器的原因。大家可以自行研究实用goproxy.cn或者去下载别人下载好的文件替换进自己的golang路径内。