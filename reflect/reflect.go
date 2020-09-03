package main

import (
	"encoding/json"
	"fmt"
)

type nameAge struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	//struct field name has json tag but is not exported
	//字段要暴露出去json包，才可以拿到，所以字段要大写。
	//过程：变量 >> json包 >> 返回main包
}

func main() {
	str := `{"name":"dew","age":22}` //str被赋予json类型的字段
	var dew nameAge
	json.Unmarshal([]byte(str), &dew)
	//当语句执行到Unmarshal时，才去判断其接受值的类型，我们输入的是一个被定义成结构体的变量，所以其读到了里面的类型。
	fmt.Println(dew.Name, dew.Age)
}
