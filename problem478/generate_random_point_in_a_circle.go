package problem478

import "math/rand"

type Solution struct {
	r, x, y float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		r: radius,
		x: x_center,
		y: y_center,
	}
}

func (this *Solution) RandPoint() []float64 {
	xFactory, yFactory := 1., 1.

	// 超过此范围 再次采样
	for xFactory*xFactory+yFactory*yFactory > 1 {
		// xFactory yFactory 范围为 [-1, 2)
		xFactory = 3*rand.Float64() - 1
		yFactory = 3*rand.Float64() - 1
	}

	x := this.x + this.r*xFactory
	y := this.y + this.r*yFactory

	return []float64{x, y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
