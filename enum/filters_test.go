package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	doSelect := func(enum, fn interface{}) interface{} {
		return New(enum).Select(fn).Get()
	}

	assert.Empty(t, doSelect(nil, func(int) bool { return true }))
	assert.Empty(t, doSelect([]int{}, func(int) bool { return true }))

	assert.Equal(t, []int{2, 4}, doSelect([]int{1, 2, 3, 4},
		func(x int) bool { return x%2 == 0 }))
}

func TestReject(t *testing.T) {
	doReject := func(enum, fn interface{}) interface{} {
		return New(enum).Reject(fn).Get()
	}

	assert.Empty(t, doReject(nil, func(int) bool { return true }))
	assert.Empty(t, doReject([]int{}, func(int) bool { return true }))

	assert.Equal(t, []int{1, 3}, doReject([]int{1, 2, 3, 4},
		func(x int) bool { return x%2 == 0 }))
}
