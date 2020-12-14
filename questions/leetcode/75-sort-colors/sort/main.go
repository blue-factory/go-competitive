package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type numbers []int

func (n numbers) Len() int           { return len(n) }
func (n numbers) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n numbers) Less(i, j int) bool { return n[i] < n[j] }

func sortColors(nums []int) {
	sort.Sort(numbers(nums))
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
