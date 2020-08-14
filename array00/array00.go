package main

import (
	"fmt"
)

func main() {
	var n1 [3]bool //[3]bool array的类型包括长度，所以n1,n2的类型是完全不同的。
	var n2 [4]bool //[4]bool

	fmt.Printf("%T %T\n", n1, n2)
	//数组初始化有数值
	fmt.Println(n1) //boolean => false int,float=>0,string=>""
	//function1
	n1 = [3]bool{true, true, true}
	fmt.Println(n1)
	//funciton2
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} //自动推算array的长度
	fmt.Println(a1)
	//function3
	a2 := [5]int{1, 2} //1 2 0 0 0
	fmt.Println(a2)
	//function4
	a3 := [5]int{0: 1, 4: 2} //1 0 0 0 2
	fmt.Println(a3)

	//数组的遍历
	citys := [...]string{"beijing", "shanghai", "guangzhou"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//for_range遍历数组
	for k, v := range citys {
		fmt.Println(k, v)
	}
	//多维数组 ex:[[1,2],[3,4],[5,6]]
	var a9 [3][2]int
	a9 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	for _, v := range a9 {
		fmt.Println(v)
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}
	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 0
	fmt.Println(b1, b2)
}
