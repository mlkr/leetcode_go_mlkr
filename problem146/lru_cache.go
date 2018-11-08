package problem146

type doubleLinkedNode struct {
	pre, next *doubleLinkedNode
	key, val  int
}

type LRUCache struct {
	len, cap    int
	first, last *doubleLinkedNode
	nodes       map[int]*doubleLinkedNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		nodes: make(map[int]*doubleLinkedNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.nodes[key]
	if !ok {
		return -1
	}

	this.moveToFirst(node)

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.nodes[key]
	if ok {
		node.val = value
		this.moveToFirst(node)
	} else {
		node = &doubleLinkedNode{
			key: key,
			val: value,
		}

		if this.len == this.cap {
			delete(this.nodes, this.last.key)
			this.deleteLast()
		} else {
			this.len++
		}

		this.insertToFirst(node)
		this.nodes[key] = node
	}
}

// 节点移动到最前(先删除,后插入)
func (this *LRUCache) moveToFirst(node *doubleLinkedNode) {
	switch node {
	case this.first:
		return
	case this.last:
		this.deleteLast()
	default:
		node.pre.next = node.next
		node.next.pre = node.pre
	}

	this.insertToFirst(node)
}

func (this *LRUCache) deleteLast() {
	if this.last.pre != nil {
		this.last.pre.next = nil
	} else {
		this.first = nil
	}

	this.last = this.last.pre
}

func (this *LRUCache) insertToFirst(node *doubleLinkedNode) {
	if this.first != nil {
		this.first.pre = node
		node.next = this.first
	} else {
		this.last = node
	}

	this.first = node
}
