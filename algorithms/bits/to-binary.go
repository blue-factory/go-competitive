package bits

import "fmt"

func toBinary(n int) {
	if n < 2 {
		fmt.Print(n)
		return
	}

	toBinary(n / 2)
	fmt.Print(n % 2)
}
