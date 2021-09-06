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
			4、var (
					name_1 type_1
					name_2 type_2
				)	// 一般用于全局声明，类型可以不同
		new 与 make：
		1、new 只分配内存，而 make 只能用于 slice、map 和 channel 的初始化
		2、new 返回指针，make 也是用于内存分配的，但是和 new 不同，它只用于 chan、map 以及 slice 的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。


	*/
	declareVar()

	optionVar()
}

func declareVar() {

	fmt.Println("---- 单变量的声明和赋值 ----------")

	// 1 先声明，后赋值
	var tVarA int
	tVarA = 1
	fmt.Println("var identifier type —— 声明，然后赋值：", tVarA)

	// 2 声明并赋值
	var tVarB = 2 // 省略类型
	fmt.Println("var identifier type = value —— 声明并赋值：", tVarB)

	// 3 仅内部使用
	tVarC := 3
	fmt.Println("identifier := value —— 内部变量声明并赋值：", tVarC)

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

	fmt.Println("------ new/make ------")
	tVarNewInt := new(int)
	fmt.Printf("new int type is : %T \n", tVarNewInt)
	tVarNewSlice := new([]int)
	fmt.Printf("new int type is : %T \n", tVarNewSlice)
	tVarMakeSlice := make([]int, 4, 4)
	fmt.Printf("make []int type is : %T \n", tVarMakeSlice)
}

func optionVar() {
	// TODO：类型转换
}
