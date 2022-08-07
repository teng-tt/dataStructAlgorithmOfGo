package main

import "fmt"

/*
翻转算法，也叫手摇算法，主要用来对数组两部分进行位置互换
比如数组： [9,8,7,1,2,3]，将前三个元素与后面的三个元素交换位置，变成 [1,2,3,9,8,7]。
再比如，将字符串 abcde1234567 的前 5 个字符与后面的字符交换位置，那么手摇后变成：1234567abcde。

如何翻转呢？
将前部分逆序
将后部分逆序
对整体逆序

示例如下：
翻转 [1234567abcde] 的前5个字符。
1. 分成两部分：[abcde][1234567]
2. 分别逆序变成：[edcba][7654321]
3. 整体逆序：[1234567abcde]
*/

func Rotation(array []int, n int) {
	// 分成2部分 1-mid/mid-r
	l := 0
	mid := n
	r := len(array)-1
	reverse(array, l, mid-1)
	fmt.Println("拍完左边后：", array)
	reverse(array, mid, r)
	fmt.Println("拍完右边后：", array)
	reverse(array, l , r)
	fmt.Println("整体拍完后：", array)
}

// 翻转
func reverse(array []int, l , r int) {
	for l < r {
		// 左右互换
		array[l], array[r] = array[r], array[l]
		// 翻转
		l++
		r--
	}
}

func main() {
	// 测试
	array := []int{1,2,3,4,5,6,7,8,9,10}
	Rotation(array, 3)
}