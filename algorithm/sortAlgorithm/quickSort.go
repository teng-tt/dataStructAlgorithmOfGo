package main

import (
	"fmt"
	"sync"
)

/*
快速排序：
	属于交换类排序，高级排序算法
	平均时间复杂度O(nlogn)
	最好时间复杂度O(nlogn)
	最坏时间复杂度O(n^2)
	空间复杂度O(1),原地交换排序
	不稳定排序算法，排序后改变原大小的排序位置
	适用于n较大时

快速排序算法介绍：
快速排序通过一趟排序将要排序的数据分割成独立的两部分，
其中一部分的所有数据都比另外一部分的所有数据都要小，
然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，
以此达到整个数据变成有序序列。

步骤如下：
1、先从数列中取出一个数作为基准数。一般取第一个数。
2、分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边。
3、再对左右区间重复第二步，直到各区间只有一个数
快速排序主要靠基准数进行切分，将数列分成两部分，一部分比基准数都小，一部分比基准数都大

算法优化：
切分的结果极大地影响快速排序的性能，比如每次切分的时候选择的基数数都是数组中最大或者最小的，
会出现最坏情况，为了避免切分不均匀情况的发生，有几种方法改进：

随机基准数选择：每次进行快速排序切分时，先将数列随机打乱，再进行切分，这样随机加了个震荡，减少不均匀的情况。
当然也可以随机选择一个基准数，而不是选第一个数。
中位数选择：每次取数列头部，中部，尾部三个数，取三个数的中位数为基准数进行切分。
方法 1 相对好，而方法 2 引入了额外的比较操作，一般情况下我们可以随机选择一个基准数。
从一个数组中随机选择一个数，或者取中位数都比较容易实现
#下文代码实现都取第一个数为基准数
*/
// 快速排序，每一次切分都维护两个下标，进行推进，最后将数列分成两部分。
// 切分函数
func partition(array []int, begin, end int) int {
	// 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
	i := begin + 1
	j := end
	// 没重合之前
	for i < j {
		if array[i] > array[begin] {
			// 比基准数大，交换
			array[i], array[j] = array[j], array[i]
			j--
		} else {
			// 不交换
			i++
		}
	}
	/* 跳出 for 循环后，i = j
	  此时数组被分割成两个部分--> array[begin+1] ~ array[i-1] < array[begin]
						 -->  array[i+1] ~ array[end] > array[begin]
	  这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置
	  最后将array[i]与array[begin]交换，进行两个分割部分的排序！
	  以此类推，直到最后i = j不满足条件就退出！
	*/
	if array[i] >= array[begin] {
		// 这里必须要取等“>=”，否则数组元素由相同的值组成时，会出现错误！
		i--
	}
	array[begin], array[i] = array[i], array[begin]
	return i
}

// QuickSort 算法实现，最普通版本
// QuickSort 普通快速排序
func QuickSort(array []int, begin, end int) {
	if begin < end {
		// 进行切分
		loc := partition(array, begin, end)
		// 对左边进行快排
		QuickSort(array, begin, loc-1)
		// 对右边进行快排
		QuickSort(array, loc+1, end)
	}
}

/*三、算法改进
快速排序可以继续进行算法改进。
1、在小规模数组的情况下，直接插入排序的效率最好，当快速排序递归部分进入小数组范围，可以切换成直接插入排序。
2、排序数列可能存在大量重复值，使用三向切分快速排序，将数组分成三部分，大于基准数，等于基准数，小于基准数，这个时候需要维护三个下标。
3、使用伪尾递归减少程序栈空间占用，使得栈空间复杂度从 O(logn) ~ log(n) 变为：O(logn)
*/

// 改进1：小规模数组使用直接插入排序
// 在小规模数组的情况下，直接插入排序的效率最好，当快速排序递归部分进入小数组范围，
// 可以切换成直接插入排序

// InsertSort 改进：当数组规模小时使用直接插入排序
func InsertSort(list []int) {
	n := len(list)
	// 进行N-1轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置
		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		for deal > list[j] {
			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j] // 某数后移，给待排序留空位
			}
			list[j+1] = deal // 结束了，待排序的数插入空位
		}
	}
}

func QuickSort1(array []int, begin, end int) {
	if begin < end {
		// 当数组小于 4 时使用直接插入排序
		// 直接插入排序在小规模数组下效率极好，我们只需将 end-begin <= 4
		// 的递归部分换成直接插入排序，这部分表示小数组排序
		if end-begin <= 4 {
			InsertSort(array[begin : end+1])
			return
		}
		// 进行切分
		loc := partition(array, begin, end)
		// 对左部分进行快排
		QuickSort1(array, begin, loc-1)
		// 对右部进行快排
		QuickSort1(array, loc+1, end)
	}
}

