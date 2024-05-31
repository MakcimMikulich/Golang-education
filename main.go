package main

import "fmt"

func main() {
	first := &ListNode{}
	second := &ListNode{}

	first.Val = 1
	second.Val = 2

	first.Next.Val = 3
	second.Next.Val = 4

	// first.Next.Next.Val = 5
	// second.Next.Next.Val = 6

	fmt.Println(first, second)

	// fmt.Println(addTwoNumbers([]int{1, 2, 3, 5, 6}, 5))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//    func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

//    }
