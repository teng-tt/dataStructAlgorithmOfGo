package main

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
	"math"
	"sync"
)

// 实现拉链哈希表（链式哈希表）
// 简单版哈希版本
/*
实现拉链哈希表有以下的一些操作：

初始化：新建一个 2^x 个长度的数组，一开始 x 较小。
添加键值：进行 hash(key) & (2^x-1)，定位到数组下标，查找数组下标对应的链表，如果链表有该键，更新其值，否则追加元素。
获取键值：进行 hash(key) & (2^x-1)，定位到数组下标，查找数组下标对应的链表，如果链表不存在该键，返回false，否则返回该值以及true。
删除键值：进行 hash(key) & (2^x-1)，定位到数组下标，查找数组下标对应的链表，如果链表不存在该键，直接返回，否则删除该键。

# 进行键值增删时如果数组容量太大或者太小，需要相应缩容或扩容。
哈希查找的速度快，主要是利用空间换时间的优点。如果哈希表的数组特别大特别大，那么哈希冲突的几率就会降低。
然而哈希表中的数组太大或太小都不行，太大浪费了空间，太小则哈希冲突太严重，所以需要对哈希表中的数组进行缩容和扩容。

如何伸缩主要根据哈希表的大小和已添加的元素数量来决定。假设哈希表的大小为 16，已添加到哈希表中的键值对数量是 8，我们称 8/16=0.5 为加载因子 factor。
可以设定加载因子 factor <= 0.125 时进行数组缩容，每次将容量砍半，当加载因子 factor >= 0.75 进行数组扩容，每次将容量翻倍。

大部分编程语言实现的哈希表只会扩容，不会缩容，因为对于一个经常访问的哈希表来说，缩容后会很快扩容，造成的哈希搬迁成本巨大，
这个成本比起存储空间的浪费还大，所以这里只实现哈希表扩容。
*/

// 扩容因子 0.75 作为扩容因子，只是它刚刚好，其它也可以
const expandFactor = 0.75

// XXHash xxhash实现哈希函数
func XXHash(key []byte) uint64 {
	hash := xxhash.New64()
	hash.Write(key)
	return hash.Sum64()
}

// HashMap hash表结构体
type HashMap struct {
	array []*KeyPairs // 哈希表数组，每个元素是一个键值对
	capacity int	  // 数组容量
	len int			  // 已添加键值对元素数量
	capacityMask int  // 容量掩码，等于 capacity-1，用来计算数组下标
	lock sync.Mutex   // 增删键值对时，需要考虑并发安全
}

// KeyPairs 键值对，连成一个链表
type KeyPairs struct {
	key string
	value interface{}
	netx *KeyPairs
}

//传入 capacity 初始化哈希表数组容量
//容量掩码 capacityMask = capacity-1 主要用来计算数组下标。
//传入容量小于默认容量 16，那么将 16 作为哈希表的初始数组大小。
//否则将第一个大于 capacity 的 2 ^ k 值作为数组的初始大小

// NewHashMap 初始化hash链表，创建大小为capacity的哈希链表
func NewHashMap(cap int) *HashMap {
	// 默认大小为19
	defaultCap := 1 << 4
	if cap <= defaultCap {
		// 如果传入的大小小于默认大小，那么使用默认大小16
		cap = defaultCap
	}else {
		// 否则，实际大小为大于 capacity 的第一个 2^k
		cap = 1 << (int(math.Ceil(math.Log2(float64(cap)))))
	}
	// 新建一个哈希表
	hash := new(HashMap)
	hash.array = make([]*KeyPairs, cap, cap)
	hash.capacity = cap
	hash.capacityMask = cap -1
	return hash
}

// Len 返回哈希表已添加元素的数量
func (h *HashMap) Len() int {
	return h.len
}

//根据公式 hash(key) & (2^x-1)，使用 xxhash 哈希算法来计算键 key 的哈希值，
//并且和容量掩码 mask 进行 & 求得数组的下标，用来定位键值对该放在数组的哪个下标下

// HashIndex 计算键的哈希值并求出数组下标
func (h *HashMap) HashIndex(key string, mask int) int {
	// 求hash值
	hashKey := XXHash([]byte(key))
	// 求下标
	index  := hashKey & uint64(mask)
	return int(index)
}

