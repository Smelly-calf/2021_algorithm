package main

import (
	"fmt"
	"math"
	"time"
)

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
/*
示例：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/

// 思路：这道题和全排列不同，不是寻求全部答案，而是寻求最优解。而动态规划就是用来寻求最优解的算法。
/*
动态规划和回溯算法的区别：
动态规划：求最优解。
回溯算法：搜索得到所有的方案。
*/

/*
maxSubArrayUseDynamicProgramming 动态规划
动态规划转移方程：f(i)=max{f(i−1)+nums[i],nums[i]}，其中f(i-1)表示以第i-1个数结尾的「连续子数组的最大和」，0<=i<=n-1.
 时间复杂度 O(n)、空间复杂度 O(n)
*/
func maxSubArrayUseDynamicProgramming(nums []int) int {
	pre, maxAns := 0, 0
	for _, x := range nums {
		pre = int(math.Max(float64(pre+x), float64(x)))       // 以i结尾的连续子数组最大和
		maxAns = int(math.Max(float64(maxAns), float64(pre))) // 全局最大和
	}
	return maxAns
}

/*
maxSubArrayWithLineSegmentTree 线段树，它不仅可以解决区间 [0, n-1][0,n−1]，还可以用于解决任意的子区间 [l,r][l,r] 的问题
1. 线段树的作用：时间复杂度不如动态规划，但其优势在于大查询。
	如果我们把 [0, n-1][0,n−1] 分治下去出现的所有子区间的信息都用堆式存储的方式记忆化下来，
	即建成一颗真正的树之后，我们就可以在 O(logn) 的时间内求到任意区间内的答案，
	我们甚至可以修改序列中的值，做一些简单的维护，之后仍然可以在 O(logn) 的时间内求到任意区间内的答案，
	对于大规模查询的情况下，这种方法的优势便体现了出来。这棵树就是上文提及的一种神奇的数据结构——线段树。

2. 解：分治思想。这个分治方法类似于「线段树求解最长公共上升子序列问题」的 pushUp 操作。
 我们定义一个操作get(a,l,r)求解a序列[l,r]区间内的最大子段和，那么最终我们要求的答案就是 get(nums, 0, nums.size() - 1)。
 如何分治实现？对区间[l,r]的左子区间[l,m]和右子区间[m+1,r]分治求解。当递归逐层深入直到区间长度缩小为1的时候，
 递归「开始回升」。这个时候我们考虑如何通过[l,m]区间的信息与[m+1,r]区间的信息合并成区间[l,r]的信息。

 对于一个区间 [l,r]，我们可以维护四个量：
	lSum：[l,r]以l为左端点的最大子段和。
    rSum：[l,r]以r为右端点的最大子段和。
    mSum：[l,r]区间的最大子段和。
	iSum：[l,r]的区间和。
 对于长度为1的区间[i,i]，四个变量的值都是nums[i]；
 对于长度大于1的区间：
	1）iSum很简单，[l,r]的iSum等于左子区间iSum + 右子区间的iSum；
	2) [l,r]的lSum分两种情况，要么等于左子区间的lSum，要么等于左子区间的iSum+右子区间的lSum，两者取最大；
    3) [l,r]的rSum分两种情况，要么等于右子区间的rSum，要么等于左子区间的rSum+右子区间的iSum，两者取最大；
    4) [l,r]的mSum分三种情况，左子区间的mSum，右子区间的mSum，左子区间的rSum+右子区间的lSum，三者取最大。
*/
func maxSubArrayWithLineSegmentTree(nums []int) int {
	type Status struct {
		iSum int
		lSum int
		rSum int
		mSum int
	}
	max := func(a int, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	pushUp := func(l Status, r Status) Status {
		iSum := l.iSum + r.iSum
		lSum := max(l.lSum, l.iSum+r.lSum)
		rSum := max(r.rSum, l.rSum+r.iSum)
		mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
		return Status{iSum, lSum, rSum, mSum}
	}

	var get func(a []int, l int, r int) Status
	get = func(a []int, l int, r int) Status {
		// 区间长度为1
		if l == r {
			return Status{a[l], a[l], a[l], a[l]}
		}
		// 区间长度大于1
		m := (l + r) >> 1
		lSub := get(a, l, m)
		rSub := get(a, m+1, r)
		return pushUp(lSub, rSub)
	}

	return get(nums, 0, len(nums)-1).mSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("动态规划：")
	start := time.Now()
	maxSubArrayUseDynamicProgramming(nums)
	fmt.Printf("耗时：%v, %d\n", time.Since(start), len(nums))

	fmt.Println("线段树：")
	start = time.Now()
	maxSubArrayWithLineSegmentTree(nums)
	fmt.Printf("耗时：%v, %d\n", time.Since(start), len(nums))
}
