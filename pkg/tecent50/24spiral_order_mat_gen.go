package main

import "fmt"

// 给正整数n，1～n^2元素，按顺时针螺旋生成n*n正方形矩阵
// 例如：n: 3，生成的mat: [[1 2 3] [8 9 4] [7 6 5]]
// 思路，按层模拟，填充数字num：1~n^2
// 终止条件：num<=n^2
func generateMatrixBySpiral(n int) [][]int {
	// 初始化n*n矩阵
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}
	// 上下左右边界
	top, left := 0, 0
	bottom, right := n-1, n-1
	// 填充
	tar := n * n
	num := 1
	for num <= tar {
		// 沿着上边界
		for i := left; i <= right; i++ {
			mat[top][i] = num
			num++
		}
		top++
		// 沿右边界
		for i := top; i <= bottom; i++ {
			mat[i][right] = num
			num++
		}
		right--
		// 沿着下边界
		for i := right; i >= left; i-- {
			mat[bottom][i] = num
			num++
		}
		bottom--
		// 沿着左边界
		for i := bottom; i >= top; i-- {
			mat[i][left] = num
			num++
		}
		left++
	}
	return mat
}

func main() {
	n := 3
	fmt.Printf("n: %d\n mat: %+v\n", n, generateMatrixBySpiral(n))
}
