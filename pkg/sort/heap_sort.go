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

// 迭代：从最后一个非叶子节点（设下标为 start）开始调整数组为大根堆，从父节点、左孩子、右孩子找最大的：比较左右孩子中较大的那个赋值给 left，如果父结点<孩子结点，交换父节点和孩子结点的值（元素下沉）。
//	递归调整左右子树中较大的树，当left>=length结束递归。当start<=0，结束循环。
// 调整：交换堆顶和堆尾元素，调整剩余堆为大根堆。
func heapSortWithRecursion(nums []int) {
	var down func(nums []int, i int, end int)
	down = func(nums []int, i int, end int) {
		if i >= end {
			return
		}
		left := 2*i + 1
		// left是孩子结点中最大的那个结点
		if left+1 <= end && nums[left] < nums[left+1] {
			left++
		}
		// 如果父节点小于孩子结点，小的元素下沉，交换父节点和孩子结点
		if i <= end && left <= end && nums[i] < nums[left] {
			swap(nums, i, left)
		}
		down(nums, left, end) // 被交换的孩子结点继续调整直到>=end
	}
	// 从最后一个非叶子节点开始通过小的元素下沉构造大根堆，2i+2==length-1 => i=(length-3)/2=length/2-1
	for i := len(nums)/2 - 1; i >= 0; i-- {
		down(nums, i, len(nums)-1)
	}
	// 固定堆顶元素到末尾
	size := len(nums)
	for size > 1 {
		swap(nums, 0, size-1)
		// 调整剩余堆：堆顶元素不断下沉
		down(nums, 0, size-2)
		size--
	}
}

func swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	a := []int{3, 5, 8, 1, 2, 10} //升序排序
	heapSortWithLoop(a)
	fmt.Printf("迭代：%#v\n", a)

	a = []int{3, 5, 8, 1, 2, 10}
	heapSortWithRecursion(a)
	fmt.Printf("递归：%#v\n", a)
}
