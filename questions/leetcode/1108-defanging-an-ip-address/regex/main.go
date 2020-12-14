package main

import (
	"bytes"
	"regexp"
)

const (
	regex = "[0-9]+.[0-9]+.[0-9]+.[0-9]+"
)

func defangIPaddr(address string) string {
	re := regexp.MustCompile(regex)
	b := re.Find([]byte(address))

	bb := bytes.ReplaceAll(b, []byte("."), []byte("[.]"))

	return string(bb)
}
