package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

type Node struct {
	name        string
	inEdge      []*Node
	outEdge     []*Node
	depth       int
	weight      int
	calcWeight  int
	fixedWeight int
}

func readFile(fpath string) (map[string]*Node) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Print(err)
	}

	m := make(map[string]*Node)

	s := string(b)
	lines := strings.Split(s, "\n")
	for i := 0; i < len(lines); i++ {
		row := strings.Fields(lines[i])
		ws := strings.Trim(strings.Trim(row[1], "("), ")")
		w, _ := strconv.Atoi(ws)
		m[row[0]] = &Node{row[0], []*Node{}, []*Node{}, -1, w, 0, 0}
	}

	for i := 0; i < len(lines); i++ {
		row := strings.Fields(lines[i])
		n, _ := m[row[0]]
		for j := 3; j < len(row); j++ {
			nodeName := strings.Trim(row[j], ",")
			dn, _ := m[nodeName]
			n.outEdge = append(n.outEdge, dn)
			dn.inEdge = append(dn.inEdge, n)
		}
	}
	return m
}

func findRoot(g map[string]*Node) (string) {
	key := ""
	for k := range g {
		key = k
		break
	}
	n := g[key]
	for ; len(n.inEdge) > 0; {
		n = g[n.inEdge[0].name]
	}
	return n.name
}

func part1(g map[string]*Node) {
	fmt.Println(findRoot(g))
}

func dfs(g map[string]*Node, root string, depth int) (int, bool) {
	n := g[root]
	n.depth = depth

	children := n.outEdge
	childWeights := make([]int, len(children))
	for i := range children {
		child := children[i]
		childWeight, done := dfs(g, child.name, depth+1)
		if done {
			return childWeight, done
		}
		childWeights[i] = childWeight
	}
	isBroken, brokenIndex := isBroken(childWeights)
	if isBroken {
		otherIndex := 0
		if brokenIndex == 0 {
			otherIndex = 1
		}
		weightDiff := children[brokenIndex].calcWeight - children[otherIndex].calcWeight
		children[brokenIndex].fixedWeight = children[brokenIndex].weight - weightDiff
		return -1, true
	}
	n.calcWeight = n.weight + sum(childWeights)
	return n.calcWeight, false
}
func isBroken(ints []int) (bool, int) {
	m := make(map[int]int)
	for i := range ints {
		m[ints[i]] = 0
	}

	for i := range ints {
		m[ints[i]]++
	}

	if len(m) <= 1 {
		return false, -1
	}

	for i := range ints {
		v, _ := m[ints[i]]
		if v == 1 {
			return true, i
		}
	}
	fmt.Println(ints, m)
	panic("IsBroken Error")
}

func sum(ints []int) (int) {
	total := 0
	for i := range ints {
		total += ints[i]
	}
	return total

}

func part2(g map[string]*Node) {
	root := findRoot(g)
	dfs(g, root, 0)
	for k := range g {
		n := g[k]
		if n.fixedWeight != 0 {
			fmt.Println(n.fixedWeight)
		}
	}
}

func main() {
	g := readFile("assets/day07.in")
	part1(g)
	part2(g)
}
