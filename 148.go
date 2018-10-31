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
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l := []*ListNode{}
	for p := head; p != nil; p = p.Next {
		l = append(l, p)
	}

	count := len(l)
	_qslist(l, 0, count-1)

	for i := 0; i < count; i++ {
		if i+1 < count {
			l[i].Next = l[i+1]
		} else {
			l[i].Next = nil
		}
	}
	return l[0]
}

func _qslist(l []*ListNode, start int, end int) {
	if start >= end {
		return
	}

	// 对 l 进行分区， q 是分区点
	q := _partition_list(l, start, end)

	// 对分区左侧进行递归
	_qslist(l, start, q-1)
	// 对分区右侧进行递归
	_qslist(l, q+1, end)
}

func _partition_list(l []*ListNode, p int, r int) int {
	pivot := l[r].Val
	i := p
	for j := p; j < r; j++ {
		if l[j].Val < pivot {
			//swap
			if i < j {
				t := l[i]
				l[i] = l[j]
				l[j] = t
			}
			i += 1
		}
	}

	//swap
	t := l[i]
	l[i] = l[r]
	l[r] = t
	return i
}

func mergesort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head1 := head
	head2 := getmiddlenode(head)

	head1 = mergesort(head1)
	head2 = mergesort(head2)
	return merge(head1, head2)
}

func getmiddlenode(head *ListNode) *ListNode {
	fast := head.Next
	slow := head.Next
	prev := head
	for {
		if fast == nil {
			break
		}
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next

		prev = slow
		slow = slow.Next
	}
	prev.Next = nil
	return slow
}

func merge(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	newhead := &ListNode{-1, nil}
	newtail := newhead

	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			newtail.Next = head1
			head1 = head1.Next
		} else {
			newtail.Next = head2
			head2 = head2.Next
		}
		newtail = newtail.Next
		newtail.Next = nil
	}
	if head1 != nil {
		newtail.Next = head1
	}
	if head2 != nil {
		newtail.Next = head2
	}
	return newhead.Next
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
	tail := addNode(head, 5)
	tail = addNode(tail, 4)
	tail = addNode(tail, 3)
	tail = addNode(tail, 2)

	//printList(sortList(head))
	printList(mergesort(head))

}
