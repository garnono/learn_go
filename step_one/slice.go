package step_one

import "fmt"

func TestSlice() {
	/*
		切片：与数组类似，但长度可变
		长度（len）：实际值的数量
		容量（cap）：可以存储的数量
		声明：参考var中的方式
	*/

	declareSlice()

	optionSlice()
}

func printSlice(s []int, num uint8, desc string) {
	fmt.Println("slice", num, s, desc, "len:", len(s), "cap:", cap(s))
}

func declareSlice() {
	var tSlice1 []int
	if tSlice1 == nil { // 数组不能通过 nil 来判断
		fmt.Println("slice 1 is nil")
	}
	printSlice(tSlice1, 1, "默认值")
	tSlice2 := make([]int, 3, 4)
	printSlice(tSlice2, 2, "make")

	// 声明数组，用于 slice 的操作
	tArray := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array:", tArray, "len:", len(tArray))
	/*
		从key的角度考虑：[k1,k2)，即包括k1，不包括k2
		k1留空，则从0开始
		k2留空，可直到最后一个元素
		容量 = 当前容量 - k1
	*/
	tSlice3 := tArray[2:5]
	fmt.Println("slice[1]:", tSlice3[1])
	printSlice(tSlice3, 3, "劫取")
}

func optionSlice() {
	fmt.Println("------ option ------")

	slice1 := make([]int, 3, 4)
	// 追加
	printSlice(slice1, 1, "刚声明")
	slice1 = append(slice1, 1)
	printSlice(slice1, 1, "append 1")
	slice1 = append(slice1, 1) // ⚠️ len, cap 都增加了：📖TODO 增加的策略是？
	printSlice(slice1, 1, "append 2")

	// 复制
	slice2 := make([]int, 3, 4)
	copy(slice2, slice1) // ⚠️ 只复制所能承受的长度，不会自动扩展
	printSlice(slice2, 2, "copy")
}
