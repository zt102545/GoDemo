package main

import "fmt"

type ListNode struct{
	val int
	next *ListNode
}


func ReverseList(node *ListNode) *ListNode {
	if node.next == nil {
		return node
	} else {
		var newHead = ReverseList(node.next)
		node.next.next = node
		node.next = nil
		return newHead
	}
}

func main() {
	head := new(ListNode)
	head.val = 1
	node := new(ListNode)
	node.val = 2
	head.next = node
	node2 := new(ListNode)
	node2.val = 3
	node.next = node2

	var newNode = ReverseList(head)

	for {
		if newNode.next != nil {
			fmt.Print(newNode.val)
			newNode = newNode.next
		} else {
			fmt.Print(newNode.val)
			break
		}
	}
}