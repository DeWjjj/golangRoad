package main

import (
	"fmt"
	"runtime"
)

func f0() {
	pc, file, line, ok := runtime.Caller(0)
	//Caller 0 >> main.f0
	//Caller 1 >> main.f1
	//Caller 2 >> main.main
	if !ok {
		fmt.Println("runtime.Caller() failed.")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(file)     ///Users/dew/go/src/golangRoad/runtime_Caller/runtime_Caller.go
	fmt.Println(line)     //9
	fmt.Println(ok)       //true
	fmt.Println(pc)       //17420475
	fmt.Println(funcName) //main.f0 main包中的f0函数。
}

func f1() {
	f0()
}

func main() {
	f1()
	// 	pc, file, line, ok := runtime.Caller(0)
	// 	if !ok {
	// 		fmt.Println("runtime.Caller() failed.")
	// 		return
	// 	}
	// 	funcName := runtime.FuncForPC(pc).Name()
	// 	fmt.Println(file)     ///Users/dew/go/src/golangRoad/runtime_Caller/runtime_Caller.go
	// 	fmt.Println(line)     //9
	// 	fmt.Println(ok)       //true
	// 	fmt.Println(pc)       //17420475
	// 	fmt.Println(funcName) //main.main main包中的main函数。
}
