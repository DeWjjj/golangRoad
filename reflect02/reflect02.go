package main

import (
	"fmt"
	"reflect"
)

func reflectSetVal(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Type())        //*int64
	fmt.Println(v.Elem().Type()) //int64
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) //Elem()，是reflect的一种特殊用法，因为reflect.ValueOf所拿到的值等于是赋值给了v。
		//而改动v是不会对其本身有影响的，所以还是需要用其特定的方法Elem()来解决反射中更改值的问题。
	}

}

func main() {
	var a int64 = 100
	reflectSetVal(&a)
	fmt.Println(a)
}
