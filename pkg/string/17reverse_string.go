package main

import "fmt"

// Easy - 反转字符串
// 双指针，遍历字符数组并交换反转前下标i 和 反转后下标 len-i-1 的值，遍历结束条件是i>=j
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

// 就离谱，这种写法性能直接下降到36ms，是上面耗时（24ms）的1.5倍，经测试是先在循环外部声明var tmp byte再赋值不如直接tmp:=s[i]性能好，引入tmp不如直接s[i],s[j]=s[j],s[i]性能好。
func reverseStringSlow(s []byte) {
	i := 0
	j := len(s) - 1
	var tmp byte
	for i < j {
		tmp = s[i]
		s[j] = s[i]
		s[i] = tmp
		i++
		j--
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Printf("%+q\n", s) //打印字符用%q，字符数组自然就是%+q

	s1 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s1)
	fmt.Printf("%+q\n", s1)
}