// Put 添加键值对
// 哈希表添加键值对,主要操作还是，链表的追加，和遍历，以及扩容（创建新的然后使用遍历赋值老的数据，最后替换老的）
func (h *HashMap) Put(key string, value interface{}) {
	// 实现并发安全
	h.lock.Lock()
	defer h.lock.Unlock()
	// 键值对要放的哈希表数组下标
	index := h.HashIndex(key, h.capacityMask)
	// 哈希表数组的下标元素
	element := h.array[index]
	// 元素为空，表示空链表，没有哈希冲突，直接赋值
	if element == nil {
		h.array[index] = &KeyPairs{
			key,
			value,
			nil,
		}
	}else {
		// 链表最后一个键值对
		var lastPairs *KeyPairs
		// 遍历链表查看元素是否存在，存在则替换值，否则找到最后一个键值对
		for element != nil {
			// 键值对存在，更新值并返回
			if element.key == key {
				element.value = value
				return
			}
			lastPairs = element
			element = element.netx
		}
		// 找不到键值对，将新键值对添加到链表尾端
		lastPairs.netx = &KeyPairs{
			key,
			value,
			nil,
		}
	}
	// 新的哈希表数量
	newLen := h.len + 1
	// 如果超出扩容因子，需要扩容
	if float64(newLen) / float64(h.capacity) >= expandFactor {
		// 新建一个原来2倍大小的哈希表
		newM := new(HashMap)
		newM.array = make([]*KeyPairs, 2*h.capacity, 2*h.capacity)
		newM.capacity = 2*h.capacity
		newM.capacityMask = 2*h.capacity - 1
		// 遍历老的哈希表，将键值对重新哈希到新哈希表
		for _, pairs := range h.array {
			for pairs != nil {
				// 直接递归PUT
				newM.Put(pairs.key, pairs.value)
				pairs = pairs.netx
			}
		}
		// 替换老的哈希表
		h.array = newM.array
		h.capacity = newM.capacity
		h.capacityMask = newM.capacityMask
	}
	h.len = newLen
}

// Get 获取哈希表键值对
func (h *HashMap) Get(key string) (value interface{}, ok bool) {
	// 实现并发安全
	h.lock.Lock()
	defer h.lock.Unlock()
	// 键值对要放的哈希表数组下标
	index := h.HashIndex(key, h.capacityMask)
	// 哈希表数组下标的元素
	element := h.array[index]

	// 遍历链表是否存在元素，存在则返回
	for element != nil {
		if element.key == key {
			return element.value, true
		}
		element = element.netx
	}
	return
}

// Delete 哈希表删除键值对
// 键值对删除时，哈希表不会缩容，此处不实现缩容
func (h *HashMap) Delete(key string) {
	// 实现并发安全
	h.lock.Lock()
	defer h.lock.Unlock()
	// 键值对要存放哈希表数组下标
	index := h.HashIndex(key, h.capacityMask)
	// 哈希表数组的下标元素
	element := h.array[index]
	// 空链表，不用删除，直接返回
	if element == nil {
		return
	}
	// 链表的第一个元素就是需要删除的元素
	if element.key == key {
		// 将第一个元素后的键值对链上
		h.array[index] = element.netx
		h.len = h.len - 1
		return
	}
	// 下一个键值对
	nextElement := element.netx
	for nextElement != nil {
		if nextElement.key == key {
			// 键值对匹配到，将该键值对从链中去掉,将下一跳的下一跳往前推一位连接上
			element.netx = nextElement.netx
			h.len = h.len -1
			return
		}
		element = nextElement
		nextElement = nextElement.netx
	}
}

// Range 哈希表变量
func (h *HashMap) Range() {
	// 实现并发安全
	h.lock.Lock()
	defer h.lock.Unlock()
	for _, pairs := range h.array {
		for pairs != nil {
			fmt.Printf("%v=%v,", pairs.key, pairs.value)
			pairs = pairs.netx
		}
	}
	fmt.Println()
}

// 测试
func main() {
	// 新建一个哈希表
	hashMap := NewHashMap(16)
	// 放35个值
	for i := 0; i < 35; i++ {
		hashMap.Put(fmt.Sprintf("%d", i), fmt.Sprintf("v%d", i))
	}
	fmt.Println("cap:", hashMap.capacity, "len:", hashMap.Len())
	// 打印全部键值对
	hashMap.Range()

	key := "4"
	value, ok := hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}
	// 删除键
	hashMap.Delete(key)
	fmt.Println("after delete cap:", hashMap.capacity, "len:", hashMap.Len())
	value, ok = hashMap.Get(key)
	if ok {
		fmt.Printf("get '%v'='%v'\n", key, value)
	} else {
		fmt.Printf("get %v not found\n", key)
	}
}

// 总结
// 哈希表查找，是一种用空间换时间的查找算法，时间复杂度能达到：O(1)，
// 最坏情况下退化到查找链表：O(n)。但均匀性很好的哈希算法以及合适空间大小的数组，
// 在很大概率避免了最坏情况。
// 哈希表在添加元素时会进行伸缩，会造成较大的性能消耗，所以有时候会用到其他的查找算法：树查找算法