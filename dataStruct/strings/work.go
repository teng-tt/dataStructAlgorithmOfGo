package main

import "fmt"

// 字符串
/*
假设要从主串 s = "goodgoogle" 中找到 t = "google" 子串
#根据我们的思考逻辑，则有：
//首先从主串 s 第 1 位开始，判断 s 的第 1 个字符是否与 t 的第 1 个字符相等
//如果不相等，则继续判断主串的第 2 个字符是否与 t 的第1 个字符相等
//直到在 s 中找到与 t 第一个字符相等的字符时，然后开始判断它之后的
//字符是否仍然与 t 的后续字符相等
#如果持续相等直到 t 的最后一个字符，则匹配成功
#如果发现一个不等的字符，则重新回到前面的步骤中
//查找 s 中是否有字符与 t 的第一个字符相等
*/

func isStrSub(sub, src string) int {
	isFind := 0
	// 只需要变量主串，与模式串之间差值的长度，去掉无效的比较节省时间
	for i :=0; i < (len(src) - len(sub) + 1); i++ {
		// 判断首位是否相等
		if src[i] == sub[0] {
			jc := 0 //标记相等的位置
			// 如果相等继续比较接下来的
			for j := 0; j < len(sub); j++ {
				// 不相等终止比较，重新开始比较
				if src[i+j] != sub[j] {
					break
				}
				jc = j
			}
			if jc == len(sub) - 1 {
				isFind = 1
			}
		}
	}
	return isFind
}


// 递归打印斐波那契数列
func fblsit(num int) int {
	if num == 1 {
		return 0
	}
	if num == 2 {
		return 1
	}
	return fblsit(num-1) + fblsit(num-2)
}


func main() {
	src := "goodgodogle"
	sub := "google"

	fmt.Println(isStrSub(sub, src))
	s := "the sky is blue"
	fmt.Println(fblsit(4))
	reverseStr(s)

}


func sum (num int) int {
	return sum(num-1) + num
}

// 字符串反转 "the sky is blue" >> "blue is sky the"
func reverseStr(srcStr string)  {
	var stack []string
	targetStr := ""

	for _, v := range srcStr{
		if v != ' ' {
			targetStr = targetStr + string(v)
		}else {
			stack = append(stack, targetStr)
			targetStr = ""
		}
	}
	for len(stack) != 0 {
		a := stack[len(stack)-1]
		fmt.Printf("%s-", a)
		stack = stack[:len(stack)-1]
	}

}