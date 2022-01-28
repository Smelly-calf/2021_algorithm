package main

import (
	"2021_algorithm/common"
	"math"
)

// MaxDepthWithDFS 深度优先搜索求解二叉树最大深度
// 递归：二叉树的深度=max(左子树深度，右子树深度)+1，root=null递归结束返回0
/* 3 -> 9 -> -1
          -> -1
	 -> 20 -> 15
		   -> 7
 */
// 递归过程：
// stack1: root=(3), root.left=(9), leftDepth := MaxDepthWithDFS((9))，由stack2得到leftDepth=1，继续递归求解root.right=(20)的深度，也是1，最终得到深度=1+1=2
// stack2: root=(9), root.left=null, leftDepth := MaxDepthWithDFS((null))，得到leftDepth=0；root.right=null，得到rightDepth=0，return max(leftD,rightD)+1=1，代入stack1
// 所以写return语句时考虑表达式右边函数的返回值，不需要考虑递归过程中的中间值存储。
func MaxDepthWithDFS(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepthWithDFS(root.Left)
	rightDepth := MaxDepthWithDFS(root.Right)
	return int(math.Max(float64(leftDepth), float64(rightDepth)) + 1)
}

// MaxDepthWithBFS 广度优先搜索求解二叉树最大深度
// 广度优先搜索的队列里存放的是「当前层的所有节点」，每次拓展下一层的时候，
// 我们需要将队列里的所有节点都拿出来进行拓展，这样能保证每次拓展完的时候队列里存放的是当前层的所有节点，
// 即我们是一层一层地进行拓展，最后我们用一个变量 ans 来维护拓展的次数，该二叉树的最大深度即为 ans。
func MaxDepthWithBFS(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*common.TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}
