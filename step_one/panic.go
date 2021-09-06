package step_one

import "fmt"

func TestPanic() {
	// 捕获
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
			return
		}
	}()

	tP1()
}

func tP1() {
	fmt.Println("before panic")
	defer fmt.Printf("defer befor panic")
	panic("err ...")
}
