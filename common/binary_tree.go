package common

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

