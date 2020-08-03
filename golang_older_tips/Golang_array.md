*Golang_array

```go
package main

import (
	"fmt"
)

func main() {
	arr1 := []string{"java", "php", "python"}
	//array := [len/none]type{"typestr1","typestr2"}
  **********************************************
	var myinfo = [4]string{1: "dew", 2: "22", 3: "music"}
	fmt.Println(myinfo)
	myinfo1 := [...]struct{
		name string
		age int8
		favor string
	}{
		{"dew",22,"music"}
	}
	fmt.Println(arr1)
}

```

这一种方法是Go语言的一维度数组，有了一维数组就自然有多维数组，在别的语言中可能此类数组被称为复合数组。



```go
	var arr1 [3][2]int //[[0 0] [0 0] [0 0]]
	arr2 := [3][2]int{{1, 1}, {2, 2}, {3, 3}} //[[1 1] [2 2] [3 3]]
	fmt.Println(arr1, arr2)
```

多维数组的第二维度不可以使用[...]，golang中的强制使用规则很多。