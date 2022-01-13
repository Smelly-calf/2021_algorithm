package main

import (
	"fmt"
	"sort"
)

// Medium - 找出n个整数数组所有的三元组，三元组数字之和为零，不能有重复的三元组

// 思路：遍历数组，第一层循环固定第一个元素，第二层循环固定第二个元素，第三层循环找到与第一个元素和第二个元素之和的相反数。
// 正确解法：排序+双指针法，首先三重循环结果中会有重复三元组，怎么去重？答：排序；而排序好的数组找三元组能否降低时间复杂度？答：可以，双指针法。
/*
分析：先把数组正序排序（O(NlogN)），考虑第二重循环，固定a和b，那么满足a+b+c=0的c是唯一的。
	当我们向右移动b时，由于b向右移动得到的b1>b（=b的情况已经跳过），为了满足a+b1+c1=0，一定有c1<c，因此随着b向右移动，c一定向左移动，
	因此第三重循环会变成一个从数组最右端开始向左移动的指针，第三重循环结束条件就是a+b+c=0或者j=third。

	当我们需要枚举数组中的两个元素时，如果随着第一个元素递增，第二个元素是递减的，就可以使用双指针法，将枚举的时间复杂度从 O(N^2) 减少至 O(N)。
	为什么说是O(N)呢？因为在每一步枚举的过程中，左指针b会向右移动一个位置，右指针c会向左移动若干个位置，但右指针在整个循环结束一共会移动的位置数为O(N)，因此均摊下来每一步枚举c移动一个位置。
	因此第二重和第三重循环总共的时间复杂度为O(N)。

	算法总时间复杂度：寻找三元组的时间复杂度为O(N^2)，排序时间复杂度为O(NlogN)，总共的时间复杂度为O(N^2)。
*/

func threeNum(nums []int) [][]int {
	// 对nums排序
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	for i := 0; i < len(nums); i++ { // 每次遍历之前跳过nums[i]和nums[i-1]重复的情况，i-1>=0
		for i > 0 && nums[i] == nums[i-1] {
			i++
		}
		a := nums[i]
		third := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			// 每次遍历j之前跳过nums[j]和nums[j-1]重复的情况，由于只有唯一的c让a+b+c=0，所以跳过a，b重复的情况，c一定不重复。
			for j > i+1 && nums[j] == nums[j-1] {
				j++
			}
			b := nums[j]
			if third > j { // 当third=j结束第三重循环，避免j和third重合
				for third-1 > j && a+b+nums[third] > 0 { // third-1>j是为了third--后不与j相等
					third--
				}
				c := nums[third]
				if a+b+c == 0 {
					res = append(res, []int{a, b, c})
				}
			}

		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4} //[[-1,-1,2],[-1,0,1]]
	nums2 := []int{}
	nums3 := []int{0}
	fmt.Printf("%+v\n", threeNum(nums))
	fmt.Printf("%+v\n", threeNum(nums2))
	fmt.Printf("%+v\n", threeNum(nums3))
}
