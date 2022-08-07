package main

import "sync"

/*
实现数组队列ArrarQueue
队列先进先出，和栈操作顺序相反
这里只实现入队，和出队操作，其他操作和栈一样
*/

// 数组队列，先进先出
type arrayQueue struct {
	array []string  // 底层切片
	size int		// 队列的元素数量
	lock sync.Mutex	// 为了并发安全使用的锁
}

// Push 入队
// 直接将元素放在数组最后面即可，和栈一样，时间复杂度为：O(n)
func (queue *arrayQueue) Push(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	// 放入切片中，后进的元素放在数组最后面
	queue.array = append(queue.array, v)
	// 队中元素+1
	queue.size = queue.size + 1
}

// 出队
func (queue *arrayQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	// 队中元素已空
	if queue.size == 0 {
		panic("empty queue")
	}
	// 队列最前面元素
	v := queue.array[0]
	// 两种方法
	// 1、 原数组缩容，但缩容后继的空间不会被释放
	queue.array = queue.array[1:]
	// 2、 创建新的数组，移动次数过多
	newArray := make([]string, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		// 从老数组的第一位开始进行数据移动，因为0位已经取出了
		newArray[i-1] = queue.array[i]
	}
	queue.array = newArray

	// 队中元素数量-1
	queue.size = queue.size - 1
	return v
}
/*
出队，把数组的第一个元素的值返回，并对数据进行空间挪位
挪位有两种：
1、原地挪位，依次补位 queue.array[i-1] = queue.array[i]，
然后数组缩容：queue.array = queue.array[0 : queue.size-1]，但是这样切片缩容的那部分内存空间不会释放。
2、创建新的数组，将老数组中除第一个元素以外的元素移动到新数组，
时间复杂度是：O(n)
*/
