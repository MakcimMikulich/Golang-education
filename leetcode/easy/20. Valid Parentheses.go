package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() {
	if s.Empty() {
		return
	}

	s.items = s.items[:len(s.items)-1]

}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

// func main() {
// 	fmt.Println(isValid("()[]{}"))
// }

func isValid(s string) bool {
	m := map[int]int{40: 41, 91: 93, 123: 125}

	stack := &Stack{}

	for _, v := range s {

		fmt.Println(int(v))

		if len(stack.items) == 0 {
			stack.Push(int(v))
			continue
		}

		lastItem := stack.items[len(stack.items)-1]

		if m[lastItem] == int(v) {
			stack.Pop()
		} else {
			stack.Push(int(v))

		}

		fmt.Println(stack, m)

	}

	return stack.Empty()
}
