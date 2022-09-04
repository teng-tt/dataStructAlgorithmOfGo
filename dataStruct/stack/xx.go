package main

import "fmt"

/*
例2
给定一个包含n个元素的数组，现在要求每K个结点一组进行反转，打印反转后的数组结果
其中k是一个正整数，且n可以被k整除
列如：数组为[1,2,3,4,5,6], k=3 则打印321654
*/

func raveres(nums []int, count int) []int {
	var newNum []int
	var resultArray []int
	lenth := len(nums)
	n := 0
	for i := 0; i < lenth; i++{
		// 按指定的反转个数入栈
		newNum = append(newNum, nums[i])
		n++
		// 判断标记数是否等于指定反转个数，不等于继续入栈，等于出栈
		if n < count {
			continue
		}else {
			// 入栈个数大于指定的反转个数，出栈
			for j := 0; j < count; j++ {
				pop := newNum[len(newNum)-1]
				newNum = newNum[:len(newNum)-1]
				// 出栈元素追加到存储翻转后元素的新切片
				resultArray = append(resultArray, pop)
			}
			n = 0
		}
	}
	// 返回按指定个数翻转后的新切片
	return resultArray
}

func main() {
	nums := []int{1,2,3,4,5,6}
	b := raveres(nums, 2)
	fmt.Printf("%#v\n", b)
}