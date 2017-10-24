package main

import (
	"fmt"
	"sort"
)

var values = []int{21, 9, 5, 8, 10, 1, 3}
var queries = []int{10, 15}

func eat(input []int, count int, query int) int {
	if len(input) > 1 {
		if input[len(input)-1] >= query {
			count++
			return eat(input[:len(input)-1], count, query)
		}
		input[len(input)-1]++
		return eat(input[1:], count, query)
	}
	if input[0] >= query {
		count++
	}
	return count
}

func main() {
	sort.Ints(values)
	count := 0
	for _, query := range queries {
		count := eat(values, count, query)
		fmt.Println("ans =", count)
	}
}
