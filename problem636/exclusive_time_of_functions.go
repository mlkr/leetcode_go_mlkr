package problem636

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	size := len(logs)
	stack := make([]Node, 0, size)

	res := make([]int, n)
	for i := 0; i < size; i++ {
		node := getNode(logs[i])
		if node.start != -1 {
			// 入栈
			stack = append(stack, node)
		} else {
			// 出栈
			startNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			time := node.end - startNode.start + 1
			if len(stack) > 0 {
				stack[len(stack)-1].cost += time
			}

			res[node.id] += time - startNode.cost
		}
	}

	return res
}

type Node struct {
	id, start, end, cost int
}

func getNode(str string) Node {
	ss := strings.Split(str, ":")
	id, _ := strconv.Atoi(ss[0])
	time, _ := strconv.Atoi(ss[2])

	node := Node{
		id:    id,
		start: -1,
		end:   -1,
	}

	if ss[1] == "start" {
		node.start = time
	}

	if ss[1] == "end" {
		node.end = time
	}

	return node
}
