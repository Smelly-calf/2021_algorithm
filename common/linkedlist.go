package common

import "fmt"

type ListNode struct {
	Num  int
	Next *ListNode
}

func CreateLinkedList(arr []int) *ListNode {
	l := &ListNode{}
	pre := l // pre指向l，每次遍历后pre后移一位
	for i := 0; i < len(arr)-1; i++ {
		pre.Num = arr[i]
		pre.Next = &ListNode{}
		pre = pre.Next
	}
	pre.Num = arr[len(arr)-1]
	return l
}

func TraverseLinkedList(l *ListNode) {
	for l != nil {
		fmt.Println(l.Num)
		l = l.Next
	}
}