// 改进2：三向切片
// 排序数列可能存在大量重复值，使用三向切分快速排序，将数组分成三部分，
// 大于基准数，等于基准数，小于基准数
// 这个时候需要维护三个下标
// 切分函数，并返回切分元素的下标
func partition3(array []int, begin, end int) (int, int) {
	lt := begin       // 左下标从第一位开始
	gt := end         // 右下标是数组的最后一位
	i := begin + 1    // 中间下标，从第二位开始
	v := array[begin] // 基准数

	// 以中间坐标为准
	for i <= gt {
		if array[i] > v {
			// 大于基准数，那么交换，右指针左移
			array[i], array[gt] = array[gt], array[i]
			gt--
		} else if array[i] < v {
			// 小于基准数，那么交换，左指针右移
			array[i], array[lt] = array[lt], array[i]
			lt++
			i++
		} else {
			i++
		}
	}
	return lt, gt
}

// QuickSort2 三切分的快速排序
//	三向切分后第一轮结果后,分成三个数列
//	(元素相同的会聚集在中间数列),中间数为与基准相等的
//	接着对第一个和最后一个数列进行递归即可
// 三切分，把小于基准数的扔到左边，大于基准数的扔到右边，相同的元素会进行聚集。
// 如果存在大量重复元素，排序速度将极大提高，将会是线性时间，因为相同的元素将会聚集在中间
// 这些元素不再进入下一个递归迭代
func QuickSort2(array []int, begin, end int) {
	if begin < end {
		// 三向切分函数，返回左边和右边下标
		lt, gt := partition3(array, begin, end)
		// 从lt到gt的部分是三切分的中间数列
		// 左边三向快排
		QuickSort2(array, begin, lt-1)
		// 右边三向快排
		QuickSort2(array, gt+1, end)
	}
}

// 改进三：伪尾递归优化
// 使用伪尾递归减少程序栈空间占用，
// 使得栈空间复杂度从 O(logn) ~ log(n) 变为：O(logn)
// 这样的快排写法是伪装的尾递归，不是真正的尾递归，因为有 for 循环
// 不是直接 return QuickSort，递归还是不断地压栈，栈的层次仍然不断地增长。
// 但因先让规模小的部分排序，栈的深度大大减少，程序栈最深不会超过 logn 层
// 这样堆栈最坏空间复杂度从 O(n) 降为 O(logn)。
// 这种优化也是一种很好的优化，因为栈的层数减少了，对于排序十亿个整数，
// 也只要：log(100 0000 0000)=29.897，占用的堆栈层数最多 30 层，
// 比不进行优化，可能出现的 O(n) 常数层好很多。

// QuickSort3 伪尾递归快速排序
func QuickSort3(array []int, begin, end int) {
	// 进行切分
	loc := partition(array, begin, end)
	// 那边元素少先排那边
	if loc-begin < end-loc {
		// 先排序左边
		QuickSort3(array, begin, loc-1)
		begin = loc + 1
	} else {
		// 先排序右边
		QuickSort3(array, loc+1, end)
		end = loc - 1
	}
}

// 非递归写法
// 非递归写法仅仅是将之前的递归栈转化为自己维持的手工栈
// 使用人工栈替代递归的程序栈，换汤不换药，速度并没有什么变化，但是代码可读性降低。

// LinkStack 定义LinkStack 链表栈，后进先出
type LinkStack struct {
	root *LinkNode  // 链表起点
	size int        // 栈的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// LinkNode 定义链表节点
type LinkNode struct {
	Next  *LinkNode
	Value int
}

// Push 入栈
func (stack *LinkStack) Push(v int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 如果栈顶为空，那么增加节点
	if stack.size == 0 {
		stack.root = new(LinkNode)
		stack.root.Value = v
	} else {
		// 否则新元素插入链表的头部
		// 原来的链表
		preNode := stack.root
		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v
		// 原来的链表链接到新元素后面
		newNode.Next = preNode
		// 将新节点放在头部
		stack.root = newNode
	}
	// 栈中元素数量+1
	stack.size = stack.size + 1
}

// Pop 出栈
func (stack *LinkStack) Pop() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}
	// 顶部元素要出栈
	topNode := stack.root
	v := topNode.Value
	// 将顶部元素的后继链接链上
	stack.root = topNode.Next
	// 栈中元素数量-1
	stack.size = stack.size - 1
	return v
}

// IsEmpty 栈是否为空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

// QuickSort5 非递归快速排序
func QuickSort5(array []int) {
	// 人工栈
	helpStack := new(LinkStack)
	// 第一次初始化栈，推入下标0，len(array)-1，表示第一次对全数组范围切分
	helpStack.Push(len(array) - 1)
	helpStack.Push(0)
	// 栈非空证明存在未排序的部分
	for !helpStack.IsEmpty() {
		// 出栈，对begin-end范围进行切分排序
		begin := helpStack.Pop() // 范围区间左边
		end := helpStack.Pop()   // 范围
		// 进行切分
		loc := partition(array, begin, end)
		// 右边范围入栈
		if loc+1 < end {
			helpStack.Push(end)
			helpStack.Push(loc + 1)
		}
		// 左边返回入栈
		if begin < loc-1 {
			helpStack.Push(loc - 1)
			helpStack.Push(begin)
		}
	}
}

