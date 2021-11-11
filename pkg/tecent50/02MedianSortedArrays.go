package main

import (
	"fmt"
	"math"
)

// Difficult - 寻找两个有序数组中位数：无需合并数组找到中位数的位置

// 解法1：写一个循环，在里面判断是否到了中位数的位置；时间复杂度O(m+n)
// 对于奇数和偶数判断合并，奇数中位数位置=length/2+1，偶数中位数位置=length/2和length/2+1，都遍历length/2+1次，
// left记录上一次遍历的值，right记录最后一次遍历的值；奇数的中位数=right，偶数中位数=(left+right)/2
func findMedianSortedArraysOmn(A []int, B []int) float64 {
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
		// 如果 (A未到最后 && A位置的数字小于B位置的数字)，A数组后移；（相等时A后移 或 B后移是一样的）
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

// 解法2 O(log(m+n)) 的二分法
// 递归寻找第K（K=length/2）小数
// 比较两个数组的第 k/2 个数字（如果k是奇数向下取整），防止数组长度小于 k/2，所以每次比较 min(k/2，len(数组) 对应的数字，
// 把小的那个对应的数组的数字排除，将两个新数组进入递归，并且 k 要减去排除的数字的个数。
// 递归出口就是当 k=1 或者其中一个数字长度是 0 了。
func findMedianSortedArraysOlogmn(A []int, B []int) float64 {
	m := len(A)
	n := len(B)
	length := m + n
	k1 := (m + n + 1) / 2 // 偶数的中位数: len/2==(len+1)/2,len/2+1；奇数的中位数: (len+1)/2 == (len+2)/2
	k2 := (m + n + 2) / 2
	fmt.Println("k1:", k1)
	fmt.Println("k2:", k2)
	if length%2 == 0 {
		return (getKth(A, 0, m-1, B, 0, n-1, k1) + getKth(A, 0, m-1, B, 0, n-1, k2)) * 0.5
	}
	return getKth(A, 0, m-1, B, 0, n-1, k1)
}

func getKth(A []int, start1 int, end1 int, B []int, start2 int, end2 int, k int) float64 {
	// 寻找第k小数字, 比较A[k/2]和B[k/2]，排除掉 k/2 个数字
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	if len1 > len2 {
		// 如果len1>len2，交换A和B，保证如果len==0一定是len1
		return getKth(B, start2, end2, A, start1, end1, k)
	}
	if len1 == 0 { // len==0 代表 中位数在B
		return float64(B[start2+k-1])
	}
	if k == 1 { // k==1 代表什么
		// ? 为什么是 start1 和 start2
		return math.Min(float64(A[start1]), float64(B[start2]))
	}

	i := start1 + int(math.Min(float64(len1), float64(k/2))) - 1
	j := start2 + int(math.Min(float64(len2), float64(k/2))) - 1
	if A[i] > B[j] {
		// 排除掉 B 位置j之前的数字，包括j，即 j+1; 并且 k 减去 B排除掉的数字个数（j-start2+1）
		return getKth(A, start1, end1, B, j+1, end2, k-(j-start2+1))
	}
	// 排除掉 A位置i之前的数字，包括i，即 i+1；并且 k 减去 A排除掉的数字个数
	return getKth(A, i+1, end1, B, start2, end2, k-(i-start1+1))
}

func main() {
	A := []int{2, 3, 5, 8}
	//B := []int{2, 3, 3, 8} // 2 2 3 3 3 5 8 8
	B := []int{2, 3, 7, 9, 10} // 2 2 3 3 5 7 8 9 10
	fmt.Println("== 解法1")
	fmt.Println(findMedianSortedArraysOmn(A, B))
	fmt.Println("== 二分法")
	fmt.Println(findMedianSortedArraysOlogmn(A, B))
}
