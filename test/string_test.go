package test

import (
	"fmt"
	"learn_go/step_one"
	"runtime"
	"testing"
)

func TestJoin1(t *testing.T) {
	sl := []string{"hello world", "what are you doing"}
	step_one.Join1(sl)
	t.Fatalf("err")

	//runtime.GC()
	b := step_one.Book{}
	b.SetAuthorName("GC ...")

	runtime.SetFinalizer(b, func() {
		name, _ := b.GetAuthorName()
		fmt.Println("gc here ...", name)
	})
}
