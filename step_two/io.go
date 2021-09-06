package step_two

import (
	"bufio"
	"bytes"
	"fmt"
)

func TestIO() {
	// 参考：http://c.biancheng.net/view/5569.html

	testRead()
	testWrite()
}

func testRead() {
	data := []byte("hello,world")

	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println("read:", string(buf[:n]), n, err)

	rd1 := bytes.NewReader(data)
	rByte := bufio.NewReader(rd1)
	c, err := rByte.ReadByte()
	fmt.Println("read byte:", string(c), err)

	rd2 := bytes.NewReader(data)
	rBytes := bufio.NewReader(rd2)
	b, err := rBytes.ReadBytes(',')
	fmt.Println("read byte:", string(b), err)

	dataReadLine := []byte("Golang is a beautiful language. \n I like it!")
	rdLine := bytes.NewReader(dataReadLine)
	rLine := bufio.NewReader(rdLine)
	l, p, err := rLine.ReadLine()
	fmt.Println("read line:", string(l), p, err)
	l1, p1, err := rLine.ReadLine()
	fmt.Println("read line:", string(l1), p1, err)
	l2, p2, err := rLine.ReadLine()
	fmt.Println("read line:", string(l2), p2, err)

	dataRune := []byte("你好")
	rdRune := bytes.NewReader(dataRune)
	rRune := bufio.NewReader(rdRune)
	ch, s, err := rRune.ReadRune()
	fmt.Println("read rune:", string(ch), s, err)
	//r.ReadSlice()
	//r.ReadString()
	//r.UnreadByte()
	//r.UnreadRune()
	//r.Buffered()

	dataPick := []byte("hello, world")
	rdPick := bytes.NewReader(dataPick)
	rPick := bufio.NewReader(rdPick)
	bl, err := rPick.Peek(5)
	fmt.Println("pick:", string(bl), err)
	bl1, err := rPick.Peek(90)
	fmt.Println("pick:", string(bl1), err)
	fmt.Println("pick:", string(bl), err)
}

func testWrite() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	fmt.Println("可用缓冲区：", w.Available())
	w.Write([]byte("hello"))
	fmt.Println("可用缓冲区：", w.Available())

	//w.Buffered()
	//w.Flush()
	// ...

}
