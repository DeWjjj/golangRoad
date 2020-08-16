package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 5, 10) //make([]<type>, <len> ,<cap>)
	fmt.Printf("S1=%v len(s1):%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := make([]int, 0, 10) //make([]<type>, <len> ,<cap>)
	fmt.Printf("S2=%v len(s2):%d cap(s2)=%d\n", s2, len(s2), cap(s2))
	s3 := make([]int, 0)
	fmt.Println(s3, len(s3), cap(s3))

	//切片就是一个选区，其选中了一块连续的内存。因为GO语言是偏底层的语言，所以其不能类似于python一样数组中多钟类型存放。
	//nil值的长度和容量都是0，但是我们不能说一个长度和容量都是0的切片一定是nil
	//切片都不保存数值

	s4 := []int{1, 3, 5} //s4数组
	s5 := s4
	fmt.Println(s4, s5)
	s4[0] = 0
	fmt.Println(s4, s5) //035 035

	//切片的遍历

	for i := 0; i < len(s4); i++ {
		fmt.Println(s4[i])
	}
	fmt.Println("")
	for i, v := range s4 {
		fmt.Println(i, v)
	}
}
