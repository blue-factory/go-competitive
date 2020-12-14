package main

// https://leetcode.com/problems/find-the-difference/

func findTheDifference(s string, t string) byte {
	ss := make(map[string]string)

	for _, c := range s {
		ss[string(c)] = ""
	}

	for _, c := range t {
		_, ok := ss[string(c)]
		if !ok {
			return byte(c)
		}

		delete(ss, string(c))
	}

	return byte(0)
}
