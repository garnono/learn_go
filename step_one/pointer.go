package step_one

import (
	"fmt"
	"reflect"
)

func TestPointer() {
	/*
		指针：即内存地址；虽然内存地址相同，但仍需要指明对应的类型——TODO
		声明：参考var中的方式，在类型前加上"*"即可，如：var p *int  指向int类型的指针变量；通过 *获取当前变量指针指向的内容; 通过 & 获取当前变量的地址
		我的理解：变量实际存储了两个值，一个是我们看到的内容；一个是存储该内容的地址；
	*/

	declarePointer()

	optionPointer()
}

func printPointer(p *int, num uint8, desc string) {
	fmt.Println("pointer", num, p, desc, "type:", reflect.TypeOf(p))
	if p != nil {
		fmt.Println("value:", *p)
	}
}

func declarePointer() {
	var tPointer1 *int
	printPointer(tPointer1, 1, "默认值")
	tInt := 123
	tPointer2 := &tInt
	printPointer(tPointer2, 2, "默认类型")
}

func optionPointer() {

}
