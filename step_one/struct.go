package step_one

import "fmt"

func TestStruct() {
	/*
		结构体：可以组合成复杂的结构
			如果对应与php的结构，struct对应class的名称和属性，struct的方法对应class的方法
		声明：
			type struct_variable_type struct {
			   member definition;
			   member definition;
			   ...
			   member definition;
			}
	*/

	declareStruct()

	optionStruct()

	testStruct()
}

func declareStruct() {
	// 局部声明使用
	type tBook struct {
		name  string
		price float32
	}

	tBook1 := tBook{"GO", 32.23}
	fmt.Println("tBook 1", tBook1)

	tBook1.price = 100.1
	fmt.Println("tBook 1`s price:", tBook1.price)

}

func optionStruct() {

}

/*
复杂的示例：全局声明，用到了 func 等相关内容
*/
const (
	Male = iota
	Female
)

type Name struct {
	name     string
	testName string
}

type People struct {
	Name         // 直接引用其他结构体，一般放在最上面
	testName int // 与Name中同名，可以通过Name.testName区分
	sex      int
}

type Book struct {
	Name
	*People // 注意：此处是指针

	author People

	price float32
}

// struct：People 的方法，相当于 php 的 class 的方法
func (p *People) SetName(name string) (err error) {
	p.name = name
	return
}

func (p *People) GetName() (name string, err error) {
	name = p.name
	return
}

func (b *Book) SetName(name string) (err error) {
	b.name = name
	return
}

func (b *Book) GetName() (name string, err error) {
	name = b.name
	return
}

func (b *Book) SetAuthorName(name string) (err error) {
	b.author.name = name
	return
}

func (b *Book) GetAuthorName() (name string, err error) {
	name = b.author.name
	return
}

func testStruct() {
	fmt.Println("------ 开始测试复杂的 struct ------")

	name := Name{"who", "haha"}
	people1 := People{name, 123, Female}
	book1 := Book{name, &people1, people1, 2.12}
	fmt.Println("简单初始化 book1:", book1)

	fmt.Println("people1 的 name：", people1.name)                     // people下仅有一个name，可以直接使用
	fmt.Println("people1 的 name 的 testName：", people1.Name.testName) // people 下存在多个 testName，需要区分使用
	fmt.Println("people1 的 testName：", people1.testName)             // people 下另一个 testName

	people1.name = "who are you" // 会影响 book1 的 People.name
	book1.SetAuthorName("gogogo")
	book1.price = 32.12

	fmt.Println("重新赋值 book1:", book1)
	fmt.Println("重新赋值 book1 的 people1.name:", book1.People.name)
}
