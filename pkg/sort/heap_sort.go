package main

import "fmt"

// 堆排序性质：数组的升序用大根堆，降序用小根堆。大根堆性质：每个结点的值大于等于左右孩子结点的值，即 arr[i]>=arr[2i+1] and arr[i]>=arr[2i+2].
// 算法：构造大顶堆实现数组升序，（数组的升序）先 元素上升 将待排序数组构造成大根堆；再固定最大值到堆尾，将剩余元素调整为大根堆.
// 插入：遍历数组下标：当index>0时，如果结点index的值大于其父结点的值（奇数父节点:(index-1)/2，偶数父节点:(index-2)/2 ），交换结点index的值和父节点的值，把父节点下标赋值给index，进入下一次循环。
// 调整：设数组长度为size，每次固定一个元素到堆尾，将剩余size-1个元素调整为大根堆，size减1，进入下一次循环。
// 循环复杂度分析：时间复杂度：插入：西格玛(i=1~N)每次循环比较log(k)次*循环次数k ~ O(Nlog(N))；调整N-1次，每次调整时间复杂度=O(Nlog(N))，总时间复杂度=O(NlogN)。空间复杂度：不需要额外空间O(1).
func heapSortWithLoop(a []int) {
	// 堆元素上升
	up := func(a []int, s int, e int) {
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
		}
	}
	// 构造大根堆
	up(a, 0, len(a)-1)
	// 固定最大值(堆顶元素)到最后一个元素，调整剩余元素为大根堆，调整范围(size)减1
	size := len(a)
	for size > 1 {
		swap(a, 0, size-1)
		up(a, 0, size-2)
		size--
	}
}

func swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	a := []int{3, 5, 8, 1, 2, 10} //升序排序
	heapSortWithLoop(a)
	fmt.Println(a)
}
