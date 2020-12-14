package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%v", &n)

	origin := "A"
	destiny := "B"
	aux := "C"

	hanoi(n, origin, destiny, aux)
}

func hanoi(n int, origin, destiny, aux string) {
	if n > 0 {
		hanoi(n-1, origin, aux, destiny)

		fmt.Println(fmt.Sprintf("Moving ring %d from %s to %s", n, origin, destiny))

		hanoi(n-1, aux, destiny, origin)
	}
}
