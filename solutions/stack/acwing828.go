package stack

import "fmt"

// AcWing828
// 828. 模拟栈
func AcWing828() {
	var n int
	fmt.Scanf("%d", &n)

	tt := -1
	stack := make([]int, 100010)
	for i := 0; i < n; i++ {
		var op string
		fmt.Scanf("%s", &op)
		if op == "push" {
			var num int
			fmt.Scanf("%d", &num)
			//push
			tt++
			stack[tt] = num

		} else if op == "pop" {
			//pop
			tt--
		} else if op == "empty" {
			if tt > 0 {
				fmt.Println("NO")
			} else {
				fmt.Println("YES")
			}
		} else if op == "query" {
			fmt.Println(stack[tt])
		}

	}

}
