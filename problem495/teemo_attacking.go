package problem495

func findPoisonedDuration(timeSeries []int, duration int) int {
	res := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i]+duration <= timeSeries[i+1] {
			res += duration
		} else {
			res += timeSeries[i+1] - timeSeries[i]
		}
	}

	res += duration

	return res
}
