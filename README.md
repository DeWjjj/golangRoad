#TITLE: GOLANG_ROAD
#AUTHOR: DEW_WANG
#START_DATE: 2020/8/3

GO语言的被称为21世纪的C语言，其语言特性因为出现的时候在2008年。
AMD公司在2005年，就出产了多核处理器，所以GO语言在设计之处就支持多线程。
且现在的处理器大多为多核，并且国内知名问答社区 知乎 通过使用GO重构，节省服务器资源达80%。
所以，我这类穷逼又不想考虑服务器访问的问题，又想省钱。
为啥不用GO语言呢？？？
首先，这不是一份十分有效教程而是本人学习go一路理顺体验的过程。
在考虑版权编辑后，考虑出书。

0.tips(之前写的一些提示)
1.fmt_00(简易使用fmt模块)
2.const_var(常量和变量的分别)
3.const
4.iota(计数器)
5.shiftOperators(移位)
6.print_type(输出占位符类型)
7.path(如何去在go中保存地址)
8.boolean(布尔值)
9.Sprintf_split(字符串常规操作(判断是否包含，判断字符在字符串中位置))
10.simpleif(得到输入，判断年龄)
11.simplefor(三种类型for语句)
12.for_range(用于遍历数组、切片、字符串、map和通道(channel)后两者不会)
13.isChinese(判断是否为中文)
14.jumpoutfor(多层循环因内循环跳出而全部结束，俗称大跳)
15.for00(continue\break\fallthrough(不建议使用))
16.goto(直接从循环中跳到另外一个标签)
17.operator(各类运算符包括给二进制使用的与或) 
18.array00(普通遍历array数组，和多维数组遍历)
19.slice00(切片的本质是在数组中取数，而并非是继承)
20.make()(make函数创造切片,切片的遍历)
21.slice_append(切片的扩容)
22.delete_slice(切片的删除)
23.slice_append00(切片的删除案例)
24.pointer(指针取地址，和取该地址的值)
25.map(map使用，类似复合数组)
26.map_delete(map按照key来删除数据) //go doc 可以查看文档，但是中国人推荐上 http://www.studygolang.com/pkgoc 查阅
27.student_score(rand函数随机生成成绩，排序)
28.total_words(判断单词在句子中的出现次数)
29.func(自建函数，int=>string转换)
30.func00(关于自建函数的补充，最特殊的一种是参数增长)
31.palindrome(回文判断练习)
32.global_var(变量位置关联)
32.func_type(函数返回可以用自建函数作为参数和返回值)
33.leetcode_922(偶数、奇数按照索引的奇偶排列)
34.anon_func(匿名函数的使用)
35.calculator(简易加减乘除计算机，闭包逻辑)
36.suffix(定义类型，后缀相同原样返回，不同则添加)
37.deferEx4(liwenzhou包括指针传值内容)
38.func_closures(函数闭包)
39.deferEx6(liwenzhou变量改变对defer后的函数取值的问题)
40.panic_recover(Go语言没有异常机制，所以我们使用其函数实现错误报错)
41.scan(扫描，得到输入的三种形式)
42.distributionCoins(分配硬币，有一e,i,o,u（不区分大小写）分别得到1，2，3，4硬币)
43.recursion(递归，阶乘案例)
44.stepsWays(递归，阶梯案例1-2步上法)
45.type(自定义类型，类型别名)
46.struct(结构体%+v,%#v,匿名结构体(临时结构体))
47.struct00(结构体通过内存地址传参)
48.struct01(new函数创建空的结构体)
49.func_pointer(指针)
50.struct02(批量结构体传参，结构体内字段内存地址连续性)
51.struct_pointer(结构体指针)
52.struct_func(结构体函数)
53.method(函数调用限制)
54.expose(简单了解标示符首字母大写暴露问题)
55.method_pointer(函数调用限制传指针)
56.type_method(自定义type，用method)