package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

func readFile(fpath string) ([]string) {

	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Print(err)
	}

	s := string(b)
	s = strings.TrimSuffix(s, "\n")
	moves := strings.Split(s, ",")
	return moves
}

func spin(programs []string, size int) []string {
	pnew := make([]string, len(programs))
	threshold := len(programs) - size
	index := 0
	for i := threshold; i < len(programs); i++ {
		pnew[index] = programs[i]
		index++
	}
	for i := 0; i < threshold; i++ {
		pnew[index] = programs[i]
		index++
	}
	return pnew

}

func exchange(programs []string, i int, j int) []string {
	programs[i], programs[j] = programs[j], programs[i]
	return programs
}

func partner(programs []string, a string, b string) []string {
	aIndex, bIndex := 0, 0
	for i := range programs {
		if programs[i] == a {
			aIndex = i
		}
		if programs[i] == b {
			bIndex = i
		}
	}
	return exchange(programs, aIndex, bIndex)
}

func part1(programs []string, moves []string) []string {
	for i := range moves {
		move := moves[i]
		if string(move[0]) == "s" {
			length, _ := strconv.Atoi(move[1:])
			programs = spin(programs, length)
		} else if string(move[0]) == "x" {
			vars := strings.Split(move[1:], "/")
			p1, _ := strconv.Atoi(vars[0])
			p2, _ := strconv.Atoi(vars[1])
			programs = exchange(programs, p1, p2)
		} else if string(move[0]) == "p" {
			vars := strings.Split(move[1:], "/")
			programs = partner(programs, vars[0], vars[1])
		} else {
			fmt.Println("Illegal Argument")
			fmt.Println(move)
		}
	}
	//fmt.Println(programs)
	return programs
}

func programToKey(pgrams []string) string {
	s := ""
	for i := range (pgrams) {
		s += pgrams[i]
	}
	return s
}

func part2(programs []string, moves []string) {
	posMap := make(map[string]int)
	startingKey := programToKey(programs)
	key := ""
	runs := 0
	for ; key != startingKey; {
		posMap[key] = runs
		programs = part1(programs, moves)
		key = programToKey(programs)
		runs += 1
	}
	for k, v := range (posMap) {
		if v == 1000000000%runs {
			fmt.Println(k)
		}

	}
}

func main() {
	moves := readFile("assets/day16.in")
	programs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
	programs = part1(programs, moves)
	fmt.Println(programs)

	programs = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
	part2(programs, moves)
}
