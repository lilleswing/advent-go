package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func pointToStr(p Point) (string) {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

func strToPoint(s string) (Point) {
	vals := strings.Split(s, ",")
	x, _ := strconv.Atoi(vals[0])
	y, _ := strconv.Atoi(vals[1])
	return Point{x, y}
}

func abs(n int) (int) {
	if n < 0 {
		return n * -1
	}
	return n
}

func part1(n int) (map[string]int) {
	board := make(map[string]int)
	loc := Point{-1, 1}
	v := 7
	directions := []Point{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}
	dirIndex := 0
	dirLength := 3
	for ; v < n; {
		direction := directions[dirIndex]
		for i := 0; i < dirLength; i++ {
			x2, y2 := loc.x+direction.x, loc.y+direction.y
			p2 := Point{x2, y2}
			loc = p2
			v += 1
			board[pointToStr(loc)] = v
		}
		dirIndex = (dirIndex + 1) % 4
		if dirIndex%2 == 0 {
			dirLength += 1
		}
	}
	return board
}

func squareValue(board map[string]int, loc Point) (int) {
	total := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			key := pointToStr(Point{loc.x + i, loc.y + j})
			if val, ok := board[key]; ok {
				total += val
			}
		}
	}
	return total
}

func part2(n int) (map[string]int) {
	board := make(map[string]int)
	loc := Point{1, 0}
	board[pointToStr(loc)] = 1
	directions := []Point{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}
	dirIndex := 0
	dirLength := 1
	v := 0
	for ; v < n; {
		direction := directions[dirIndex]
		for i := 0; i < dirLength; i++ {
			x2, y2 := loc.x+direction.x, loc.y+direction.y
			p2 := Point{x2, y2}
			loc = p2
			v = squareValue(board, loc)
			board[pointToStr(loc)] = v
			if v > n {
				return board
			}
		}
		dirIndex = (dirIndex + 1) % 4
		if dirIndex%2 == 0 {
			dirLength += 1
		}
	}
	return board
}

func displayBoard(board map[string]int) {
	s := ""
	for i := -5; i < 5; i++ {
		for j := -5; j < 5; j++ {
			key := pointToStr(Point{j, i})
			if val, ok := board[key]; ok {
				s += strconv.Itoa(val)
			} else {
				s += "0"
			}
			s += "\t"
		}
		s += "\n"
	}
	fmt.Println(s)
}

func maxValue(board map[string]int) (int) {
	max := -1
	for _, v := range board {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	n := 312051
	board := part1(n)
	for k, v := range board {
		if v == n {
			p := strToPoint(k)
			fmt.Println(p.x, p.y, abs(p.x)+abs(p.y))
		}
	}

	board = part2(n)
	retval := strconv.Itoa(maxValue(board))
	fmt.Println(retval)
}
