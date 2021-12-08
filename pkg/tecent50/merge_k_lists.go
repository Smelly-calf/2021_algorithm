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
// 设共有k个链表，首先遍历所有链表的第一个节点，将(节点，k)入队；初始化虚拟头结点和prev，循环：当队列不为空，pop队列元素加入到prev.Next，pop元素的next元素推入到队列；最后返回虚拟头结点.next。
// 复杂度分析：最开始遍历数组k次 + (循环k*n次) * 每次操作插入和删除时间复杂度O(logK) = O(KN*logK)
// 空间复杂度：优先级队列O(k)
func MergeKLists(lists []*common.ListNode) *common.ListNode {
	p := common.PriorityQueue{}
	heap.Init(&p)
	for i, list := range lists {
		heap.Push(&p, common.NewEntry(list, list.Num, i))
	}
	head := &common.ListNode{}
	prev := head
	for p.Len() > 0 {
		entry := heap.Pop(&p).(common.Entry)
		prev.Next = entry.GetKey()
		next := entry.GetKey().Next
		if next != nil {
			heap.Push(&p, common.NewEntry(next, next.Num, entry.GetIndex()))
		}
		prev = prev.Next
	}
	return head.Next
}

func main() {
	arrs := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	lists := []*common.ListNode{}
	for _, arr := range arrs {
		lists = append(lists, common.CreateLinkedList(arr))
	}
	ret := MergeKLists(lists)
	retArr := common.TraverseLinkedList(ret)
	fmt.Printf("%#v", retArr)
}
