package main

import (
	"fmt"
	"log"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init()  {
	fmt.Println("init")
}
func main() {
	//log.SetOutput(os.Stdout)

	p := message.NewPrinter(language.English)
	a := 12346543
	b := float64(a) / 1000
	fmt.Println(p.Sprintf("%.2f", b))

	log.Println("Log……")
}
