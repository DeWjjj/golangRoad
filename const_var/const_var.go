package main //GO语言中推荐使用驼峰不建议使用_(下划线)。

import (
	"fmt"
)

const ( //常量定义了之后不能修改
	country = "china"
	sexy    = "man" // 全局变量可以声明之后不进行使用
)

var birthYear int = 1998

//现在的GO语言代码已经有了逆推的功能所以int在实际运用中可以剔除。
//GO语言声明的变量不赋值，会被赋予默认的数值。
//并且输出的类型各有不同，某些特定的场景下会存在数据精度不足的问题。

func main() { //一个文件夹下只能有一个main()函数
	name := "dew" //函数内的变量必须使用
	age := "22"
	fmt.Printf("My name is %v.I come from %v.I am %v years old!\n", name, country, age)
	fmt.Printf("My name is %v.\n", name)
	fmt.Printf("I come from %v.\n", country)
	fmt.Printf("I am %v years old.\n", age)
	//这里要强调一下，为什么我们会使用Printf，因为Println意味输出一行，Printf意味输出语句。
	//但是这样输出一行，并不好看，所以此时我们可以使用换行符\n，然后我们将第一句拆成三句。
}
