package main

import (
	"fmt"
)

func main() {
	s1 := []string{"beijing", "shanghai", "guangzhou"} //len = 3 cap = 3
	fmt.Println(s1)
	s1[0] = "北京"
	//s1[3] = "shenzhen" index out of range,错误的追加写法，编译越界。
	s1 = append(s1, "shenzhen")             //调用append元素必须使用变量接受返回值。
	s2 := append(s1, "hangzhou", "chengdu") //连续加两个值
	fmt.Println(s2)
	s3 := []string{"wuhan", "xian"}
	s2 = append(s2, s3...) //...表示拆开
	fmt.Println(s2)
	//其实现逻辑是，将s1数组的位置重新搬移。实际上，变相的等于调用出其数据，删除原来的数据，找一个新的内存地址赋值上去。
	// new s1 len = 4 cap = 6
	//fmt.Printf("S1:%v len(S1):%d cap(S1):%d\n", s1, len(s1), cap(s1))
	//不强调cap的计算扩容方式，因为其计算策略常有更改。
	//fmt.Println(s1)
}
