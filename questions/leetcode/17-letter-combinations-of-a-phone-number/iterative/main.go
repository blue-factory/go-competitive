package main

var (
	combs = map[rune][]rune{
		'2': []rune{'a', 'b', 'c'},
		'3': []rune{'d', 'e', 'f'},
		'4': []rune{'g', 'h', 'i'},
		'5': []rune{'j', 'k', 'l'},
		'6': []rune{'m', 'n', 'o'},
		'7': []rune{'p', 'q', 'r', 's'},
		'8': []rune{'t', 'u', 'v'},
		'9': []rune{'w', 'x', 'y', 'z'},
	}
)

func letterCombinations(digits string) []string {
	results := []string{}

	if len(digits) == 0 {
		return results
	}

	type candidateNode struct {
		candidate string
		digits    string
	}

	stack := []candidateNode{candidateNode{"", digits}}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(top.digits) == 0 {
			results = append(results, top.candidate)
			continue
		}

		key := rune(top.digits[0])
		letters := combs[key]

		for _, l := range letters {
			c := top.candidate + string(l)
			stack = append(stack, candidateNode{c, top.digits[1:]})
		}
	}

	return results
}
