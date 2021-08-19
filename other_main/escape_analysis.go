package main

import "fmt"

/*
变量逃逸：
参考文档：http://c.biancheng.net/view/22.html
*/

func dummy(b int) int {
	// 声明一个变量c并赋值
	var c int
	c = b
	return c
}

// 空函数, 什么也不做
func void() {
}

// 声明空结构体测试结构体逃逸情况
type Data struct {
}

func dummy2() *Data {
	// 实例化c为Data类型
	var c Data //  moved to heap: c
	//返回函数局部变量地址
	return &c
}

func main() {
	// 1
	var a int
	void()
	fmt.Println(a, dummy(0)) // ... argument does not escape; a escapes to heap;  dummy(0) escapes to heap

	// 2
	fmt.Println(dummy2()) // ... argument does not escape
}
