package enum

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChains(t *testing.T) {
	ints := New([]int{1, 2, 3, 4, 5, 6})

	strs := ints.Map(func(n int) string { return fmt.Sprintf("%d", n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6"}, strs.Get(), "strs")

	doubles := ints.Map(func(n int) int { return 2 * n })
	assert.Equal(t, []int{2, 4, 6, 8, 10, 12}, doubles.Get(), "doubles")

	plusones := doubles.Map(func(n int) int { return n + 1 })
	assert.Equal(t, []int{3, 5, 7, 9, 11, 13}, plusones.Get(), "plusones")

	filter := func(n int) bool { return n > 8 }
	assert.Equal(t, []int{9, 11, 13}, plusones.Select(filter).Get(), "plusones > 8")
	assert.Equal(t, []int{3, 5, 7}, plusones.Reject(filter).Get(), "plusones not > 8")
	assert.Equal(t, []int{}, plusones.Select(filter).Reject(filter).Get(), "plusones not > 8 and also > 8")

	assert.True(t, plusones.All(func(n int) bool { return n > 2 }), "all > 2")
	assert.False(t, plusones.Any(func(n int) bool { return n < 2 }), "any < 2")
	assert.Equal(t, 21, strs.Map(func(s string) int { n, _ := strconv.Atoi(s); return n }).Reduce(0, func(a, b int) int { return a + b }), "strs -> ints -> sum")
}
