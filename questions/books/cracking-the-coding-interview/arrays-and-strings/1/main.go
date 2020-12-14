package main

func isUnique(word string) bool {
	visited := map[string]string{}

	for _, c := range word {
		if _, exists := visited[string(c)]; exists {
			return false
		}

		visited[string(c)] = string(c)
	}

	return true
}

func isUnique2(word string) bool {
	visited := []bool{}

	for _, c := range word {
		if visited[c] {
			return false
		}

		visited[c] = true
	}

	return true
}

func isUnique3(word string) bool {
	for i := 0; i < len(word); i++ {
		for j := 0; j < len(word); j++ {
			if word[i] == word[j] && i != j {
				return false
			}
		}
	}

	return true
}
