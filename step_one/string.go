package step_one

import (
	"fmt"
	"strings"
)

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

	// 反引号 定义的字符串可以保持格式
	tString3 := `----- hello
world -------
`
	printString(tString3, 3, "反引号")

	// byte：代表了 ASCII 码的一个字符，byte 类型是 uint8 的别名
	var tByte1 byte = 65 // A
	fmt.Printf("test byte: %c \n", tByte1)

	// rune：代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型，rune 类型等价于 int32 类型
	var tRune1 int = '\u0041' // A
	fmt.Printf("test rune: %c \n", tRune1)
}

func optionString() {
	fmt.Println("------ option ------")

	// 字符串的连接：+ or join
	strSlice := []string{"hello world", "what are you doing"}

	// +
	str1, _ := Join1(strSlice)
	fmt.Println("str1 len:", len(str1))

	// join
	str2, _ := Join2(strSlice)
	fmt.Println("str2 len:", len(str2))
}

func Join1(sl []string) (s string, err error) {
	for i := 0; i < len(sl); i++ {
		s += sl[i]
	}

	return
}

func Join2(sl []string) (s string, err error) {
	s = strings.Join(sl, "")
	return
}
