package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if A[i] < A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}

	return A
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
