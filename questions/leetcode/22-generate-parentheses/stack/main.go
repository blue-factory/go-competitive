package main

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	type withParams struct {
		candidate string
		open      int
		close     int
	}

	stack := []withParams{withParams{"", n, n}}
	results := []string{}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.open == 0 && top.close == 0 {
			results = append(results, top.candidate)
			continue
		}

		if top.open > 0 {
			stack = append(stack, withParams{top.candidate + "(", top.open - 1, top.close})
		}

		if top.close > top.open {
			stack = append(stack, withParams{top.candidate + ")", top.open, top.close - 1})
		}
	}

	return results
}
