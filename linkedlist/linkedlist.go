package main

type ListNode struct {
	num  int
	next *ListNode
}

func CreateLinkedList(arr []int) *ListNode {
	pre := &ListNode{}
	l := &ListNode{}
	pre = l // pre保存l的head节点
	for i := 0; i < len(arr)-1; i++ {
		l.num = arr[i]
		l.next = &ListNode{}
		l = l.next
	}
	l.num = arr[len(arr)-1]
	return pre
}
