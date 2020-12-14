package main

// https://leetcode.com/problems/decode-string/

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	re := regexp.MustCompile(`\d+\[\w+\]`)

	for {
		matches := re.FindIndex([]byte(s))
		if len(matches) == 0 {
			break
		}

		first := matches[0]
		last := matches[1]

		m := re.FindAllString(s, -1)

		subStr := m[0]

		index := strings.LastIndex(subStr, "[")
		num, _ := strconv.Atoi(subStr[0:index])
		txt := subStr[index+1 : len(subStr)-1]

		newStr := ""
		for i := 0; i < num; i++ {
			newStr += txt
		}

		s = fmt.Sprintf("%s%s%s", s[0:first], newStr, s[last:len(s)])
	}

	return s
}
