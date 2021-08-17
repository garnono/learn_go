package step_one

import "fmt"

func TestSentence() {

	tiaojian() // 条件

	xunhuan() // 循环
}

func tiaojian() {

	// if ... else ...
	ifVar := true
	if ifVar {
		fmt.Println("if is true")
	} else {
		fmt.Println("else is false")
	}

	// switch
	sw := 1
	switch sw != 1 {
	case false:
		fmt.Println("switch case false")
		// 不需要break，自动停止
	case true:
		fmt.Println("switch case true")
	default:
		fmt.Println("switch default")
	}

	// select
	/*
	   每个case都必须是一个通信
	   所有channel表达式都会被求值
	   所有被发送的表达式都会被求值
	   如果任意某个通信可以进行，它就执行；其他被忽略。
	   如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
	   否则：
	       如果有default子句，则执行该语句。
	       如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
	*/
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received", i1, "from c1")
	case c2 <- i2:
		fmt.Println("sent", i2, "to c2")
	case i3, ok := (<-c3): // same as i3,ok:=<-c3
		if ok {
			fmt.Println("received", i3, "from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}

}

func xunhuan() {

	/*
		三种形式：
		for init; condition; post { }
		for condition { }
		for { }
	*/
	a1 := []int{1, 2, 3, 4, 5}

	for ak := 0; ak < len(a1); ak += 1 {
		fmt.Println(ak, "=", a1[ak])
	}

	for ak, av := range a1 {
		fmt.Println(ak, "=", av)
	}

	ak := 0
	for {
		if ak < len(a1) {
			fmt.Println(ak, "=", a1[ak])
			ak++
		} else {
			break
		}
	}

	ak2, ak3 := 0, 0
	for ak2 < 3 {
		for {
			if ak3 == 1 {
				ak3++
				continue // 跳过当前循环，继续下一个
			}

			if ak3 == 3 {
				break // 终止当前循环，还可用于switch
			}

			fmt.Println("for ak 3", ak3)
			ak3++
		}

		if ak2 == 2 {
			goto breakHere // 跳到指定位置继续执行
		}

		fmt.Println("for ak 2", ak2)
		ak2++
	}

breakHere:
	fmt.Println("goto here")
}
