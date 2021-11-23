package main

import "2021_algorithm/common"

// Medium - 链表两两结点交换
/* 解法：递归的两两交换链表中的结点，
用 head 表示原始链表的头节点，新的链表的第二个节点，用 newHead 表示新的链表的头节点，原始链表的第二个节点，则原始链表中的第三个结点是 newHead.next。
令 head.next = swapPairs(newHead.next)，表示将第三个结点到最后一个结点进行两两交换，交换后的新的头节点为 head 的下一个节点。
然后令 newHead.next = head，即完成了所有节点的交换。最后返回新的链表的头节点 newHead。

递归 head=[1 2 3 4]
step1: head=[1 2 3 4] newHead=[2 3 4]
step2: head=[3 4] newHead=[4] newHead.next=nil
step3: head=nil，返回nil

递归调用上面的部分是递归，递归调用下面的部分是回溯。

回溯:
step2: swapPairs结果nil, head=[3], newHead.next=[3], newHead=[4 3] return 返回头结点4
step1: 上一步的值赋值给head.next，head=[1 4 3]，newHead.next=[1 4 3]，newHead=[2 1 4 3]，return 头结点2
*/
func swapPairs(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}


