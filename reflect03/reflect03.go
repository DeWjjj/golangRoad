package main

import (
	"fmt"
	"reflect"
)

type nameAge struct {
	name string
	age  int
}

func main() {
	var a *int
	fmt.Println("a is IsNil:", reflect.ValueOf(a).IsNil())         //判断指针是否为空
	fmt.Println("nil is IsValid:", reflect.ValueOf(nil).IsValid()) //判断返回值是否有效

	b := nameAge{
		name: "dew",
		age:  22,
	}
	fmt.Printf("b_struct have \"name\" key: %v.\n", reflect.ValueOf(b).FieldByName("name").IsValid())
	fmt.Printf("b_struct have \"abc\" key: %v.\n", reflect.ValueOf(b).FieldByName("abc").IsValid())

	d := struct{}{}
	fmt.Printf("anon_struct have \"abc\" key: %v.\n", reflect.ValueOf(d).FieldByName("abc").IsValid())
	fmt.Printf("anon_struct have \"abc\" method: %v.\n", reflect.ValueOf(d).MethodByName("abc").IsValid())

	c := map[string]int{}
	c = make(map[string]int, 10)
	c["dew"] = 22
	fmt.Printf("c_map have \"dew\" key: %v.\n", reflect.ValueOf(c).MapIndex(reflect.ValueOf("dew")).IsValid())
	fmt.Printf("c_map have \"abc\" key: %v.\n", reflect.ValueOf(c).MapIndex(reflect.ValueOf("abc")).IsValid())
}
