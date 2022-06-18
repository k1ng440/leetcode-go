package sum3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_three_sum(t *testing.T) {
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
