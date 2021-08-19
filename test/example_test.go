package test

import (
	"fmt"
	"learn_go/step_one"
	"testing"
)

/* ------ 被测试的函数 ------ */
func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}

func Perm(num int) (out []int) {
	for i := 0; i < num; i++ {
		out = append(out, i)
	}
	return
}

// 用于测试 help
// tag 是否调用helper
// level 调用的层级深度
func tHelper(t *testing.T, tag bool, level uint8) {
	if level > 0 {
		level--
		tHelper(t, tag, level)
	} else {
		if tag {
			t.Helper() // 将输出调用者的信息
		}

		t.Fatalf("----- test helper %v ------", tag)
	}
}

/* ------ 单元测试的示例 ------ */

// 第一个示例：注意命名规则与参数
func TestHello(t *testing.T) {
	fmt.Println("hello")
}

// 标记错误，但不退出
func TestFail(t *testing.T) {
	t.Fail()
	fmt.Println("... test fail ...") // 会打印
}

// 标记错误，并且，退出
func TestFailNow(t *testing.T) {
	t.FailNow()
	fmt.Println("... test fail_now ...") // 不会打印
}

// 标记错误、记录日志、退出
func TestFatal(t *testing.T) {
	t.Fatal("... test fatal ...")
	//t.Fatalf("test %s ...", "fatalf")

	//t.Error()
	//t.Errorf()
}

// 判断是否有错误标记
func TestFailed(t *testing.T) {
	TestFail(t)
	if t.Failed() {
		fmt.Println("... test failed ...")
	}
}

// 打印输出
func TestLog(t *testing.T) {
	fmt.Println("... test log ...")
	t.Log("... test log here ...")
	//t.Logf("... test %s ...", "logf")
	fmt.Println("...... log here")
}

// log + fail
func TestError(t *testing.T) {
	t.Error("... test error ...")
	//t.Errorf("... test %s ...", "errorf")
}

func TestSkip(t *testing.T) {
	t.SkipNow() // 忽略 & 退出

	//t.Skip("... test skip ...") // log + SkipNow
	//t.Skipf("... test %s ...", "skipf")

	//t.Skipped() // 判断是否存在退出
}

func TestHelperNotHave(t *testing.T) {
	tHelper(t, false, 3)
}

// 打印调用函数的位置
func TestHelperHave(t *testing.T) {
	tHelper(t, true, 3)
}

// 包含多个单元测试用例的函数：run
func TestChildren(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}

// 对比是否与 Output: 定义的内容相同，不能用于Test开头的函数中
func ExampleHello() {
	fmt.Println("hello")
	// Output: hello world
}

// 对比是否与 Unordered output: 定义的内容相同，注意不包括顺序，不能用于Test开头的函数中
func ExamplePerm() {
	for _, value := range Perm(5) {
		fmt.Println(value)
	}

	// Unordered output: 4
	// 2
	// 1
	// 3
	// 0
}

// 基准点：以Benchmark开头
func BenchmarkStringJoin1(b *testing.B) {
	sl := []string{"0123456789", "abcdefghijklmnopqrstuvwxyz"}
	for i := 0; i < b.N; i++ {
		step_one.Join1(sl)
	}
}
func BenchmarkStringJoin2(b *testing.B) {
	//b.StopTimer()	// 可以跳过下面代码执行的计时时间
	//time.Sleep(10 * time.Second)
	//b.StartTimer()

	sl := []string{"0123456789", "abcdefghijklmnopqrstuvwxyz"}
	for i := 0; i < b.N; i++ {
		step_one.Join2(sl)
	}
}
