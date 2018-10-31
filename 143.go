package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
** 思路：
**     遍历链表，对于第 x 个节点，
**        如果是中间节点，则修改其Next指向前一个节点
**        如果是首节点，则暂存
**        如果是尾节点，则修改其Next指向后，
**
** 需要：
**     一个 tail0，指向第一个链的尾部
**     一个 tmp， 指向首节点
**     一个 head2， 指向第三个链得头部
 */

// prev 指向前一个节点，需要用到，因此，只能在使用之后，才修改 prev
//
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	l := []*ListNode{}

	count := 0
	for p := head; p != nil; p = p.Next {
		count += 1
		l = append(l, p)
	}

	for i := 0; i < count; i++ {
		p := l[i]

		if i < count/2 {
			j := count - i - 1
			p.Next = l[j]
		} else if i > count/2 {
			j := count - i
			p.Next = l[j]
		} else { // i== count/2
			p.Next = nil
		}
	}
}

func printList(head *ListNode) {

	count := 0
	for p := head; p != nil; p = p.Next {
		if count >= 10 {
			break

		}
		count += 1
		fmt.Printf("%d\n", p.Val)
	}
}

func addNode(tail *ListNode, v int) *ListNode {
	n := &ListNode{
		v,
		nil,
	}
	tail.Next = n
	return n
}

func main() {
	head := &ListNode{1, nil}
	tail := addNode(head, 2)
	tail = addNode(tail, 3)
	tail = addNode(tail, 4)
	tail = addNode(tail, 5)

	reorderList(head)
	printList(head)
}
