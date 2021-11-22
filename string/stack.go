package main

import (
	"fmt"
	"reflect"
)

// 栈: 数组，push和pop，create方法
type stack struct {
	elems []interface{}
	len   int
}

func (s *stack) push(e interface{}) {
	s.len++
	s.elems = append(s.elems, e)
}

func (s *stack) pop() interface{} {
	if s.len == 0 {
		return -1
	}
	last := s.elems[s.len-1]
	s.len--
	s.elems = s.elems[:s.len]
	return last
}

func createStack(arr interface{}) *stack {
	st := stack{}
	if reflect.TypeOf(arr).Kind() != reflect.Slice {
		panic("wrong input type")
	}
	realArr := reflect.ValueOf(arr)
	for i := 0; i < realArr.Len(); i++ {
		st.push(realArr.Index(i))
	}
	return &st
}

func (s *stack) traverse() {
	for s.len > 0 {
		fmt.Println(s.pop())
	}
}
