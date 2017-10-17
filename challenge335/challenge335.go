package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInts(line []string) []int {
	var x []int

	for i := 0; i < len(line); i++ {
		num, err := strconv.Atoi(line[i])
		check(err)
		x = append(x, num)
	}

	return x
}

func consecutive(start int, end int) bool {
	if end == (start+1) || end == (start-1) {
		return true
	}
	return false
}

func solveLine(seq []int) int {
	total := 0

	for i := 0; i < len(seq); i++ {
		count := 0
		for j := i + 1; j < len(seq); j++ {
			count++
			if consecutive(seq[i], seq[j]) {
				total = total + count
				break
			}
		}
	}
	return total
}

func main() {
	file, err := os.Open("challenge335_input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sequences [][]int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		sequences = append(sequences, getInts(line))
	}
	check(scanner.Err())

	for _, val := range sequences {
		fmt.Println(solveLine(val))
	}

}
