package main

import (
	"fmt"
	"reflect"
)

type name struct { //name结构体
	name string
}

func (n name) Sleep() string { //使用name结构体为方法的Sleep函数
	msg := "Sleep early is help to evade cancer."
	fmt.Println(msg)
	return msg
}

func (n name) Wake() string { //使用name结构体为方法的Wake函数
	msg := "Wake up early is help to earn money."
	fmt.Println(msg)
	return msg
}

func main() {
	dewName := name{ //创建以name结构体为原型的结构体dewName
		name: "dew",
	}
	s := reflect.TypeOf(dewName) //通过反射方法得到其类型
	//w := reflect.ValueOf(dewName)
	fmt.Println(s.NumMethod()) //输出以name结构体为类型方法的数量
	for i := 0; i < s.NumMethod(); i++ {
		methodType := s.Method(i).Type //赋予变量methodType，以Method()调用出来的字段的类型
		fmt.Printf("Method name: %v.\n", s.Method(i).Name)
		fmt.Printf("Method: %s.\n", methodType) //这里是查询方法所以是func类型
	}
}
