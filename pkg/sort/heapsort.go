package main

import "fmt"

// 堆排序：数组的升序排序用大根堆，降序排序用小根堆，大根堆性质：每个结点的值大于等于左右孩子结点的值；arr[i]>=arr[2i+1]&&arr[i]>=arr[2i+2]
// 大根堆排序：先 元素上升 将待排序数组构造成大根堆，再固定最大值（元素下沉），将剩余元素调整为大根堆；堆排序主要的方法：使用元素上升构造大顶堆；次要方法：固定堆顶元素到当前堆最后一个元素，剩余元素重新构造大根堆。
// 堆排序是基于完全二叉树（每个非叶子结点左右孩子都存在）的选择排序：从i开始，与2i+1与2i+2选择最大索引，与父结点i交换；再以2i+1开始构造大顶堆。
// 堆排序时间复杂度：不稳定的排序，时间复杂度

// 堆排序可以是非递归，也可以是递归

// 非递归
func heapSort(a []int) {
	// 构造大根堆
	heapInsert(a, 0, len(a)-1)
	// 固定最大值(堆顶元素)到最后一个元素，调整剩余元素为大根堆，调整范围(size)减1
	size := len(a)
	for size > 1 {
		swap(a, 0, size-1)
		heapInsert(a, 0, size-2)
		size--
	}
}

// 核心：构造大根堆，从数组第2*0+1=1个元素开始，每个待排序元素与父结点比较，叶子节点大的上升至堆顶；每插入一个新的元素与父结点比较直到堆顶。
func heapInsert(a []int, s int, e int) {
	for i := s; i <= e; i++ {
		index := i
		fmt.Println("index: ", index)
		// 2parent+1=i => parent=(i-1)/2; 2parent+2=i => parent=(i-2)/2
		for index > 0 {
			var parent int
			if index%2 != 0 { // 奇数下标的父结点
				parent = (index - 1) / 2
			} else { //偶数下标的父结点
				parent = (index - 2) / 2
			}
			// 如果新插入元素大于父结点, 交换新插入结点的值和父结点的值；index上升至parent，直到 index 为0
			if a[index] > a[parent] {
				swap(a, index, parent)
			}
			index = parent
		}
		fmt.Println(a)
	}
}

func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func main() {
	a := []int{3, 5, 8, 1, 2, 10}
	heapSort(a)
	fmt.Println(a)
}
