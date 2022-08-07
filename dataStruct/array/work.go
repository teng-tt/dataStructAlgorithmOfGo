package main

import "fmt"

/*
//例题，数组存储了 5 个评委对 1 个运动员的打分，且每个评委的打分都不相等
#现在需要你：
1、用数组按照连续顺序保存，去掉一个最高分和一个最低分后的 3 个打分样本
2、计算这 3 个样本的平均分并打印
#要求是，不允许再开辟 O(n) 空间复杂度的复杂数据结构。
*/

func arrayCount(array []int) {
	maxTemp := 0
	maxIndex := 0
	minTemp := 0
	minIndex := 0

	// 获取到最大值，最小值的下标索引值
	for i, v := range array {
		if v > maxTemp {
			maxIndex = i
			maxTemp = v
		}else if v < minTemp {
			minIndex = i
			minTemp = i
		}
	}
	// 从最大值最小值的索引开始分别遍历去掉，最大值，最小值,
	// 通过移动下标实现使整体的数值都向前移动一位
	for i := maxIndex; i < len(array)-1; i++ {
		array[i] = array[i+1]
	}
	fmt.Println("去掉最大值后：", array)
	for i := minIndex; i < len(array)-1; i++ {
		array[i] = array[i+1]
	}
	fmt.Println("去掉最小值后：", array)
	// 获取总分
	sumScore := 0
	for i := 0; i < len(array)-2; i++{
		sumScore += array[i]
	}
	fmt.Println("去掉最大最小后总分：", sumScore)
	// 求平均分
	avgScore := sumScore/3
	fmt.Println("平均值：",avgScore)
}

/*
例题2：给定一个重复的数组，返回去掉重复后的数组，与新数组不重复的长度
，要求空间复杂度为O(1)，就使不在开辟新的内存空间 ，在原有数组上做操作

分析，可以利用数组的索引下标做操作在远数组上进行操作删除
*/
func arrayCounts(array []int) {
	tempValue := 0
	valueIndex := 0
	vauleCount := 0

	// 获取到最大值，最小值的下标索引值
	for i := 0; i < len(array); i++ {
		for i, v := range array {
			// 存在重复的
			if v == tempValue {
				valueIndex = i
				// 从重复值的索引开始,通过移动下标实现使整体的数值都向前移动一位,
				// 所有重复的值都会以索引最后一位的值填充，填充的数量就是重复交换发的数量
				for i := valueIndex; i < len(array)-1; i++ {
					array[i] = array[i+1]
				}
				vauleCount++
			} else {
				// 不存在重复的从这个不重复的值开始，重新开始统计,数组不做操作
				tempValue = v
				valueIndex = 0
				vauleCount = 0
			}
		}
	}
	fmt.Println("去重后：", array, vauleCount)
}

func main() {
	array := []int{2,4,6,8,10}
	array1 := []int{0,0,1,1,1,2,2,3,3,4,4,5}
	arrayCount(array)
	arrayCounts(array1)
}


