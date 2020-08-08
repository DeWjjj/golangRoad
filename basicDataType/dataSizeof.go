package main

import (
	"fmt"
	"unsafe"
)

const (
	n0 int   = 1 //不同int类型所占字节和实际覆盖范围不同，虽然都是1，但是实际上代表的是00000001 0000000000000001。
	n1 int8  = 1
	n2 int16 = 1
	n3 int32 = 1
)

func main() {
	fmt.Println(unsafe.Sizeof(n0))
	fmt.Println(unsafe.Sizeof(n1))
	fmt.Println(unsafe.Sizeof(n2))
	fmt.Println(unsafe.Sizeof(n3))
}
