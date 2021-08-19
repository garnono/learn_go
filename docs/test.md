# tool：单元测试

### 简介
查看帮助
```
go help test
go help testfunc
```

文件命名
1. 需要以 "_test.go" 结尾
2. 以 "_" or "." 开头的文件会被忽略

函数命名
1. 以 Test 开头，后接大写开头的名称，如：TestHello，不能是Testhello
2. 以 Example 开头
3. 以 Benchmark 开头

### 示例——参看 test/example_test.go 

### 常用命令——需要在单元测试用例的目录下执行，或者指定对应的目录
```
# 执行所有的单元测试用例

# 在项目根目录下时
go test ./test
# 在 test 目录下时——下面都按照在 test 目录下执行
go test

# 展示执行的步骤信息——带上"-v"参数即可
go test -v

# 指定文件
go test example_test.go
# 指定文件中的函数——注意：run选项支持正则匹配
go test -run TestHello example_test.go
# 指定文件中函数的子用例
go test -run TestChildren/neg example_test.go

# 执行基准用例
go test -bench=. -run=^$

# 基准命令参数说明：
    -cpu: 设置测试最大 cpu 逻辑数(也就是 GPM 中 P, 最大并行执行的 gorouting 数量, 默认等于 cpu 核数)
    -count: 设置执行测试函数的次数, 默认为 1
    -run: 执行功能测试函数, 支持正则匹配, 可以选择测试函数或者测试文件来仅测试单个函数或者单个文件
    -bench: 执行基准测试函数, 支持正在匹配
    -benchtime: 基准测试最大探索时间上限-------?
    -parallel: 设置同一个被测试代码包中的功能测试函数的最大并发执行数------?
    -v: 是展示测试过程信息

```

### 参考文档
[官方文档](https://golang.google.cn/cmd/go/#hdr-Testing_flags)
[参考1](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter09/09.1.html)
[参考2](https://geektutu.com/post/quick-go-test.html)
[参考3](https://www.cnblogs.com/jssyjam/p/10728801.html)
[参考4](http://c.biancheng.net/view/124.html)