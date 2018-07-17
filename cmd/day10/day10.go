package main

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/lilleswing/advent-go/pkg/knothash"
)

func part1(lengths []int, size int) {
	l := knot(lengths, size, 1)
	fmt.Println(l[0]*l[1], l)
}

func knot(lengths []int, size int, rounds int) []int {
	l := make([]int, size)
	for i := range l {
		l[i] = i
	}
	zeroPos, skipSize := 0, 0
	for round := 0; round < rounds; round++ {
		for i := range lengths {
			revSize := lengths[i]
			l = reverse(l, revSize)

			shiftSize := (skipSize + revSize) % size
			l = shift(l, shiftSize)
			zeroPos = (zeroPos + size - shiftSize) % size
			skipSize += 1
		}
	}
	l = shift(l, zeroPos)
	return l
}

func shift(l []int, size int) []int {
	l2 := make([]int, len(l))
	for i := size; i < len(l2); i++ {
		l2[i-size] = l[i]
	}
	for i := 0; i < size; i++ {
		l2[len(l2)-size+i] = l[i]
	}
	return l2
}

func reverse(l []int, size int) []int {
	l2 := make([]int, len(l))
	for i := 0; i < size; i++ {
		l2[size-1-i] = l[i]
	}
	for i := size; i < len(l2); i++ {
		l2[i] = l[i]
	}
	return l2

}

func parseLengths(s string) []int {
	fields := strings.Split(s, ",")
	l := make([]int, len(fields))
	for i := range l {
		v, _ := strconv.Atoi(fields[i])
		l[i] = v
	}
	return l
}

func main() {
	slengths := "227,169,3,166,246,201,0,47,1,255,2,254,96,3,97,144"
	lengths := parseLengths(slengths)
	part1(lengths, 256)
	fmt.Println(knothash.KnotHash(slengths))
}
