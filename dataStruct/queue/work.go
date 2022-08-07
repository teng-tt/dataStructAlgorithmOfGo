package main

import "fmt"

/*
例题：约瑟夫环是一个数学的应用问题，具体为:
已知 n 个人（以编号 1，2，3...n 分别表示）围坐在一张圆桌周围
从编号为 k 的人开始报数，数到 m 的那个人出列
他的下一个人又从 1 开始报数，数到 m 的那个人又出列
依此规律重复下去，直到圆桌周围的人全部出列
*/

func ring(n, m int) {
	var cqueue []int
	// 全部入队列
	for i := 1; i < n; i++ {
		cqueue = append(cqueue, i)
	}

	element := 0  // 记录出队列的元素
	//k := 2        // 报数起始变量 ，指定的k
	i := 1        // 计算变量 i < k
	// 从指定的K开始，出队列,并记录队列元素
	//for ; i < k; i++ {
	//	element = cqueue[0]
	//	cqueue = cqueue[1:]
	//	// i < m 还需要如队列
	//	cqueue = append(cqueue, element)
	//}
	// 从 1 开始,出队列,并记录队列元素
	i = 1 // 计算变量 i < m
	for len(cqueue) > 0 {
		// 出元素
		element = cqueue[0]
		cqueue = cqueue[1:]
		if i < m {
			// i < m 还需要如队列
			cqueue = append(cqueue, element)
			i++
		}else {
			// 输出元素，重新开始
			i = 1
			fmt.Println(element)
		}
	}
}

func main() {
	ring(10, 3)
}