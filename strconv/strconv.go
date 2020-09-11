package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 97
	// ret := string(i)
	// fmt.Println(ret) //a
	ret1 := fmt.Sprintf("%d", i)
	fmt.Println(ret1) //97

	//strconv.ParseInt()
	str := "10000"
	// ret2, err := strconv.ParseInt(str, 10, 8)
	ret2, err := strconv.ParseInt(str, 10, 64)
	//如果函数内bitSize值被设置成了一个错误类型，代码就会报错。
	//例如bitSize是8，int8类型的字符串范围是 -127 ～ 128
	// The bitSize argument specifies the integer type
	// that the result must fit into. Bit sizes 0, 8, 16, 32, and 64
	// correspond to int, int8, int16, int32, and int64.
	// If bitSize is below 0 or above 64, an error is returned.
	//这段话的意思是值，0的话对应int、8对应int8，并且会有一个在8附近就归类到8，在16附近就归类到16
	if err != nil {
		fmt.Println("pasrseint failed, err:", err)
		return
	}
	fmt.Printf("%#v %T\n", ret2, ret2) //10000 int

	//strconv.Atoi() 将整型转化成字符串型
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt) //10000 int

	//strconv.Itoa() 将字符串型转化成整型
	ret3 := strconv.Itoa(i)
	fmt.Printf("%#v %T\n", ret3, ret3) //"97" string

	//strconv.ParseBool() 将字符串转换成布尔型
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue) //true bool

	//strconv.Float
	floatStr := "1.2345"
	floatValue, _ := strconv.ParseFloat(floatStr, 2) //这里的2实际上会运算成float32，函数中会4舍5入。
	fmt.Printf("%#v %T\n", floatValue, floatValue)   //1.234 float64
}
