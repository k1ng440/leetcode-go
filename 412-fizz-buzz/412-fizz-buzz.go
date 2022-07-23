package fizzbuzz

import (
	"fmt"
)

func fizzBuzz(n int) []string {
	res := make([]string, n+1)

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			res[i] = "FizzBuzz"
		case i%3 == 0:
			res[i] = "Fizz"
		case i%5 == 0:
			res[i] = "Buzz"
		default:
			res[i] = fmt.Sprintf("%d", i)
		}
	}

	return res[1:]
}
