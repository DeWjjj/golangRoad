package main

import (
	"fmt"
)

//panic&recover

func funcA() {
	fmt.Println("A")
}
func funcB() {
	defer func() { //defer一定要在引发panic之前定义
		err := recover() //recover可以使得程序恢复执行后退出（recover必须搭配defer使用）
		if err != nil {
			fmt.Println("Recover in B")
			fmt.Println(err) //Worry!
		}
	}()
	panic("Worry!") //panic可以使得程序崩溃退出
}
func funcC() {
	fmt.Println("C")
}
func main() {
	funcA()
	funcB()
	funcC()

}
