package step_one

import (
	"fmt"
	"unsafe"
)

func TestConst() {
	/*
		常量：运行时不能修改
		声明：const identifier [type] = value
			const 是关键字
			identifier 是名称，首字母一般大写，驼峰式
			type 可省略
			value 只能是基础类型的值 或 len(), cap(), unsafe.Sizeof() 等内置函数计算表达式的值
		iota：特殊常量
			在 const关键字出现时将被重置为 0(const 内部的第一行之前)
			const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)
	*/

	declareConst()

	optionConst()
}

func declareConst() {
	const TConst1 int = 1                              // 完整的声明 & 赋值
	const TConst2 = 2                                  // 省略类型
	const TConst3, TConst4 = 3, unsafe.Sizeof(TConst2) // 多常量，内置函数计算，注意，内置函数不能用于计算同行声明的常量
	fmt.Println(TConst1, TConst2, TConst3, TConst4)

	const ( // 合并声明，一般用于相关性强的逻辑上
		TConstString    = "string"
		TConstStringLen = len(TConstString)
	)
	fmt.Println("TConstString:", TConstString, "len:", TConstStringLen)

	// iota
	const (
		Ta1 = iota      // iota = 0
		Ta2             // iota = 1，自增1
		Ta3 = 100       // iota = 2，重新赋值 100
		Ta4             // iota = 3，延续上面的值 100
		Ta5 = 1 << iota // iota = 4，1<<4 = 10000 = 2^4 = 16
		Ta6             // iota = 5，延续上面的表达式，1<<5 = 32
		Ta7 = iota      // iota = 6，重新赋值 iota
	)

	const Ta8 = iota // 重新归0

	fmt.Println("iota:", Ta1, Ta2, Ta3, Ta4, Ta5, Ta6, Ta7, Ta8)
}

func optionConst() {

}
