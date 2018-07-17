package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

func readFile(fpath string) ([][]string) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Print(err)
	}

	s := string(b)
	lines := strings.Split(s, "\n")
	var table [][]string
	for i := 0; i < len(lines); i++ {
		row := strings.Fields(lines[i])
		table = append(table, row)
	}

	return table
}

func part1(t [][]string) {
	regs := make(map[string]int)
	for i := range t {
		row := t[i]
		j, k := 0, 0
		regs[row[0]] = j
		regs[row[4]] = k
	}
	for i := range t {
		row := t[i]
		update(regs, row)
	}

	largest := regs["g"]
	for k := range regs {
		v := regs[k]
		if v > largest {
			largest = v
		}
	}
	fmt.Println(largest)
}

func part2(t [][] string) {
	regs := make(map[string]int)
	for i := range t {
		row := t[i]
		regs[row[0]] = 0
		regs[row[4]] = 0
	}

	largest := 0
	for i := range t {
		row := t[i]
		update(regs, row)
		v, _ := regs[row[0]]
		if v > largest {
			largest = v
		}
	}
	fmt.Println(largest)

}

func update(regs map[string]int, row []string) {
	if !checkCondition(regs, row) {
		return
	}
	runUpdate(regs, row)
}

func runUpdate(regs map[string]int, row []string) {
	v, _ := strconv.Atoi(row[2])
	if row[1] == "dec" {
		newVal := regs[row[0]] - v
		regs[row[0]] = newVal
		return
	}
	newVal := regs[row[0]] + v
	regs[row[0]] = newVal
}

func checkCondition(regs map[string]int, row []string) (bool) {
	v1 := regs[row[4]]
	comp := row[5]
	v2, _ := strconv.Atoi(row[6])
	if comp == "==" {
		return v1 == v2

	} else if comp == "!=" {
		return v1 != v2
	} else if comp == ">=" {
		return v1 >= v2
	} else if comp == "<=" {
		return v1 <= v2
	} else if comp == ">" {
		return v1 > v2
	} else if comp == "<" {
		return v1 < v2
	}
	fmt.Println(comp)
	panic("No Comparator")
}

func main() {
	t := readFile("assets/day08.in")
	part1(t)
	part2(t)

}
