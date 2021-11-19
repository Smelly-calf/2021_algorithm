package main

// 合并两个有序数组
// 题目：将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 思路：类似于寻找两个有序数组中位数，不过这里是两个链表。
// 最原始的思路：初始化虚拟头结点，prev指向虚拟头结点，同时遍历链表l1和l2，比较两个链表头结点值：较小的头结点插入prev结点，prev后移，循环结束条件是其中一个链表头节点为null。

// 递归解法：
// 递归：
// 1. 什么是递归？函数在运行时调用自己，这个函数就叫做递归函数，调用的过程叫做递归。
// 2. 什么构成了递归？1）终止条件；2）递归函数。
// 总结：递归是函数先不断调用自身，遇到终止条件后回溯，最终返回答案。

/*
递归函数: 合并两个链表中较小头结点与剩下元素的 merge 操作结果合并。
{ list1[0]+merge(list1[1:],list2) list1[0]<list[2]
  list2[0]+merge(list1,list2[1:]) otherwise }
终止条件: 如果两个链表有一个为空，递归结束。
*/

// 容易犯的错：递归写成了循环；递归没有return语句
/*
详解递归过程：递归调用过程不断去掉l1或者l2的头结点，直到l1或者l2为空；
递归：mergeTwoLists(l1, l2)
	l1=[1,2,4]
	l2=[1,3,4]
	第一次: l1=1, l2=1，不满足l1<l2，l2.next=merge(l1,l2.next) 去掉l2的头结点进入下一次递归: l1=[1,2,4]，l2=[3,4]
	第二次: l1=1, l2=3，满足l1<l2，l1.next=merge(l1.next, l2) 去掉l1的头结点进入下一次递归: l1=[2,4], l2=[3,4]
	第三次: l1=2, l2=3，l1<l2，l1.next=merge(l1.next, l2) => l1=[4]，l2[3,4]
	第四次：l1=4, l2=3，不满足l1<l2，l2.next=merge(l1, l2.next) => l1=[4]，l2[4]
	第五次：l1=4, l2=4，不满足l1<l2, l2.next=merge(l1, l2.next) => l1=4, l2=null
	l2==null，返回l1
回溯：
	第一次
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil && l2 != nil {
		if l1.num < l2.num {
			// l1.next=merge(l1.next, l2)
			remainEleMerged := mergeTwoLists(l1.next, l2)
			l1.next = remainEleMerged
			return l1
		} else {
			// l2.next=merge(l1, l2[1:])
			remainEleMerged := mergeTwoLists(l1, l2.next)
			l2.next = remainEleMerged
			return l2
		}
	}
	if l1 != nil {
		return l1
	} else {
		return l2
	}
}
