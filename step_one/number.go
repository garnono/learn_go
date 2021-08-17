package step_one

import (
	"fmt"
	"reflect"
	"unsafe"
)

func TestNumber() {
	/*
		数字类型：
			整数（num：8、16、32、64位）：
				无符号：uint[num]，如unit8
				有符号：int[num],如int16
				其他：byte 类似 uint8
					 rune 类似 int32
					 uint uint32 or unint64
					 int  int32 or int64
					 uintptr 无符号，存放指针
			浮点：float32，float64
			复数：complex64，complex128
		声明：参考var中的方式
	*/

	declareNumber()

	optionNumber()
}

func printNumber(p interface{}, num uint, strType string) {
	fmt.Println("number", num, ":", p, "type:", strType, reflect.TypeOf(p), unsafe.Sizeof(p)) // ⚠️ TODO：unsafe.Sizeof 编译时求值，与类型的结构有关，并不返回实际运行时内存使用情况
}

func declareNumber() {
	fmt.Println("------ 整数 ------")

	var tNumber1 uint8 // 默认值 0
	printNumber(tNumber1, 1, "int")

	// 省略类型
	var tNumber2, tNumber3 = 3, unsafe.Sizeof(tNumber1) // ⚠️ TODO：tNumber5 的值与 tNumber1 在打印是的 unsafe.Sizeof 不同
	printNumber(tNumber2, 2, "默认类型")
	printNumber(tNumber3, 3, "默认类型")

	fmt.Println("------ 浮点 ------")
	var tFloat1 float64 // 0
	printNumber(tFloat1, 1, "默认值")
	tFloat2 := 3.14
	printNumber(tFloat2, 2, "默认类型")

	fmt.Println("------ 复数 ------")
	var tComplex1 complex64
	printNumber(tComplex1, 1, "默认值")
	tComplex2 := complex(1, 2)
	printNumber(tComplex2, 2, "默认类型")
	tComplex3 := complex(2.1, 3.11)
	printNumber(tComplex3, 3, "默认类型")
}

func optionNumber() {

}
