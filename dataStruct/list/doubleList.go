package main

import (
	"fmt"
	"sync"
)

// 双端队列，此处使用双向链表实现
// 缓存数据库 Redis 的 列表List 基本类型就是用它来实现的
// DoubleList 双端链表，双端队列

// DoubleList 双端队列结构体
type DoubleList struct {
	head *ListNode  // 指向链表头部
	tail *ListNode  // 指向链表尾部
	len  int        // 列表长度
	lock sync.Mutex // 为了进行并发安全的pop弹出操作
}

// ListNode 列表节点
type ListNode struct {
	pre   *ListNode // 前驱节点
	nex   *ListNode // 后继节点
	value string    // 值
}

// 列表节点普通操作
// 以下是对节点结构体 ListNode 的操作，主要判断节点是否为空
// 有没有后驱和前驱节点，返回值等，时间复杂度都是 O(1)

// GetValue 获取节点值
func (node *ListNode) GetValue() string {
	return node.value
}

// GetPre 获取前驱节点
func (node *ListNode) GetPre() *ListNode {
	return node.pre
}

// GetNext 获取后继节点
func (node *ListNode) GetNext() *ListNode {
	return node.nex
}

// IsPre 是否存在前驱节点
func (node *ListNode) IsPre() bool {
	return node.pre != nil
}

// IsNext 是否存在后继节点
func (node *ListNode) IsNext() bool {
	return node.nex != nil
}

// IsNil 节点是否为空
func (node *ListNode) IsNil() bool {
	return node == nil
}

// Len 返回列表长度
func (list *DoubleList) Len() int {
	return list.len
}

/*
从头部开始某个位置前插入新节点
参考数组下标，下标从0开始。从双端列表的头部，插入新的节点
大部分时间花在遍历位置上，如果 n=0，那么时间复杂度为 O(1)，否则为 O(n)
定位到的节点有三种情况，我们需要在该节点之前插入新节点：
第一种情况，判断定位到的节点 node 是否为空，
	如果为空，表明列表没有元素，将新节点设置为新头部和新尾部即可
	否则，我们要插入新的节点到非空的列表上。
	我们找到定位到的节点的前驱节点：pre := node.pre，
	我们要把新节点变成定位到的节点的前驱节点，之前的前驱节点 pre 要往前顺延。
第二种情况，如果前驱节点为空：pre.IsNil()，表明定位到的节点 node 为头部，
	那么新节点要取代它，成为新的头部
	新节点成为新的头部，需要将新节点的后驱设置为老头部：newNode.next = node，
	老头部的前驱为新头部：node.pre = newNode，并且新头部变化：list.head = newNode
第三种情况，如果定位到的节点的前驱节点不为空，表明定位到的节点 node 不是头部节点，
	么我们只需将新节点链接到节点 node 之前即可：
	先将定位到的节点的前驱节点和新节点绑定，因为现在新节点插在前面了，
	把定位节点的前驱节点的后驱设置为新节点：pre.next = newNode，
	新节点的前驱设置为定位节点的前驱节点：newNode.pre = pre。
	同时，定位到的节点现在要链接到新节点之后，所以新节点的后驱设置为：newNode.next = node，
	定位到的节点的前驱设置为：node.pre = newNode。
最后：当然插入新节点的最后，我们要将链表长度加一
*/

// AddNodeFormHead 从头部开始，添加节点到第N+1个元素之前，
// N=0表示添加到第一个元素之前，表示新节点成为新的头部，
// N=1表示添加到第二个元素之前，以此类推
func (list *DoubleList) AddNodeFormHead(n int, v string) {
	// 并发加锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 如果索引超过或等于列表长度，一定找不到，直接panic
	if n != 0 && n >= list.len {
		panic("index out")
	}
	// 先找出t头部
	node := list.head
	// 从头部节点往后遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.nex
	}
	// 新节点
	newNode := new(ListNode)
	newNode.value = v
	// 如果定位到的节点为空，表示列表为空，将新节点设置为新头部和新尾部
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	}else {
		// 定位到的节点，它的前驱节点
		pre := node.pre
		// 如果定位到的节点前驱为nil，那么定位到的节点为链表头部，需要换头部
		if pre.IsNil() {
			// 将新节点链接在老头部之前
			newNode.nex = node
			node.pre = newNode
			// 将新节点链接在老头部之前
			list.head = newNode
		}else {
			// 将新节点插入到定位到的节点之前
			// 定位到的节点的前驱节点 pre 现在链接到新节点上
			pre.nex = newNode
			newNode.pre = pre
			// 定位到的节点链接到新的结点之后
			newNode.nex = node
			node.pre= newNode
		}
	}
	// 列表长度+1
	list.len = list.len + 1
}

// 从尾部开始某个位置后插入新节点
// AddNodeFromTail 从尾部开始，添加节点到第N+1个元素之后，
// N=0表示添加到第一个元素之后，表示新节点成为新的尾部，
// N=1表示添加到第二个元素之后，以此类推
// 与从头部一样也存在三种情况，找到的结点为空，为尾部节点，不为尾部节点，处理方法一直

func (list *DoubleList) AddNodeFromTail(n int, v string) {
	// 并发加锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 如果索引超过或等于列表长度，一定找不到，直接panic
	if n != 0 && n >= list.len {
		panic("index out")
	}
	// 先找出尾部
	node := list.tail
	//  往前遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	// 新节点
	newNode := new(ListNode)
	newNode.value = v
	// 如果定位到的节点为空，表示列表为空，将新节点设置为新头部和新尾部
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	}else {
		// 定位到的节点，它的后驱
		next := node.nex
		// 如果定位到的节点后驱为nil，那么定位到的节点为链表尾部，需要换尾部
		if next.IsNil() {
			// 将新节点链接在老尾部之后
			node.nex = newNode
			newNode.pre = node
			// 新节点成为尾部
			list.tail = newNode
		}else {
			// 将新节点插入到定位到的节点之后
			// 新节点链接到定位到的节点之后
			newNode.pre = node
			node.nex = newNode
			// 定位到的节点的后驱节点链接在新节点之后
			next.pre = newNode
			newNode.nex = next
		}
	}
	// 列表长度加1
	list.len = list.len + 1
}

