package main

// https://leetcode.com/problems/find-the-difference/

func findTheDifference(s string, t string) byte {
	var ss int
	for _, c := range s {
		ss ^= int(c)
	}

	var tt int
	for _, c := range t {
		tt ^= int(c)
	}

	return byte(ss ^ tt)
}
