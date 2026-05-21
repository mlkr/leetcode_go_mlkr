package problem535

import (
	"math/rand"
	"strings"
)

type Codec struct {
	domain string
	urls   map[string]string
}

func Constructor() Codec {
	return Codec{"https://leetcode.com/", make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	key := ""
	ok := true
	for ok {
		key = generateKey()
		_, ok = this.urls[key]
	}

	this.urls[key] = longUrl
	return this.domain + key
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	key := shortUrl[len(this.domain):]
	return this.urls[key]
}

func generateKey() string {
	key := strings.Builder{}

	for i := 0; i < 6; i++ {
		n := byte(rand.Intn(62))

		if n < 26 {
			key.WriteByte(n + 'A')
		} else if n < 52 {
			key.WriteByte(n - 26 + 'a')
		} else {
			key.WriteByte(n - 52 + '0')
		}
	}

	return key.String()
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
