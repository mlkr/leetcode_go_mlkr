package problem385

import "strconv"

func deserialize(s string) *NestedInteger {
	size := len(s)
	if size == 0 {
		return nil
	}

	res := &NestedInteger{}
	if isDigit(s[0]) {
		num, _ := strconv.Atoi(s)
		res.SetInteger(num)
		return res
	}

	start := 0
	stack := 0
	for i := 1; i < size; i++ {
		if isDigit(s[i]) {
			start = i
			i++
			for i < size && isDigit(s[i]) {
				i++
			}

			num, _ := strconv.Atoi(s[start:i])
			tmp := NestedInteger{}
			tmp.SetInteger(num)
			res.Add(tmp)
		}

		if s[i] == '[' {
			start = i

			stack++
			i++
			for stack > 0 {
				if s[i] == '[' {
					stack++
				}

				if s[i] == ']' {
					stack--
				}

				i++
			}

			res.Add(*deserialize(s[start:i]))
		}
	}

	return res
}

func isDigit(b byte) bool {
	return ('0' <= b && b <= '9') || b == '-'
}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	num int
	ns  []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
// func (n NestedInteger) IsInteger() bool {
// 	return n.ns != nil
// }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
// func (n NestedInteger) GetInteger() int {
// 	return n.num
// }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.num = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.ns = append(n.ns, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
// func (n NestedInteger) GetList() []*NestedInteger {
// 	return n.ns
// }
