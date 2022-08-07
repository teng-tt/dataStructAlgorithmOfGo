package main

import "fmt"


// fb 递归实现菲波那切数列
func fb(num int) int {
	if num == 1 {
		return 0
	}else if num == 2 {
		return 1
	}
	return fb(num-2) + fb(num-1)
}

func main() {
	c:= fb(9)
	fmt.Println(c)
	num := []int{1,2,3,4,5,6,8,9,11,16}
	b := TwoFind(9, num)
	d := TwoFind(19, num)
	fmt.Println(b, d)
}

// TwoFind 二分查找
func TwoFind(target int, srcArray []int) bool {
	middle := 0
	low := 0
	high := len(srcArray) - 1
	isFind := false
	for ;low <= high ;{
		middle = (high + low) / 2
		if srcArray[middle] == target {
			fmt.Println(target, "在数组中,下标值为: ", middle)
			isFind = true
			return isFind
		}else if srcArray[middle] > target {
			// 说明该数在low~middle之间
			high = middle -1
		}else {
			// 说明该数在middle~high之间
			low = middle + 1
		}
	}
	fmt.Println("数组不含 ", target)
	return isFind
}

func containsDuplicate(nums []int) bool {
	// 定义map变量
	numsMap := make(map[int]int)
	// 遍历数组将值添加到map中，key:数值， value:值出现的次数
	for _, val := range nums {
		// 如果存在，说明重复了
		if _, ok := numsMap[val]; ok {
			return true
		}
		// 不存在,将值加入map
		numsMap[val] += 1
	}
	return false
}

/*
时间复杂度O(N), N 为数组的长度
空间复杂度O(N), 使用了额外的存储空间map数据结构，N 为数组的长度
*/


