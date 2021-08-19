package step_one

import (
	"fmt"
	"strconv"
)

func TestTypeConversion() {
	/*
		类型转换
	*/

	// 一般情况：只有相同底层类型的变量之间可以进行相互转换
	var tByte1 byte = 65
	tInt1 := int(tByte1) // 显示转换
	fmt.Printf("byte to int: %d \n", tInt1)

	// int ——》 string
	tString2 := "2"
	tInt2, _ := strconv.Atoi(tString2)
	fmt.Printf("string to int: %d \n", tInt2)

	// string ——》 int
	tInt3 := 3
	tString3 := string(tInt3)
	fmt.Printf("int to string: %d \n", tString3)
}
