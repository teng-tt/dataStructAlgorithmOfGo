package main

import "fmt"

// 输入a = [1, 2, 3, 4, 5],输出[5, 4, 3, 2, 1]
// 方法一 时间复杂度是O(n) + O(n) 也就是O(n)，空间复杂度是O(n)
func s1_1() {
	a := []int{1,2,3,4,5}
	b := make([]int, 5)
	for i := 0; i < len(a); i++{
		b[i] = a[i]
	}
	// 利用b的逆序索引顺序实现，把a正序索引值的数据赋值
	for j := 0; j < len(a); j++{
		b[len(a)-j-1] = a[j]
	}
	fmt.Printf("%v\n", b)
}

func s1_2(){
	a := []int{1,2,3,4,5}
	tmp := 0
	// len(a)/2 因为一个数组只需要比较1/2次就可以出结果
	for i := 0; i < len(a)/2; i++{
		// 保存正序遍历的临时值
		tmp = a[i]
		// 最后一个赋值给第一个
		a[i] = a[len(a)-i-1]
		// 第一个赋值给最后一个
		a[len(a)-i-1] = tmp
	}
	fmt.Printf("%v\n", a)
}

//在一个数组中找出出现次数最多的那个元素的值
//例如：输入数组a = [1, 2, 3, 4, 5, 5, 6]
func numCount() {
	a := []int{1,2,3,4,5,5,6,6,6,2,2,2,2}
	b := map[int]int{}
	// o(n)
	for _, value := range a {
		b[value]++
	}
	maxCount := 0
	maxKey := 0
	// o(n)
	for key, value := range b {
		if value > maxCount {
			maxCount = value
			maxKey = key
		}
	}
	// o(n) + o(n) = o(n)
	fmt.Println(maxKey, ":",b[maxKey])
}

// 求阶乘常规递归算法，使用栈保存，计算过程先进后出，需要消耗大量的栈空间，如果没有
// 终止条件，栈的层数无线增加，会栈溢出
func rescuvie(n int) int {
	if n == 0 {
		return 1
	}
	return rescuvie(n-1) * n
}

// 使用尾递归，节约堆栈的空间，减少堆栈层数
// 函数在调用自身后直接传回其值，而不对其再加运算，效率将会极大的提高。
// 如果一个函数中所有递归形式的调用都出现在函数的末尾，我们称这个递归函数是尾递归的
// 当递归调用是整个函数体中最后执行的语句且它的返回值不属于表达式的一部分时，这个递归调用就是尾递归
func rescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}
	return rescuvieTail(n-1, n*a)
}

// 尾递归求解斐波拉切数列,n为第几位斐波拉切数
// 所有的结果都由a，b = 1, 1初始值累加而来
// 当 n=5 的递归过程如下:
/* F(5,1,1)
F(4,1,1+1)=F(4,1,2)
F(3,2,1+2)=F(3,2,3)
F(2,3,2+3)=F(2,3,5)
F(1,5,3+5)=F(1,5,8)
F(0,8,5+8)=F(0,8,13)
8
*/
func fblist(n, a, b int) int {
	if n == 0 {
		return a
	}
	return fblist(n-1, b, a+b)
}

// 二分查找
func binarySearch(array []int, target, left, right int) int {
	if left > right {
		// 遍历完毕，出界了找不到
		return -1
	}
	// 从中间开始查找
	mid := (left + right) / 2
	// 获取中间值
	middleNum := array[mid]
	// 如果相等，返回找到了
	if target == middleNum {
		return mid
	}else if target > middleNum {
		// 中间值比目标值还小，从右边区间开始查找
		return binarySearch(array, target, mid+1, right)
	}else {
		// 中间值比目标值还大，从左边边区间开始查找
		return binarySearch(array, target, 0, mid-1)
	}
}

// 很多计算机问题都可以用递归来简化求解，
// 理论上，所有的递归方式都可以转化为非递归的方式，不过使用递归，代码的可读性更高
// 非递归实现
func binarySerach2(array []int, target, l, r int) int {
	templ := l
	tempr := r
	for {
		// 判断是否查找完毕越界
		if templ > templ {
			return -1
		}
		// 从中间开始查找
		mid := (tempr+templ) / 2
		midNum := array[mid]
		if target == midNum {
			return mid // 找到了
		}else if target > midNum {
			// 中间的数比目标还小，从右边找
			templ = mid + 1
		}else {
			// 中间的数比目标还大，从左边找
			tempr = mid - 1
		}
	}
}

func main() {
	s1_1()
	s1_2()
	numCount()

	array := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 500
	result := binarySerach2(array, target, 0, len(array)-1)
	fmt.Println(target, result)

	target = 189
	result = binarySerach2(array, target, 0, len(array)-1)
	fmt.Println(target, result)
}






