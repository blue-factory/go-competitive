package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	if len(nums) == 2 {
		if nums[0] < nums[1] {
			return
		}

		nums[0], nums[1] = nums[1], nums[0]
		return
	}

	colors := make([]int, 3)

	for i := 0; i < len(nums); i++ {
		colors[nums[i]] = colors[nums[i]] + 1
	}

	indexColor := 0
	index := 0

	for index < len(nums) {
		if colors[indexColor] == 0 {
			indexColor++
			continue
		}

		nums[index] = indexColor
		colors[indexColor] = colors[indexColor] - 1
		index++
	}
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

	sortColors(ns)

	log.Println(fmt.Sprintf("input: %v", nums))
	log.Println(fmt.Sprintf("output: %v", ns))
}
