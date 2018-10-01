package problem125

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	ast := assert.New(t)

	questions := []struct {
		s       string
		isValid bool
	}{
		{
			"0P",
			false,
		},

		{
			".,",
			true,
		},

		{
			"A man, a plan, a canal: Panama",
			true,
		},

		{
			"race a car",
			false,
		},
	}

	for _, q := range questions {
		ast.Equal(isPalindrome(q.s), q.isValid, "输入 %v", q)
		ast.Equal(isPalindrome2(q.s), q.isValid, "输入 %v", q)
	}
}
