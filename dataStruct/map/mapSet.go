package main

import (
	"fmt"
	"sync"
)

/*
字典是存储键值对的数据结构，把一个键和一个值映射起来，一一映射，键不能重复
在某些教程中，这种结构可能称为符号表，关联数组或映射
Golang 提供了这一数据结构：map，并且要求键的数据类型必须是可比较的，
因为如果不可比较，就无法知道键是存在还是不存在
字典的实现有两种方式：
	哈希表 HashTable
	红黑树 RBTree
Golang 语言中字典 map 的实现由哈希表实现，具体可参考标准库 runtime 下的 map.go 文件
 */
/*
实现不可重复集合 Set
不可重复集合 Set 存放数据，特点就是没有数据会重复，会去重
放一个数据进去，再放一个数据进去，如果两个数据一样，那么只会保存一份数据

集合 Set 可以没有顺序关系，也可以按值排序，算一种特殊的列表。
因为我们知道字典的键是不重复的，所以只要我们不考虑字典的值，
就可以实现集合，我们来实现存整数的集合 Set
 */

// Set 集合结构体
type Set struct {
	m map[int]struct{} // 用字典来实现，因为字段键不能重复
	len int			  // 集合的大小
	sync.RWMutex	  // 锁，实现并发安全
}

// NewSet 新建一个空集合,初始化一个集合
/*
使用一个容量为 cap 的 map 来实现不可重复集合
map 的值不使用，所以值定义为空结构体 struct{}
因为空结构体不占用内存空间
空结构体的内存地址都一样，并且不占用内存空间*/
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp,
	}
}

// Add 添加一个元素
/*
往结构体 s *Set 里面的内置 map 添加该元素：item，元素作为字典的键，会自动去重
同时，集合大小重新生成。时间复杂度等于字典设置键值对的复杂度，
哈希不冲突的时间复杂度为：O(1)，否则为 O(n)*/
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{} // 实际往字典添加这个键
	s.len = len(s.m) // 重新计算元素数量
}

// Remove 移除一个元素
/*
删除 map 里面的键：item 时间复杂度等于字典删除键值对的复杂度，
哈希不冲突的时间复杂度为：O(1)，否则为 O(n)*/
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	// 集合没有元素直接返回
	if s.len == 0 {
		return
	}
	delete(s.m, item) // 实际从字典删除这个键
	s.len = len(s.m)  // 重新计算元素数量
}

// Has 查看是否存在元素
// 时间复杂度等于字典获取键值对的复杂度，哈希不冲突的时间复杂度为：O(1)，否则为 O(n)
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// Len 查看集合大小 时间复杂度：O(1)
func (s *Set) Len() int {
	return s.len
}

// IsEmpty 集合是够为空 时间复杂度：O(1)
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// Clear 清除集合所有元素
// 将原先的 map 释放掉，并且重新赋一个空的 map
// 时间复杂度：O(1)
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{}  // 字典重新赋值
	s.len = 0  				// 大小归零
}

// List 将集合转换为列表，时间复杂度为O(n)
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// 测试
func main() {
	// 初始化一个容量为5的不可重复集合
	s := NewSet(5)

	s.Add(1)
	s.Add(1)
	s.Add(2)
	fmt.Println("list of all items", s.List())

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("empty")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)
	if s.Has(2) {
		fmt.Println("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items", s.List())
}