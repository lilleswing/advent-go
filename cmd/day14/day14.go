package main

import (
	"fmt"
	"strings"
	"github.com/lilleswing/advent-go/pkg/knothash"
	"encoding/hex"
)

type Point struct {
	r int
	c int
}

func part1(board [][]string) {
	total := 0
	for i := range board {
		for j := range board[i] {
			c := board[i][j]
			if c == "#" {
				total += 1
			}
		}
	}
	fmt.Println(total)
}

func toBinaryString(b byte) string {
	codes := []byte{128, 64, 32, 16, 8, 4, 2, 1}
	s := ""
	for i := range codes {
		if (b & codes[i]) != 0 {
			s += "#"
		} else {
			s += "."
		}
	}
	return s
}

func makeBoard(key string) [][]string {
	board := make([][]string, 128)
	for i := 0; i < 128; i++ {
		newKey := fmt.Sprintf("%s-%d", key, i)
		hash := knothash.KnotHash(newKey)
		hexHash, _ := hex.DecodeString(hash)
		row := ""
		for j := range hexHash {
			row += toBinaryString(hexHash[j])
		}
		board[i] = strings.Split(row, "")
	}
	return board
}

func pointToString(point Point) string {
	return fmt.Sprintf("%d,%d", point.r, point.c)
}

func dfs(board [][]string, used map[string]int, point Point, regionNum int) {
	if point.r < 0 || point.r >= len(board) {
		return
	}
	if point.c < 0 || point.c >= len(board[0]) {
		return
	}
	if board[point.r][point.c] == "." {
		return
	}
	key := pointToString(point)
	_, exists := used[key]
	if exists {
		return
	}
	used[key] = regionNum
	for dr := -1; dr < 2; dr += 2 {
		p2 := Point{point.r + dr, point.c}
		dfs(board, used, p2, regionNum)

		p2 = Point{point.r, point.c + dr}
		dfs(board, used, p2, regionNum)
	}

}

func part2(board [][]string) {
	used := make(map[string]int)
	regionNum := 0
	for r := range board {
		for c := range board[0] {
			if board[r][c] == "." {
				continue
			}
			p := Point{r, c}
			key := pointToString(p)
			_, exists := used[key]
			if exists {
				continue
			}
			dfs(board, used, p, regionNum)
			regionNum++
		}
	}
	fmt.Println(regionNum)

}

func main() {
	board := makeBoard("uugsqrei")
	part1(board)
	part2(board)
}
