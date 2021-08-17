package step_one

import (
	"fmt"
	"reflect"
)

var forClosure = 0

func TestFunc() {
	/*
		func：函数、方法
			函数：func name (参数) (返回值) {函数体}
			方法：func (结构体) name (参数) (返回值) {函数体}
		参数：无、单个、多个、不定；值传递 or 引用传递；作为参数；
		返回：无、单个、多个；
		方法：继承、重载
	*/

	declareFunc()

	optionFunc()
}

// 无参数、无返回值
func declareFunc() {
	// ------ 参数 & 返回值 ------
	fmt.Println("无参数、无返回值：比如本函数")
	fmt.Println("单参数、单返回值:", t1(1))
	tFunc1, tFunc2 := t2(1, 2)
	fmt.Println("多参数、多返回值:", tFunc1, tFunc2)
	fmt.Println("不定参数、多返回值:")
	t3("test", 2, 3, 4)

	// 作为参数
	t5(t4, 3)

	// 闭包
	cl1 := closureCount()
	fmt.Println("------ 闭包1 ------")
	fmt.Println(cl1())
	fmt.Println(cl1())
	cl2 := closureCount()
	fmt.Println("------ 闭包2 ------")
	fmt.Println(cl2())
	fmt.Println(cl2())

	cl3 := closureTwo(0, 0)
	fmt.Println("------ 闭包3 ------")
	fmt.Println(cl3())
	fmt.Println(cl3())

	cl4 := closureTwo(0, 0)
	fmt.Println("------ 闭包4 ------")
	fmt.Println(cl4())
	fmt.Println(cl4())
	fmt.Println(cl4())

	cl5 := closureMore(0, 0)
	cl51 := cl5()
	cl52 := cl5()
	fmt.Println("------ 闭包51 ------")
	fmt.Println(cl51())
	fmt.Println(cl51())
	fmt.Println(cl51())
	fmt.Println("------ 闭包52 ------")
	fmt.Println(cl52())
	fmt.Println(cl52())
	fmt.Println(cl52())

	cl6 := closureMore(0, 0)
	cl61 := cl6()
	cl62 := cl6()
	cl63 := cl6()
	fmt.Println("------ 闭包61 ------")
	fmt.Println(cl61())
	fmt.Println(cl61())
	fmt.Println("------ 闭包62 ------")
	fmt.Println(cl62())
	fmt.Println(cl62())
	fmt.Println(cl62())
	cl63()

	// 方法：关联了struct的函数
	color := Color{"红色", 222}
	apple := Apple{color, "红富士"}
	fmt.Println("------ 方法 ------")
	fmt.Println("调用：", apple.getName())

	// 继承
	fmt.Println("继承：", apple.getColor())

	// 重写
	fmt.Println("重写：", apple.getNum())
}

func optionFunc() {

}

// 单参数、单返回值
func t1(a int) (b int) {
	b = a + 1
	return // 省略参数名称，会按照定义的参数进行返回
}

// 多参数、多返回值
func t2(a, b int) (c int, d int) {
	c = a
	d = b
	return b, c // 显示返回参数
}

// 不定参数、多返回值
func t3(a string, b ...int) (c, d int) {
	fmt.Println("t3 a:", a)
	fmt.Println("type:", reflect.TypeOf(b))
	for k, v := range b { // 获取参数
		fmt.Println(k, "=", v)
	}
	return
}

func t4(a int) int {
	return a * a
}

func t5(a func(b int) int, b int) {
	fmt.Println("------ 作为参数", a(b))
}

/*
闭包：就是能够读取其他函数内部变量的函数
1、返回的是函数
2、变量的范围
3、应用：私有
*/
// 计数器
func closureCount() func() int {
	i := 0 // 私有

	return func() int {
		i += 1
		fmt.Println("--", i)
		return i
	}
}

/*
参数的范围：
1、在闭包外部定义，内部修改，可以保存状态的变更：b,d
2、其他情况，不会保存参数的状态：a,c,e
*/
func closureTwo(a, b int) func() int {
	c, d := 0, 0
	a++ // 与当前方法调用次数无关
	c++ // 与当前方法调用次数无关

	return func() int {
		e := 0
		b++ // 与闭包的调用次数有关
		d++ // 与闭包的调用次数有关
		e++ // 内部变量，与调用次数无关
		forClosure++
		fmt.Println(a, b, c, d, forClosure)
		return a * b * c * d
	}
}

/*
多层闭包 的参数的范围：
1、在闭包外部定义，内部修改，可以保存状态的变更：b,d,f
2、其他情况，不会保存参数的状态：a,c,e
*/
func closureMore(a, b int) func() func() int {
	c, d := 0, 0

	return func() func() int {
		e, f := 0, 0

		a++ // 与 当前 闭包的调用次数有关，在子级闭包中的值不变
		c++ // 与 当前 闭包的调用次数有关，在子级闭包中的值不变
		e++ // 无
		forClosure++

		return func() int {
			b++ // 与 当前及上层 闭包的调用次数有关
			d++ // 与 当前及上层 闭包的调用次数有关
			f++ // 与 当前 闭包的调用次数有关
			fmt.Println(a, b, c, d, e, f, forClosure)
			return a * b * c * d * e * f
		}
	}
}

type Color struct {
	color string
	num   int
}

// 方法
func (c *Color) getColor() (color string) {
	return c.color
}

func (c *Color) getNum() (num int) {
	return c.num
}

type Apple struct {
	Color

	name string
}

// 方法
func (a *Apple) getName() (name string) {
	return a.name
}

// 重写
func (a *Apple) getNum() (num int) {
	return a.num * 10
}
