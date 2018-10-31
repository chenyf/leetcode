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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m == n {
		return head
	}

	var tail0 *ListNode = nil
	var prev *ListNode = nil
	var newhead *ListNode = head
	var tmp *ListNode = nil

	index := 1
	for next := head; next != nil; {

		current := next
		next = next.Next

		fmt.Printf("handle %d\n", current.Val)

		if index < m {
			index += 1
			prev = current
			continue
		}

		if index == m {
			// found head
			tail0 = prev
			tmp = current
			index += 1
			prev = current
			continue
		}
		// update tail0
		if index == n {
			//found end
			// 先将悬空的节点链接起来
			tmp.Next = current.Next

			// 再让第一个链指向我
			if tail0 != nil {
				tail0.Next = current
			} else {
				newhead = current
			}

			// 最后找到我的下一个
			current.Next = prev
			break
		} else {
			current.Next = prev
			prev = current
			index += 1
		}
	}
	return newhead
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

	printList(reverseBetween(head, 3, 4))
}
