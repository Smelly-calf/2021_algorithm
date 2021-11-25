package binary_tree

// TreeNode 二叉树的结点
// 二叉树的增删改，遍历，遍历包括前序、中序、后序、层序，遍历方法两种：深度优先搜索dfs、广度优先搜索bfs
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CreateBinaryTree 递归创建一颗二叉树
// arr[0]为根结点，i=0，根结点左孩子为arr[2i+1]构成的二叉树，根结点右孩子为arr[2i+2]构成的二叉树
// 递归函数：先递归创建左子树：根结点.left=CreateBinaryTree(arr[2i+1:len]);i+=1, 再递归创建右子树：根结点.right=CreateBinaryTree(arr[2i+2:len])，i+=1
// 终止条件: i=len时，返回结点
func CreateBinaryTree(nums []int) *TreeNode {
	return recursionCreate(nums, 0)
}

func recursionCreate(nums []int, i int) *TreeNode {
	if i >= len(nums) {
		return nil
	}
	if nums[i] == -1 {
		return nil
	}
	root := &TreeNode{Val: nums[i]}
	root.Left = recursionCreate(nums, 2*i+1)
	root.Right = recursionCreate(nums, 2*i+2)
	return root
}

// LevelOrderWithDfs 二叉树的层序遍历之深度优先搜索
// 初始化res二维数组的第一维，二维在递归时动态分配，level表示层。
// 深度优先：
// 递归遍历左子树把值append到res[level]指向的数组，接着遍历level+1；
// 再递归遍历右子树把结点值append到res[level]对应的二维数组；
// 递归结束条件是root为空。
func LevelOrderWithDfs(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	// dfsHelper 一个最简单的深度优先搜索，注意函数是有返回值的
	var dfsHelper func(res [][]int, root *TreeNode, level int) [][]int
	dfsHelper = func(res [][]int, root *TreeNode, level int) [][]int {
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
