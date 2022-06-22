package floodfill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloodfill(t *testing.T) {
	values := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	expected := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}

	assert.Equal(t, expected, floodFill(values, 1, 1, 2))
}

func TestFloodfill2(t *testing.T) {
	values := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}

	expected := [][]int{
		{2, 2, 2},
		{2, 2, 2},
	}

	assert.Equal(t, expected, floodFill(values, 1, 0, 2))
}
