package test

import (
	"learn_go/step_one"
	"testing"
)

func getStrSlice() (sl []string) {
	sl = []string{"hello world", "what are you doing"}
	return
}

func BenchJoin1(b *testing.B) {
	sl := getStrSlice()
	for i := 0; i < b.N; i++ {
		step_one.Join1(sl)
	}
}

func BenchJoin2(b *testing.B) {
	sl := getStrSlice()
	for i := 0; i < b.N; i++ {
		step_one.Join2(sl)
	}
}
