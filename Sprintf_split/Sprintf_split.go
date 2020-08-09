package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "DeW"
	fmt.Println(len(name)) //len()方法是计算出字符的长度

	//分割
	nameSplit := strings.Split(name, "e") //[D E]
	fmt.Println(nameSplit)

	//组合
	nameUnsplit := strings.Join(nameSplit, "+") //[D+W]
	fmt.Println(nameUnsplit)
	//判断包含
	fmt.Println(strings.Contains(name, "e")) //true
	fmt.Println(strings.Contains(name, "1")) //false

	//前缀和后缀判断
	fmt.Println(strings.HasPrefix(name, "D")) //true
	fmt.Println(strings.HasSuffix(name, "W")) //true

	//判断在字符串中位置
	fmt.Println(strings.Index(name, "W"))     // DeW W:2 [0 1 2]
	fmt.Println(strings.LastIndex(name, "W")) //从后面查找到的第一个W，在字符串中的位置
	status := "Handsome"
	nameStatus := name + ":" + status
	nameStatus1 := fmt.Sprintf("%s:%s", name, status) //Sprintf方法是把字符串组合后返回
	fmt.Printf("%s:%s\n", name, status)
	fmt.Println(nameStatus)
	fmt.Println(nameStatus1)
}
