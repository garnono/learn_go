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

	// int => string
	var tInt2 int = 65
	tStr21 := string(tInt2)
	fmt.Println("int to string:", tStr21)

	// ------ 可以通过 strconv 包进行转化 ------

	// string => int : Atoi
	tStr3 := "3"
	tInt31, _ := strconv.Atoi(tStr3)
	fmt.Printf("string to int: %d \n", tInt31)
	// int => string : Itoa
	tInt4 := 3
	tStr41 := strconv.Itoa(tInt4)
	fmt.Printf("int to string: %d \n", tStr41)

	// string => 其他 : Parse 系列
	tStr5 := "522"
	tInt5, _ := strconv.ParseInt(tStr5, 10, 64)
	fmt.Printf("string to int by ParseInt: %d \n", tInt5)

	// 其他 => string : Format 系列
	tInt6 := 3
	tStr61 := strconv.FormatInt(int64(tInt6), 10)
	fmt.Printf("int64 to string by FormatInt: %d \n", tStr61)

	tStr62 := strconv.FormatUint(uint64(tInt6), 10)
	fmt.Printf("uint64 to string by FormatUint: %d \n", tStr62)
	tFloat1 := 3.14
	tStr63 := strconv.FormatFloat(tFloat1, 'e', 10, 32)
	fmt.Printf("float to string by FormatFloat: %d \n", tStr63)

	// []byte 追加其他 : Append 系列
	tByte2 := []byte{'a', 'b'}
	tByte3 := strconv.AppendBool(tByte2, false)
	fmt.Printf("append bool: %d \n", tByte3)

	tByte4 := strconv.AppendInt(tByte2, 4, 2)
	fmt.Printf("append int: %d \n", tByte4)

	// 单双引号 ：Quote 系列
	tStrQuote1 := strconv.Quote(`string sss`)
	fmt.Printf("quote : %d \n", tStrQuote1)
	tStrQuote2 := `123456
 22`
	fmt.Printf("CanBackquote: %d \n", strconv.CanBackquote(tStrQuote2))
}
