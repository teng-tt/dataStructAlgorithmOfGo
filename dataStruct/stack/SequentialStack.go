package main

import "fmt"

// 顺序栈

type Stack struct {
	Value []int
}

// Push 进栈
func (s *Stack) Push(value int) {
	s.Value = append(s.Value, value)
}

// Pop 出栈
func (s *Stack) Pop()(int, bool) {
	lenth := len(s.Value)

	// 空栈
	if lenth == 0 {
		return 0, false
	}
	value := s.Value[lenth-1]
	s.Value = s.Value[:lenth-1]
	return value, true
}

// Traverse 栈的遍历
func (s *Stack) Traverse() {
	lenth := len(s.Value)
	if lenth == 0 {
		fmt.Println("空栈!")
		return
	}
	for i := 0; i < lenth; i++ {
		fmt.Printf("%d ->", s.Value[i])
	}
	fmt.Println()
}

func main() {
	stack := Stack{}
	fmt.Println(stack)
	// 遍历
	stack.Traverse()
	// 入栈
	stack.Push(1)
	// 再次遍历
	stack.Traverse()
	// 多次入栈
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	// 遍历
	stack.Traverse()
	//出栈
	v1, _ := stack.Pop()
	fmt.Println("弹出：",v1)
	// 遍历
	stack.Traverse()
	v2, _ := stack.Pop()
	v3, _ := stack.Pop()
	fmt.Println("弹出：",v2, v3)
	// 遍历
	stack.Traverse()
}