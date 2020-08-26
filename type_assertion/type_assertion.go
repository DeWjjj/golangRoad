package main

import (
	"fmt"
)

func assertionType(x interface{}) {
	switch t := x.(type) { //用.(type)的形式赋值，这中方法值能在switch中这样写。
	case string:
		fmt.Printf("%v is %T.", t, t)
	case int:
		fmt.Printf("%v is %T.", t, t)
	case bool:
		fmt.Printf("%v is %T", t, t)
	default:
		fmt.Println("Unsupport type.")
	}

}

func main() {
	var x interface{}
	x = "Dew"
	v, err := x.(string)
	if err {
		fmt.Println(v)
	} else {
		fmt.Println("Assertion_failed.")
	}

	assertionType(x)          //Dew is string.
	assertionType(int64(200)) //Unsupport type.
}
