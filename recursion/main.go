package main

import (
	"2021_algorithm/common"
	"2021_algorithm/common/binary_tree"
	"fmt"
)

// https://leetcode-cn.com/problems/merge-two-sorted-lists/solution/yi-kan-jiu-hui-yi-xie-jiu-fei-xiang-jie-di-gui-by-/
func hanotaTest() {
	A, B, C := []int{2, 1, 0}, []int{}, []int{}
	hanota(A, B, C)
	//A = []int{1, 0}
}

func mergeTwoListsTest() {
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	l1 := common.CreateLinkedList(arr1)
	l2 := common.CreateLinkedList(arr2)
	common.TraverseLinkedList(mergeTwoLists(l1, l2))
}

func swapPairsTest() {
	a := []int{1, 2, 3, 4}
	l := common.CreateLinkedList(a)
	newL := swapPairs(l)
	common.TraverseLinkedList(newL)
}

func createBinaryTreeTest() {
	tree := binary_tree.CreateBinaryTree([]int{3, 9, 20, -1, -1, 15, 7})
	levelRes := binary_tree.LevelOrderWithDfs(tree)
	fmt.Printf("%+v", levelRes)
}

func main() {
	//hanotaTest()
	//mergeTwoListsTest()
	//swapPairsTest()
	createBinaryTreeTest()
}
