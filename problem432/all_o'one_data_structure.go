package problem432

// 参看
// https://zhuanlan.zhihu.com/p/111861747
type Node struct {
	val       int
	keys      map[string]struct{}
	pre, next *Node
}

type AllOne struct {
	m                 map[string]*Node
	begin, head, tail *Node
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	begin := &Node{0, make(map[string]struct{}), nil, nil}
	return AllOne{
		make(map[string]*Node),
		begin,
		begin,
		begin,
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	node, ok := this.m[key]
	if !ok {
		if this.begin.next == nil || this.begin.next.val != 1 {
			// 插入节点
			keys := make(map[string]struct{})
			keys[key] = struct{}{}
			node = &Node{1, keys, this.begin, this.begin.next}

			if this.begin.next != nil {
				this.begin.next.pre = node
			}
			this.begin.next = node
			this.head = node
		} else {
			// 更新节点
			node = this.head
			node.keys[key] = struct{}{}
		}

		if this.tail == this.begin {
			this.tail = this.head
		}

	} else {
		val := node.val + 1
		delete(node.keys, key)
		if len(node.keys) == 0 {
			// 删除节点

			// 更新 head tail
			if this.head == node {
				if this.head.next != nil {
					this.head = this.head.next
				} else {
					this.head = node.pre
				}
			}

			if this.tail == node {
				this.tail = node.pre
			}

			temp := node.pre
			temp.next = node.next
			if node.next != nil {
				node.next.pre = temp
			}
			node = temp
		}

		if node.next == nil || node.next.val != val {
			// 插入节点
			keys := make(map[string]struct{})
			keys[key] = struct{}{}
			temp := &Node{val, keys, node, node.next}

			if node.next != nil {
				node.next.pre = temp
			}
			node.next = temp
			node = node.next
		} else {
			// 更新节点
			node.next.keys[key] = struct{}{}
			node = node.next
		}

		// 更新 head tail
		if this.head == this.begin || this.head.val > node.val {
			this.head = node
		}

		if this.tail.val < node.val {
			this.tail = node
		}

	}

	this.m[key] = node
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	node, ok := this.m[key]
	if !ok {
		return
	}

	val := node.val - 1
	preNode := node.pre

	delete(node.keys, key)
	if len(node.keys) == 0 {
		// 删除节点

		// 更新 head tail
		if this.head == node {
			if this.head.next != nil {
				this.head = this.head.next
			} else {
				this.head = node.pre
			}
		}

		if this.tail == node {
			this.tail = node.pre
		}

		preNode.next = node.next
		if node.next != nil {
			node.next.pre = preNode
		}
	}

	if preNode.val != val {
		// 插入节点
		keys := make(map[string]struct{})
		keys[key] = struct{}{}
		temp := &Node{val, keys, preNode, preNode.next}

		if preNode.next != nil {
			preNode.next.pre = temp
		}
		preNode.next = temp
		node = temp
	} else {
		// 更新节点
		if preNode != this.begin {
			preNode.keys[key] = struct{}{}
		}
		node = preNode
	}

	// todo 更新 head tail
	if (this.head == this.begin || this.head.val > node.val) &&
		node != this.begin {

		this.head = node
	}

	if this.tail.val < node.val {
		this.tail = node
	}

	this.m[key] = node
	if node == this.begin {
		delete(this.m, key)
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.tail == this.begin {
		return ""
	}

	res := ""
	for i := range this.tail.keys {
		res = i
		break
	}

	return res
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.head == this.begin {
		return ""
	}

	res := ""
	for i := range this.head.keys {
		res = i
		break
	}

	return res
}
