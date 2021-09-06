package step_one

import (
	"fmt"
	"reflect"
)

func TestNil() {

	/*
		nil：预定义标识符
	*/

	// 不能比较
	fmt.Println("compare:", compareNil())

	// 没有默认类型
	fmt.Printf("nil type %T \n", nil)
	//fmt.Println(reflect.TypeOf(nil).String())

	// 不同类型的 nil 是不能比较的
	var nil1 []int
	var nil2 map[int]string
	if nil1 == nil {
		fmt.Println("[]int nil")
	}
	if nil2 == nil {
		fmt.Println("map is nil")
	}
	fmt.Println("cc:", nil1, nil2)
	//if nil1 == nil2 {
	//
	//}

	// 不是关键字或保留字——可以作为变量名称，但不建议
	nil := 1
	fmt.Println("nil as int: ", nil)
	fmt.Printf("nil type %T \n", nil)
	fmt.Println(reflect.TypeOf(nil))
}

func compareNil() (err error) {
	//if nil == nil {
	//}

	return
}
