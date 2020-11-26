package problem460

import (
	"container/heap"
)

type LFUCache struct {
	pq        PQ
	m         map[int]*Entry
	cap, time int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		pq:  make(PQ, 0, capacity),
		m:   make(map[int]*Entry, capacity),
		cap: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	this.time++

	entry, ok := this.m[key]
	if !ok {
		return -1
	}

	val := entry.val
	this.pq.update(entry, this.time)

	return val
}

func (this *LFUCache) Put(key int, value int) {
	this.time++

	if this.cap <= 0 {
		return
	}

	entry, ok := this.m[key]
	if ok {
		entry.val = value
		this.pq.update(entry, this.time)
		return
	}

	if len(this.pq) == this.cap {
		temp := heap.Pop(&this.pq)
		delete(this.m, temp.(*Entry).key)
	}

	entry = &Entry{key: key, val: value, time: this.time}
	this.m[key] = entry
	heap.Push(&this.pq, entry)

}

type Entry struct {
	key, val, index, freq, time int
}

type PQ []*Entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	if pq[i].freq == pq[j].freq {
		return pq[i].time < pq[j].time
	}

	return pq[i].freq < pq[j].freq
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	entry := x.(*Entry)
	entry.index = n
	*pq = append(*pq, entry)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	entry := old[n-1]
	entry.index = -1
	*pq = old[:n-1]

	return entry
}

func (pq *PQ) update(entry *Entry, time int) {
	entry.freq++
	entry.time = time
	heap.Fix(pq, entry.index)
}
