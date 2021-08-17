package step_one

import "fmt"

func TestArray() {
	/*
		数组：
			key-value键值对；
			key从0开始，自增1；
			value类型相同；
			长度固定
		声明：参考var中的方式
		多纬数组：类型要一致
	*/

	declareArray()

	optionArray()
}

func printArray(a [3]int, num uint8, desc string) {
	fmt.Println("array", num, a, desc, "len", len(a), "cap", cap(a))
}

func declareArray() {
	var tArray1 [3]int // 不指定长度，则为slice
	//if tArray1 == nil {
	//	fmt.Println("array 1 is nil")
	//}
	printArray(tArray1, 1, "默认值")
	tArray1[0] = 1
	printArray(tArray1, 1, "赋值")

	// 多维数组
	tArray2 := [2][3]int{}
	tArray2[0] = [3]int{}
	tArray2[0][0] = 1
	fmt.Println("多维数组", tArray2)
}

func optionArray() {

}
