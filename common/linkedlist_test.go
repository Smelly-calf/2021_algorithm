package common

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	head := CreateLinkedList(nums)
	newHead := ReverseLinkedList(head)
	fmt.Printf("reversedList: %+v\n", TraverseLinkedList(newHead))
}
