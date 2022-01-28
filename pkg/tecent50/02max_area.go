package main

import "fmt"

// 题： Med - 盛最多水的容器
// 	给 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai)，在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为(i,ai) 和 (i, 0) 。
//	找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
/**
思路：双指针法，设两指针i和j，指向的数字分别为h[i]和h[j]，此状态下水槽面积为S(i,j)。可容纳水的高度由短板决定，所以S(i,j)=min(h[i],h[j])*(j-i)。
在每个状态下，无论向内移动长板还是短板，水槽底边都变小；
	向内移动短板，水槽的短板 min(h[i],h[j]) 可能变大，因此下一个水槽面积可能变大；
	向内移动长板，min(h[i],h[j])变小或不变，因此下一个水槽面积一定变大。
因此初始状态两指针分别指向左右两端，循环每轮短板向内移动一格，更新面积最大值，直到两指针相遇跳出循环，即可得到面积最大值。

正确性证明：水槽两板围成面积 S(i, j)S(i,j) 的状态总数为 C(n, 2)。
假设状态 S(i,j) 下 h[i]<=h[j]，在向内移动短板至 S(i+1,j)，则相当于消去了以 i 为短板的状态集合：S(i,j-1),S(i,j-2),...,S(i,i+1)。而所有消去状态的面积都一定小于S(i,j)，因为这些状态：
	- 短板高度：min(h[i],h[j1])小于等于min(h[i],h[j])，因为如果h[j1]<=h[i]则短板高度小于等于h[i]，如果h[j1]>h[i]，则短板高度等于h[i]。
	- 底边宽度相比S(i,j)状态一定变短。
所以消去的状态一定不会导致面积最大值丢失。证毕。
*/

// 提示:
// 	n == height.length
// 	2 <= n <= 10^5; 0 <= height[i] <= 10^4

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var res int
	for i < j {
		if height[i] <= height[j] {
			area := height[i] * (j - i)
			if area > res {
				res = area
			}
			i++
		} else {
			area := height[j] * (j - i)
			if area > res {
				res = area
			}
			j--
		}
	}
	return res
}
func main() {
	// [1,8,6,2,5,4,8,3,7] => 49
	// [1,1] => 1，不包含y轴
	// [4,3,2,1,4] => 16
	// [1,2,1] => 2
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	h1 := []int{1, 1}
	h2 := []int{4, 3, 2, 1, 4}
	h3 := []int{1, 2, 1}
	fmt.Println(maxArea(h))
	fmt.Println(maxArea(h1))
	fmt.Println(maxArea(h2))
	fmt.Println(maxArea(h3))
}
