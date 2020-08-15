package main

import (
	"fmt"
)

func main() {
	var maps map[string]int
	maps = make(map[string]int, 10)
	maps["wangchao"] = 22
	maps["dew"] = 18
	maps["jackie"] = 14
	fmt.Println(maps) //map[dew:18 jackie:14 wangchao:22]
	delete(maps, "dew")
	fmt.Println(maps)        //map[jackie:14 wangchao:22]
	delete(maps, "xiaoming") //提供不存在的key，去删除值
	fmt.Println(maps)        //提供空key删除时，会no-op(不操作)
}
