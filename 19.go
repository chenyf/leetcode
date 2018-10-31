package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	a := []*ListNode{}
	next := head
	index := 0
	for next != nil {
		a[index] = next
		next = next.Next
		index += 1
	}
	target := len(a) - n
	a[target-1].Next = a[target].Next
	return head
}

func main() {

}
