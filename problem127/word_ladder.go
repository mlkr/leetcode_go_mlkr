package problem127

import (
	"sync"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]bool, len(wordList))
	for _, i := range wordList {
		dict[i] = true
	}
	delete(dict, beginWord)

	nextWords := make([]string, 0, len(wordList))

	trans := func(word string) bool {
		b := []byte(word)
		for i := 0; i < len(b); i++ {
			diff := b[i]

			for j := byte(97); j <= 122; j++ {
				if j == diff {
					continue
				}
				b[i] = j

				if dict[string(b)] {
					if string(b) == endWord {
						return true
					}

					nextWords = append(nextWords, string(b))
					delete(dict, string(b))
				}
			}

			b[i] = diff
		}

		return false
	}

	nextWords = append(nextWords, beginWord)
	stepCount := 1
	for len(nextWords) != 0 {
		size := len(nextWords)
		for i := 0; i < size; i++ {
			word := nextWords[0]
			nextWords = nextWords[1:]
			if trans(word) {
				return stepCount + 1
			}
		}

		stepCount++
	}

	return 0
}

// 并发版
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	dict := &sync.Map{}
	for _, i := range wordList {
		dict.Store(i, true)
	}
	dict.Delete(beginWord)

	queue := make([]string, 0, len(wordList))
	nextQueue := make([]string, 0, len(wordList))
	rw := &sync.RWMutex{}
	wg := &sync.WaitGroup{}
	var hit bool

	trans := func(word string) {
		b := []byte(word)
		for i := 0; i < len(b); i++ {
			if hit {
				break
			}

			diff := b[i]

			for j := byte(97); j <= 122; j++ {
				if hit {
					break
				}

				if j == diff {
					continue
				}
				b[i] = j

				_, ok := dict.Load(string(b))
				if ok {
					rw.Lock()
					dict.Delete(string(b))
					rw.Unlock()

					if string(b) == endWord {
						rw.Lock()
						hit = true
						rw.Unlock()
						break
					}

					rw.Lock()
					nextQueue = append(nextQueue, string(b))
					rw.Unlock()
				}
			}

			b[i] = diff
		}

		wg.Done()
	}

	queue = append(queue, beginWord)
	step := 1
	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			wg.Add(1)
			go trans(queue[i])
		}

		step++
		wg.Wait()

		if hit {
			return step
		}

		queue = nextQueue
		nextQueue = nextQueue[len(nextQueue):]
	}

	return 0
}
