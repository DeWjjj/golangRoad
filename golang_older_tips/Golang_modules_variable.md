Go语言的命名规则：

Go语言目前赋值不需要自己去添加string，因为目前的版本已经可以逆推数值类型。

example:

//var name string = "dew"

var name = "dew"

Go语言中，双使用单引号会输出该字符的ASCII码值。如果非要使用就需要用string函数去转换。

Example:

//var name = string('a')

Go的函数、变量、常量、自定义类型、包的命名方式遵循以下规则：

首字符可以是Unicode字符或者下划线，不可以是数字，剩余字符可以加上数字，且剩余长度不限制。

Go语言有25个基础的关键字，类似if、break、else、go、for之流。并且，Go还拥有37个保留字，类似于true、false、int、float。	

使用范围，可以定义在函数内部，也就作为函数内部的自定义值来使用。或者定义在全局状态下，函数中同样也可以调用。

然后，具有四种主要的声明方式，如：var（变量）、const（常量：不变量）、type（声明类型）、func（声明函数）。

Go语言中变量声明的格式：

var 变量名 变量类型

example：

//var name string

//var age int

*Golang中可以使用批量声明

example：

//var(

​	name string

​	age int

)

*Golang中一次初始化多个变量

example：

//var name, age = "DeW", 22

*Golang中函数可以用短变量声明（:=）

example：

//func main(){

​	age := 22

}

*Golang中经常可能会用到函数来给变量赋值的情况，有的时候我们只需要拿到前一个值或者后一个值。

example：

//func pval() (string, int) {

​	return "dew", 22

}

func main() {

​	name, _ =pval()

​	_, age = pval()

​	fmt.Println("name:" ,name)

​	fmt.Println("age:", age)

}

匿名变量可以用_来表示，并且用于占位，表示自己忽略的值，并且可以重复声明。所有的局部变量和所有的全局变量，不可以重复声明，但可以赋值运算。

*Golang中的常量（const）

const pi = 3.14

const e = 2.7

Const (

​	pi = 3.14

​	e = 2.7

​	eCopy

)

//e_copy会沿用上面e的数值,pi = 3.14 e = 2.7 = eCopy

*Go语言的计数器，iota

Example:

//const (

​	a0 = iota

​	a1 = iota

​	a2 = iota

​	a3 = iota

)

//a0 = 0 a1 = 1 a2 = 2 a3 = 3

*Go语言二进制偏移（<<）

var numA = 1 << 1

var numB = 1 << 10

实际上是将 1 先转化成二进制数”00000001“然后再将其左偏移十位变成，”00000010“，偏移十位则是”10000000000“

128、64、32、16、8、4、2、1

1024、512、256、128、64、32、16、8、4、2、1

*Go语言的基本类型

| Type   | Length(byte) | Default              |
| ------ | ------------ | -------------------- |
| bool   | 1            | false（ex:00000001） |
| string |              | ""                   |
| float  | 4            | 0.0                  |

并非用于数模所以，类型并不需要掌握太多。

*Golang赋值多行，可以用\n，但是不太方便所以使用``

Ex:

//var strA = ``

Println(strA)

//

var strB = `

1

2

3`

Println(strB)

//1

2

3

Tips:可能运用此方式会造成误差和空格歧义，所以在使用中需要去测试一下。实际情况下，使用情况少。

*Go语言len用法

Ex:

​	var a = "123"

​	var lenA=len(a)

//lenA = 3

<>