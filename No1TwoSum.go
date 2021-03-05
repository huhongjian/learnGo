package main

import (
	"fmt" // 导入 fmt 包，打印字符串是需要用到
)

func twoSum(nums []int, target int) []int {
	var res []int
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				res = append(res, i)
				res = append(res, j)
				return res
			}
		}
	}
	return res
}

func main() { // 声明 main 主函数
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
