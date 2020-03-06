package problem352

type SummaryRanges struct {
	sr [][]int
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (this *SummaryRanges) AddNum(val int) {
	if this.sr == nil {
		this.sr = [][]int{
			[]int{val, val},
		}

		return
	}

	// 二分插入
	lo, hi := 0, len(this.sr)-1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if this.sr[mid][0] <= val && this.sr[mid][1] >= val {
			return
		}

		if this.sr[mid][1] < val {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	if lo == 0 {
		if this.sr[0][0] == val+1 {
			this.sr[0][0] = val
			return
		}

		this.sr = append(this.sr, []int{})
		copy(this.sr[1:], this.sr)
		this.sr[0] = []int{val, val}

		return
	}

	size := len(this.sr)
	if lo == size {
		if this.sr[size-1][1] == val-1 {
			this.sr[size-1][1] = val
			return
		}

		this.sr = append(this.sr, []int{val, val})

		return
	}

	if this.sr[lo-1][1] == val-1 && this.sr[lo][0] == val+1 {
		this.sr[lo-1][1] = this.sr[lo][1]
		copy(this.sr[lo:], this.sr[lo+1:])
		this.sr = this.sr[:size-1]
		return
	}

	if this.sr[lo-1][1] < val-1 && this.sr[lo][0] > val+1 {
		this.sr = append(this.sr, []int{})
		copy(this.sr[lo+1:], this.sr[lo:])
		this.sr[lo] = []int{val, val}
		return
	}

	if this.sr[lo-1][1] == val-1 {
		this.sr[lo-1][1] = val
	} else {
		this.sr[lo][0] = val
	}

}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.sr
}
