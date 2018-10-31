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
func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil || head.Next.Next == nil { // 0, 1, 2
		return head
	}

	var odd_head *ListNode = head
	var odd_tail *ListNode = head
	var even_head *ListNode = head.Next
	var even_tail *ListNode = head.Next

	index := 2
	for next := even_head.Next; next != nil; {
		node := next
		next = next.Next

		if index%2 == 0 {
			//odd
			odd_tail.Next = node
			odd_tail = node
		} else {
			//even
			even_tail.Next = node
			even_tail = node
		}
		index += 1
	}

	odd_tail.Next = even_head
	even_tail.Next = nil
	return odd_head
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

	printList(oddEvenList(head))
}
