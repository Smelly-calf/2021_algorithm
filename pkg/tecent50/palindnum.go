package main

import "fmt"

// Easy - 回文整数
// 思路：负数肯定不是回文；先反转x，得到rx，如果x与rx是相同的，那么这个数字就是回文数，反转时推入rx时考虑溢出.
// 官方解法：反转一半的x，不用考虑溢出，终止条件是x小于等于反转后的数字，最后判断x=rx（偶数位）或者x=rx/10（奇数位）。
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	rx := 0
	for x > rx { // 100 rx=1 x=0，符合x==rx/10，但不是回文，因此对于末尾是0的需要单独判断，但0本身是回文；102 rx=20 x=1 不符合x==rx/10，因此考虑末尾是0的即可。
		rx = rx*10 + x%10
		x /= 10
	}
	return x == rx || x == rx/10
}

func main() {
	x := 121
	x1 := -121
	x2 := 10
	x3 := -101
	fmt.Println(isPalindrome(x))
	fmt.Println(isPalindrome(x1))
	fmt.Println(isPalindrome(x2))
	fmt.Println(isPalindrome(x3))
}
