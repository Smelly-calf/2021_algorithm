package main

import "fmt"

// https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/

// [1,2,3] 数组全排列
// 思路：回溯算法，思路类型于深度优先搜索
/* 1. 深度优先搜索 算法（英语：Depth-First-Search，DFS）是一种用于遍历或搜索树或图的算法。
	1)深度遍历：这个算法会 尽可能深 的搜索树的分支。
	2)回溯：当结点 v 的所在边都己被探寻过，搜索将 回溯 到发现结点 v 的那条边的起始结点。这一过程一直进行到已发现从源结点（不是结点v）可达的所有结点为止。
	如果还存在未被发现的结点，则选择其中一个作为源结点并重复以上过程，整个进程反复进行直到所有结点都被访问为止。
   2. 回溯算法与深度优先搜索不同点：
	1） 需要使用不断变化的变量存储
    2） 布尔数组标记数字是否被选择过。回溯时重置被选择过的数字的状态。
*/
/*
「全排列」问题的树形结构：选择和撤销
			             		[]
		|选择1|撤销1  			|选择2|撤销2   				|选择3|撤销3
 		[1]          	 		[2]           				[3]
	|选择2|撤销2  |选择3|撤销3     |选择1|撤销1  |选择3|撤销3
  [1 2]        [1 3]		  [2 1]        [2 3]
   |选择3|撤销3   |选择2|撤销2	   |选择3|撤销3   |选择1|撤销1
[1 2 3]        [1 3 2]		  [2 1 3]        [2 3 1]
 */
// 编程：路径path保存每次遍历结果，写入res。状态数组used存储正在遍历路径每个结点的状态。
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

				res = dfs(nums, n, depth+1, path, res) // 寻找到路径 0->1->2，i位于路径的尽头的结点2
				used[i] = false // 回溯：重置结点2状态
				path = path[0:depth] // 紧接着回溯到结点1，发现2还未被搜索，继续从1开始遍历直到所有结点都被访问。回溯结点0，从1为源结点开始遍历。回溯结点1，从2为源结点开始遍历。
			}
		}
		return res
		// 当结点 v 的所在边都己被探寻过，搜索将 回溯 到发现结点 v 的那条边的起始结点。
		// 这一过程一直进行到已发现从源结点可达的所有结点为止。如果还存在未被发现的结点，则选择其中一个作为源结点并重复以上过程，整个进程反复进行直到所有结点都被访问为止
	}

	res = dfs(nums, n, 0, path, res)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Printf("%+v", res)
}
