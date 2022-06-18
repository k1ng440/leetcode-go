package integertoroman

import (
	"strings"
)

var (
	val = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

func intToRoman(num int) string {
	res := make([]string, 0)
	for i := 0; num != 0; i++ {
		for num >= val[i] {
			if num == val[i] {
				res = append(res, sym[i])
			}
		}
	}

	return strings.Join(res, "")
}
