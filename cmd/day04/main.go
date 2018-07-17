package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"sort"
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

func isUnique(l []string) (bool) {
	d := make(map[string]bool)
	for i := range l {
		if _, exists := d[l[i]]; exists {
			return false
		}
		d[l[i]] = true
	}
	return true
}

func part1(table [][]string) {
	total := 0
	for i := range table {
		row := table[i]
		if isUnique(row) {
			total += 1
		}
	}
	fmt.Println(total)
}

func sortWordlist(l []string) ([]string) {
	sorted := make([]string, len(l))
	for i := range l {
		vals := strings.Split(l[i], "")
		sort.Strings(vals)
		sorted[i] = strings.Join(vals, "")
	}
	return sorted
}

func part2(table [][]string) {
	total := 0
	for i := range table {
		row := sortWordlist(table[i])
		if isUnique(row) {
			total += 1
		}
	}
	fmt.Println(total)
}

func main() {
	table := readFile("assets/day04.in")
	part1(table)
	part2(table)
}
