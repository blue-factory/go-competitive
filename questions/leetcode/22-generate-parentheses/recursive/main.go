package main

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	results := []string{}
	generateParenthesisImpl(&results, "", n, n)

	return results
}

func generateParenthesisImpl(results *[]string, candidate string, open, close int) {
	if open == 0 && close == 0 {
		*results = append(*results, candidate)
		return
	}

	if open > 0 {
		generateParenthesisImpl(results, candidate+"(", open-1, close)
	}

	if close > open {
		generateParenthesisImpl(results, candidate+")", open, close-1)
	}
}

type Node struct {
	Value    int
	Children []*Node
}

func getNumberOfUniqueTrees(root *Node) int {
	count := 0
	_ = helper(root, &count)

	return count
}

func helper(root *Node, count *int) bool {
	for _, v := range root.Children {
		if !helper(v, count) || v.Value != root.Value {
			return false
		}
	}
	*count++
	return true
}
