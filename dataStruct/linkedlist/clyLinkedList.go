package main

import (
	"fmt"
)

// 循环链表

// 定义循环链表
type clyNode struct {
	value int
	pre   *clyNode
	next  *clyNode
}

// 初始化空的链表
// 此时前驱和后驱节点为自己，没有循环，时间复杂度为：O(1)
func (c *clyNode) init() *clyNode {
	// 初始化，时前倾后继节点都指向自己
	c.pre = c
	c.next = c
	return c
}

// New 创建N个空节点的循环链表
// 会连续绑定前驱和后驱节点，时间复杂度为：O(n)
func New(n int) *clyNode {
	if n <= 0 {
		return nil
	}
	// new 2个 空结点
	c := new(clyNode)
	p := c
	// 循环添加空节点
	for i := 0; i < n; i++ {
		p.next = &clyNode{pre: p}
		p = p.next
	}
	// 回到头部
	p.next = c
	c.pre = p
	return c
}

// Next 分别获取循环链表的上一个节点和下一个结点
// 获取前驱或后驱节点，时间复杂度为：O(1)
// 获取下一个节点
func (c *clyNode) Next() *clyNode {
	if c.next == nil {
		// 下一个节点为空，初始化，返回前驱后继都指向自己的空链表
		return c.init()
	}
	return c.next
}

// Prev 获取上一个节点
func (c *clyNode) Prev() *clyNode {
	if c == nil {
		// 下一个节点为空，初始化，返回前驱后继都指向自己的空链表
		return c.init()
	}
	return c.pre
}

// Move 获取第n个节点,需要遍历 n 次，所以时间复杂度为：O(n)
// 因为链表是环的，当n为负数，表示从前面我那个前的遍历，否则往后面遍历
func (c *clyNode) Move(n int) *clyNode {
	if c.next == nil {
		return nil
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			c = c.pre
		}
	case n > 0:
		for ; n > 0; n-- {
			c = c.next
		}
	}
	return c
}

// Link 添加节点， 往节点A，链接一个节点，并且返回之前节点A的后驱节点
// 如果节点 s 是一个新的节点。
// 那么也就是在 r 节点后插入一个新节点 s，而 r 节点之前的后驱节点，
// 将会链接到新节点后面，并返回 r 节点之前的第一个后驱节点 n
func (c *clyNode) Link(s *clyNode) *clyNode {
	n := c.Next()
	if s != nil {
		p := s.Prev()
		c.next = s
		s.pre = c
		n.pre = p
		p.next = n
	}
	return n
}

// 测试
func LinkNewTest() {
	// 第一个节点
	r := &clyNode{value: -1}
	r = r.init()
	// 链接新的五个节点
	r.Link(&clyNode{value: 1})
	r.Link(&clyNode{value: 2})
	r.Link(&clyNode{value: 3})
	r.Link(&clyNode{value: 4})
	r.Link(&clyNode{value: 5})

	node := r
	for {
		// 打印节点值
		fmt.Println(node.value)
		// 移到下一个结点
		node = node.Next()
		// 如果结点回到了起点，结束
		if node == r {
			return
		}
	}
}

func main() {
	LinkNewTest()
}
