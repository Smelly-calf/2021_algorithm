package main

import (
	"2021_algorithm/common"
	"fmt"
)

// 给你一个链表的头结点 head，旋转链表，将链表每个结点向右移动 k 个位置。
// 1. 旋转开始的位置：与头结点距离为k的结点，需要先使用 pre 遍历链表定位k的位置
// 2. 从第k个结点把链表断开，pre表示k-n结点组成的链表，然后把 pre 链表反转（单链表反转）
// 3. 把反转后链表的尾部与head拼接，完成旋转。
func rotateList(head *common.ListNode, k int) *common.ListNode {
	pre := head
	for i := 0; i < k; i++ {
		// 链表长度小于k，直接返回
		if pre == nil {
			return head
		}
		pre = pre.Next
	}
	if pre == nil {
		return head
	}
	// 反转pre，分别用newHead、next表示新链表头结点、原始链表的next结点
	var newHead *common.ListNode
	next := pre.Next
	for pre != nil {
		pre.Next = newHead
		newHead = pre
		pre = next
		if next != nil {
			next = next.Next
		}
	}
	// newHead尾部指向head
	return newHead
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := common.CreateLinkedList(nums)
	newHead := rotateList(head, 3)
	fmt.Printf("%+v", common.TraverseLinkedList(newHead))
}
