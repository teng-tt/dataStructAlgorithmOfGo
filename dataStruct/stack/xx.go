package main

import "fmt"

var nums = []int{1,2,3,4,5,6}
func raveres(count int) []int {
	lenth := len(nums)
	var newNum []int
	var c []int
	n := 0
	for _, vaule := range nums {
		newNum = append(newNum, vaule)
		n ++
		if n > count{
			for j := 0; j < count; j++ {
				pop := newNum[len(newNum)-1]
				newNum = newNum[:len(newNum)-1]
				c = append(c, pop)
			}

		}

	}
	return c
}

func main() {
	b := raveres(3)
	fmt.Printf("%#v\n", b)
}