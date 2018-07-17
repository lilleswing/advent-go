package main

import (
	"io/ioutil"
	"fmt"
)

func readFile(fpath string) (string) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Print(err)
	}

	s := string(b)
	return s
}

func part1(s string) {
	score, gCount := 0, 0
	isGarbage, isCancelled := false, false
	for i := range s {
		c := string(s[i])
		if isCancelled {
			isCancelled = false
			continue
		}
		if c == "!" {
			isCancelled = true
			continue
		}
		if isGarbage && c == ">" {
			isGarbage = false
			continue
		}
		if isGarbage {
			continue
		}
		if c == "<" {
			isGarbage = true
			continue
		}
		if c == "{" {
			gCount += 1
			continue
		}
		if c == "}" {
			score += gCount
			gCount -= 1
			continue
		}
	}
	fmt.Println(score)
}

func part2(s string) {
	score, gCount, nGarbage := 0, 0, 0
	isGarbage, isCancelled := false, false
	for i := range s {
		c := string(s[i])
		if isCancelled {
			isCancelled = false
			continue
		}
		if c == "!" {
			isCancelled = true
			continue
		}
		if isGarbage && c == ">" {
			isGarbage = false
			continue
		}
		if isGarbage {
			nGarbage += 1
			continue
		}
		if c == "<" {
			isGarbage = true
			continue
		}
		if c == "{" {
			gCount += 1
			continue
		}
		if c == "}" {
			score += gCount
			gCount -= 1
			continue
		}
		panic("State Machine Broken")
	}
	fmt.Println(nGarbage)
}

func main() {
	s := readFile("assets/day09.in")
	part1(s)
	part2(s)
}
