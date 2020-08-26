package main

import (
	"fmt"
)

func showType(x interface{}) {
	fmt.Printf("%T-%v\n", x, x)
}

//空接口
//interface 是一种关键字
//interface{} 是一种空接口类型
func main() {
	var x interface{}
	x = 123
	fmt.Printf("%v-%T", x, x)
	fmt.Println("")
	x = "123"
	fmt.Printf("%v-%T", x, x)
	fmt.Println("")

	var y map[string]interface{}
	y = make(map[string]interface{}, 20)
	y["name"] = "dew"
	y["age"] = 22
	y["gender"] = "male"
	y["isAdult"] = true
	fmt.Println(y)

	showType(true)                       //bool-true
	showType(123)                        //int-123
	showType("dew")                      //string-dew
	showType(map[string]int{"dew": 767}) //map[string]int-map[dew:767]

}
