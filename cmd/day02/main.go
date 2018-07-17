package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func to_int(l []string) ([]int) {
	var ns []int
	for i := 0; i < len(l); i++ {
		v, _ := strconv.Atoi(l[i])
		ns = append(ns, v)
	}
	return ns
}

func min(l []int) (int) {
	v := l[0]
	for i := 1; i < len(l); i++ {
		if l[i] < v {
			v = l[i]
		}
	}
	return v
}

func max(l []int) (int) {
	v := l[0]
	for i := 1; i < len(l); i++ {
		if l[i] > v {
			v = l[i]
		}
	}
	return v
}

func readFile(fpath string) ([][]int) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Print(err)
	}

	s := string(b)
	lines := strings.Split(s, "\n")
	var table [][]int
	for i := 0; i < len(lines); i++ {
		row := to_int(strings.Fields(lines[i]))
		table = append(table, row)
	}
	return table
}

func part2RowCheckSum(l []int) (int) {
	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			v1, v2 := l[i], l[j]
			if v2 > v1 {
				v1, v2 = v2, v1
			}
			if v1%v2 == 0 {
				return v1 / v2
			}
		}
	}
	return -1
}

func part1(table [][]int) {
	checksum := 0
	for i := 0; i < len(table); i++ {
		row := table[i]
		checksum += max(row) - min(row)
	}
	fmt.Println(checksum)
}

func part2(table [][]int) {
	checksum := 0
	for i := 0; i < len(table); i++ {
		row := table[i]
		checksum += part2RowCheckSum(row)
	}
	fmt.Println(checksum)
}

func main() {
	data := readFile("assets/day02.in")
	part1(data)
	part2(data)
}
