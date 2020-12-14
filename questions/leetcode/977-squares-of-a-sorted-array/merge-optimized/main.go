package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sortedSquares(A []int) []int {
	left, right := 0, 0

	for right < len(A) && A[right] < 0 {
		right++
	}
	left = right - 1

	res := []int{}
	for left >= 0 && right < len(A) {
		if A[left]*A[left] < A[right]*A[right] {
			res = append(res, A[left]*A[left])
			left--
		} else {
			res = append(res, A[right]*A[right])
			right++
		}
	}

	for left >= 0 {
		res = append(res, A[left]*A[left])
		left--
	}

	for right < len(A) {
		res = append(res, A[right]*A[right])
		right++
	}

	return res
}

func main() {
	input := os.Args[1]
	nums := strings.Split(input, ",")

	ns := []int{}
	for i := 0; i < len(nums); i++ {
		n, err := strconv.Atoi(nums[i])
		if err != nil {
			log.Fatalln(err)
		}
		ns = append(ns, n)
	}

	res := sortedSquares(ns)

	log.Println(fmt.Sprintf("input: %v", nums))
	log.Println(fmt.Sprintf("output: %v", res))
}
