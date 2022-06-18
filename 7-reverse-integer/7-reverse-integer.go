package reverseinteger

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	ret := 0
	for x != 0 {
		pop := x % 10
		px := ret * 10

		fmt.Printf("x: %d \t\t\t ret: %d \t\t\t ret*10: %d \t\t\t x%%10: %d \t\t\t x/10: %d \t\t\t int32 rat: %d\n", x, ret, px, pop, x/10, int32(ret))

		if ret > math.MaxInt32/10 {
			return math.MaxInt32
		}

		if ret < math.MinInt32/10 {
			return math.MinInt32
		}

		ret = pop + px // (ret * 10) + x % 10
		x /= 10
	}

	return ret
}
