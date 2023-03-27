package main

import "fmt"

/*
选择排序：
	属于选择类排序，初级排序算法
	平均时间复杂度O(n^2)
	最好时间复杂度O(n^2)
	最坏时间复杂度O(n^2)
	空间复杂度O(1),原地交换排序
	不稳定排序算法，排序改变原大小的排序位置
	一般不建议使用,效率较低，适用于n较小时
*/

// SelectSort 普通版本
// 每进行一轮迭代，我们都会维持这一轮最小数：minValue 和最小数的下标：minValueIndex
// 然后开始扫描，如果扫描的数比该数小，那么替换掉最小数和最小数下标
// 扫描完后判断是否应交换，然后交换：num[i], num[minValueIndex] = num[minValueIndex], num[i]
func SelectSort(num []int) {
	lenth := len(num) // 长度

	// 进行N-1轮次迭代
	for i := 0; i < lenth-1; i++ {
		minValueIndex := i // 最小值索引
		minValue := num[i] // 最小值
		for j := i + 1; j < lenth; j++ {
			// 获取最小值，比最小还小就是新的最小值，和最小值下标
			if num[j] < minValue {
				minValue = num[j]
				minValueIndex = j
			}
		}
		// 如果这轮最小值的下标不等于最开始的下标，说明有新的最小值，交换元素
		if i != minValueIndex {
			num[i], num[minValueIndex] = num[minValueIndex], num[i]
		}
	}
}

// SelectGoodSort 优化版本
// 上面的算法需要从某个数开始，一直扫描到尾部
// 可以优化算法，使得复杂度减少一半，每一轮，除了找最小数之外，还找最大数
// 然后分别和前面和后面的元素交换，这样循环次数减少一半
func SelectGoodSort(num []int) {
	lenth := len(num) // 获取数组长度
	// 只循环一半
	for i := 0; i < lenth/2; i++ {
		minIndex := i // 最小值下标
		maxIndex := i // 最大值下标
		// 在这一轮迭代中要找到最大值和最小值的下标
		for j := i + 1; j < lenth; j++ {
			// 找到最大值下标
			if num[j] > num[maxIndex] {
				maxIndex = j // 这一轮这个是大的，直接 continue
				continue
			}
			// 找到最小值下标
			if num[j] < num[minIndex] {
				minIndex = j
			}
		}
		// num[i] 为正向递增元素，num[lenth-i-1] 为最后一个元素（迭代n-1次所以-1）
		if maxIndex == i && minIndex != lenth-i-1 {
			// 如果最大值是开头的元素，而最小值不是最尾的元素
			// 先将最大值和最尾的元素交换
			num[lenth-i-1], num[maxIndex] = num[maxIndex], num[lenth-i-1]
			// 然后最小的元素放在最开头
			num[i], num[minIndex] = num[minIndex], num[i]

		} else if maxIndex == i && minIndex == lenth-i-1 {
			// 如果最大值在开头，最小值在结尾，直接交换
			num[maxIndex], num[minIndex] = num[minIndex], num[maxIndex]

		} else {
			// 否则先将最小值放在开头，再将最大值放在结尾
			num[i], num[minIndex] = num[minIndex], num[i]
			num[lenth-i-1], num[maxIndex] = num[maxIndex], num[lenth-i-1]
		}

	}
}

func main() {
	// 测试
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectSort(list)
	fmt.Println(list)

	list = []int{5}
	SelectGoodSort(list)
	fmt.Println(list)

	list1 := []int{5, 9}
	SelectGoodSort(list1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1}
	SelectGoodSort(list2)
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	SelectGoodSort(list3)
	fmt.Println(list3)

	list4 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6}
	SelectGoodSort(list4)
	fmt.Println(list4)
}
