/*
数组实现：能快速随机访问存储的元素，通过下标 index 访问，支持随机访问，查询速度快
	但存在元素在数组空间中大量移动的操作，增删效率低。
链表实现：只支持顺序访问，在某些遍历操作中查询速度慢，但增删元素快。
*/
package main

import (
	"fmt"
	"sync"
)

// 实现数组栈 ,数组形式的下压栈，后进先出:
// 主要使用可变长数组来实现。

// 定义数据结构
type ArrayStack struct {
	array []string  // 底层切片
	size int  // 栈的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// Push 入栈
/*
将元素入栈，会先加锁实现并发安全。
入栈时直接把元素放在数组的最后面，然后元素数量加 1
性能损耗主要花在切片追加元素上，切片如果容量不够会自动扩容，
底层损耗的复杂度我们这里不计，所以时间复杂度为 O(1)
*/
func (stack *ArrayStack) Push (v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	stack.array = append(stack.array, v)
	// 栈中元素数量+1
	stack.size = stack.size + 1
}

// Pop 出栈
func (stack *ArrayStack) Pop () string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 栈中元素为空
	if stack.size == 0 {
		panic("empty")
	}
	// 栈顶元素
	v := stack.array[stack.size-1]
	// 切片收缩，但可能占用空间越来越大
	// stack.array = stack.array[:size-1]
	// 创建新的数组，空间占用不会越来越大，但可能移动元素次数过多
	newArry := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArry[i] = stack.array[i]
	}
	stack.array = newArry
	// 栈中元素减一
	stack.size = stack.size - 1
	return  v
}
/*
元素出栈，会先加锁实现并发安全。
如果栈大小为0，那么不允许出栈，否则从数组的最后面拿出元素。
元素取出后:
如果切片偏移量向前移动 stack.array[0 : stack.size-1]，表明最后的元素已经不属于该数组了，数组变相的缩容了
此时，切片被缩容的部分并不会被回收，仍然占用着空间，所以空间复杂度较高，但操作的时间复杂度为：O(1)
如果我们创建新的数组 newArray，然后把老数组的元素复制到新数组，就不会占用多余的空间，但移动次数过多
时间复杂度为：O(n)
最后元素数量减一，并返回值
*/

// Peek 获取栈顶元素
// 获取栈顶元素，但不出栈。和出栈一样，时间复杂度为：O(1)
func (stack *ArrayStack) Peek() string {
	// 栈中元素以空
	if stack.size == 0 {
		panic("empty")
	}
	// 栈顶元素值
	v := stack.array[stack.size-1]
	return v
}

// Size 获取栈大小和判定是否为空,时间复杂度都O(1)
// 栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// IsEmpty 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

// 测试
func main() {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
}