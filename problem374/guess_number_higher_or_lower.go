package problem374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	return guessRange(1, n)
}

func guessRange(lo, hi int) int {
	point := (lo + hi) / 2
	res := guess(point)

	if res == -1 {
		return guessRange(lo, point-1)
	}

	if res == 1 {
		return guessRange(point+1, hi)
	}

	return point
}
