package main

import (
	"fmt"
)

func addTwoNumbersTest() {
	arr1 := []int{9, 9, 9, 9, 9, 9, 9}
	arr2 := []int{9, 9, 9, 9}

	l1 := CreateLinkedList(arr1)
	l2 := CreateLinkedList(arr2)
	fmt.Println("create finished")
	l3 := addTwoNums(l1, l2)
	for l3 != nil { // 输出结果: 0742
		fmt.Print(l3.num)
		l3 = l3.next
	}
}

func mergeTwoListsTest() {
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	l1 := CreateLinkedList(arr1)
	l2 := CreateLinkedList(arr2)
	TraverseLinkedList(mergeTwoLists(l1, l2))
}

func main() {
	mergeTwoListsTest()
}
