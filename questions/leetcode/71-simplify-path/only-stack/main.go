package main

import "strings"

// https://leetcode.com/problems/simplify-path

func simplifyPath(path string) string {
	stack := []string{}
	word := ""

	for i, s := range path + "/" {
		if i == 0 && string(s) == "/" {
			continue
		}

		if string(s) != "/" {
			word += string(s)
			continue
		}

		if len(word) == 0 || word == "." {
			word = ""
			continue
		}

		if word == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			word = ""
			continue
		}

		stack = append(stack, word)
		word = ""
	}

	return "/" + strings.Join(stack, "/")
}
