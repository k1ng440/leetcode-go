package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleTest(t *testing.T) {

	assert.Equal(t, search([]int{1, 2, 3, 4, 5, 6, 7, 8}, 5), 4)
}
