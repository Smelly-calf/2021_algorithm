package main

import "fmt"

// day10(1.25) 螺旋矩阵（顺时针螺旋遍历矩阵）
/*
方法一：模拟
可以模拟螺旋矩阵的路径。初始位置是矩阵的左上角，初始方向是向右，当路径超出界限或者进入之前访问过的位置时，顺时针旋转，进入下一个方向。
判断路径是否进入之前访问过的位置需要使用一个与输入矩阵大小相同的辅助矩阵 visited，其中的每个元素表示该位置是否被访问过。
当一个元素被访问时，将 visited 中的对应位置的元素设为已访问。
如何判断路径是否结束？由于矩阵中的每个元素都被访问一次，所以order元素个数等于矩阵元素个数时，表示访问结束。
*/
/*
时间复杂度：遍历次数total=m*n，时间复杂度是 O(mn)
空间复杂度：visited额外占用m*n空间，空间复杂度O(mn)
*/
func spiralOrder(matrix [][]int) []int {
	// 定义order数组存储遍历后元素，整体思想就是遍历矩阵，把元素不断追加到order数组。但是添加的顺序有一定规则：按照顺时针螺旋顺序。因此遍历不是按照矩阵行列。
	// 遍历结束条件是order数组元素等于矩阵元素，所以循环条件设定为矩阵元素总和total。
	// 按照顺时针螺旋顺序遍历：向右->向下->向左->向上->向右->再次进入螺旋，
	//	当螺旋需要变换方向时，需要 1) 确定何时变换方向；2）确定变换何种方向；
	//	第1）点，预计算nextRow和nextColumn，最外层螺旋边界条件，超出行和列范围；内层螺旋边界：需要定义与矩阵相同大小的visited辅助矩阵，其中每个元素表示该位置是否被访问过；当遍历时发现某个元素被访问过，说明到达边界了。
	//  第2）点，需要定义directions二维数组，它包含4个元素，每个元素表示行和列的方向，分别是{0 1}表示向右，{1 0}表示向下，{0, -1}表示向左，{-1, 0}表示向上。directionIndex标记变换方向次数，当超过4次需要和4取余。
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m := len(matrix)
	n := len(matrix[0])
	total := m * n
	order := make([]int, 0, total)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	row, column := 0, 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	directionIndex := 0
	for i := 0; i < total; i++ {
		order = append(order, matrix[row][column])
		visited[row][column] = true
		// 预计算nextRow和nextColumn
		nextRow := directions[directionIndex][0] + row
		nextColumn := directions[directionIndex][1] + column
		// 判断是否到达变换方向时机，最外层边界条件：超出行和列范围；内层螺旋边界条件：visited对应元素是否被访问过。
		if nextRow < 0 || nextRow > m-1 || nextColumn < 0 || nextColumn > n-1 || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4 //变换方向次数加1
		}
		// 实际的下一次的row和column
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}

/*
方法二：按层模拟
把矩阵分层，先遍历最外层，再遍历次外层。
定义上下左右边界：top, bottom, left, right
1）沿着上边界从left到right遍历顶点，top++
2）如果上下边界交错，返回结果集；再沿着右边界从top到bottom遍历顶点，right--
3）如果左右边界交错，返回结果集；沿着下边界从right到left遍历顶点，bottom--
4）如果上下边界交错，返回结果集；沿着左边界从bottom到top，left++
沿着左右边界时顶点为matrix[i][当前边界]
沿着上下边界时顶点为matrix[当前边界][i]
*/
func spiralOrderByLevel(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m := len(matrix)
	n := len(matrix[0])
	order := make([]int, 0, m*n)
	// 定义上下左右边界
	top := 0
	bottom := m - 1
	left := 0
	right := n - 1
	// 螺旋
	// 1 2 3
	// 4 5 6
	// 7 8 9
	for true {
		// 沿着上边界从左向右遍历
		for i := left; i <= right; i++ {
			order = append(order, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}
		// 沿着右边界从上到下
		for i := top; i <= bottom; i++ {
			order = append(order, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		// 沿着下边界从右向左
		for i := right; i >= left; i-- {
			order = append(order, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}
		// 沿着左边界从下向上
		for i := bottom; i >= top; i-- {
			order = append(order, matrix[i][left])
		}
		left++
		if right < left {
			break
		}
	}
	return order
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	order := spiralOrder(matrix)
	fmt.Printf("matrix: %+v\n order: %+v\n", matrix, order)

	order = spiralOrderByLevel(matrix)
	fmt.Printf("matrix: %+v\n order: %+v\n", matrix, order)
}
