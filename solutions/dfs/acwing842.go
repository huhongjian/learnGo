package dfs

import "fmt"

var n int
var path []int
var st []bool // 标记是否写过了

// AcWing842
// 842. 排列数字
func AcWing842() {
	fmt.Scan(&n)
	path = make([]int, n)
	st = make([]bool, n+1)

	dfs(0)
}

func dfs(u int) {
	// 终止条件
	if u == n {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", path[i])
		}
		fmt.Println()
	}
	for i := 1; i <= n; i++ {
		if st[i] == false {
			path[u] = i
			st[i] = true
			dfs(u + 1)
			// 恢复现场
			// path[u] = 0
			st[i] = false
		}
	}
}
