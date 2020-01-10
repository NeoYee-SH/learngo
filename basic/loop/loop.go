package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loop() {
	sum := 0
	for i := 1; i <= 10; i++ {
		fmt.Print(i)
		if i != 10 {
			fmt.Print("+")
		} else {
			fmt.Print("=")
		}
		sum += i
	}
	fmt.Print(sum)
}

func convertToBin(v int) string {
	if v == 0 {
		return "0"
	}
	var result string
	for ; v > 0; v /= 2 {
		lsb := v % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Print("aaa")
	}
}

func main() {
	//loop()
	fmt.Println(
		convertToBin(0),
		convertToBin(10),
		convertToBin(13),
		convertToBin(33333),
		convertToBin(13243),
	)

	printFile("/Users/yihuaiyuan/golib/src/github.com/yihuaiyuan/learngo/basic/loop/a.txt")
}
