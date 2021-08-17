package main

import (
	"fmt"
	"learn_go/step_one"
	"learn_go/step_two"
)

func main() {
	fmt.Println("hello world")

	// 首先，我们获取下命令行的参数
	param, _ := step_one.TestParamOS()

	// 根据输入的参数，动态执行相应的代码
	switch {
	case param.Step == "one":
		// 基础
		step_one.TestRun(param.Command)
	case param.Step == "two":
		// 进阶
		step_two.TestRun(param.Command)
	default:
		fmt.Sprintf("传入的 step：%s 不支持", param.Step)
	}

	fmt.Println("执行完毕")
}
