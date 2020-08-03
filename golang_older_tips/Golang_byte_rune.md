*Golang byte_rune

Golang中原样输出字符

ex:

```go
package main

import (
    "fmt"
)

func main() {
    var a = 'a'
    fmt.Printf("值:%v", a) //输出'a'的ASCII码值
  	fmt.Printf("值:%c", a) //输出'a'的原样值
}
```

'a'的值实际是a的ASCII码值，也就是97。

```go
	var strA = "this"
	fmt.Printf("值:%v 原样输出:%c 类型:%T", strA[2], strA[2], strA[2])
//值:105 原样输出:i 类型:uint8
	var strB = "this this this this"
fmt.Println(unsafe.Sizeof(strA))//unsafe.Sizeof()没有办法查询string类型所占的字符空间，但是len()可以查询长度
	//strA and strB print 16
```

Golang中汉字所使用的是utf-8的编码，编码后的值就是该文字的int值。如果想原样输出的话，%c就可以原样输出。utf-8的出现是因为多个国家的不文字，如日语韩语。

| 类型  | 表示            |
| ----- | --------------- |
| unit8 | ASCII的一个字符 |
| rune  | UTF-8的一个字符 |

因此，我们在循环语句中要拿出中文的时候，需要使用range(rune)，len(byte)。

```go
func main() {
	// var a = 'a'
	// fmt.Printf("值:%v", a) //输出'a'的ASCII码值
	// fmt.Printf("值:%c", a) //输出'a'的原样值
	// var strA = "this this this this"
	// fmt.Println(unsafe.Sizeof(strA)) //unsafe.Sizeof()没有办法查询string类型所占的字符空间
	// fmt.Printf("值:%v 原样输出:%c 类型:%T", strA[2], strA[2], strA[2])
	var str1 = "大家好，我叫白露！"
	for _, r := range str1 {
		fmt.Printf("%v(%c)", r, r)
	}
}	
/*
22823(大)23478(家)22909(好)65292(，)25105(我)21483(叫)30333(白)38706(露)65281(！)
*/
```

*Golang修改字符串

Golang中修改字符串，需要先将其转换成[]rune []byte，完成后再转换成string。无论哪种转换，都会重新分配内存，并赋值字节数组。

```go
package main

import (
	"fmt"
)

func main() {
	str1 := "big"
	bytestr1 := []byte(str1)
	bytestr1[0] = 'p'
	fmt.Println(string(bytestr1))
	str2 := "白露"
	runestr2 := []rune(str2)
	runestr2[1] = '鹭'
	fmt.Println(string(runestr2))
}
/*pig
白鹭*/
```

