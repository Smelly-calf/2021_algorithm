package common

import (
	"fmt"
	"reflect"
)

// 栈: 数组，push和pop，create方法
type Stack struct {
	elems []interface{}
	len   int
}

func (s *Stack) Push(e interface{}) {
	s.len++
	s.elems = append(s.elems, e)
}

func (s *Stack) Pop() interface{} {
	if s.len == 0 {
		return -1
	}
	last := s.elems[s.len-1]
	s.len--
	s.elems = s.elems[:s.len]
	return last
}

func (s *Stack) Empty() bool{
	if s.len == 0 {
		return true
	}
	return false
}

func CreateStack(arr interface{}) *Stack {
	st := Stack{}
	if reflect.TypeOf(arr).Kind() != reflect.Slice {
		panic("wrong input type")
	}
	realArr := reflect.ValueOf(arr)
	for i := 0; i < realArr.Len(); i++ {
		st.Push(realArr.Index(i))
	}
	return &st
}

func (s *Stack) Traverse() {
	for s.len > 0 {
		fmt.Printf("%v ", s.Pop())
	}
	fmt.Println()
}
