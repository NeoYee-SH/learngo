package main

import "fmt"

type Element interface {
}
type Stack struct {
	list [10]Element //容量
	top  int         //栈顶
}

func (s *Stack) empty() bool {
	return s.top == 0
}

func (s *Stack) push(v Element) {
	s.list[s.top] = v
	s.top += 1
}

func (s *Stack) pop() (Element, bool) {
	if s.empty() {
		return 0, false
	}
	s.top -= 1
	return s.list[s.top], true
}
func main() {
	var s Stack

	fmt.Println(s.empty())
	s.push(1)
	s.push(2)
	s.push(3)
	ret, ok := s.pop()
	fmt.Println(ret, ok)
	fmt.Println(s.empty())
	ret, ok = s.pop()
	fmt.Println(ret, ok)
	ret, ok = s.pop()
	fmt.Println(ret, ok)
	fmt.Println(s.empty())
}

/*
output :

true
3 true
false
2 true
1 true
true
*/
