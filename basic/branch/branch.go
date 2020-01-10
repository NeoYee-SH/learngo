package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func branch(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

func readFile() {
	const filename = "a.txt"

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

}

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator :" + op)
	}
	return result
}

func grade(score int) string {
	var result string
	switch {
	case score < 0 || score > 100:
		panic("illegal score: " + strconv.Itoa(score))
	case score < 60:
		result = "F"
	case score < 70:
		result = "D"
	case score < 80:
		result = "C"
	case score < 90:
		result = "B"
	case score < 100:
		result = "A"
	}

	return result
}
func main() {
	value := 23
	fmt.Println(branch(value))
	readFile()

	a, b := 12, 4
	fmt.Printf("%d+%d=%d\n", a, b, eval(a, b, "+"))
	fmt.Printf("%d-%d=%d\n", a, b, eval(a, b, "-"))
	fmt.Printf("%d*%d=%d\n", a, b, eval(a, b, "*"))
	fmt.Printf("%d/%d=%d\n", a, b, eval(a, b, "/"))
	//fmt.Printf("%dx%d=%d\n", a, b, eval(a, b, "x"))

	fmt.Println(
		grade(10),
		grade(66),
		grade(71),
		grade(79),
		grade(91),
		grade(59),
	)

}
