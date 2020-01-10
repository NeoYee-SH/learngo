package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//函数类定义的变量作用域在函数内， 函数外也可以定义变量，是包内部变量，但是不可以短赋值
//var aa = 33
//var sb = "stupid"

//也可以使用括号简写
var (
	aa = 33
	sb = "stupid"
)

//go语言的变量定义变量名在前，类型在后； 变量具有初值
func variableZeroValue() {
	var a int
	var s string
	fmt.Println("variableZeroValue: ", a, s)
}

//变量定义时也可以赋值
func variableInitialValue() {
	var a, b int = 3, 4
	var s, ss string = "Hello", "World"

	fmt.Println("variableInitialValue: ", a, b, s, ss)
}

//变量类型可以推断
func variableTypeDeduction() {
	var a, b, s, ss = 3, true, "Hello", "World"
	fmt.Println("variableTypeDeduction: ", a, b, s, ss)
}

//在函数中变量短赋值
func variableShorter() {
	a, b, s, ss := 3, true, "Hello", "World"
	a = 4
	fmt.Println("variableShorter: ", a, b, s, ss)
}

//欧拉公式
//e^{ix} = cos x+i*sin x
func euler() {
	ret := cmplx.Exp(math.Pi*1i) + 1
	fmt.Printf("%.3f\n", ret)
}

func triangle() {
	a, b := 3, 4
	var c int
	//c = math.Sqrt(a*a + b*b)
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	c := math.Sqrt(a*a + b*b)

	fmt.Println(filename, c)

}

const (
	cpp = iota
	_
	python
	golang
	javascript
)

func enums() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)

}
func main() {
	//用fmt包的Println函数打印出Hello World
	fmt.Println("main: Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println("package var: ", aa, sb)
	euler()
	triangle()
	consts()
	enums()
}
