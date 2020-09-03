package main

import (
	"fmt"
	"reflect"
)

type testType struct {
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type: %v.\n", t)
	fmt.Printf("type name: %v type kind: %v.\n", t.Name(), t.Kind())
}

func main() {
	var a float32 = 3.14
	var b int64 = 100
	var c = testType{}
	reflectType(a) //type: float32 | type name: float32 type kind: float32.
	reflectType(b) //type: int64 | type name: int64 type kind: int64.
	reflectType(c) //type: main.testType | type name: testType type kind: struct.
}
