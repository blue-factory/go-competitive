package main

func isValid(str string) bool {
	stack := []string{}

	for _, s := range str {
		pos := string(s)
		if len(stack) == 0 && (pos == "]" || pos == "}" || pos == ")") {
			return false
		}

		if pos == "{" || pos == "[" || pos == "(" {
			stack = append(stack, pos)
			continue
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pos == "]" && top == "[" {
			continue
		}

		if pos == "}" && top == "{" {
			continue
		}

		if pos == ")" && top == "(" {
			continue
		}

		return false
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
