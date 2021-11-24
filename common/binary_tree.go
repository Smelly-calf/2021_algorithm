package common

// TreeNode 二叉树的结点
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
	return recursionCreate(nums)
}

var i = -1

func recursionCreate(nums []int) *TreeNode {
	i++
	if i >= len(nums) {
		return nil
	}
	if nums[i] == -1 {
		return nil
	}
	root := &TreeNode{Val: nums[i]}
	root.Left = recursionCreate(nums)
	root.Right = recursionCreate(nums)
	return root
}

// LevelOrder 二叉树的层序遍历
// 初始化res二维数组的第一维，二维在递归时动态分配，level表示一维数组的下标
// 递归遍历左子树把值append到res[level]指向的数组，并level++；再递归遍历右子树把结点值append到res[level]对应的二维数组，level++；
// 递归结束条件：root==nil{return res}
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	levelOrderFunc(res, root, 0)
	return res
}

func levelOrderFunc(res [][]int, root *TreeNode, level int) {
	if len(res) < level+1 {
		res = append(res, make([]int, 0))
	}
	res[level] = append(res[level], root.Val)
	if root.Left != nil {
		levelOrderFunc(res, root.Left, level+1)
	}
	if root.Right != nil {
		levelOrderFunc(res, root.Right, level+1)
	}
}
