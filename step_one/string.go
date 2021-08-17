package step_one

import "fmt"

func TestString() {
	/*
		字符串：由 长度固定 的 （使用UTF-8编码标识Unicode文本）单个字节 连接组成的字符序列
		声明：参考var中的方式
	*/

	declareString()

	optionString()
}

func printString(s string, num uint8, desc string) {
	fmt.Println("string", num, s, desc, "len:", len(s))
}

func declareString() {
	var tString1 string
	printString(tString1, 1, "默认值")

	tString2 := "hello"
	printString(tString2, 2, "默认类型")

	tString2 = "hello world"
	printString(tString2, 2, "重新赋值")
}

func optionString() {

}
