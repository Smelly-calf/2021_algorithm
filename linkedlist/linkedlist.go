package main

import "fmt"

type ListNode struct {
	num  int
	next *ListNode
}

func CreateLinkedList(arr []int) *ListNode {
	l := &ListNode{}
	pre := l // pre指向l，每次遍历后pre后移一位
	for i := 0; i < len(arr)-1; i++ {
		pre.num = arr[i]
		pre.next = &ListNode{}
		pre = pre.next
	}
	pre.num = arr[len(arr)-1]
	return l
}

func TraverseLinkedList(l *ListNode) {
	for l != nil {
		fmt.Println(l.num)
		l = l.next
	}
}
