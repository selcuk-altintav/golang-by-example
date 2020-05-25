/**
 * @author Selçuk Altıntav
 * @email s.altintav@gmail.com
 * @create date 2020-05-25 02:02:52
 * @modify date 2020-05-25 02:02:52
 * @desc [description]
 */

/**
 * Question #2
 * Link : https://leetcode.com/problems/add-two-numbers/
 */

package main

import "fmt"

// ListNode node for linkedlist
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var firstArray = []int{2, 4, 3}
	var secondArray = []int{5, 6, 4}

	var l1 *ListNode = constructLinkedNumber(firstArray)
	var l2 *ListNode = constructLinkedNumber(secondArray)

	printLinkedList(l1)
	printLinkedList(l2)

	printLinkedList(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	for looper, sum := head, 0; l1 != nil || l2 != nil || sum > 0; {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		looper.Next = &ListNode{
			Val: sum % 10,
		}
		looper = looper.Next
		sum /= 10
	}
	return head.Next

}

func constructLinkedNumber(vals []int) *ListNode {
	var head *ListNode
	for i := len(vals) - 1; i >= 0; i-- {
		head = insertElementToLinkedList(head, vals[i])
	}
	return head
}

func insertElementToLinkedList(head *ListNode, val int) *ListNode {
	if head == nil {
		return createNodePointer(nil, val)
	}
	return createNodePointer(head, val)
}

func createNodePointer(head *ListNode, val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: head,
	}
}

func printLinkedList(list *ListNode) {
	for list.Next != nil {
		fmt.Printf("%d -> ", list.Val)
		list = list.Next
	}
	fmt.Printf("%d -> \\\n", list.Val)
}
