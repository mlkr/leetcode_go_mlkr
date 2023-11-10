package problem535

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var urls = []string{"https://leetcode.com/problems/design-tinyurl"}

func Test_Codec(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor()
	for _, longUrl := range urls {
		shortUrl := obj.encode(longUrl)
		ast.Equal(longUrl, obj.decode(shortUrl))
	}
}

func Benchmark_Codec(b *testing.B) {
	obj := Constructor()
	for i := 0; i < b.N; i++ {
		for _, longUrl := range urls {
			shortUrl := obj.encode(longUrl)
			obj.decode(shortUrl)
		}
	}
}
