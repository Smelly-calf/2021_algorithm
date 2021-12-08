package main

import "2021_algorithm/common"

// Easy - 合并两个有序链表
// 题目：将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 递归解法：
// 递归：
// 1. 什么是递归？函数在运行时调用自己，这个函数就叫做递归函数，调用的过程叫做递归。
// 2. 什么构成了递归？1）终止条件；2）递归函数。
// 总结：递归是函数先不断调用自身，遇到终止条件后回溯，最终返回答案。

// 容易犯的错：递归写成了循环；递归没有return语句
// 写递归需要注意：写结束条件时考虑递归结束，写return语句时考虑回溯。

/*
递归函数: 合并两个链表中较小头结点与剩下元素的merge操作的结果。
{ list1[0]+merge(list1[1:],list2) list1[0]<list[2]
  list2[0]+merge(list1,list2[1:]) otherwise }
终止条件: 如果两个链表有一个为空，递归结束。
*/

/*
详解递归过程：递归调用过程不断去掉l1或者l2的头结点，直到l1或者l2为空；
递归：mergeTwoLists(l1, l2)
	step1: l1=[1,2,4]，l2=[1,3,4]，l1头结点的值(=1)=l2头结点的值(=1)，不满足l1<l2，调用 l2.Next=merge(l1, l2.Next)
	step2: l1=[1,2,4]，l2=[3,4]， 调用 l1.Next=merge(l1.Next, l2)
	step3: l1=[2,4]，l2=[3,4]， 调用 l1.Next=merge(l1.Next, l2)
	step4：l1=[4]，l2[3,4]， 调用 l2.Next=merge(l1, l2.Next)
	step5：l1=[4]，l2[4]，调用 l2.Next=merge(l1, l2.Next)
	step6：l1=[4]，l2=null，返回[4]

回溯：
	不断将调用结果代入到递归函数，直到回溯到递归最开始。
	1.递归的step6 返回了[4]
	2：代入step5，得到l2.Next=[4]，l2=[4,4]，返回[4,4]
	3：代入step4，得到l2.Next=[4,4]，此时l2指向3，返回 l2=[3,4,4]
	4：代入step3，得到l1.Next=[3,4,4]，此时l1指向2，返回 l1=[2,3,4,4]
	5：代入step2，得到l1.Next=[2,3,4,4]，此时l1指向1，返回 l1=[1,2,3,4,4]
	6：代入step1，得到 l2.Next=[1,2,3,4,4]，此时l2指向1，返回 l2=[1,1,2,3,4,4]
	函数最终返回结果就是l2=[1,1,2,3,4,4]。
*/

func mergeTwoLists(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	if l1 != nil && l2 != nil {
		if l1.Num < l2.Num {
			// l1.Next=merge(l1.Next, l2)
			remainEleMerged := mergeTwoLists(l1.Next, l2)
			l1.Next = remainEleMerged
			return l1
		} else {
			// l2.Next=merge(l1, l2[1:])
			remainEleMerged := mergeTwoLists(l1, l2.Next)
			l2.Next = remainEleMerged
			return l2
		}
	}
	if l1 != nil {
		return l1
	} else {
		return l2
	}
}

// 循环解法：初始化虚拟头结点，prev指向虚拟头结点，同时遍历链表l1和l2，比较两个链表头结点值：较小的头结点插入prev结点，prev后移，循环结束条件是其中一个链表头节点为null。最后返回虚拟头结点.next
func mergeTwoListsLoop(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	l3 := &common.ListNode{} //设l3的地址是1，保留链表初始的地址
	prev := l3               // prev此时的地址是1，在l3地址的基础上开始向后移动，但从l3出发可以遍历全部链表
	for l1 != nil && l2 != nil {
		if l1.Num < l2.Num {
			prev.Next = &common.ListNode{l1.Num, &common.ListNode{}} //必须是给 prev.Next 赋值，如果直接赋值给 prev 会丢失与 l3 的关联。
			l1 = l1.Next                                             //l1向后移动，丢弃头结点
		} else {
			prev.Next = &common.ListNode{l2.Num, &common.ListNode{}}
			l2 = l2.Next //l2向后移动，丢弃头结点
		}
		prev = prev.Next //prev地址移动到2
	}
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}
	return l3.Next
}

