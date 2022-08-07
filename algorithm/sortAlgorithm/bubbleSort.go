package main

import "fmt"

/*
冒泡排序：
	属于交换类排序，初级排序算法
	平均时间复杂度O(n^2)
	最好时间复杂度O(n)
	最坏时间复杂度O(n^2)
	空间复杂度O(1),原地交换排序
	稳定排序算法，排序不改变原大小的排序位置
	一般不建议使用,效率较低，适用于n较小时
*/

// BubbleSort 常规
func BubbleSort(num []int) {
	// 计算切片长度
	lenth := len(num)
	for i := 0; i < lenth-1; i++ {
		// lenth-i-1 优化，减少无效遍历，因为每排序一次，就可以少一次
		for j := 0; j < (lenth - i - 1); j++ {
			if num[j] > num[j+1] {
				// 从小到大排序，把大的元素放右边
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
}

// BubbleSortFlag 优化到最优，使用标志位记录是否交换，
// 如果一次也没交换说明数组有序，直接退出
func BubbleSortFlag(num []int) {
	// 标志
	var Flag = false
	// 计算切片长度
	lenth := len(num)
	for i := 0; i < lenth-1; i++ {
		// lenth-i-1 优化，减少无效遍历，因为每排序一次，就可以少一次
		for j := 0; j < (lenth - i - 1); j++ {
			if num[j] > num[j+1] {
				// 从小到大排序，把大的元素放右边
				num[j], num[j+1] = num[j+1], num[j]
				Flag = true
			}
		}
		// flag没改变，说明一次也没交换，原数组有序
		if !Flag {
			return
		}
	}
}

func main() {
	// 验证
	num := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	BubbleSort(num)
	fmt.Println(num)
	BubbleSortFlag(num)
	fmt.Println(num)
}
