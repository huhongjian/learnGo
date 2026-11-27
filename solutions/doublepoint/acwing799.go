package doublepoint

import "fmt"

// AcWing799
// 799. 最长连续不重复子序列
func AcWing799() {
	var n int
	fmt.Scanf("%d", &n)
	q := make([]int, n+1)
	// 每个数出现的次数
	// 起到一个map的作用
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	res := 0
	for i, j := 0, 0; i < n; i++ {
		s[q[i]]++
		for j < i && s[q[i]] > 1 {
			s[q[j]]--
			j++
		}
		// max
		if res < i-j+1 {
			res = i - j + 1
		}
	}
	fmt.Println(res)
}
