package main

import "fmt"

// 最长回文子串
// 解法1：用动态规划寻找两个字符串最长公共子串的方法，还需要判断该字符串倒置后下标和倒置前是否匹配，只用比较末尾下标maxJ和倒置前是否一样就好。
// 已知maxI表示最长子串的末尾下标，maxJ倒置前下标加最长子串长度应该也是子串的末尾，应该等于maxI，即满足 maxI = Len-1-maxJ+arr[maxI][maxJ]-1 就说明是回文串。
// 至于为什么只比较末尾，先这么理解吧：如果不是回文串，末尾下标一定不匹配。
// case: caba，abc435cba，babad，cbbd，a，ac
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
	var maxLen int // maxLen 保存最大公共子串长度
	var maxEnd int // maxEnd 保存最大公共子串的末尾下标
	for i := 0; i < lens; i++ {
		for j := 0; j < lens; j++ {
			if s[i] == reverse[j] { // 相等的时候才加1
				if i == 0 || j == 0 {
					d[i][j] = 1
				} else {
					d[i][j] = d[i-1][j-1] + 1
				}
			}
			if d[i][j] > maxLen {
				// 保存maxLen时判断边界条件: Len-1-maxJ+arr[maxI][maxJ]-1=maxI, sf=Len-1-maxJ 表示maxJ反转前的下标，sf+arr[maxI][maxJ]-1表示maxJ反转前下标加上最长公共子串的长度对应的下标，
				// 等于maxI则表示最长公共子串就是要找的回文子串.
				sf := lens - 1 - j //sf意思为reverse下标j对应的s的下标
				if sf+d[i][j]-1 == i {
					maxLen = d[i][j]
					maxEnd = i
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
	s := "caba"
	fmt.Println("最长回文子串:", longestPalindrome(s))
}
