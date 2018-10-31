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

// 思路： 链表总个数，除以k， 得到每个部分的大小
//   10/3 == 3..1  [4, 3, 3]
//   3/5  == 0..3  [1, 1, 1, 0, 0]
//	 3/3  == 1..0
func splitListToParts(root *ListNode, k int) []*ListNode {
	result := []*ListNode{}
	if root == nil {
		for i := 0; i < k; i++ {
			result = append(result, nil)
		}
		return result
	}

	count := 0
	for p := root; p != nil; p = p.Next {
		count += 1
	}

	m := count / k
	n := count % k

	start := root
	for i := 0; i < k; i++ {
		subcount := m
		if n > 0 {
			subcount += 1
		}
		n -= 1

		if subcount == 0 {
			result = append(result, nil)
			continue
		}

		j := subcount - 1
		end := start
		for j > 0 {
			end = end.Next
			j -= 1
		}
		result = append(result, start)
		start = end.Next
		end.Next = nil
	}
	return result
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

	result := splitListToParts(head, 6)
	for index := range result {
		fmt.Printf("part %d\n", index)
		printList(result[index])
	}
}
