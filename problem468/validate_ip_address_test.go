package problem468

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	IP  string
	ans string
}{
	{
		"172.16.254.1",
		"IPv4",
	},

	{
		"172.016.254.1",
		"Neither",
	},

	{
		"2001:0db8:85a3:0:0:8A2E:0370:7334",
		"IPv6",
	},

	{
		"02001:0db8:85a3:0:0:8A2E:0370:7334",
		"Neither",
	},

	{
		"2001:0dg8:85a3:0:0:8A2E:0370:7334",
		"Neither",
	},

	{
		"256.256.256.256",
		"Neither",
	},

	{
		"2001:0db8:85a3:0:0:8A2E:0370:7334:",
		"Neither",
	},

	{
		"1e1.4.5.6",
		"Neither",
	},
}

func Test_validIPAddress(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, validIPAddress(q.IP), "输入 %v", q.IP)
	}
}

func Benchmark_validIPAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			validIPAddress(q.IP)
		}
	}
}
