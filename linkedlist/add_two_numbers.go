package main

import (
	"fmt"
)

// Medium - 两逆序整数链表相加
// 两个非空链表表示两个非负整数，每位数字按照逆序方式存储，每个节点存一个数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。时间复杂度O(m+n)

// 思路1：将长度较短的链表在末尾补零使得两个连表长度相等，再一个一个元素对其相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var len1, len2 int
	l11 := l1
	l22 := l2
	for l11 != nil {
		len1 += 1
		l11 = l11.next
	}
	for l22 != nil {
		len2 += 1
		l22 = l22.next
	}

	if len1 < len2 { // l1末尾补零
		l11 = l1
		for i := len1; i < len2; i++ {
			l11.next = &ListNode{}
			l11 = l11.next
		}
	} else {
		l22 = l2
		for l22.next != nil { //l22移动到l2的末尾
			l22 = l22.next
		}
		for i := len2; i < len1; i++ {
			l22.next = &ListNode{}
			l22 = l22.next
		}
	}

	// 初始化新的链表和链表头指针pre，遍历链表节点相加
	pre := &ListNode{}
	cur := pre
	var carry int
	for l1 != nil && l2 != nil {
		i := carry + l1.num + l2.num
		cur.next = &ListNode{num: i % 10} //初始化cur.next并赋第一个值
		cur = cur.next

		carry = i / 10 // carry=当前节点数字之和除10
		l1 = l1.next
		l2 = l2.next
	}
	if carry > 0 {
		cur.next = &ListNode{num: carry}
	}
	return pre.next
}

// 思路2: 不对齐补零，sum(代表每个位的和的结果)加上，考虑进位。
func addTwoNums(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	var carry int
	for l1 != nil || l2 != nil {
		var sum int
		if l1 != nil {
			sum += l1.num
			l1 = l1.next
		}
		if l2 != nil {
			sum += l2.num
			l2 = l2.next
		}
		sum += carry
		fmt.Println(sum)
		cur.next = &ListNode{num: sum % 10}
		carry = sum / 10
		cur = cur.next
	}
	if carry > 0 {
		cur.next = &ListNode{num: carry}
	}
	return pre.next
}

func main() {
	arr1 := []int{9, 9, 9, 9, 9, 9, 9}
	arr2 := []int{9, 9, 9, 9}

	l1 := CreateLinkedList(arr1)
	l2 := CreateLinkedList(arr2)
	fmt.Println("create finished")
	l3 := addTwoNums(l1, l2)
	for l3 != nil { // 输出结果: 0742
		fmt.Print(l3.num)
		l3 = l3.next
	}
}
