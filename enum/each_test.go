package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEach(t *testing.T) {
	examples := map[string]struct {
		input  interface{}
		verify func(tt *testing.T, args []interface{})
	}{
		"nil": {
			input: nil,
			verify: func(tt *testing.T, args []interface{}) {
				assert.Len(tt, args, 0, "expect no calls")
			},
		},

		"empty": {
			input: []int{},
			verify: func(tt *testing.T, args []interface{}) {
				assert.Len(tt, args, 0, "expect no calls")
			},
		},

		"ints": {
			input: []int{1, 2, 3},
			verify: func(tt *testing.T, args []interface{}) {
				assert.Equal(tt, []interface{}{1, 2, 3}, args)
			},
		},

		"strings": {
			input: []string{"a", "b", "c"},
			verify: func(tt *testing.T, args []interface{}) {
				assert.Equal(tt, []interface{}{"a", "b", "c"}, args)
			},
		},

		"mixed": {
			input: []interface{}{"a", 1},
			verify: func(tt *testing.T, args []interface{}) {
				assert.Equal(tt, []interface{}{"a", 1}, args)
			},
		},
	}

	for name, example := range examples {
		t.Run(name, func(tt *testing.T) {
			var args []interface{}
			New(example.input).Each(func(arg interface{}) {
				args = append(args, arg)
			})
			example.verify(tt, args)
		})
	}

	t.Run("matches fn type", func(tt *testing.T) {
		var args []int
		New([]int{1, 2, 3}).Each(func(arg int) {
			args = append(args, arg)
		})
		assert.Equal(tt, []int{1, 2, 3}, args)
	})
}
