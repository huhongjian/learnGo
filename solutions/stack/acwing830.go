package stack

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner *bufio.Scanner

// AcWing830
// 830. 单调栈
func AcWing830() {
	// bufio的性能比scanf好差不多25倍
	scanner = bufio.NewScanner(os.Stdin)
	bs := make([]byte, 20000*1024)
	scanner.Buffer(bs, len(bs))
	scanner.Split(bufio.ScanWords)

	const N = 100010
	stack := make([]int, N) // 单调栈
	var tt int              // 栈顶指针
	var res []string        // 定义结果数组

	var n = readWord() // 读取第一行操作次数
	for i := 0; i < n; i++ {
		var x = readWord()
		for tt > 0 && stack[tt] >= x {
			tt--
		}
		if tt != 0 {
			res = append(res, strconv.Itoa(stack[tt]))
		} else {
			res = append(res, "-1")
		}

		tt++
		stack[tt] = x
	}
	fmt.Println(strings.Join(res, " "))
}

func readWord() int {
	scanner.Scan()
	var res int
	res, _ = strconv.Atoi(scanner.Text())
	return res
}
