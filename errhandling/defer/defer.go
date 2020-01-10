package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yihuaiyuan/learngo/funcs/fib"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func tryRecover() {
	defer func() {
		r := recover()

		if err, ok := r.(error); ok {
			fmt.Println("Error occurred : ", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do for %v", r))
		}
	}()

	//error := errors.New("this is error")
	error := "123"
	panic(error)
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	//err = errors.New("this is a custom error")
	if err != nil {
		if pathErr, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathErr.Op,
				pathErr.Path,
				pathErr.Err)
		}
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
func main() {
	//for i := 0; i < 100; i++ {
	//	defer fmt.Println(i) //9 8 7 6 5 4 3 2 1 0
	//	if i == 9 {
	//		panic("printed too many")
	//	}
	//}

	//writeFile("ab.txt")

	tryRecover()
}
