package main

import (
	"encoding/json"
	"fmt"
)

type information struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	//序列化，把go语言中的结构体变量 >> json格式的字符串。
	dew := information{
		Name: "dew",
		Age:  22,
	}
	x, err := json.Marshal(dew)
	if err != nil {
		fmt.Printf("marshal failed,err:%v", err)
		return
	}
	fmt.Printf("%+v\n", string(x))
	//反序列化：json格式的字符串 >> go语言中能识别的结构体变量。
	str := `{"name":"dew","age":22}`
	var dew0 information
	json.Unmarshal([]byte(str), &dew0)
	//传指针是为了可以让函数json.Unmarshal能内部修改变量(dew0)的值。
	//否则要改成赋值形式，让另外一个变量得到其中的值。
	//ex: dew1 = json.Unmarshal([]byte(str), dew0)
	fmt.Printf("%+v\n", dew0)
}
