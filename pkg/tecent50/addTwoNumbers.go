package main

import "fmt"

// 两数相加: 两个非空链表表示两个非负整数，每位数字按照逆序方式存储，每个节点存一个数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。时间复杂度O(m+n)

// 思路1：将长度较短的链表在末尾补零使得两个连表长度相等，再一个一个元素对其相加
type Node struct {
	num  int
	next *Node
}

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	var len1, len2 int
	l11 := l1
	l22 := l2
	for l11 != nil {
		len1 += 1
		l11 = l11.next
	}
	for l22 != nil {
		len2 += 1
		l22 = l22.next
	}

	if len1 < len2 { // l1末尾补零
		l11 = l1
		for i := len1; i < len2; i++ {
			l11.next = &Node{}
			l11 = l11.next
		}
	} else {
		l22 = l2
		for l22.next != nil { //l22移动到l2的末尾
			l22 = l22.next
		}
		for i := len2; i < len1; i++ {
			l22.next = &Node{}
			l22 = l22.next
		}
	}

	// 初始化新的链表和链表头指针pre，遍历链表节点相加
	cur := &Node{}
	pre := &Node{}
	pre = cur
	var carry int
	for l1 != nil && l2 != nil {
		cur.num = (carry + l1.num + l2.num) % 10
		if l1.next != nil && l2.next != nil {
			cur.next = &Node{0, nil} // l1.next和l2.next不为nil时，再初始化cur.next
			cur = cur.next
		}
		carry = (carry + l1.num + l2.num) / 10 // carry=当前节点数字之和除10
		l1 = l1.next
		l2 = l2.next
	}
	if carry > 0 {
		cur.next = &Node{num: carry}
	}
	return pre
}

// 思路2: 不对齐补零，每次遍历用sum(代表每个位的和的结果)加，节点的值=sum%10
func addTwoNums(l1 *Node, l2 *Node) *Node {
	cur := &Node{}
	pre := &Node{}
	pre = cur
	var carry int
	var sum int
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			sum += carry + l1.num + l2.num
			carry = sum / 10
			cur.num = (carry + l1.num + l2.num) % 10
		}
		if l1 == nil {

		}
		cur.num = (carry + l1.num + l2.num) % 10
		if l1.next != nil && l2.next != nil {
			cur.next = &Node{0, nil} // l1.next和l2.next不为nil时，再初始化cur.next
			cur = cur.next
		}
		carry = (carry + l1.num + l2.num) / 10 // carry=当前节点数字之和除10
		l1 = l1.next
		l2 = l2.next
	}
	if carry > 0 {
		cur.next = &Node{num: carry}
	}
	return pre
}

func CreateNode(arr []int) *Node {
	pre := &Node{}
	l := &Node{}
	pre = l // pre保存l的head节点
	for i := 0; i < len(arr)-1; i++ {
		l.num = arr[i]
		l.next = &Node{}
		l = l.next
	}
	l.num = arr[len(arr)-1]
	return pre
}

func main() {
	arr1 := []int{9, 9, 9, 9, 9, 9, 9}
	arr2 := []int{9, 9, 9, 9}

	l1 := CreateNode(arr1)
	l2 := CreateNode(arr2)
	fmt.Println("create finished")
	l3 := addTwoNumbers(l1, l2)
	for l3 != nil { // 输出结果: 0742
		fmt.Print(l3.num)
		l3 = l3.next
	}
}
