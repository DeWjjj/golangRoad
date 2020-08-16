package main

import (
	"fmt"
)

func main() {
	var flag = '&' //定义一个特定字符，用于判断内循环是否停止（内循环终止时写入改变其数值）
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = '*'
				break
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag == '*' { //判断其数值是否为我们预设的值，如果是则再跳出外循环
			break
		}
	}
}
