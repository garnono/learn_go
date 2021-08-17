package step_one

import (
	"errors"
	"fmt"
)

func TestInterface() {
	/*
		接口：方法的集合，实现了所有的方法，则认为实现了该接口
		定义：
		type interface_name interface {
			method_name1 [return_type]
			method_name2 [return_type]
		}
		接口作为类型，可以接收实现该接口的实例，在类型处理方面更便捷——如方法的传参的类型
		类型判断的方式
	*/

	declareInterface()

	optionInterface()
}

func declareInterface() {

	iphone := IPhone{name: "iphone"}
	android := Android{name: "android"}

	fmt.Println("------ 调用接口定义的方法 ------")
	iphone.call()
	android.call()

	// 存储的值——任意类型的值
	fmt.Println("------ 接口类型接收实现该接口的实例 ------")
	var mobile Phone
	mobile = iphone
	mobile.call()
	mobile = android
	mobile.call()

	fmt.Println("------ 判断类型 ------")
	if _, ok := mobile.(Android); ok {
		fmt.Println("type is android")
	}

	switch v := mobile.(type) {
	case IPhone:
		fmt.Println("is iphone", v)
	case Android:
		fmt.Println("is android", v)
	}
}

func optionInterface() {

}

// 定义接口
type Phone interface {
	call() (int, error)
	send(msg string)
}

type IPhone struct {
	name string
}

func (p IPhone) call() (int, error) {
	fmt.Println(p.name, "to call")
	return 1, errors.New("nothing")
}

func (p IPhone) send(msg string) {
	fmt.Println(p.name, "send message...", msg)
}

type Android struct {
	name string
}

func (p Android) call() (int, error) {
	fmt.Println(p.name, "to call")
	return 1, errors.New("nothing")
}

func (p Android) send(msg string) {
	fmt.Println(p.name, "send message...", msg)
}
