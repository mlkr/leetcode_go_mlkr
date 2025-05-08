package problem632

import "sort"

func smallestRange(nums [][]int) []int {
	elements := make([]Element, 0)
	for i := range nums {
		for _, num := range nums[i] {

			el := Element{
				num:    num,
				listId: i,
			}

			elements = append(elements, el)
		}
	}

	sort.Slice(elements, func(i, j int) bool {
		return elements[i].num < elements[j].num
	})

	size := len(nums)
	list := make([]int, size)

	begin := 0
	minDiff := 1<<31 - 1
	res := make([]int, 2)
	count := make(map[int]bool)
	for i := range elements {
		listId := elements[i].listId
		list[listId]++
		count[listId] = true

		// 每个list 都包含
		for len(count) == size {

			if minDiff > elements[i].num-elements[begin].num {
				res[0], res[1] = elements[begin].num, elements[i].num
				minDiff = res[1] - res[0]
			}

			begin++
			list[elements[begin].listId]--

			if list[elements[begin].listId] == 0 {
				delete(count, elements[begin].listId)
			}
		}
	}

	return res
}

type Element struct {
	num    int
	listId int
}

// 更好的解法
func smallestRange2(nums [][]int) []int {
	size := len(nums)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = nums[i][0]
	}

	st := getTree(arr)
	res := []int{st.getMin(), st.getMax()}

	curr := make([]int, size)
	for {
		minId := st.getMinId()
		curr[minId]++
		if curr[minId] == len(nums[minId]) {
			break
		}

		st.update(nums[minId][curr[minId]], minId)
		res = getMinRes(res, []int{st.getMin(), st.getMax()})
	}

	return res
}

func getMinRes(curr, new []int) []int {
	if curr[1]-curr[0] > new[1]-new[0] {
		return new
	}

	return curr
}

func getTree(arr []int) SegmentTree {
	pow := 1
	for pow < len(arr) {
		pow *= 2
	}

	nodes := make([]Node, pow*2)
	for i := 0; i < len(arr); i++ {
		id := pow + i
		nodes[id] = Node{max: arr[i], min: arr[i], minId: i}
	}

	for i := len(arr); i < pow; i++ {
		id := pow + i
		nodes[id] = Node{minId: -1}
	}

	for i := pow - 1; i > 0; i-- {
		nodes[i].max = getMax(nodes[i*2], nodes[i*2+1])
		nodes[i].min, nodes[i].minId = getMin(nodes[i*2], nodes[i*2+1])
	}

	return SegmentTree{nodes: nodes, pow: pow}
}

type SegmentTree struct {
	pow   int
	nodes []Node
}

func (st SegmentTree) getMax() int {
	return st.nodes[1].max
}

func (st SegmentTree) getMin() int {
	return st.nodes[1].min
}

func (st SegmentTree) getMinId() int {
	return st.nodes[1].minId
}

func (st SegmentTree) update(val, id int) {
	idx := st.pow + id
	st.nodes[idx].max = val
	st.nodes[idx].min = val
	st.nodes[idx].minId = id

	idx /= 2
	for idx > 0 {
		st.nodes[idx].max = getMax(st.nodes[idx*2], st.nodes[idx*2+1])
		st.nodes[idx].min, st.nodes[idx].minId = getMin(st.nodes[idx*2], st.nodes[idx*2+1])

		idx /= 2
	}
}

type Node struct {
	max, min, minId int
}

func getMax(a, b Node) int {
	if b.minId == -1 || a.max > b.max {
		return a.max
	}

	return b.max
}

func getMin(a, b Node) (int, int) {
	if b.minId == -1 || a.min < b.min {
		return a.min, a.minId
	}

	return b.min, b.minId
}
