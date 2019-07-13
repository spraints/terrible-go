package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	assert.True(t, New(nil).All(func(interface{}) bool { return false }),
		"nil enum is All true")
	assert.True(t, New([]string{}).All(func(interface{}) bool { return false }),
		"empty enum is All true")

	assert.True(t,
		New([]int{2, 4, 6}).All(func(n int) bool { return n%2 == 0 }))
	assert.False(t,
		New([]int{2, 4, 6}).All(func(n int) bool { return n < 6 }))
	assert.False(t,
		New([]int{2, 4, 6}).All(func(n int) bool { return n > 2 }))
}
