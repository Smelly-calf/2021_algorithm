package main

import "fmt"

// 快排：记下基准的下标mid和值tmp，双指针法i和j初始指向数组开头和末尾；当i<j时，如果j的值>=tmp，向左移动j，交换i和j的值；接着向右移动i，直到i的值>tmp，交换i和j的值；
//	直到i和j相遇循环结束，tmp=arr[mid]，分别递归排序mid左边的数组和右边的数组，当开头和结尾距离=0时结束递归。
func quickSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	mid, i := start, start
	tmp := arr[mid]
	j := end
	for i < j {
		for arr[j] >= tmp && i < j {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
		for arr[i] <= tmp && i < j {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	mid = i
	tmp = arr[mid]
	quickSort(arr, start, mid-1)
	quickSort(arr, mid+1, end)
}

func main() {
	nums := []int{0, 1, 2, 5, 7, 3, 3, 4}
	quickSort(nums, 0, len(nums)-1)
	fmt.Printf("%#v", nums)
}
