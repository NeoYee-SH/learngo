package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

func op(a, b int, op string) (ret int, err error) {
	switch op {
	case "+":
		ret, err = a+b, nil
	case "-":
		ret, err = a-b, nil
	case "*":
		ret, err = a*b, nil
	case "/":
		ret, err = a/b, nil
	default:
		ret, err = 0, fmt.Errorf("unsupported operator "+op)
	}
	return
}
func apply(op func(a, b int) int, a int, b int) {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("opName = %s, ret = %s\n", opName, strconv.Itoa(op(a, b)))
}

func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}
func main() {
	apply(func(a, b int) int {
		return a + b
	}, 3, 4)

	fmt.Println(sum(1, 3, 5, 7, 8))
}
