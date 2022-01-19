package common

import "container/heap"

// 优先级队列，其实现原理是小顶堆
// This example demonstrates a priority queue built using the heap interface.
// Entry 是 priorityQueue 中的元素
type Entry struct {
	key      *ListNode
	priority int
	// index 是 entry 在 heap 中的索引号
	// entry 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 entry.priority 一直不变的话，可以删除 index
	index int
}

func NewEntry(key *ListNode, p int, index int) Entry {
	return Entry{key, p, index}
}

func (e *Entry) GetKey() *ListNode {
	return e.key
}

func (e *Entry) GetIndex() int {
	return e.index
}

// PQ implements heap.Interface and holds entries.
type PriorityQueue []Entry

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 entry
func (pq *PriorityQueue) Push(x interface{}) {
	temp := x.(Entry)
	temp.index = len(*pq)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PriorityQueue) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	temp.index = -1 // for safety
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

// update modifies the priority and value of an entry in the queue.
func (pq *PriorityQueue) update(entry *Entry, value *ListNode, priority int) {
	entry.key = value
	entry.priority = priority
	heap.Fix(pq, entry.index)
}
