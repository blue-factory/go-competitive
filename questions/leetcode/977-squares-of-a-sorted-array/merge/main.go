package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func merge(A []int, B []int) []int {
	left, right := 0, 0
	result := make([]int, 0)

	for left < len(A) && right < len(B) {
		if A[left] < B[right] {
			result = append(result, A[left])
			left++
			continue
		}

		result = append(result, B[right])
		right++
	}

	result = append(result, A[left:]...)
	result = append(result, B[right:]...)

	return result
}

func mergeSort(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	middle := len(A) / 2
	left := mergeSort(A[:middle])
	right := mergeSort(A[middle:])

	return merge(left, right)
}

func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}

	return mergeSort(A)
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
