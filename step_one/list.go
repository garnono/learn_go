package step_one

import (
	"container/list"
	"fmt"
)

func TestList() {

	l := list.New()

	e1 := l.PushBack("尾部添加")
	e2 := l.PushFront("头部添加")

	//l.PushBackList(l)
	//l.PushFrontList(l)
	//l.Init() // clear

	fmt.Println("len:", l.Len())
	//l.Remove(e1)

	// 从头部开始循环
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println("---", i.Value)
	}

	l.MoveAfter(e2, e1)
	fmt.Println("move after")
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println("---", i.Value)
	}
}
