package tanxin

import (
	"fmt"
	"sort"
)

type Section struct {
	l, r int
}

const N = 100001

var s [N]*Section

// AcWing905
// 905.区间选点
func AcWing905() {
	var n int
	fmt.Scan(&n)
	var l, r int
	for i := 0; i < n; i++ {
		fmt.Scan(&l, &r)
		s[i] = &Section{l, r}
	}
	// 根据区间右端点从小到大排序
	// 每次都取区间右端点
	// 如果当前取的点不在下一个区间内，则取下一个区间的右端点
	sort.Slice(s[0:n], func(i, j int) bool {
		return s[i].r < s[j].r
	})
	cnt := 1
	p := s[0].r
	for i := 1; i < n; i++ {
		if s[i].l > p {
			cnt++
			p = s[i].r
		}
	}
	fmt.Println(cnt)
}
