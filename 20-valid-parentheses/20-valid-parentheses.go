package validparentheses

type runes []rune

func (p *runes) pop() rune {
	if len(*p) == 0 {
		return 0
	}

	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func isValid(s string) bool {
	stack := make(runes, 0)
	for _, x := range s {
		switch x {
		case ')':
			if stack.pop() != '(' {
				return false
			}
		case '}':
			if stack.pop() != '{' {
				return false
			}
		case ']':
			if stack.pop() != '[' {
				return false
			}
		default:
			stack = append(stack, x)
		}
	}

	return len(stack) == 0
}
