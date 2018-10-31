package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
** 思路：
**     遍历链表，对于每个节点，将其插入到新链表中
**         如果 p.Val >= x， 插入到新链表的结尾
**         否则，插入到链表中的指定插入点
** 需要：
**     一个newhead 指向新链表
**     一个insertpos 指向新链表的插入点
**     一个newtail 指向新链表的尾部
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newhead *ListNode = head
	var newtail *ListNode = head
	var insertpos *ListNode = nil

	if head.Val < x {
		insertpos = head
	}

	count := 1
	for p := head.Next; p != nil; {
		count += 1
		if p.Val >= x {
			newtail.Next = p
			newtail = p
			p = p.Next
			continue
		}
		// insert this node
		q := p.Next //backup

		if insertpos == nil {
			// insert into head
			p.Next = newhead
			insertpos = p
			newhead = p
		} else {
			// insert into
			if newtail == insertpos {
				newtail = p
			}
			p.Next = insertpos.Next
			insertpos.Next = p
			insertpos = p
		}
		p = q // loop continue
	}
	newtail.Next = nil
	return newhead
}

func printList(head *ListNode) {

	for p := head; p != nil; p = p.Next {
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
	tail := addNode(head, 4)
	tail = addNode(tail, 3)
	tail = addNode(tail, 2)
	tail = addNode(tail, 5)
	tail = addNode(tail, 2)

	printList(partition(head, 3))
}
