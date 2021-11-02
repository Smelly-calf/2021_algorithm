package main

import "fmt"

// 最长回文字串
// 解法1：用动态规划寻找两个字符串最长公共子串的方法，寻找最长回文子串，边界条件是 maxI=Len-1-maxJ+arr[maxI][maxJ]-1，maxI表示最长子串末尾下标
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	// 寻找s 和 reverse的最长公共字串
	reverse := reverseS(s)
	fmt.Println("s=", s)
	fmt.Println("reverse=", reverse)
	lens := len(s)
	d := make([][]int, lens)
	for i := 0; i < lens; i++ {
		d[i] = make([]int, lens)
	}
	// d[i][j]保存00到ij的最大公共子串长度
	var maxLen int // maxLen保存最大公共子串长度
	var maxEnd int // maxEnd保存最大公共子串的末尾下标
	for i := 0; i < lens; i++ {
		for j := 0; j < lens; j++ {
			if s[i] == reverse[j] {
				if i == 0 || j == 0 {
					d[i][j] = 1
				} else {
					d[i][j] = d[i-1][j-1] + 1
				}
			}
			if d[i][j] > maxLen {
				// 保存maxLen时判断边界条件
				sf := lens - 1 - j //sf意思为reverse下标j对应的s的下标
				if sf+d[i][j]-1 == i {
					maxLen = d[i][j]
					maxEnd = i
					fmt.Println(maxEnd)
					fmt.Println(maxLen)
				}
			}

		}
	}
	return s[maxEnd-maxLen+1 : maxEnd+1]
}

func reverseS(s string) string {
	lens := len(s)
	reverse := make([]uint8, lens)
	for i := 0; i <= len(s)/2; i++ {
		reverse[i] = s[len(s)-1-i]
		reverse[len(s)-1-i] = s[i]
	}
	return string(reverse)
}

func main() {
	s := "aacdefcaa"
	fmt.Println("最长回文子串:", longestPalindrome(s))
}
