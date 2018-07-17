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
	lines := strings.Split(s, "\t")
	var table []int
	for i := 0; i < len(lines); i++ {
		v, _ := strconv.Atoi(lines[i])
		table = append(table, v)
	}
	return table
}

func redistribute(l []int) {
	index := getLargest(l)
	v := l[index]
	l[index] = 0

	for ; v > 0; v-- {
		index++
		if index >= len(l) {
			index = 0
		}
		l[index]++
	}
}

func getLargest(l []int) (int) {
	ind := 0
	for i := 1; i < len(l); i++ {
		if l[i] > l[ind] {
			ind = i
		}
	}
	return ind

}

func makeKey(l []int) (string) {
	s := ""
	for i := range l {
		s += strconv.Itoa(l[i])
		s += ","
	}
	return s
}

func part1(l []int) {
	d := make(map[string]bool)
	steps := 0
	key := makeKey(l)
	_, ok := d[key]
	for ; !ok; {
		d[key] = true
		redistribute(l)
		key = makeKey(l)
		_, ok = d[key]
		steps += 1
	}
	fmt.Println(steps)
}

func part2(l []int) {
	d := make(map[string]int)
	steps := 0
	key := makeKey(l)
	_, ok := d[key]
	for ; !ok; {
		d[key] = steps
		redistribute(l)
		key = makeKey(l)
		_, ok = d[key]
		steps += 1
	}
	fmt.Println(steps - d[key])
}

func main() {
	l := readFile("assets/day06.in")
	part1(l)

	l = readFile("assets/day06.in")
	part2(l)
}
