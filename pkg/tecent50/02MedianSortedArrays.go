package main

import (
	"fmt"
)

// 寻找两个有序数组中位数：无需合并数组找到中位数的位置
// 思路：写一个循环，在里面判断是否到了中位数的位置；
// 对于奇数和偶数判断合并，奇数中位数位置=length/2+1，偶数中位数位置=length/2和length/2+1，都遍历length/2+1次，
// left记录上一次遍历的值，right记录最后一次遍历的值；奇数的中位数=right，偶数中位数=(left+right)/2
// todo O(m+n) 解法：每次排除一个
func findMedianSortedArrays(A []int, B []int) float64 {
	m := len(A)
	n := len(B)
	length := m + n
	i := 0 // A的指针
	j := 0 // B的指针
	left := -1
	right := -1
	//k: 合并数组指针
	k := 1
	for k <= length/2+1 {
		// 如果 (A未到最后 && A位置的数字小于B位置的数字)，A数组后移；（相等时A后移或B后移是一样的）
		// 考虑B数组到最后的情况，B数组到最后(j==n，因为B移动后j会+1，所以是j==n而不是n-1)或者A位置数字小于B位置数字(j==n 一定命中前者，因此不会越界)，A后移
		left = right
		if i < m && (j == n || A[i] <= B[j]) {
			fmt.Printf("right: A[%d]=%d\n", i, A[i])
			right = A[i]
			i++
		} else {
			fmt.Printf("right: B[%d]=%d\n", j, B[j])
			right = B[j]
			j++
		}
		k++
	}
	if length%2 == 0 {
		return float64(left+right) / 2
	} else {
		return float64(right)
	}
}

// todo O(log(m+n)) 的二分法：递归和非递归方法，等价于 寻找第K小的数字（K=length/2）和第K+1小的数字；奇数: k=K+1，求两次相同的 K
// 方法：每次排除掉一半的数字，第1~第K/2
func getKth(A []int, start1 int, end1 int, B []int, start2 int, end2 int, k int) {

}

func main() {
	A := []int{2, 3, 5, 8}
	B := []int{2, 3, 3, 8} // 2 2 3 3 3 5 8 8
	//B := []int{2, 3, 7, 9, 10} // 2 2 3 3 5 7 8 9 10
	fmt.Println(findMedianSortedArrays(A, B))
}
