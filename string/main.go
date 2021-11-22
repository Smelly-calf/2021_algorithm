package main

import "fmt"

func isValidTest() {
	//s := "()"
	s1 := "()[]{}"
	//s2 := "(]"
	//s3 := "([)]"
	//s4 := "{[]}"
	//fmt.Println(s, " is ", myIsValid(s))
	fmt.Println(s1, " is ", myIsValid(s1))
	//fmt.Println(s2, " is ", myIsValid(s2))
	//fmt.Println(s3, " is ", myIsValid(s3))
	//fmt.Println(s4, " is ", myIsValid(s4))
	//
	//s5 := "["
	//fmt.Println(s5, " is ", myIsValid(s5))
	//
	//s6 := "]"
	//fmt.Println(s6, " is ", myIsValid(s6))

	s7 := "([]"
	fmt.Println(s7, " is ", myIsValid(s7))
}

func longestCommonPrefixTest() {
	strs := []string{"flower", "flow", "flight"}
	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs), longestCommonPrefix(strs) == "")
	fmt.Println(longestCommonPrefix(strs2), longestCommonPrefix(strs2) == "")
}

func longestPalindromeTest() {
	s := "caba"
	fmt.Println("最长回文子串:", longestPalindrome(s))
}


func myAtoiTest() {
	s := "42"               // 预期: 42
	s1 := "   -42"          // -42
	s2 := "4193 with words" // 4193
	s3 := "words and 987"   // 0
	s4 := "-91283472332"    //  -2^31 = -2147483648
	fmt.Println(myAtoi(s))
	fmt.Println(myAtoi(s1))
	fmt.Println(myAtoi(s2))
	fmt.Println(myAtoi(s3))
	fmt.Println(myAtoi(s4))
}

func reverseStringTest() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Printf("%+q\n", s) //打印字符用%q，字符数组自然就是%+q

	s1 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s1)
	fmt.Printf("%+q\n", s1)
}

func hanotaTest() {
	A, B, C := []int{2, 1, 0}, []int{}, []int{}
	hanota(A, B, C)
	fmt.Printf("A=%v,B=%v,C=%v\n", A, B, C)
	A = []int{1, 0}
}

func main() {
	hanotaTest()
}
