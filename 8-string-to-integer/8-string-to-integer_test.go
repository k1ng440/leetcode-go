package stringtointeger

import (
	"testing"

	"gotest.tools/assert"
)

func Test_myAtoi(t *testing.T) {
	// assert.Equal(t, -123456, myAtoi("-123456"))
	// assert.Equal(t, 123456, myAtoi("123456 asdfg"))
	// assert.Equal(t, 0, myAtoi("words and 987"))
	// assert.Equal(t, 3, myAtoi("3.4535345"))
	// assert.Equal(t, 2147483647, myAtoi("2147483646"))
	// assert.Equal(t, -2147483648, myAtoi("-91283472332"))
	assert.Equal(t, 2147483647, myAtoi("21474836460"))

}
