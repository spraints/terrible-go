package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	assert.Equal(t, 0,
		New(nil).Reduce(0, func(int, int) int { return 1 }))
	assert.Equal(t, 0,
		New([]int{}).Reduce(0, func(int, int) int { return 1 }))
	assert.Equal(t, 1,
		New([]int{99}).Reduce(0, func(int, int) int { return 1 }))

	assert.Equal(t, "{}abcdef",
		New([]string{"a", "b", "c", "d", "e", "f"}).Reduce("{}", func(memo, obj string) string {
			return memo + obj
		}))
	assert.Equal(t, 6,
		New([]string{"abcd", "ef"}).Reduce(0, func(memo int, obj string) int {
			return memo + len(obj)
		}))
}
