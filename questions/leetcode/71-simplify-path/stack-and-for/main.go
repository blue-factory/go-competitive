package main

import "strings"

// https://leetcode.com/problems/simplify-path

func simplifyPath(path string) string {
	singles := strings.Split(path[1:], "/")
	stack := []string{}

	for _, s := range singles {
		if len(s) == 0 || s == "." {
			continue
		}

		if s == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}

		stack = append(stack, s)
	}

	return "/" + strings.Join(stack, "/")
}
