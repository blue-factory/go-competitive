package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	n := 2

	res := binarySearch(arr, 0, len(arr)-1, n)

	fmt.Println(res)
}

func binarySearch(arr []int, start, end, n int) bool {
	if start > end {
		return false
	}

	mid := (end-start)/2 + start

	if arr[mid] == n {
		return true
	}

	if arr[mid] < n {
		return binarySearch(arr, mid+1, end, n)
	}

	return binarySearch(arr, start, mid-1, n)
}
