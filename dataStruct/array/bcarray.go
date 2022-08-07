package main

import (
	"fmt"
	"sync"
)

// 自定义实现一个可变成数组，类似go中的slice

// Arrar 可变成数组
type Arrar struct {
	array []int // 固定大小的数组，用满容量和满大小的切片来代替
	len int   // 真正长度
	cap int  // 容量
	lock sync.Mutex  // 为了并发安全使用的锁
}

// Make 数组初始化
// 创建一个 len 个元素，容量为 cap 的可变长数组
// 新建一个可变长数组
/*
主要利用满容量和满大小的切片来充当固定数组，结构体 Array 里面的字段 len 和 cap 来控制值的存取
不允许设置 len > cap 的可变长数组。时间复杂度为：O(1)，因为分配内存空间和设置几个值是常数时间
*/
func Make(len, cap int) *Arrar {
	s := new(Arrar)
	if len > cap {
		panic("len large than cap")
	}
	// 把切片当数组使用
	array := make([]int, cap, cap)
	// 元数据
	s.array = array
	s.cap = cap
	s.len = 0
	return s
}

// Append 增加一个元素
// 耗时主要在老数组中的数据移动到新数组，时间复杂度为：O(n)
// 如果容量够的情况下，时间复杂度会变为：O(1)
func (a *Arrar) Append(element int) {
	// 并发锁
	a.lock.Lock()
	defer a.lock.Unlock()

	// 大小等于容量，表示没有多余位置了
	if a.len == a.cap {
		// 没容量，数组要扩容，扩容到2倍
		newCap := 2 * a.len
		// 如果之前的容量为0，那么新增的容量为1
		if a.cap == 0 {
			newCap = 1
		}
		// 生成一个新数组
		newArray := make([]int, newCap, newCap)
		// 把旧数组的数据挪移到新的数组
		for k, v := range a.array {
			newArray[k] = v
		}
		// 替换数组
		a.array = newArray
		a.cap = newCap
	}
	// 把元素放在数组里,放在最后一位，数组的末尾
	a.array[a.len] = element
	// 真实长度+1
	a.len = a.len + 1
}

// AppendMany AppendAny 添加多个元素
// 简单遍历一下，调用 Append 函数
// ...int 是 Golang 的语言特征，表示多个函数变量
func (a *Arrar) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

// Get 获取指定下标元素
// 可变长数组的真实大小为0，或者下标 index 超出了真实长度 len ，将会 panic 越界
// 因为只获取下标的值，所以时间复杂度为 O(1)
func (a *Arrar) Get (index int) int {
	// 越界了
	if a.len == 0 || index >= a.len {
		panic("index over len")
	}
	return a.array[index]
}

// Len 获取真实长度容量 时间复杂度为 O(1)
func (a *Arrar) Len() int {
	return a.len
}

// Cap 获取真实容量
func (a *Arrar) Cap() int {
	return a.cap
}

// Print 测试辅助打印
func Print(array *Arrar) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		if i == 0 {//第一个元素
		result = fmt.Sprintf("%s%d", result, array.Get(i))
		continue
		}
		result = fmt.Sprintf("%s %d", result, array.Get(i))
	}
	result = result + "]"
	return result
}

// 测试自定义可变数组
func main() {
	// 创建一个容量为3的动态数组
	a := Make(0, 3)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加一个元素
	a.Append(10)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加一个元素
	a.Append(9)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	// 增加多个元素
	a.AppendMany(8, 7)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
}

