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
	fmt.Println(desc, "slice", num, s, "len:", len(s), "cap:", cap(s))
}

func declareSlice() {
	var tSlice1 []int
	if tSlice1 == nil { // 数组不能通过 nil 来判断
		fmt.Println("slice 1 is nil")
	}
	printSlice(tSlice1, 1, "仅声明：")
	tSlice1 = make([]int, 0) // 初始化
	printSlice(tSlice1, 1, "初始化：")
	if tSlice1 != nil {
		fmt.Println("slice 1 is not nil")
	}
	//tSlice1[1] = 1 // 修改，超出范围会报错
	tSlice1 = append(tSlice1, 1)
	printSlice(tSlice1, 1, "赋值后")
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

	// 动态填充
	tSlice4 := make([]int, 5, 5)
	tSlice4 = append(tSlice4, 1)

	// 多维切片
	tSlice5 := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("slice 5", tSlice5, "多维", "len:", len(tSlice5), "cap:", cap(tSlice5))
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

	// 删除：获取需要的内容，再重新组合
	slice3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice4 := append(slice3[:2], slice3[3:5]...) // 删除第三个元素：2
	slice5 := append(slice4, slice3[9])          // 删除：5，6，7，8
	printSlice(slice5, 5, "delete")

}
