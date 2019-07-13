package enum

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert.Empty(t, New(nil).Map(nil).Get(), "nil maps to empty")
	assert.Empty(t, New([]int{}).Map(nil).Get(), "empty maps to empty")

	assert.Equal(t, []int{2, 4, 6},
		New([]int{1, 2, 3}).Map(func(n int) int { return n * 2 }).Get(),
		"map works")
	assert.Equal(t, []int{3, 5, 7},
		New([]int{1, 2, 3}).Map(func(n int) int { return n * 2 }).Map(func(n int) int { return n + 1 }).Get(),
		"chained map works")
	assert.Equal(t, []string{"a", "aa", "aaa"},
		New([]int{1, 2, 3}).Map(func(n int) string { return strings.Repeat("a", n) }).Get(),
		"chained map works")
}
