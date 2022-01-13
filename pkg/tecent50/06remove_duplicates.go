package main

import "fmt"

// 题目：原地删除数组中重复项
// 数组nums，通过修改数组元素原地删除重复项，输出新数组长度。

// 方法1：让k=1，迭代数组下标i，每次迭代移动k直到k不等于k-1，把k-1的值赋值给i，k+1（因为下一次迭代比较k-1和k的值），并且i+1进入下一次迭代，直到k遍历到数组末尾提前跳出迭代，或者i=length-1，返回i+1。
// 时间复杂度分析：总共遍历了N次，平均每次迭代比较O(1)，时间复杂度O(N).
// [0,0,1,1,1,2,2,3,3,4]
func removeDuplicates(nums []int) int {
	k := 1
	i := 0
	for ; i < len(nums); i++ {
		for k < len(nums) && nums[k] == nums[k-1] {
			k++
		}
		nums[i] = nums[k-1]
		// 提前结束循环
		if k == len(nums) {
			return i+1
		}
		k++
	}
	return i+1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	newLength := removeDuplicates(nums)
	fmt.Printf("%+v", nums[:newLength])
}
