package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(arr []int) *ListNode {
	l := &ListNode{}
	pre := l // pre指向l，每次遍历后pre后移一位
	for i := 0; i < len(arr)-1; i++ {
		pre.Val = arr[i]
		pre.Next = &ListNode{}
		pre = pre.Next
	}
	pre.Val = arr[len(arr)-1]
	return l
}

func TraverseLinkedList(l *ListNode) []int {
	arr := []int{}
	for l != nil {
		arr = append(arr, l.Val)
		l = l.Next
	}
	return arr
}

// ReverseLinkedList 链表反转
func ReverseLinkedList(head *ListNode) *ListNode{
	var newHead *ListNode
	next := head.Next
	for head != nil {
		head.Next = newHead
		newHead = head
		head = next
		if next != nil {
			next = next.Next
		}
	}
	return newHead
}
