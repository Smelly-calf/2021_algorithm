package main

import "2021_algorithm/common"

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/solution/er-cha-shu-de-zui-da-shen-du-by-leetcode-solution/

// LevelOrderWithDfs 二叉树的层序遍历之深度优先搜索
// 初始化res二维数组的第一维，二维在递归时动态分配，level表示层。
// 深度优先：
// 递归遍历左子树把值append到res[level]指向的数组，接着遍历level+1；
// 再递归遍历右子树把结点值append到res[level]对应的二维数组；
// 递归结束条件是root为空。
func LevelOrderWithDfs(root *common.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	// dfsHelper 一个最简单的深度优先搜索，注意函数是有返回值的
	var dfsHelper func(res [][]int, root *common.TreeNode, level int) [][]int
	dfsHelper = func(res [][]int, root *common.TreeNode, level int) [][]int {
		if root == nil {
			return res
		}
		if len(res) < level+1 {
			res = append(res, make([]int, 0))
		}
		res[level] = append(res[level], root.Val)
		res = dfsHelper(res, root.Left, level+1)  //每次递归需要保留本次递归的res值，否则的话回溯会最终返回最初的空数组
		res = dfsHelper(res, root.Right, level+1) //每次递归需要保留本次递归的res值，否则的话回溯会最终返回最初的空数组
		return res
	}
	res = dfsHelper(res, root, 0)
	return res
}

// TODO LevelOrderWithBFS 层序遍历的广度优先搜索，从根结点开始，每次将一层的所有节点加入队列中。
// 当遍历完某个节点左右孩子，需要回到它的父节点，遍历父节点右孩子的左右孩子，直到回到根结点，再遍历根结点右孩子。
func LevelOrderWithBFS(root *common.TreeNode) [][]int {
	return nil
}
