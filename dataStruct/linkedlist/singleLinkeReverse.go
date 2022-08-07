package main

import (
	"fmt"
)

// 单向链表翻转

// ListNode 定义链表
type ListNode struct {
	Value int
	Next *ListNode
}

// 翻转单链表
func reverseList(t * ListNode) *ListNode {
	cur := t
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

// 链表新增
func addListNode(t *ListNode, v int) int {
	if t == nil {
		t = &ListNode{v, nil}
		return 0
	}
	if v == t.Value {
		fmt.Println("结点已存在!")
		return -1
	}
	// 如果当前节点下一节点为空
	if t.Next == nil {
		t.Next = &ListNode{v, nil}
		return -2
	}
	// 如果当前节点下一节点不为空
	return addListNode(t.Next, v)
}

// 遍历链表
func traverse(t *ListNode) {
	if t == nil {
		fmt.Println("-> 空链表!")
	}
	for t != nil {
		fmt.Printf("%d ->", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {

	var head = new(ListNode)
	fmt.Println(head)
	// 遍历链表
	traverse(head)
	// 新增结点
	addListNode(head, 1)
	// 遍历链表
	traverse(head)
	// 新增结点
	addListNode(head,2)
	addListNode(head, 3)
	addListNode(head, 4)
	addListNode(head, 5)
	addListNode(head, 6)
	// 新增重复结点
	addListNode(head, 2)
	// 遍历链表
	traverse(head)
	// 链表反转
	newHead := reverseList(head)
	// 遍历链表
	traverse(newHead)
}

