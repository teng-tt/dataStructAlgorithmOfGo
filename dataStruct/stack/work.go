package main

import (
	"fmt"
	"sync"
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
	} else {
		return 2
	}
}

func isPair(p, curr string) int {
	if p == "{" && curr == "}" || p == "(" && curr == ")" || p == "[" && curr == "]" {
		return 1
	} else {
		return 0
	}
}

func isLegal(s string) string {
	var stack []string
	for i := 0; i < len(s); i++ {
		currs := s[i]
		curr := string(currs)
		if isLeft(curr) == 1 {
			// 入栈
			stack = append(stack, curr)
		} else {
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
	} else {
		return "非法"
	}
}

func IsLegalByStack(str string) bool {
	stack := make([]string, 0)

	for i := 0; i < len(str); i++ {
		val := string(str[i])
		// 左括号，入栈
		switch val {
		case "{":
			stack = append(stack, "}")
		case "[":
			stack = append(stack, "]")
		case "(":
			stack = append(stack, ")")
		}
		// 还没遍历完，栈就没元素了，栈为空，说明没有匹配到相应的左边元素
		if len(stack) == 0 {
			return false
		}
		// 右括号且与栈顶元素相等，出栈
		if len(stack) > 0 && val == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
	}
	// 遍历完毕，栈内元素为空，说明都匹配消除了
	if len(stack) == 0 {
		return true
	}
	return false

}

/*
例2
给定一个包含n个元素的链表，现在要求每K个结点一组进行反转，打印反转后的链表结果
其中k是一个正整数，且n可以被k整除
列如：链表为1->2->3->4->5->6, k=3 则打印321654
*/

type LinkedStack struct {
	root *LinkNodes
	size int
	lock sync.Mutex
}
type LinkNodes struct {
	Next  *LinkNodes
	Value int
	Lock  sync.Mutex
}

// Push 入栈
func (stack *LinkedStack) Push(v int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.root == nil {
		stack.root = new(LinkNodes)
		stack.root.Value = v
	} else {
		preNode := stack.root
		newNode := &LinkNodes{preNode, v, sync.Mutex{}}
		stack.root = newNode
	}
	stack.size = stack.size + 1
}

// Pop 出栈
func (stack *LinkedStack) Pop() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size == 0 {
		return -1
	}
	topNode := stack.root
	v := topNode.Value
	stack.root = topNode.Next
	stack.size = stack.size - 1
	return v
}

// Size 获取栈的大小
func (stack *LinkedStack) Size() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.size
}

// AddLinkNodes 添加LinkNodes的元素
func (node *LinkNodes) AddLinkNodes(v int) *LinkNodes {
	node.Lock.Lock()
	defer node.Lock.Unlock()
	if node == nil {
		node = &LinkNodes{nil, v, sync.Mutex{}}
		return node
	}
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &LinkNodes{nil, v, sync.Mutex{}}
	return node
}

// GetLinkNodesLen 获取LinkNodes的大小
func (node *LinkNodes) GetLinkNodesLen() int {
	node.Lock.Lock()
	defer node.Lock.Unlock()
	n := 0
	for node != nil {
		n++
		node = node.Next
	}
	return n
}

// GetIndexNodeValue 获取LinkNodes指定节点值
func (node *LinkNodes) GetIndexNodeValue(n int) int {
	node.Lock.Lock()
	defer node.Lock.Unlock()
	if node == nil {
		return 0
	}
	for i := 0; i < n; i++ {
		node = node.Next
	}
	return node.Value
}

// Traverse 遍历链表
func (node *LinkNodes) Traverse() {
	node.Lock.Lock()
	defer node.Lock.Unlock()
	if node == nil {
		fmt.Println("-> 空链表!")
		return
	}
	for node != nil {
		fmt.Printf("%d ->", node.Value)
		node = node.Next
	}
	fmt.Println()
}

func raveress(node *LinkNodes, count int) *LinkedStack {
	var newStack *LinkedStack
	var resultStack *LinkedStack
	lenth := node.GetLinkNodesLen()
	n := 0
	for i := 0; i < lenth; i++ {
		// 按指定的反转个数入栈
		fmt.Println(lenth)
		newStack.Push(node.GetIndexNodeValue(i))
		n++
		// 判断标记数是否等于指定反转个数，不等于继续入栈，等于出栈
		if n < count {
			continue
		} else {
			// 入栈个数大于指定的反转个数，出栈
			for j := 0; j < count; j++ {
				pop := newStack.Pop()
				// 出栈元素追加到存储翻转后元素的新切片
				resultStack.Push(pop)
			}
			n = 0
		}
	}
	// 返回按指定个数翻转后的新切片
	return resultStack
}

func main() {
	s := "{[()()]}"
	s1 := "{[(]}"
	fmt.Println(IsLegalByStack(s))
	fmt.Println(IsLegalByStack(s1))
	node := new(LinkNodes)
	for i := 1; i <= 6; i++ {
		node.AddLinkNodes(i)
	}
	node.Traverse()
	newNod := raveress(node, 3)
	fmt.Println(newNod)
}
