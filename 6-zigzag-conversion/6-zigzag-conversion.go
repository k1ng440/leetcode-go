package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	runes := []rune(s)
	res := make([]rune, 0, len(runes))
	increment := 2 * (numRows - 1)

	for row := 0; row < numRows; row++ {
		for c := row; c < len(runes); c += increment {
			res = append(res, runes[c])
			if row > 0 && row < numRows-1 && c+increment-2*row < len(runes) {
				res = append(res, runes[c+increment-2*row])
			}
		}
	}

	return string(res)
}

//PAHN
//PYAIH
//RNPA
//HN
//
//PAHN
//APLSI
//IGYIR
