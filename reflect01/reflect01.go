package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("Type is int64,value is %d.\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("Type is float32,value is %f.\n", float32(v.Float()))
	}
}

func main() {
	var (
		a int64   = 100
		b float32 = 3.14
	)
	reflectValue(a)
	reflectValue(b)
	//Type is int64,value is 100.
	//Type is float32,value is 3.140000.
}
