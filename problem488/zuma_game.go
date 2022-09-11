package problem488

import "math"

type Key struct {
	board string
	dict  [26]uint8
}

func findMinStep(board string, hand string) int {
	dict := [26]uint8{}
	for _, v := range hand {
		dict[v-'A']++
	}

	mem := make(map[Key]int)

	res := dfs(board, dict, mem)
	if res == math.MaxInt32 {
		res = -1
	}

	return res
}

func dfs(board string, dict [26]uint8, mem map[Key]int) int {
	board = dedup(board)

	key := Key{board, dict}
	if res, ok := mem[key]; ok {
		return res
	}

	size := len(board)
	if size == 0 {
		return 0
	}

	res := math.MaxInt32

	for k, v := range dict {
		if v == 0 {
			continue
		}

		ch := byte(k + 'A')
		for i := 0; i < size; i++ {
			// 如果字母相同，那在前面插入
			if i > 0 && board[i-1] == ch {
				continue
			}

			// 有两种方式插入字母
			// 1.在相同字母的前面插入
			// 2.如果字母不同，在不同字母的中间插入，比如RBR，起到阻断作用
			sameColor := board[i] == ch
			betweenColor := false

			if i > 0 && board[i-1] == board[i] {
				betweenColor = true
			}

			if sameColor || betweenColor {
				dict[k]--
				newBoard := board[:i] + string(ch) + board[i:]
				res = min(res, 1+dfs(newBoard, dict, mem))
				dict[k]++
			}

		}
	}

	mem[key] = res

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 去掉三个以上重复的字母
func dedup(board string) string {
	size := len(board)
	res := make([]byte, 0, size)
	isdup := false

loop:
	for i := 0; i < size; {
		ch := board[i]
		cnt := 0
		for i < size && board[i] == ch {
			i++
			cnt++
		}

		for cnt < 3 {
			res = append(res, ch)
			cnt--

			if cnt == 0 {
				continue loop
			}
		}

		// 有三个以上重复的字母
		isdup = true
	}

	if isdup {
		return dedup(string(res))
	}

	return string(res)
}
