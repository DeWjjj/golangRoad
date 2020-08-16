package main

import (
	"fmt"
)

//切片的底层就是数组
func main() {
	var sli1 []int    //定义一个名为sli1，存放int类型元素的切片。
	var sli2 []string //定义一个数组名为sli2，存放string类型元素的切片。
	fmt.Println(sli1, sli2)
	fmt.Println(nil)         //<nil>
	fmt.Println(sli1 == nil) //true
	fmt.Println(sli2 == nil) //true nil = []

	sli1 = []int{1, 2, 3}
	fmt.Printf("%T\n", sli1)
	sli2 = []string{"北京", "上海", "广州"}
	fmt.Println(sli1, sli2)
	fmt.Printf("len(sli1):%d cap(sli2):%d\n", len(sli1), cap(sli1))
	fmt.Printf("len(sli1):%d cap(sli2):%d\n", len(sli1), cap(sli1))

	nl1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	nl2 := nl1[0:4] //基于一个数组进行取出数值，左边包含，右边不包含（如同数学中的开区间）。
	//1,2,3,4
	fmt.Printf("len(nl2):%d cap(nl2):%d\n", len(nl2), cap(nl2))
	//切片的cap是指容量，因为从数组转化成切片之后。其原先的容量，还是被切片保留的。
	//ex: 0 0 0 0 0 0 0 0 => 1 2 3 4 5 6 7 8
	//ex:nl2[0:4]
	//ex: 1 2 3 4 0 0 0 0 其cap为8。
	//ex: nl1[1:6]
	//ex 1 2 3 4 5 6 7 8 => 2 3 4 5 6 0 0 其cap值为7，头部的cap就被删去了。
	nl3 := nl1[1:6] //2,3,4,5,6
	nl4 := nl1[1:]  //2,3,4,5,6,7,8
	nl5 := nl1[:]   //1,2,3,4,5,6,7,8
	nl6 := nl1[:4]  //1,2,3,4
	fmt.Println(nl2, nl3, nl4, nl5, nl6)
	nl1[0] = 0       //0 2 3 4 5 6 7 8
	fmt.Println(nl2) //0 2 3 4
	//切片是一种引用类型，源数组被改之后切片也会被改变。切片并非是继承，所以其cap值也是一直改动的。
}
