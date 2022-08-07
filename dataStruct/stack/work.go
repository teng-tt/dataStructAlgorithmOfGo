package main

import (
	"fmt"
)

/*
例1
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
有效字符串需满足：左括号必须与相同类型的右括号匹配，左括号必须以正确的顺序匹配
例如{ [ ( ) ( ) ] } 是合法的，而 { ( [ ) ] } 是非法的。
*/

func isLeft(c string) int {
	if c == "{" || c == "(" || c == "[" {
		return 1
	}else {
		return 2
	}
}

func isPair(p, curr string) int {
	if p == "{" && curr == "}" ||  p == "(" && curr == ")" ||  p == "[" && curr == "]" {
		return 1
	}else {
		return 0
	}
}

func isLegal(s string) string{
	var stack []string
	for i := 0; i < len(s); i++ {
		currs := s[i]
		curr := string(currs)
		if isLeft(curr) == 1 {
			// 入栈
			stack = append(stack, curr)
		}else {
			if len(stack) == 0 {
				return "非法"
			}
			// 出栈
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if isPair(p, curr) == 0 {
				return "非法"
			}
		}
	}
	if len(stack) == 0 {
		return "合法"
	}else {
		return "非法"
	}
}

/*
例2
给定一个包含n个元素的链表，现在要求每K个结点一组进行反转，打印反转后的链表结果
其中k是一个正整数，且n可以被k整除
列如：链表为1->2->3->4->5->6, k=3 则打印321654
*/

type Node struct{
	Value int
	Next *Node
}

// 初始化栈
var size = 0
var stack = new(Node)

// Push 入栈
func Push(v int) bool {
	if stack == nil {
		stack = &Node{v, nil}
		size =1
		return true
	}
	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp
	size++
	return true
}

// Pop 出栈
func Pop(t *Node)(int, bool) {
	if t == nil {
		return 0, false
	}
	if size == 1 {
		size = 0
		stack = nil
		return t.Value, true
	}
	stack = stack.Next
	size--
	return t.Value, true

}

// Traveres 遍历
func Traveres(t *Node) {
	if size == 0 {
		fmt.Println("空栈")
		return
	}
	for t != nil {
		fmt.Printf("%d ->", t.Value)
		t = t.Next
	}
	fmt.Println()
}

// 添加结点
func addNode(t *Node, v int) int {
	if t == nil {
		t = &Node{v, nil}
		return 0
	}
	if v == t.Value {
		fmt.Println("结点已存在", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}
	return addNode(t.Next, v)
}


func main() {
	s := "{[()()]}"
	s1 := "{[(]}"
	fmt.Println(isLegal(s))
	fmt.Println(isLegal(s1))
}