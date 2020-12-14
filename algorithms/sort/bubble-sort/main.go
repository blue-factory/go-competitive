package main

import "fmt"

func bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	arr := []int{6, 8, 2, 3, 1, 5}

	bubble(arr)

	fmt.Println(arr)
}
