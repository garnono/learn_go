package step_one

import "fmt"

func TestBool() {
	/*
		布尔类型：true、false（默认）
		声明：参考var中的方式
	*/

	declareBool()

	optionBool()
}

func printBool(b bool, num uint) {
	fmt.Println("bool", num, b)
}

func declareBool() {
	// 以此为例，加强变量声明、赋值的方式的使用

	fmt.Println("------ 单变量 ------")

	// 1 先声明，后赋值
	var tBool1 bool
	printBool(tBool1, 1) // 默认值
	tBool1 = true
	printBool(tBool1, 1) // 赋值后

	// 2 声明并赋值，可省略类型
	var tBool2 = true // var tBool2 bool = true
	printBool(tBool2, 2)

	// 3 仅限内部方式声明并赋值，可省略类型
	tBool3 := true
	printBool(tBool3, 3)

	fmt.Println("")
	fmt.Println("------ 多变量 ------")

	// 1 先声明，后赋值
	var tBoolM1, tBoolM2 bool
	tBoolM2 = true
	printBool(tBoolM1, 1) // 默认false
	printBool(tBoolM2, 2)

	// 2 声明并赋值，可省略类型
	var tBoolM3, tBoolM4 = true, false // 类型必须一样
	printBool(tBoolM3, 3)
	printBool(tBoolM4, 4)

	// 3 仅限内部方式声明并赋值，可省略类型
	tBoolM5, tBoolM6 := false, true // 类型可以不同
	printBool(tBoolM5, 5)
	printBool(tBoolM6, 6)

	// 合并方式
	var (
		tBoolM7 bool   // 仅声明
		tBoolM8 = true // 声明并赋值
	)
	printBool(tBoolM7, 7)
	printBool(tBoolM8, 8)
}

func optionBool() {

}
