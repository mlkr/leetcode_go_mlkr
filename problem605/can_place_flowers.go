package problem605

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0

	size := len(flowerbed)
	if size == 1 {
		if flowerbed[0] == 0 {
			count++
		}
		return n <= count
	}

	for i := 0; i < size; i++ {
		if flowerbed[i] == 1 {
			continue
		}

		if flowerbed[i] == 0 {
			if ((i-1 >= 0 && flowerbed[i-1] == 0) || i-1 == -1) &&
				((i+1 < size && flowerbed[i+1] == 0) || i+1 == size) {
				count++
				i++
			}

		}
	}

	return count >= n
}
