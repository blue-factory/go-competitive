package main

// https://leetcode.com/problems/find-the-difference/

func findTheDifference(s string, t string) byte {
	var ss rune
	for _, c := range s {
		ss += c
	}

	var tt rune
	for _, c := range t {
		tt += c
	}

	return byte(tt - ss)
}
