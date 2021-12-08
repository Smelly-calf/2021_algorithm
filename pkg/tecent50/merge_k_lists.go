package main

import (
	"2021_algorithm/common"
	"container/heap"
	"fmt"
)

// 合并K个链表
// 总体有两种方法：1.优先级队列(用heap堆实现的)，2.k个链表两两合并

// 1.优先级队列
// 需要定义一个优先级队列: 在java中实现Comparable接口定义它的compareTo方法；go语言有标准容器库 "container/heap"，用来实现优先级队列
// 设共有k个链表，首先遍历所有链表的第一个节点，将(节点，k)入队；初始化虚拟头结点和prev，循环：当队列不为空，pop队列元素加入到prev.Next，并每次拓展链表的一个元素到优先级队列，直到队列为空，返回虚拟头结点.next。
// 复杂度分析：最开始遍历数组k次 + (循环k*n次) * 每次操作插入和删除时间复杂度O(logK) = O(KN*logK)
// 空间复杂度：优先级队列O(k)
func MergeKLists(lists []*common.ListNode) *common.ListNode {
	p := common.PriorityQueue{}
	heap.Init(&p)
	for i, list := range lists {
		heap.Push(&p, common.NewEntry(list, list.Val, i))
	}
	head := &common.ListNode{}
	prev := head
	for p.Len() > 0 {
		entry := heap.Pop(&p).(common.Entry)
		prev.Next = entry.GetKey()
		next := entry.GetKey().Next
		if next != nil {
			heap.Push(&p, common.NewEntry(next, next.Val, entry.GetIndex()))
		}
		prev = prev.Next
	}
	return head.Next
}

// 2. 链表两两合并
func MergeKListsWithMergeTwoLists(lists []*common.ListNode) *common.ListNode {
	mergeTwoLists := func(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
		l3 := &common.ListNode{} //设l3的地址是1，保留链表初始的地址
		prev := l3               // prev此时的地址是1，在l3地址的基础上开始向后移动，但从l3出发可以遍历全部链表
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				prev.Next = &common.ListNode{l1.Val, &common.ListNode{}} //必须是给 prev.Next 赋值，如果直接赋值给 prev 会丢失与 l3 的关联。
				l1 = l1.Next                                             //l1向后移动，丢弃头结点
			} else {
				prev.Next = &common.ListNode{l2.Val, &common.ListNode{}}
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

	// 递归：二分法：l和r分别指向数组的开头和结尾，计算中间值mid，合并后链表=MergeTwoLists(merge(mid左边数组,merge(mid右边数组))，当l=r结束递归返回lists[l]
	var merge func(lists []*common.ListNode, l int, r int) *common.ListNode
	merge = func(lists []*common.ListNode, l int, r int) *common.ListNode {
		if l == r {
			return lists[l]
		}
		mid := (l + r) / 2
		return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
	}

	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func main() {
	arrs := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}

	lists := []*common.ListNode{}
	for _, arr := range arrs {
		lists = append(lists, common.CreateLinkedList(arr))
	}
	ret := MergeKLists(lists)
	retArr := common.TraverseLinkedList(ret)
	fmt.Printf("优先级队列：%#v\n", retArr)

	lists2 := []*common.ListNode{}
	for _, arr := range arrs {
		lists2 = append(lists2, common.CreateLinkedList(arr))
	}
	ret2 := MergeKListsWithMergeTwoLists(lists2)
	retArr2 := common.TraverseLinkedList(ret2)
	fmt.Printf("分治：%#v\n", retArr2)
}
