package step_one

import "fmt"

func TestVar() {
	fmt.Println("test var")

	/*
		变量：在运行时可以改变
		单声明：
			1、var identifier type	声明，var是关键字，identifier是变量名称，type是变量值的类型
			2、var identifier type = value 声明并赋值，此时 type 可以省略
			3、identifier := value	声明并赋值，不能用户全局，仅可用于局部变量

		多变量的声明（前三个与单声明类似）：
			1、var name_1, name_2 type		类型需要一样
			2、var name_1, name_2 type = value_1,value_2	类型省略可以不一样
			3、name_1, name_2 := value_1,value_2	类型可以不一样
			4、var {
					name_1 type_1
					name_2 type_2
				}	// 一般用于全局声明，类型可以不同
		注意：
			某些类型的变量需要使用new、make
	*/
	declareVar()

	optionVar()
}

func declareVar() {
	// 单变量的声明
	fmt.Println("---- 单变量的声明和赋值 ----------")

	// 1 先声明，后赋值
	var tVarA int
	tVarA = 1
	fmt.Println("var identifier type —— 声明，然后赋值：", tVarA)

	// 2 声明并赋值
	var tVarB = 2
	fmt.Println("var identifier type = value —— 声明并赋值：", tVarB)

	// 3 仅内部使用
	tVarC := 3
	fmt.Println("identifier := value —— 内部变量声明并赋值：", tVarC)

	// 多变量的声明
	fmt.Println("---- 多变量的声明和赋值 ----------")

	// 1 先声明，后赋值
	var tVarA1, tVarA2 int
	tVarA1, tVarA2 = 1, 2
	fmt.Println("声明，然后赋值：", tVarA1, tVarA2)

	// 2 声明并赋值
	var tVarB1, tVarB2 = 1, "str"
	fmt.Println("声明并赋值：", tVarB1, tVarB2)

	// 3 仅内部使用
	tVarC1, tVarC2 := 1, "string"
	fmt.Println("内部变量声明并赋值：", tVarC1, tVarC2)

	// 4 合并方式
	var (
		tVarD1 = 1
		tVarD2 string
	)
	tVarD2 = "string"
	fmt.Println("var()的方式 声明，然后赋值：", tVarD1, tVarD2)

	// TODO：new、make
}

func optionVar() {
	// TODO：类型转换
}
