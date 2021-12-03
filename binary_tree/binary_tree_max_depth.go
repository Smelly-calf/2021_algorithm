package main

import (
	"2021_algorithm/common"
	"fmt"
)

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

func main() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	tree := common.CreateBinaryTree(nums)
	fmt.Println("二叉树最大深度：", MaxDepthWithBFS(tree))
}
