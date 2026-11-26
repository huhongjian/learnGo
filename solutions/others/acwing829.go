package others

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner *bufio.Scanner

// AcWing829
// 829. 模拟队列
func AcWing829() {
	scanner = bufio.NewScanner(os.Stdin)
	bs := make([]byte, 20000*1024)
	scanner.Buffer(bs, len(bs))

	const N = 100010
	q := make([]int, N)
	hh, tt := 0, -1

	var m int
	x := readLine()[0] // 读取第一行操作次数
	m, _ = strconv.Atoi(x)
	for i := 0; i < m; i++ {
		x := readLine()
		switch x[0] {
		case "push":
			val, _ := strconv.Atoi(x[1])
			tt++
			q[tt] = val
		case "pop":
			hh++
		case "empty":
			if hh > tt {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		case "query":
			fmt.Println(q[hh])
		}
	}

}

func readLine() []string {
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	res := make([]string, len(line))
	for i, l := range line {
		res[i] = l
	}
	return res
}
