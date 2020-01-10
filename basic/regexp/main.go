package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is.niceliar0316@gmail.com!
email1 is dfjkg_kf@qq.com.cn
email2 is dfjkgkf.cfdf@163.com
email3 is.4fdfd33@qq.com.cn
email3 is @ali.com
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
