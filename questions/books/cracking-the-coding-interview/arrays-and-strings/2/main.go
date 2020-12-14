package main

func checkPermutation(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	visited := map[string]int{}
	for _, char := range word1 {
		if _, exists := visited[string(char)]; !exists {
			visited[string(char)] = 0
		}

		visited[string(char)]++
	}

	for _, char := range word2 {
		if _, exists := visited[string(char)]; !exists {
			return false
		}

		visited[string(char)]--
	}

	for _, v := range visited {
		if v != 0 {
			return false
		}
	}

	return true
}
