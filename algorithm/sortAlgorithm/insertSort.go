package main

import "fmt"

/*
插入排序：
	属于插入类排序，初级排序算法
	平均时间复杂度O(n^2)
	最好时间复杂度O(n)
	最坏时间复杂度O(n^2)
	空间复杂度O(1),原地交换排序
	稳定排序算法，排序不改变原大小的排序位置
	适用于数组较有序或n较小时

其实，就是通过将右边的一个数 deal（待排序的数） ，找到它的左边 已排好序的数列 位置，然后插进去。
数组规模 n 较小的大多数情况下，我们可以使用插入排序，它比冒泡排序，选择排序都快，甚至比任何的排序算法都快。
数列中的有序性越高，插入排序的性能越高，因为待排序数组有序性越高，插入排序比较的次数越少。
一般很少使用冒泡、直接选择，直接插入排序算法，因为在有大量元素的无序数列下，这些算法的效率都很低。
*/

func InsertSort(num []int) {
	lenth := len(num) // 获取长度
	// 进行N-1轮迭代
	for i := 1; i <= lenth-1; i ++ {
		deal := num[i]  //待排序的数
		j := i - 1 // 待排序数左边的第一个数的位置
		// 如果第一次比较，比左边已排好序的第一个数小，那么进入处理
		if deal < num[j] {
			// 一直往左边查找，比待排序大的数都往后挪，腾空位置给待排序插入
			for ; j >= 0 && deal < num[j]; j-- {
				num[j+1] = num[j] // 某数后移，给待排序留空位
			}
			num[j+1] = deal // 结束了，待排序的数插入空位
		}
	}
}


func main() {
	// 测试
	list := []int{5}
	InsertSort(list)
	fmt.Println(list)

	list1 := []int{5, 9}
	InsertSort(list1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	InsertSort(list2)
	fmt.Println(list2)
}