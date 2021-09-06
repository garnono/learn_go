package step_one

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func TestChannel() {
	/*
		渠道：用于goroutine数据间的传递
		声明：var c [<-]chan[<-] int <-代表了方向，接受or发送，不写则是双向
		内容：
			1、方向：双向、（相对chan而言的）接受、发送
			2、双向：发送与接受必须匹配，否则将挂起——发送的多，则多余的不能发送；接受的多，多余的不能接受
			3、无缓冲区：同步
			4、有缓冲区：异步
			TODO
				发送，接受 —— 如何保证顺序
		      	超过缓冲区的数据，将如何处理
	*/
	declareChannel()

	optionChannel()
}

func declareChannel() {

	var wg sync.WaitGroup

	// 一般方式：声明 & make 初始化，二者缺一不可
	var chCache chan int
	chCache = make(chan int, 2) // 加上缓存

	// 局部方式：声明 & 初始化
	// chan<- 仅能向渠道发送数据
	// <-chan 仅能从渠道接受数据
	chNoCache := make(chan int)
	chSend, chReceive := make(chan<- int, 3), make(<-chan int, 3)

	// 延迟关闭
	defer func() {
		wg.Wait()
		fmt.Println("------- wait -------------")
		close(chCache)
		close(chNoCache)
		close(chSend)
		//close(chReceive) // 用于接受的渠道不需要关闭
	}()

	// 通过改变m、n的大小关系，查看发送、接受的情况
	m1, n1 := 5, 3
	m2, n2 := 2, 2

	// 双向——有缓存
	wg.Add(1)
	go func() {
		for i := 0; i < m1; i++ { // 超出缓存：先接受缓存容量的数据，在发送后，继续接受
			chCache <- i
			fmt.Println("ch cache send:", i)
		}
		wg.Done()
	}()

	// 接受缓存的数据：按发送的顺序
	for i := 0; i < n1; i++ {
		if chCacheValue, ok := <-chCache; ok {
			fmt.Println("ch cache receive:", chCacheValue)
		}
	}

	// 双向——无缓存
	wg.Add(1)
	go func() {
		for i := 0; i < m2; i++ { // 超出缓存
			chNoCache <- i
			fmt.Println("ch no cache send:", i)
		}
		wg.Done()
	}()
	for i := 0; i < n2; i++ {
		if chNoCacheValue, ok := <-chNoCache; ok {
			fmt.Println("ch no cache receive:", chNoCacheValue)
		}
	}

	wg.Add(1)
	go func() {
		for i := 0; i < 3; i++ {
			select {
			case chSend <- 3: // 接受
				fmt.Println("send data by ch send")
			case i3 := <-chReceive: // 发送
				fmt.Println("receive date from chReceive:", i3)
				//default:
				//	fmt.Println("default chan")
			}
		}
		// 必须有可执行的通道，否则，会死锁
		wg.Done()
	}()

	fmt.Println("------- done -------------")
}

func optionChannel() {
	fmt.Println("------------ 活锁 ---------------")

	runtime.GOMAXPROCS(3)
	cv := sync.NewCond(&sync.Mutex{}) // 条件锁
	go func() {
		for range time.Tick(1 * time.Second) { // 通过tick控制两个人的步调
			cv.Broadcast() // 唤醒条件锁
		}
	}()
	takeStep := func() {
		cv.L.Lock()
		cv.Wait()
		cv.L.Unlock()
	}
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %+v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()                      //走上一步
		if atomic.LoadInt32(dir) == 1 { //走成功就返回
			fmt.Fprint(out, ". Success!")
			return true
		}
		takeStep() // 没走成功，再走回来
		atomic.AddInt32(dir, -1)
		return false
	}
	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("向左走", &left, out)
	}
	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("向右走", &right, out)
	}
	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer walking.Done()
		defer func() { fmt.Println(out.String()) }()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v is tried!", name)
	}
	var trail sync.WaitGroup
	trail.Add(2)
	go walk(&trail, "男人") // 男人在路上走
	go walk(&trail, "女人") // 女人在路上走
	trail.Wait()
}
