package main

import (
	"fmt"
	"math"
	"sort"
)

// Medium - 最接近目标的三数之和
// 解法：排序+双指针，相似题目：三数之和为0，盛最多水的容器（选择丢弃左边界或右边界减少枚举范围）
// 提示：3 <= nums.length <= 1000; -1000 <= nums[i] <= 1000; -10^4 <= target <= 10^4

// 思路：题目改为求与target距离最近的三数之和，排序+双指针：初始化距离dist为 nums[0]+nums[1]+nums[len-1]与target距离，每次固定a和b，将c左移x位直到 新的a+b+c与target距离小于dist，更新dist，继续b+1，最后返回三数之和。
// 复杂度：O(N^2)

/*
精选解法：用每次遍历抛弃一个边界的做法：先固定a，b和c看作左右两个边界，b++或者c--比较a+b+c与target大小，如果和大于target，c左移抛弃右边的c'；如果和小于target，b右移抛弃左边的b'，直到b和c相邻：
	当a+b+c>target时，让c向左移动1，因为b向右只会让a+b+c变大从而远离target，优化：c向左移动x次直到和c-1不等；
	当a+b+c<target时，让b向右移动1，a+b+c变大从而更接近target，优化：b向右移动x次直到和b-1不等；
	当a+b+c=target时，直接返回a+b+c的值。
复杂度：第一重遍历N，第二重遍历N，O(N^2)
1.res初始应该设为什么值？res初始为1e7的原因是：
2.res的值应该在每次移动b和c之前更新，这样可以把初始的值也纳入进来
*/

func threeSumClosest(nums []int, target int) int {
	// 对nums排序
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := int(1e7)
	// 更新res
	update := func(sum int) {
		//如果sum和target距离< res和target距离，更新res
		if math.Abs(float64(sum-target)) < math.Abs(float64(res-target)) {
			res = sum
		}
	}

	for i := 0; i < len(nums); i++ {
		// i=1开始跳过nums[i]和nums[i-1]重复的情况
		if i > 0 && nums[i] == nums[i-1] {
			continue // for{i++}和if{continue}的区别：if{continue}不用再写一次for循环并且还需要判断边界i<len(nums)等可能导致溢出的情况
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			// update res
			update(sum)
			if sum > target {
				// k向左移动1，优化：k移动到下一个与c（nums[k-1]）不相同的数
				for j < k-1 && nums[k-1] == nums[k] {
					k--
				}
				k--
			} else {
				// j向右移动1，优化：j移动到下一个与b（nums[j-1]）不相同的数
				for j < k-1 && nums[j-1] == nums[j] {
					j++
				}
				j++
			}
		}
	}
	return res
}
func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
	nums1 := []int{0, 0, 0}
	target1 := 1
	fmt.Println(threeSumClosest(nums1, target1))
}
