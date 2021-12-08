package main

import "fmt"

// 搜索旋转排序数组
// 数组nums在某个下标k处发生了旋转，旋转后数组下标k之前和之后的数字都大于k，从数组中搜索目标值并返回下标，没有返回-1

// 方法1：旋转数组的性质：数组二分后一定有一部分是有序的。初始l和r分别指向数组的开头和末尾，当l<=r时，
//	如果nums[l]<=nums[mid]，且 target在l,mid 之间，让r=mid-1缩小查找范围，否则让l=mid+1在[mid+1,r]中查找；
//  否则如果nums[l]>nums[mid]，则[mid,r]一定有序，且target在[mid+1,r]之间，则让l=mid+1缩小查找范围，否则让r=mid-1在[l,mid-1]中查找。
// case: [1,3], 2; [1,3,5], 2
func search(nums []int, target int) int {
	l, r, n := 0, len(nums)-1, len(nums)
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] { //target < nums[mid]不加等于条件因为等于在循环开始判断。
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] { //nums[mid] < target不加等于条件因为等于在循环开始判断。
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 3))

	nums = []int{1}
	fmt.Println(search(nums, 0))

	nums = []int{1, 3}
	fmt.Println(search(nums, 2))
}
