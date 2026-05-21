package problem519

import (
	"math/rand"
)

type Solution struct {
	rec  map[int]int
	m, n int
	sum  int
}

func Constructor(m int, n int) Solution {
	return Solution{
		rec: make(map[int]int),
		m:   m,
		n:   n,
		sum: m * n,
	}
}

// 0 1 2
// 3 4 5
// index = 1, x = index / cols, y = index & cols

// 使用一个变量 size，初始化为矩形的长乘以宽，然后还是只生成一个随机数id，
// 并使用另一个变量 val 来记录它。接下来给 size 自减1，
// 由于 rand() % size 得到的随机数的范围是 [0, size-1]，
// 那么假如第一次随机出了 size-1 后，此时 size 自减1之后，
// 下一次不必担心还会随机出 size-1，因为此时的 size 比之前减少了1。
// 如果第一次随机出了0，假设最开始 size=4，那么此时自减1之后，size=3，
// 此时将0映射到3。那么下次如果再次随机出了0，此时 size 自减1之后，size=2，
// 现在0有映射值，所以将 id 改为其映射值3，然后再将0映射到2，
// 这样下次就算再摇出了0，还可以改变id值。大家有没有发现，
// 映射值都是没有没使用过的数字，这也是为啥开始先检测 id 是否被使用了，
// 若已经被使用了，则换成其映射值，然后再更新之前的 id 的映射值，
// 找到下一个未被使用的值即可
func (this *Solution) Flip() []int {
	if this.sum == 0 {
		return nil
	}

	index := rand.Intn(this.sum)
	redire := index
	if val, ok := this.rec[index]; ok {
		redire = val
	}

	x := redire / this.n
	y := redire % this.n
	this.sum--

	this.rec[index] = this.sum
	if val, ok := this.rec[this.sum]; ok {
		this.rec[index] = val
	}

	return []int{x, y}

}

func (this *Solution) Reset() {
	this.sum = this.m * this.n
	this.rec = make(map[int]int)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
