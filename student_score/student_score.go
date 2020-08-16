package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("Stu%02d", i) //给key赋予Stu开头两位数字结尾的名字
		value := rand.Intn(100)          //给value随机一个0-99的数字
		scoreMap[key] = value            //给该key的map值赋值
	}
	fmt.Println(scoreMap)

	var keys = make([]string, 0, 200) //make函数制造0len 200cap的切片

	for key := range scoreMap { //遍历scoreMap的key后给其keys
		keys = append(keys, key) //切片中一直添加scoreMap循环的key值
	}
	fmt.Println(keys)
	sort.Strings(keys) //有序排列keys
	fmt.Println(keys)
	for _, val := range keys { //通过遍历切片keys的值，因为其值就是scoreMap的key
		fmt.Println(val, scoreMap[val])
	}

}
