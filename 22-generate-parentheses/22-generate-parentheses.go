package generateparentheses

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}

	res := make([]string, 0)
	backtrack(&res, []rune{}, n, 0, 0)
	return res
}

func backtrack(res *[]string, stack []rune, n, pOpen, pClose int) {
	if len(stack) == n*2 {
		*res = append(*res, string(stack))
		return
	}

	if pOpen < n {
		backtrack(res, append(stack, '('), n, pOpen+1, pClose)
	}

	if pClose < pOpen {
		backtrack(res, append(stack, ')'), n, pOpen, pClose+1)
	}
}
