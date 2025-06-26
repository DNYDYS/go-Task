package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	//136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
	//可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
	//例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
	fmt.Println("第一题------------------------------------------------------------")
	var arr = []int{1, 2, 3, 2, 1}
	oneNum := getOneNum(arr)
	if oneNum > 0 {
		fmt.Println("只出现一次的数字为：", oneNum)
	} else {
		fmt.Println("未找到只出现一次的数字")
	}

	fmt.Println("第二题------------------------------------------------------------")
	//回文数 考察：数字操作、条件判断 题目：判断一个整数是否是回文数
	getPalindrome(12358321)

	fmt.Println("第三题------------------------------------------------------------")
	//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效1
	var brackeStr = "{[()]}"
	isTrue := getBracke(brackeStr)
	if isTrue {
		fmt.Println("字符串有效：", brackeStr)
	} else {
		fmt.Println("字符串无效：", brackeStr)
	}

	fmt.Println("第四题------------------------------------------------------------")
	//查找字符串数组中的最长公共前缀
	strArr := []string{"lamasd中文", "lamb", "lambda", "lambda"}

	shortCommon := getCommonPrefix(strArr)
	fmt.Println("字符串数组中的最长公共前缀为：", shortCommon)

	fmt.Println("第五题------------------------------------------------------------")
	/*
		第五题
		给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
		这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
	*/
	var digits = []int{9, 9, 9, 9, 9}
	digits = getAddOneNum(digits)
	fmt.Println("+1后的结果：", digits)

	fmt.Println("第六题------------------------------------------------------------")
	/*
		删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
		返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
		可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，
		当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
	*/
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	newNums := removeDuplicates(nums)
	fmt.Println("去重后的长度为：", newNums)

	fmt.Println("第七题------------------------------------------------------------")
	// 7.合并区间
	intervals := [][]int{{1, 2}, {3, 6}, {7, 11}, {10, 15}}
	newIntervals := mergeIntervals(intervals)
	fmt.Printf("合并前: %v 合并后: %v\n", intervals, newIntervals)

	fmt.Println("第八题------------------------------------------------------------")
	//两数之和  给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
	nums1 := []int{3, 2, 4}
	target := 6
	newNums1 := getSum(nums1, target)
	fmt.Printf("两数之和: nums = %v, target = %d 结果为: %v\n", nums1, target, newNums1)

}

// 第一题
func getOneNum(arr []int) int {
	var map1 = make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if 0 == map1[arr[i]] {
			map1[arr[i]] = 1
		} else {
			map1[arr[i]] = map1[arr[i]] + 1
		}
	}

	for key, value := range map1 {
		if value == 1 {
			return key
		}
	}
	return -1
}

// 第二题 回文数判断
func getPalindrome(palindrome int) {
	if palindrome < 0 {
		fmt.Println("负数不是回文数", palindrome)
		return
	}

	str := strconv.Itoa(palindrome)

	var newPalindrome string
	for i := len(str); i > 0; i-- {
		newPalindrome = newPalindrome + string(str[i-1])
	}

	if newPalindrome == strconv.Itoa(palindrome) {
		fmt.Println("是回文数", palindrome)
	} else {
		fmt.Println("不是回文数", newPalindrome)
	}

}

// 第三题 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func getBracke(str string) bool {
	// 括号成对出现
	if len(str)%2 == 1 {
		return false
	}
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range str {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) > 0 && stack[len(stack)-1] != bracketMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}

	return false
}

// 第四题 查找字符串数组中的最长公共前缀
func getCommonPrefix(strArr []string) string {
	if len(strArr) == 0 {
		return ""
	}
	// 先获取长度最短的字符串
	var shortPrefix string
	for i := 0; i < len(strArr); i++ {
		if len(shortPrefix) > 0 && len(shortPrefix) < len(strArr[i]) {
			shortPrefix = shortPrefix
		} else {
			shortPrefix = string(strArr[i])
		}
	}
	if len(shortPrefix) == 0 {
		return ""
	}
	var longCommonPrefix string

	// 便利长度最短的字符串和集合中的每一个比对
	for i := 0; i < len(shortPrefix); i++ {
		for _, j := range strArr {
			if shortPrefix[i] != j[i] {
				if i == 0 {
					return ""
				}
				longCommonPrefix = shortPrefix[0:i]
				return longCommonPrefix
			}
		}
	}

	return longCommonPrefix
}

/*
第五题
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
*/
func getAddOneNum(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// 从后往前检查是否需要进位
		if digits[i] == 9 {
			//需要进位
			digits[i] = 0
		} else {
			//不用进位加一即可
			digits[i] += 1
			return digits
		}
		//如果进位到首位，则首元素加一位
		if i == 0 {
			digits = append(digits, 0)
			digits[0] = 1
		}
	}
	return digits

}

/*
第六题
删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，
当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/
func removeDuplicates(nums []int) int {
	var j int
	j = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

/*
第7题
合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	newIntervals := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := newIntervals[len(newIntervals)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			newIntervals = append(newIntervals, current)
		}
	}

	return newIntervals
}

// 第八题 两数之和  给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func getSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		otherNum := target - num

		if idx, exists := numMap[otherNum]; exists {
			return []int{idx, i}
		}

		numMap[num] = i
	}

	return []int{}
}
