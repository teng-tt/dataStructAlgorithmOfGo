package main

import (
	"fmt"
	"sync"
)

// LinkStack 实现链表形式栈，后进先出
// 链表栈，后进先出
type LinkStack struct {
	root *LinkNode // 链表起点
	size int // 栈的元素数量
	lock sync.Mutex // 为了并发安全使用锁
}

// LinkNode 链表节点
type LinkNode struct {
	Next *LinkNode
	Value string
}

// Push 入栈
func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 如果栈顶为空，新建节点并设置为链表起点
	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
	}else {
		// 否则新元素插入链表的头部
		// 原来的链表
		preNode := stack.root
		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v
		// 原来的链表链接到新元素后面
		newNode.Next = preNode
		// 将新节点放在头部
		stack.root = newNode
	}
	// 栈汇总元素 +1
	stack.size = stack.size + 1
}
/*
将元素入栈，会先加锁实现并发安全。
如果栈里面的底层链表为空，表明没有元素，那么新建节点并设置为链表起点：stack.root = new(LinkNode)。
否则取出老的节点：preNode := stack.root，新建节点：newNode := new(LinkNode)，
然后将原来的老节点链接在新节点后面： newNode.Next = preNode，
最后将新节点设置为链表起点 stack.root = newNode。
时间复杂度为：O(1)
*/

// Pop 出栈
/*
元素出栈,如果栈大小为0，那么不允许出栈。
直接将链表的第一个节点 topNode := stack.root 的值取出，
然后将表头设置为链表的下一个节点：stack.root = topNode.Next，
相当于移除了链表的第一个节点
时间复杂度为：O(1)
*/
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 栈中元素一空
	if stack.size == 0 {
		panic("empty stack")

	}
	// 顶部元素出栈
	topNode := stack.root
	v := topNode.Value
	// 将顶部元素的后继链接链上
	stack.root = topNode.Next
	// 栈中元素数量-1
	stack.size = stack.size - 1
	return v
}

// Peek 获取栈顶元素
// 获取栈顶元素，但不出栈。和出栈一样，时间复杂度为：O(1)
func (stack *LinkStack) Peek() string {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}
	// 顶部元素
	v := stack.root.Value
	return v
}

// Size 获取栈大小判定是否为空
// 栈大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// IsEmpty 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

// 测试
func main() {
	linkStack := new(LinkStack)
	linkStack.Push("cat")
	linkStack.Push("dog")
	linkStack.Push("hen")
	fmt.Println("size:", linkStack.Size())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Size())
	linkStack.Push("drag")
	fmt.Println("pop:", linkStack.Pop())
}

