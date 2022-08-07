package main

import "sync"

// LinkQueue 实现链式队列
// 链表队列，先进先出
type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

// LinkNode 链表节点
type LinkNode struct {
	Value string
	Next *LinkNode
}

// Push 入队
// 将元素放在链表的末尾，所以需要遍历链表，时间复杂度为：O(n)
func (queue *LinkQueue) Push(v string) {
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

// Remove 出队
// 链表第一个节点出队即可，时间复杂度为：O(1)
func (queue *LinkQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 队中元素已空
	if queue.size == 0 {
		panic("empty")
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