// First 返回列表链表头结点
func (list *DoubleList) First() *ListNode {
	return list.head
}

// Last 返回列表链表尾结点
func (list *DoubleList) Last() *ListNode {
	return list.tail
}

// IndexFromHead 从头部开始某个位置获取列表节点
// 如果索引超出或等于列表长度，那么找不到节点，返回空。
// 否则从头部开始遍历，拿到节点。
// 时间复杂度为：O(n)
// IndexFromHead 从头部开始往后找，获取第N+1个位置的节点，索引从0开始
func (list *DoubleList) IndexFromHead(n int) *ListNode {
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}
	// 获取头部节点
	node := list.head
	// 往后遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.nex
	}
	return node
}

// IndexFromTail 从尾部开始某个位置获取列表节点
// IndexFromTail 从尾部开始往前找，获取第N+1个位置的节点，索引从0开始。
// 如果索引超出或等于列表长度，那么找不到节点，返回空。
// 否则从尾部开始遍历，拿到节点。
// 时间复杂度为：O(n)
func(list *DoubleList) IndexFromTail(n int) *ListNode {
	if n >= list.len {
		return nil
	}
	// 获取头部节点
	node := list.tail
	// 往前遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	return node
}

// PopFromHead 从头部开始移除并返回某个位置的节点
// 获取某个位置的元素，并移除它
// PopFromHead 从头部开始往后找，获取第N+1个位置的节点，并移除返回
// 定位到的并要移除的节点有三种情况发生，移除的是头部，尾部或者中间节点
// 主要的耗时用在定位节点上，其他的操作都是链表链接，可以知道时间复杂度为：O(n)
func (list *DoubleList) PopFromHead(n int) *ListNode {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}
	// 获取头部
	node := list.head
	// 往后遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.nex
	}
	// 移除的节点的前驱和后驱
	pre := node.pre
	next := node.nex
	// 如果前驱和后驱都为nil，那么移除的节点为链表唯一节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	}else if pre.IsNil() {
		// 表示移除的是头部节点，那么下一个节点成为头节点
		list.head = next
		next.pre = nil
	}else if next.IsNil() {
		// 表示移除的是尾部节点，那么上一个节点成为尾节点
		list.tail = pre
		pre.nex = nil
	}else {
		// 移除的是中间节点
		pre.nex = next
		next.pre = pre
	}
	// 节点减一
	list.len = list.len - 1
	return node
}

// PopFromTail 从尾部开始移除并返回某个位置的节点
// PopFromTail 从尾部开始往前找，获取第N+1个位置的节点，并移除返回
// 定位到的并要移除的节点有三种情况发生，移除的是尾部，头部或者中间节点
// 主要的耗时用在定位节点上，其他的操作都是链表链接，可以知道时间复杂度为：O(n)
func (list *DoubleList) PopFromTail(n int) *ListNode {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}
	// 取出尾部节点
	node := list.tail
	// 往前遍历拿到第 N+1 个位置的元素
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	// 移除的节点的前驱和后驱
	pre  := node.pre
	next := node.nex
	// 如果前驱和后驱都为nil，那么移除的节点为链表唯一节点
	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	}else if pre.IsNil() {
		// 表示移除的是头部节点，那么下一个节点成为头节点
		list.head = next
		next.pre = nil
	}else if next.IsNil() {
		// 表示移除的是尾部节点，那么上一个节点成为尾节点
		list.tail = pre
		pre.nex = nil
	}else {
		// 移除的是中间节点
		pre.nex = next
		next.pre = pre
	}
	// 节点减一
	list.len = list.len - 1
	return node
}

func main() {
	// 测试
	list := new(DoubleList)
	// 在列表头部插入新元素
	list.AddNodeFormHead(0, "I")
	list.AddNodeFormHead(0, "Love")
	list.AddNodeFormHead(0, "you")
	// 在列表尾部插入新元素
	list.AddNodeFromTail(0, "may")
	list.AddNodeFromTail(0, "happy")
	// 在N之前后之后插入
	list.AddNodeFromTail(list.Len()-1, "begin second")
	list.AddNodeFormHead(list.Len()-1, "end second")

	// 正常遍历，比较慢，因为内部会遍历拿到值返回
	//for i := 0; i < list.len; i++ {
	//	// 从头部开始索引
	//	node := list.IndexFromHead(i)
	//	// 节点为空不可能，因为list.Len()使得索引不会越界
	//	if !node.IsNil() {
	//		fmt.Println(node.GetValue())
	//	}
	//}
	//fmt.Println("----------")

	// 正常遍历，特别快，因为直接拿到的链表节点
	// 先取出第一个元素
	first := list.First()
	for !first.IsNil() {
		// 如果非空就一直遍历
		fmt.Println(first.GetValue())
		// 接着下一个节点
		first = first.GetNext()
	}
	fmt.Println("----------")

	// 元素一个个 POP 出来
	for {
		node := list.PopFromHead(0)
		if node.IsNil() {
			// 没有元素了，直接返回
			break
		}
		fmt.Println(node.GetValue())
	}
	fmt.Println("----------")
	fmt.Println("len", list.Len())
}