package main

// Easy - 字符串数组最长公共前缀

// 一个指针i，每次轮询数组[2:]的第i个字符与第一个字符串的第i个字符进行比较，如果有一个与第一个字符串的第i个字符不相同，返回最长前缀或者空字符串，否则 i++，最长前缀拼接第i个字符。
func longestCommonPrefix(strs []string) string {
	str0 := strs[0]
	var res string
	for i := 0; i < len(str0); i++ {
		for _, str := range strs[1:] {
			if str0[i] != str[i] {
				return res
			}
		}
		res += string(str0[i])
	}
	return res
}
