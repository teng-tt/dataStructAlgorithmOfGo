package main

import "fmt"

// Node 定义链式队列
type Node struct {
	Value int
	Next *Node
}

// 初始化队列
var size = 0
var queue = new(Node)

// Push 入队（从对头插入）
func Push(t *Node, v int) bool {
	// 如果队列为空
	if queue == nil {
		queue = &Node{v, nil}
		size++
		return true
	}
	// 队列不为空
	t = &Node{v, nil}
	t.Next = queue
	queue = t
	size++
	return true
}

// Pop 出队列（从队尾删除）
func Pop(t *Node) (int, bool) {
	if size == 0 {
		fmt.Println("空队列!")
		return 0, false
	}
	// 如果只有一个元素
	if size == 1 {
		queue = nil
		size--
		return t.Value, true
	}
	// 否则迭代队列，直到队尾
	temp := t
	for t.Next != nil {
		temp = t
		t = t.Next
	}
	v := temp.Next.Value
	temp.Next = nil
	size--
	return v, true

}

// Traverse 遍历队列
func Traverse(t *Node) {
	if size == 0 {
		fmt.Println("空队列!")
		return
	}
	for t != nil {
		fmt.Printf("%d ->", t.Value)
		t = t.Next
	}
	fmt.Println()

}

func main() {
	queue = nil
	// 入队列
	Push(queue, 1)
	fmt.Println("Size:", size)
	// 遍历
	Traverse(queue)
	// 出队列
	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	// 批量入队列
	for i := 0; i < 5; i++ {
		Push(queue, i)
	}
	//再次遍历
	Traverse(queue)
	fmt.Println("Size:", size)
	// 出队
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	// 再次出队
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	// 再次遍历
	Traverse(queue)
}