package main

import "fmt"

/**
503. 下一个更大元素 II

其实是一个单调栈的题目
*/
func nextGreaterElements1(nums []int) []int {
	var len = len(nums)
	var res []int
	for i := 0; i < len; i++ {
		res = append(res, -1)
	}
	if len <= 1 {
		return res
	}
	i := 0
	j := 1
	for i < len {
		for j < len+i {
			if nums[i] < nums[j%len] {
				res[i] = nums[j%len]
				break
			}
			if res[j%len] != -1 && nums[i] < res[j%len] {
				res[i] = res[j%len]
				break
			}
			j++
		}
		i++
		j = i + 1
	}
	return res
}

/**
单调栈解法
*/
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	stack := []int{}
	for i := 0; i < n*2-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return ans
}

func main() { // 声明 main 主函数
	//fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	//fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
	fmt.Println(nextGreaterElements([]int{-1, -2, -3, 0}))
}
