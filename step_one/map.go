package step_one

import "fmt"

func TestMap() {
	/*
		集合：使用hash表实现的无序键值对
		声明：参考var中的方式
		make && new
	*/

	declareMap()

	optionMap()
}

func printMap(m map[string]string, num uint8, desc string) {
	fmt.Println("map", num, m, desc, "len:", len(m))
}

func declareMap() {
	var tMap1 map[string]string
	printMap(tMap1, 1, "声明")
	tMap1 = make(map[string]string) // 必须初始化后才能使用，声明和初始化可以合并：var tMap1 = make(map[string]string)
	printMap(tMap1, 1, "make 初始化")
	tMap1["test_01"] = "go"
	tMap1["test_02"] = "php"
	printMap(tMap1, 1, "赋值")

	tMap2 := map[string]string{"t1": "GGG", "t2": "PPP"}
	printMap(tMap2, 2, "直接初始化")
}

func optionMap() {
	fmt.Println("------ option ------")

	// 无序
	map1 := make(map[uint8]string)
	map1[1] = "go"
	map1[2] = "php"
	map1[3] = "python"
	map1[4] = "java"
	map1[5] = "c/c++"

	// 循环1
	fmt.Println("无须循环1：")
	for k, v := range map1 {
		fmt.Println(k, "=", v)
	}

	// 循环2
	fmt.Println("无须循环2：")
	for k, v := range map1 {
		fmt.Println(k, "=", v)
	}
}
