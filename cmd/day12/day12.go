package main

import (
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
)

func readFile(fpath string) ([][]int) {
	graph := make([][]int, 2000)
	for i := 0; i < 2000; i++ {
		v := make([]int, 2000)
		graph[i] = v
	}

	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Print(err)
	}

	s := string(b)
	lines := strings.Split(s, "\n")
	for i := 0; i < len(lines); i++ {

		line := strings.Split(lines[i], " <-> ")[1]
		neighbors := strings.Split(line, ", ")
		for j := 0; j < len(neighbors); j++ {
			v, _ := strconv.Atoi(neighbors[j])
			graph[i][v] = 1
			graph[v][i] = 1
		}
	}
	return graph
}

func part1(graph [][]int) {
	seen := make(map[int]bool)
	dfs(graph, seen, 0, 800)
	fmt.Println(len(seen))

}

func part2(graph [][]int) {
	seen := make(map[int]bool)
	count := 0
	for i := 0; i < 2000; i++ {
		if _, exists := seen[i]; !exists {
			dfs(graph, seen, i, 800)
			count++
		}
	}
	fmt.Println(count)

}

func dfs(g [][]int, seen map[int]bool, start int, depth int) {
	if depth == 0 {
		return
	}
	if _, exists := seen[start]; exists {
		return
	}
	seen[start] = true

	for i := 0; i < 2000; i++ {
		if g[start][i] == 1 {
			dfs(g, seen, i, depth-1)
		}
	}
}

func main() {
	g := readFile("assets/day12.in")
	part1(g)
	part2(g)
}
