package main

import (
	"2021_algorithm/common"
	"fmt"
)

// Medium - 两逆序整数链表相加
// 两个非空链表表示两个非负整数，每位数字按照逆序方式存储，每个节点存一个数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。时间复杂度O(m+n)

// 思路1：将长度较短的链表在末尾补零使得两个连表长度相等，再一个一个元素对其相加
func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	var len1, len2 int
	l11 := l1
	l22 := l2
	for l11 != nil {
		len1 += 1
		l11 = l11.Next
	}
	for l22 != nil {
		len2 += 1
		l22 = l22.Next
	}

	if len1 < len2 { // l1末尾补零
		l11 = l1
		for i := len1; i < len2; i++ {
			l11.Next = &common.ListNode{}
			l11 = l11.Next
		}
	} else {
		l22 = l2
		for l22.Next != nil { //l22移动到l2的末尾
			l22 = l22.Next
		}
		for i := len2; i < len1; i++ {
			l22.Next = &common.ListNode{}
			l22 = l22.Next
		}
	}

	// 初始化新的链表和链表头指针pre，遍历链表节点相加
	pre := &common.ListNode{}
	cur := pre
	var carry int
	for l1 != nil && l2 != nil {
		i := carry + l1.Val + l2.Val
		cur.Next = &common.ListNode{Val: i % 10} //初始化cur.next并赋第一个值
		cur = cur.Next

		carry = i / 10 // carry=当前节点数字之和除10
		l1 = l1.Next
		l2 = l2.Next
	}
	if carry > 0 {
		cur.Next = &common.ListNode{Val: carry}
	}
	return pre.Next
}

// 思路2: 不对齐补零，sum(代表每个位的和的结果)加上，考虑进位。
func addTwoNums(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	pre := &common.ListNode{}
	cur := pre
	var carry int
	for l1 != nil || l2 != nil {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		fmt.Println(sum)
		cur.Next = &common.ListNode{Val: sum % 10}
		carry = sum / 10
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &common.ListNode{Val: carry}
	}
	return pre.Next
}

func main() {
	arr1 := []int{9, 9, 9, 9, 9, 9, 9}
	arr2 := []int{9, 9, 9, 9}

	l1 := common.CreateLinkedList(arr1)
	l2 := common.CreateLinkedList(arr2)
	fmt.Println("create finished")
	l3 := addTwoNums(l1, l2)
	for l3 != nil { // 输出结果: 0742
		fmt.Print(l3.Val)
		l3 = l3.Next
	}
}
