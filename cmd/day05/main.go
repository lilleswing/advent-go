package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

func readFile(fpath string) ([]int) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Print(err)
	}

	s := string(b)
	lines := strings.Split(s, "\n")
	var table []int
	for i := 0; i < len(lines); i++ {
		val, _ := strconv.Atoi(lines[i])
		table = append(table, val)
	}
	return table
}

func part1(l []int) {
	index, steps := 0, 0
	for ; index < len(l) && index >= 0; {
		newIndex := index + l[index]
		l[index]++
		index = newIndex
		steps += 1
	}
	fmt.Println(steps)
}

func part2(l []int) {
	index, steps := 0, 0
	for ; index < len(l) && index >= 0; {
		newIndex := index + l[index]
		if l[index] >= 3 {
			l[index]--
		} else {
			l[index]++
		}
		index = newIndex
		steps += 1
	}
	fmt.Println(steps)
}

func main() {
	l := readFile("assets/day05.in")
	part1(l)

	l = readFile("assets/day05.in")
	part2(l)
}
