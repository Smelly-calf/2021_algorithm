package main

import (
	"2021_algorithm/common"
	"fmt"
)

// Easy - 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效：左括号必须用相同类型的右括号闭合，左括号必须以正确的顺序闭合。

// 思路：遍历字符串，用栈：遇到左括号入栈，遇到右括号时左括号出栈：如果与右括号匹配继续遍历下一个字符，结束条件: 遇到无法消去的右括号或者字符串遍历结束。
// 注意 
//	1.入栈是要判断是否是最后一个字符，如果是则直接返回false；
//	2.注意pop时判断溢出；
// 	3.遍历结束后栈长度如果>0返回false。

func myIsValid(s string) bool {
	st := common.Stack{}
	isLeft := func(c byte) bool {
		if c == '(' || c == '{' || c == '[' {
			return true
		}
		return false
	}
	isMatch := func(a byte, b byte) bool {
		if a == '(' && b == ')' || a == '{' && b == '}' || a == '[' && b == ']' {
			return true
		}
		return false
	}
	for i := 0; i < len(s); i++ {
		if isLeft(s[i]) {
			if i == len(s)-1 {
				return false
			}
			st.Push(s[i]) // 考虑当字符串长度为奇数
			continue
		}
		// last是左括号，s[i]是右括号
		last := st.Pop() // 考虑栈溢出
		if !isMatch(last.(byte), s[i]) {
			return false
		}
	}
	// 判断是否还有左括号
	if !st.Empty() {
		return false
	}
	return true
}

func main() {
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