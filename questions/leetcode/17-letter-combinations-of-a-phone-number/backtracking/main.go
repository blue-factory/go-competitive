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

	letterCombinationImpl(&results, "", digits)

	return results
}

func letterCombinationImpl(results *[]string, candidate string, digits string) {
	if len(digits) == 0 {
		*results = append(*results, candidate)
		return
	}

	key := rune(digits[0])
	letters := combs[key]
	for _, l := range letters {
		c := candidate + string(l)
		letterCombinationImpl(results, c, digits[1:])
	}
}
