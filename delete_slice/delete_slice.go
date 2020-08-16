package main

import (
	"fmt"
)

func main() {
	s1 := [...]int{1, 2, 3, 4, 5} //数组
	fmt.Printf("%T\n", s1)
	x1 := s1[:] //切片
	fmt.Printf("%T\n", x1)
	fmt.Println(x1)
	fmt.Println(s1, cap(s1))
	x1 = append(x1[:2], x1[3:]...) //当使用了append()去修改切片的时候，实际上会修改底层数组。
	fmt.Println(x1)
	fmt.Println(s1, cap(s1))
	//虽然切片是去调用一片连续的数组，但是用append函数去重新组织切片所指向的内存值的时候实际上就是在重新安排数组的内存值。
	//并且虽然切片的值看上去删除了地址，实际上只是将其后面的值赋值到了删除的值上，没有波及的值未变。

	//ex：s1 := [...]int{1,2,3}
	//x1 := s1[:] 获得所有所选范围的内存值
	//x1 = append(x[0],x[2])等于是把x[0]的数值保留，直接开始续上x[2]的数值。原先的x2并不会变化
	//x1 = 1 3
	//s1 = 1 3 3
}
