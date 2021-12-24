package main

import (
	"fmt"
	"unsafe"
)

// 字符串相乘
// 方法一：设nums1和nums2长度分别为m,n，则m+n-1<=乘积后结果长度<=m+n。
// 把每一位乘积结果使用数组arr存放起来，nums1[i]*nums2[j]的结果位于arr[i+j+1]，进位位于arr[i+j]。
func multiply(nums1 string, nums2 string) string {
	m, n := len(nums1), len(nums2)
	ret := make([]byte, m+n) // byte数组存储两两字符相乘和相加后结果
	for i := m - 1; i >= 0; i-- {
		a := nums1[i] - '0' // 字符 nums1[i] 对应的整数，单位 byte
		for j := n - 1; j >= 0; j-- {
			b := nums2[j] - '0'
			ret[i+j+1] += a * b
		}
	}
	for i := len(ret) - 1; i > 0; i-- {
		ret[i-1] += ret[i] / 10
		ret[i] %= 10
	}
	if ret[0] == 0 {
		ret = ret[1:]
	}
	fmt.Println(ret)
	fmt.Println(*(*string)(unsafe.Pointer(&ret)))
	return *(*string)(unsafe.Pointer(&ret))
}

func main() {
	num1, num2 := "2", "3"
	fmt.Println("multiply:", multiply(num1, num2))
	fmt.Printf("%s*%s=%s\n", num1, num2, multiply(num1, num2))
	num1, num2 = "123", "456"
	fmt.Printf("%s*%s=%s\n", num1, num2, multiply(num1, num2))
}
