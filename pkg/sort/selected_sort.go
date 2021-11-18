package main

import "fmt"

// 选择排序: 将数组R 划分为有序和无序两个区域，一开始无序区域为 R[0...N-1]，有序区域为空；
// 第一次排序假设索引0为最小元素，R[1...N-1]的每个元素与第0个元素比较，找到最小元素的索引，让索引0与最小索引的元素进行交换；（每次循环只交换一次）
// 第i次排序数组 R[0...i]为有序区域，从i+1到N-1找到比i小的最小元素的索引，交换i与最小索引对应元素；i遍历到 i=n-2 即可，因为第n-1个元素一定是最大的。
// 时间复杂度：双层循环 O(n^2)；排序后相等元素的相对顺序会发生改变，不稳定的算法。

func selectSort(r []int) {
	for i := 0; i < len(r)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(r); j++ {
			if r[j] < r[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			tmp := r[i]
			r[i] = r[minIndex]
			r[minIndex] = tmp
		}
	}
}

func main() {
	a := []int{10, 8, 9, 57, 28, 39, 4, 100}
	selectSort(a)
	fmt.Println(a)
}
