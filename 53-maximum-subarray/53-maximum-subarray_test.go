package maximumsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxmimumSubArray(t *testing.T) {
	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, -1, maxSubArray([]int{-2, -1}))
}
