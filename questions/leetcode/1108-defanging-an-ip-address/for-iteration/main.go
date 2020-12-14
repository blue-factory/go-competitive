package main

// https://leetcode.com/problems/defanging-an-ip-address/

func defangIPaddr(address string) string {
	res := ""
	for _, c := range address {
		if string(c) == "." {
			res += "[.]"
			continue
		}

		res += string(c)
	}

	return res
}
