package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type maze struct {name string; gameBoard []string}
type explorer struct {x, y int; facing string}

func main() {
	channel := make(chan maze)

	go getMaze("./mazes/", channel)
	for i := range channel {
		explore(i)
	}
}

func getMaze(dir string, channel chan maze) {
	files, err := ioutil.ReadDir(dir)
	check(err)

	for _, file := range files {
		name := file.Name()
		content := readFile(dir + name)
		m := maze{name, buildGameBoard(content)}
		channel <- m
	}
	close(channel)
}

func readFile(path string) string {
	content, err := ioutil.ReadFile(path)
	check(err)

	return string(content)
}

func buildGameBoard(content string) []string {
	gameBoard := strings.Split(content, "\n")
	return gameBoard
}

func explore(m maze) {
	fmt.Println(m.name)
	displayBoard(m)
	
	e := findExplorer(m)

}

func displayBoard(m maze) {
	for _, line := range m.gameBoard {
		fmt.Println(line)
	}	
}

func findExplorer(m maze ) explorer{
	for i, line := range m.gameBoard {
		for j, char := range line {
			if avatar(string(char)) {
				e := explorer{j, i, char}
			}
		}
	}	
}

func avatar(char string) bool{

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
