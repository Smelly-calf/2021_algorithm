package main

import "fmt"

// [1,2,3] 数组全排列
// 思路：回溯算法，深度优先遍历的思想
// path保存每次遍历的栈，当path的深度depth等于数组长度，表示一次搜索完成。状态数组used存储path元素状态，需要重置。
// https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
func permute(nums []int) [][]int {
	res := make([][]int, 0) // 存储全排列结果
	n := len(nums)
	path := make([]int, 0, n) //深度优先搜索结果
	used := make([]bool, n)   //初始化数组被搜索的状态为false
	for i := 0; i < n; i++ {
		used[i] = false
	}

	var dfs func(nums []int, n int, depth int, path []int, res [][]int) [][]int
	dfs = func(nums []int, n int, depth int, path []int, res [][]int) [][]int{
		if depth == n {
			// 排列结果添加到res，深拷贝
			c := make([]int, 0, n)
			for _, e := range path {
				c = append(c, e)
			}
			res = append(res, c)
			return res
		}
		// 在还未选择的数中依次选择一个元素作为下一个位置的元素，这显然得通过一个循环实现
		for i, num := range nums {
			if !used[i] {
				path = append(path, num)
				used[i] = true

				res = dfs(nums, n, depth+1, path, res)

				// 回溯
				used[i] = false
				path = path[0:depth]
			}
		}
		return res
	}

	res = dfs(nums, n, 0, path, res)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Printf("%+v", res)

}
