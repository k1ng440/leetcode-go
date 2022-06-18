package reverseinteger

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	assert.Equal(t, 0, reverse(int(math.Pow(2, 31)-1)))
	// assert.Equal(t, 0, reverse(1534236469))
}
