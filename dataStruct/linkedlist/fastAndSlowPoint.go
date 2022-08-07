package main

import "fmt"

// 快慢指针

type ListNode struct {
	Value int
	Next *ListNode
}

// 判断是否有环,true有，false无
func isCycle(t *ListNode) bool {
	// 如果为空或者只有一个结点，肯定无环
	if t == nil || t.Next == nil {
		return false
	}
	slow := t // 慢指针
	fast := t.Next // 快指针
	// 不重合执行循环
	for slow != fast{
		if fast == nil || fast.Next == nil {//到链表尾部无环
			return false
		}
		// 慢指针走一步，快指针走两边，为什么要1、2,因为这样时间复杂度最短
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 重合返回true
	return true
}

// 有环，判断首次相遇时，slow的位置
func detectCycle(t *ListNode) *ListNode {
	fast, slow := t, t
	// 获取首次相遇时，slow的位置
	for fast != nil || fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 如果相等保留当前位置
		if slow == fast {
			break
		}
		// 如果快指针走到尽头没环
		if fast == nil || fast.Next == nil {
			return nil
		}
	}
	// 快指针重新出发，相遇位置就是入口位置
	fast = t
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// 链表新增
func addListNode(t *ListNode, v int) int {
	if t == nil {
		t = &ListNode{v, nil}
		return 0
	}

	if v == t.Value {
		fmt.Println("结点已存在: ", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &ListNode{v, nil}
		return -2
	}
	return addListNode(t.Next, v)

}

// traverse 链表遍历
func traverse(t *ListNode) {
	if t == nil {
		fmt.Println("-> 空链表!")
		return
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
	traverse(head)
	addListNode(head, 1)
	addListNode(head,2)
	addListNode(head, 3)
	traverse(head)
	ok := isCycle(head)
	slow := detectCycle(head)
	fmt.Println(ok, slow)
}