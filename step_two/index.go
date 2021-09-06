package step_two

import (
	"fmt"
	"reflect"
)

var commands = map[string]func(){
	"io": TestIO,
}

func call(name string, params ...interface{}) {
	f := reflect.ValueOf(commands[name])
	if len(params) != f.Type().NumIn() {
		return
	}

	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}

	f.Call(in)
}

func TestRun(c string) {
	fmt.Println("-- step two start to test ---------")
	call(c)
}
