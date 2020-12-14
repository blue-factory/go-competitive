package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Val int
}

func (p Node) String() string {
	return fmt.Sprintf("%d", p.Val)
}

type ByVal []Node

func (a ByVal) Len() int           { return len(a) }
func (a ByVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVal) Less(i, j int) bool { return a[i].Val < a[j].Val }

func main() {
	nodes := []Node{31, 42, 17, 26}

	fmt.Println(nodes)

	sort.Sort(ByVal(nodes))
	fmt.Println(nodes)

}
