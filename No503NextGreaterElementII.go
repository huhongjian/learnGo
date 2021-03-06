package main

import "fmt"

/**
503. 下一个更大元素 II
*/
func nextGreaterElements(nums []int) []int {
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

func main() { // 声明 main 主函数
	//fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	//fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
	fmt.Println(nextGreaterElements([]int{-1, -2, -3, 0}))
}
