package main

import "2021_algorithm/common"

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/

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

// LevelOrderWithBFS 广度优先搜索层序遍历二叉树，借助两个队列
// 方法一：哈希表，key为level，value为level对应节点值组成的数组
// 方法二：
// 存储：需要两个优先级队列q和p，分别存储第k层元素S_k(元素按照从左到右顺序排列好的)和第k+1层元素S_k+1
// 算法：首先根元素入队列，当队列不为空: 求q的长度S_k，依次从队列中取S_k个元素进行拓展：即把S_k的元素加入到ret[k]，并且把S_k每个元素的左右孩子节点加入到p，最后把p赋值给q，进入下一次迭代，直到队列为空返回ret
func LevelOrderWithBFS(root *common.TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	q := []*common.TreeNode{root} // q表示第一层
	k := 0 //层下标
	for len(q) > 0 { // 当叶子结点层拓展完成后，q为空，表示广度优先搜索完成
		ret = append(ret, []int{}) //初始化
		var p []*common.TreeNode   //存储拓展后元素
		// 从q拓展下一层元素，p存储拓展后元素，一层结点均拓展完成把p赋值给q
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[k] = append(ret[k], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
		k++   // 继续拓展下一层
	}
	return ret
}
