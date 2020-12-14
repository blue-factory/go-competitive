package main

// https://leetcode.com/problems/decode-string

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func nplicate(n int, txt string) string {
	if n == 0 || n == 1 {
		return txt
	}

	return nplicate(n-1, txt)
}

func decodeString(s string) string {
	re := regexp.MustCompile(`\d+\[\w+\]`)
	if len(s) == 0 || !re.MatchString(s) {
		return s
	}

	matches := re.FindIndex([]byte(s))
	firstIndex := matches[0]
	lastIndex := matches[1]

	subStr := re.FindAllString(s, -1)[0]

	index := strings.LastIndex(subStr, "[")
	num, _ := strconv.Atoi(subStr[0:index])
	txt := subStr[index+1 : len(subStr)-1]

	newStr := nplicate(num, txt)

	return decodeString(fmt.Sprintf("%s%s%s", s[0:firstIndex], newStr, s[lastIndex:len(s)]))
}
