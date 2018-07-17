package main

import (
	"fmt"
	"strings"
	"strconv"
	"pkg/knothash/knothash"
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

func parsePart2Lengths(s string) [] int {
	l := make([]int, len(s), len(s)+5)
	for i := range s {
		l[i] = int(s[i])
	}
	l = append(l, 17, 31, 73, 47, 23)
	return l
}

func part2(lengths []int, size int) {
	spHash := knot(lengths, size, 64)
	dHash := denseHash(spHash)
	s := ""
	for i := range dHash {
		s += toHex(dHash[i])
	}
	fmt.Println(s)
}
func toHex(i int) string {
	s := fmt.Sprintf("%x", i)
	if len(s) == 1 {
		return "0" + s
	}
	return s
}

func denseHash(l []int) []int {
	retval := make([]int, 16)
	for i := 0; i < 16; i++ {
		offset := i * 16
		v := l[offset]
		for j := 1; j < 16; j++ {
			v = v ^ l[offset+j]
		}
		retval[i] = v
	}
	return retval
}

func main() {
	slengths := "227,169,3,166,246,201,0,47,1,255,2,254,96,3,97,144"
	lengths := parseLengths(slengths)
	part1(lengths, 256)
	lengths = parsePart2Lengths(slengths)
	part2(lengths, 256)
	knothash.KnotHash(slengths)
}
