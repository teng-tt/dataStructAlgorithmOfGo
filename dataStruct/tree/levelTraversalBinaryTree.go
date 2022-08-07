package main

import (
	"fmt"
	"sync"
)

/*
层次遍历二叉树，使用广度优先算法
层次遍历较复杂，用到一种名叫广度遍历的方法，需要使用辅助的先进先出的队列。

先将树的根节点放入队列
从队列里面 remove 出节点，先打印节点值，
如果该节点有左子树节点，左子树入栈，
如果有右子树节点，右子树入栈
重复2，直到队列里面没有元素
*/

// TreeNode 详细代码实现如下
type TreeNode struct {
	Data string
	Left *TreeNode
	Right *TreeNode
}

// LinkNode 链表节点
type LinkNode struct {
	Next *LinkNode
	Value *TreeNode
}

// LinkQueue 链表队列，先进先出
type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

// Push 入队
func (queue *LinkQueue) Push(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	// 如果栈顶为空，那么增加节点
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	}else {
		// 否则新元素插入链表的末尾
		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v
		// 一直遍历到链表尾部
		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		// 新节点放在链表尾部
		nowNode.Next = newNode
	}
	// 队中元素数量+1
	queue.size = queue.size + 1
}

// Pop 出队
func (queue *LinkQueue) Pop() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	// 队中元素已空
	if queue.size == 0 {
		panic("over limit")
	}
	// 顶部元素要出队
	topNode := queue.root
	v := topNode.Value
	// 将顶部元素的后继链接链上
	queue.root = topNode.Next
	// 队中元素数量-1
	queue.size = queue.size - 1
	return v
}

// Size 队列中元素数量
func (queue *LinkQueue) Size() int {
	return queue.size
}

// LayerOrder 层次遍历,使用广度优先算法，队列辅助
func LayerOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	// 新建队列
	queue := new(LinkQueue)
	// 根节点入队列
	queue.Push(tree)
	for queue.size > 0 {
		// 不断出队列
		element := queue.Pop()
		// 先打印结点值
		fmt.Print(element.Data, " ")
		// 左子树非空，入队列
		if element.Left != nil {
			queue.Push(element.Left)
		}
		// 右子树非空，入队列
		if element.Right != nil {
			queue.Push(element.Right)
		}
	}
}

func main() {
	// 测试
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}

	fmt.Println("\n层次排序")
	LayerOrder(t)
}