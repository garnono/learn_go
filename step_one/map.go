package step_one

import (
	"fmt"
	"sync"
)

func TestMap() {
	/*
		集合：使用hash表实现的无序键值对
		声明：参考var中的方式

		对比：
					长度		多维		顺序		类型
			数组：	固定		支持		有序		一致
			切片：	不固定	支持		有序		一致
			集合：	不固定	支持		无序		一致
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

	tMap2 := map[string]string{}
	tMap2["t1"] = "GOGOGO"
	printMap(tMap2, 2, "直接初始化")
}

func optionMap() {
	fmt.Println("------ option ------")

	// 无序
	map1 := make(map[uint8]string, 10) // 对于大容量 or 会快速增加的，建议指定容量
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

	// 删除
	delete(map1, 2)
	fmt.Println("删除：2")
	for k, v := range map1 {
		fmt.Println(k, "=", v)
	}

	// 清空：重新make一个
	map1 = make(map[uint8]string)
	fmt.Println("清空：")
	for k, v := range map1 {
		fmt.Println(k, "=", v)
	}

	// 并发——需要使用sync.Map
	map2 := map[int]string{}

	go func() {
		for {
			map2[1] = "a"
		}
	}()

	go func() {
		for {
			_ = map2[1]
		}
	}()
	// fatal error: concurrent map read and map write
	//for {
	//
	//}
	// ------------ end ------------

	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}
