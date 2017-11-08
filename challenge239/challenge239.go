package main

import (
	"fmt"
	"math/rand"
)

func threes(n int) string {

	if n == 1 {
		return "done/n"
	}

	remainder := n % 3
	switch remainder {
	case 0:
		fmt.Println(n, "0")
	case 1:
		fmt.Println(n, "-1")
		n--
	case 2:
		fmt.Println(n, "+1")
		n++
	}

	return threes(n / 3)
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(threes(rand.Intn(100)))
	}
}
