package linkedlist

import "fmt"

type link struct {
	head int   // 头节点下标
	e    []int // 数组 e 表示节点 i 的值
	ne   []int // 表示节点 i 的 next 指针
	idx  int   // 存储当前已经用到哪个点，相当于指针
}

// AcWing826
// 826. 单链表
func AcWing826() {
	var m int
	fmt.Scan(&m) // 第一行包含整数M，表示操作次数。
	l := new()
	for i := 0; i < m; i++ { // 接下来M行，每行包含一个操作命令
		var op string
		fmt.Scan(&op)
		switch op {
		case "H":
			var x int
			fmt.Scanf("%d", &x)
			l.addToHead(x)
		case "I":
			var k, x int
			fmt.Scanf("%d%d", &k, &x)
			l.add(k-1, x)
		case "D":
			var k int
			fmt.Scanf("%d", &k)
			if k == 0 {
				l.head = l.ne[l.head]
			} else {
				l.remove(k - 1)
			}
		}
	}
	res := l.traverse(l.head)
	for i := range res {
		fmt.Printf("%d ", res[i])
	}
}

// 初始化
func new() *link {
	return &link{
		head: -1,
		idx:  0,
		e:    make([]int, 100010),
		ne:   make([]int, 100010),
	}
}

// 将 x 插入头节点
func (l *link) addToHead(x int) {
	l.e[l.idx] = x
	l.ne[l.idx] = l.head
	l.head = l.idx
	l.idx++
}

// 将 x 插入下标是 k 的节点后面
func (l *link) add(k, x int) {
	l.e[l.idx] = x
	l.ne[l.idx] = l.ne[k]
	l.ne[k] = l.idx
	l.idx++
}

// 将下标是 k 的节点后面的点删掉
func (l *link) remove(k int) {
	l.ne[k] = l.ne[l.ne[k]]
}

// 遍历链表
func (l *link) traverse(head int) []int {
	var res []int
	for i := head; i != -1; i = l.ne[i] {
		res = append(res, l.e[i])
	}
	return res
}
