package stringtointeger

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	validInts := map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
	ret := 0
	signMultiplier := 1

	if s[0] == '-' {
		signMultiplier = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for i, x := range s {
		if x < '0' || x > '9' {
			return ret * signMultiplier
		}

		if v, ok := validInts[byte(x)]; ok {
			if i == 0 {
				ret = v
				continue
			}

			ret = (ret * 10) + v

			if ret*signMultiplier <= math.MinInt32 {
				return math.MinInt32
			} else if ret*signMultiplier >= math.MaxInt32 {
				return math.MaxInt32
			}
		}
	}

	return ret * signMultiplier
}
