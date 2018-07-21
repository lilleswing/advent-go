package main

import (
	"fmt"
)

func insert(a []int, i int, v int) []int {
	a = append(a[:i], append([]int{v}, a[i:]...)...)
	return a
}
func part1(stepSize int) {
	l := []int{0}
	index := 0
	for i := 1; i <= 2017; i++ {
		index = (index + stepSize) % len(l)
		l = insert(l, index, i)
		index++
	}
	fmt.Println(l[index])
}

func part2(stepSize int) {
	index, answer, length  := 0, 1, 1
	for i := 1; i <= 50000000; i++ {
		index = (index + stepSize) % length
		length += 1
		if index == 0 {
			answer = i
		}
		index++
	}
	fmt.Println(answer)
}

func main() {
	stepSize := 371
	part1(stepSize)
	part2(stepSize)
}
