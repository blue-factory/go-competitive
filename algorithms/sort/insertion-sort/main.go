package main

import "fmt"

func insertion(arr []int) {
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		/* Move elements of arr[0..i-1], that are greater than key, to one position ahead of their current position */
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{6, 8, 2, 3, 1, 5}

	insertion(arr)

	fmt.Println(arr)
}
