package step_one

import (
	"fmt"
	"os"
	"strconv"
)

// 命令行参数的处理

type paramRS struct {
	Step    string
	Command string
}

// 系统方式
func TestParamOS() (p paramRS, err error) {
	fmt.Println("os 方式接受打印 cli 参数 ----- start ------")

	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}

	p.Step = os.Args[1]
	p.Command = os.Args[2]

	fmt.Println("os 方式接受打印 cli 参数 ----- end ------")

	return
}
