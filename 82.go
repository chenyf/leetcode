package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	m := map[int]int{}

	for p := head; p != nil; {
		_, ok := m[p.Val]
		if ok {
			m[p.Val] += 1
		} else {
			m[p.Val] = 1
		}
		p = p.Next
	}

	var newhead *ListNode = nil
	var prev *ListNode = nil

	for p := head; p != nil; {
		if m[p.Val] > 1 {
			p = p.Next
			continue
		}
		// found valid node
		if newhead == nil {
			newhead = p
			prev = newhead
		} else {
			prev.Next = p
			prev = p
		}
		p = p.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	return newhead
}

func main() {
	fmt.Println(deleteDuplicates(nil))
	fmt.Println(deleteDuplicates(&ListNode{2, nil}))
}
