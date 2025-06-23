package main

import "fmt"

// 指针第一题 num+10
func addTen(num *int) {
	if num == nil {
		return
	}
	*num = *num + 10
}

// 指针第二题 *2
func twoMultiply(nums []int) {
	if nums == nil {
		return
	}
	for i, num := range nums {
		nums[i] = num * 2
	}
}

func main() {
	// 指针第一题 num+10
	var num *int
	// 必须先初始化
	num = new(int)
	*num = 10
	addTen(num)
	fmt.Println(*num)

	// 指针第二题 *2
	var nums = []int{1, 2, 3}
	twoMultiply(nums)
	fmt.Println(nums)
}