// 本来需要进行递归的数组范围 begin,end，不使用递归，
// 依次推入自己的人工栈，然后循环对人工栈进行处理。
// 可以看到没有递归，程序栈空间复杂度变为了：O(1)，但额外的存储空间产生了。
// 辅助人工栈结构 helpStack 占用了额外的空间，存储空间由原地排序的 O(1)
// 变成了 O(logn) ~ log(n)。
// 参考上面伪尾递归版本，继续优化，先让短一点的范围入栈，这样存储复杂度可以变为：O(logn)

// QuickSort6 非递归版本优化
func QuickSort6(array []int) {
	// 人工栈
	helpStack := new(LinkStack)
	// 第一次初始化栈，推入下标0，len(array)-1，表示第一次对全数组范围切分
	helpStack.Push(len(array) - 1)
	helpStack.Push(0)
	// 栈非空证明存在未排序的部分
	for !helpStack.IsEmpty() {
		// 出栈，对begin-end范围进行切分排序
		begin := helpStack.Pop() // 范围区间左边
		end := helpStack.Pop()   // 范围
		// 进行切分
		loc := partition(array, begin, end)
		// 切分后右边范围大小
		rSize := -1
		// 切分后左边范围大小
		lSize := -1
		// 右边范围入栈
		if loc+1 < end {
			rSize = end - (loc + 1)
		}
		// 左边返回入栈
		if begin < loc-1 {
			lSize = loc - 1 - begin
		}
		// 两个范围，让范围小的先入栈，减少人工栈空间
		if rSize != -1 && lSize != -1 {
			if lSize > rSize {
				helpStack.Push(end)
				helpStack.Push(loc + 1)
				helpStack.Push(loc - 1)
				helpStack.Push(begin)
			} else {
				helpStack.Push(loc - 1)
				helpStack.Push(begin)
				helpStack.Push(end)
				helpStack.Push(loc + 1)
			}
		} else {
			if rSize != -1 {
				helpStack.Push(end)
				helpStack.Push(loc + 1)
			}
			if lSize != -1 {
				helpStack.Push(loc - 1)
				helpStack.Push(begin)
			}
		}
	}
}

// 测试
func main() {
	list := []int{5}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)

	list1 := []int{5, 9}
	QuickSort(list1, 0, len(list1)-1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1}
	QuickSort(list2, 0, len(list2)-1)
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(list3, 0, len(list3)-1)
	fmt.Println(list3)

	list4 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort5(list4)
	fmt.Println(list4)

	list5 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort6(list5)
	fmt.Println(list5)
}


/*
补充：内置库使用快速排序的原因
	首先堆排序，归并排序最好最坏时间复杂度都是：O(nlogn)，
	而快速排序最坏的时间复杂度是：O(n^2)，
	但是很多编程语言内置的排序算法使用的仍然是快速排序，这是为什么？
	这个问题有偏颇，选择排序算法要看具体的场景，
	Linux 内核用的排序算法就是堆排序，而 Java 对于数量比较多的复杂对象排序，内置排序使用的是归并排序，
	只是一般情况下，快速排序更快。
	归并排序有两个稳定，
		第一个稳定是排序前后相同的元素位置不变，
		第二个稳定是，每次都是很平均地进行排序，读取数据也是顺序读取，
		能够利用存储器缓存的特征，比如从磁盘读取数据进行排序。因为排序过程需要占用额外的辅助数组空间，
		所以这部分有代价损耗，但是原地手摇的归并排序克服了这个缺陷。
	复杂度中，大 O 有一个常数项被省略了，堆排序每次取最大的值之后，都需要进行节点翻转，重新恢复堆的特征，
	做了大量无用功，常数项比快速排序大，大部分情况下比快速排序慢很多。
	但是堆排序时间较稳定，不会出现快排最坏 O(n^2) 的情况，且省空间，不需要额外的存储空间和栈空间。
	当待排序数量大于16000个元素时，使用自底向上的堆排序比快速排序还快，可见此：https://core.ac.uk/download/pdf/82350265.pdf。
	快速排序最坏情况下复杂度高，主要在于切分不像归并排序一样平均，而是很依赖基准数的现在，
	我们通过改进，比如随机数，三切分等，这种最坏情况的概率极大的降低。
	大多数情况下，它并不会那么地坏，大多数快才是真的块。
	归并排序和快速排序都是分治法，排序的数据都是相邻的，
	而堆排序比较的数可能跨越很大的范围，导致局部性命中率降低，不能利用现代存储器缓存的特征，加载数据过程会损失性能。
	对稳定性有要求的，要求排序前后相同元素位置不变，可以使用归并排序，
	Java 中的复杂对象类型，要求排序前后位置不能发生变化，所以小规模数据下使用了直接插入排序，
	大规模数据下使用了归并排序。
	对栈，存储空间有要求的可以使用堆排序，比如 Linux 内核栈小，快速排序占用程序栈太大了，
	使用快速排序可能栈溢出，所以使用了堆排序
*/