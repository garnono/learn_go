package step_one

import (
	"fmt"
	"time"
)

func TestSentence() {

	// 条件
	testIfElse()
	testSwitch()
	testSelect()

	testFor() // 循环
}

func testIfElse() {
	fmt.Println("------ if...else... ------")

	ifVar := true
	if ifVar {
		fmt.Println("if is true")
	} else {
		fmt.Println("else is false")
	}
}

func testSwitch() {
	fmt.Println("------ switch ------")

	sw := 1
	switch {
	case sw < 3: // 可以用表达式
		fmt.Println("表达式...")
		fallthrough // 继续执行下一个，不会判断是否满足case的条件
	case sw > 3:
		fmt.Println("switch case true by fallthrough")
		// 不需要break，自动停止
	case sw == 3:
		fmt.Println("switch case false")
	default:
		fmt.Println("switch default")
	}

	// type-switch
	fmt.Println("------ type-switch ------")
	var x interface{} // 不能是具体类型
	x = 1

	switch i := x.(type) { // i 可以省略
	case nil:
		fmt.Printf(" x 的类型 :%T \n", i)
	case int:
		fmt.Printf("x 是 int 型 \n")
	case float64:
		fmt.Printf("x 是 float64 型 \n")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型 \n")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型 \n")
	default:
		fmt.Printf("未知型 \n")
	}
}

func testSelect() {
	fmt.Println("------ select ------")
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
	c1, c2, c3 := make(<-chan int, 1), make(chan<- int, 10), make(chan int, 2)
	defer func() {
		//close(c1)	// 发送的不需要关闭
		close(c2)
		close(c3)
	}()

	i2 := 0
	c3 <- 1

	select {
	case i1 := <-c1:
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

func testFor() {
	fmt.Println("------ test ------")
	/*
		三种形式：
		for init; condition; post { }
		for condition { }
		for { }
	*/
	a1 := []int{1, 2, 3, 4, 5}

	fmt.Println("------ for init; condition; post { } ------")
	start1 := time.Now()
	for ak := 0; ak < len(a1); ak += 1 {
		fmt.Println(ak, "=", a1[ak])
	}
	fmt.Println("执行时间：", time.Since(start1))

	fmt.Println("------ for condition { } ------")
	start2 := time.Now()
	ak := 0
	for ak < len(a1) {
		fmt.Println(ak, "=", a1[ak])
		ak++
	}
	fmt.Println("执行时间：", time.Since(start2))

	fmt.Println("------ for { } ------")
	start3 := time.Now()
	for {
		if ak < len(a1) {
			fmt.Println(ak, "=", a1[ak])
			ak++
		} else {
			break
		}
	}
	fmt.Println("执行时间：", time.Now().Sub(start3))

	fmt.Println("------ range ------")
	start4 := time.Now()
	for ak, av := range a1 {
		fmt.Println(ak, "=", av)
	}
	fmt.Println("执行时间：", time.Now().Sub(start4))

	ak1, ak2, ak3 := 0, 0, 0

	for ak1 < 3 {
		ak1++
	BreakLoop:
		for ak2 < 5 {
			ak2++
			ak3 = 0
		OuterLoop:
			for {
				ak3++
				if ak3 == 2 {
					continue // 跳过当前循环，继续下一个
				}

				if ak3 == 3 {
					continue OuterLoop
				}

				if ak3 > 3 {
					break BreakLoop // 终止当前循环，还可用于switch
				}

				fmt.Println("for ak:", ak1, ak2, ak3)
			}

			if ak2 == 3 {
				goto breakHere // 跳到指定位置继续执行
			}

			fmt.Println("for ak2:", ak2)
		}
	}

breakHere:
	fmt.Println("goto here")

	/*
		range：返回的是副本
	*/
	slice1 := []string{"a", "b", "c"}
	for k, v := range slice1 {
		// k,v的地址不会变，存储slice1对应的副本数据
		fmt.Println("range:k-", &k, "v-", &v, "value-", &slice1[k])
	}

}
