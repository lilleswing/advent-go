package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

type Layer struct {
	myDepth   int
	myRange   int
	curLoc    int
	direction int
}

func readFile(fpath string) ([]Layer) {

	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Print(err)
	}

	s := string(b)
	lines := strings.Split(s, "\n")
	firewall := make([]Layer, len(lines))
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], ": ")
		depth, _ := strconv.Atoi(line[0])
		myRange, _ := strconv.Atoi(line[1])
		layer := Layer{myDepth: depth, myRange: myRange, curLoc: 0, direction: -1}
		firewall[i] = layer
	}
	return firewall
}

// Mutates layers --> proceed with care
func part1(layers []Layer) {
	myLoc, endDepth := -1, layers[len(layers)-1].myDepth;
	score := 0
	for ; myLoc <= endDepth; {
		myLoc++
		for i := range layers {
			if myLoc == layers[i].myDepth && layers[i].curLoc == 0 {
				score += layers[i].myDepth * layers[i].myRange
			}
			updateLayer(&layers[i])
		}
	}
	fmt.Println(score)
}

func updateLayer(layer *Layer) {
	if layer.curLoc+1 == layer.myRange {
		layer.direction = layer.direction * -1
	}
	if layer.curLoc == 0 {
		layer.direction = layer.direction * -1
	}
	layer.curLoc += layer.direction
}

func part2(layers []Layer) {
	timesteps := 0
	layersCopy := make([]Layer, len(layers))
	for ; true; {
		for i := range layers {
			layersCopy[i] = deepCopyLayer(layers[i])
		}
		if passClean(layersCopy) {
			break
		}
		timesteps++
		for i := range layers {
			updateLayer(&layers[i])
		}

	}
	fmt.Println(timesteps)
}

// Go is pass by value on structs so passing it to a function does a deep copy!
// NEAT!
func deepCopyLayer(layer Layer) Layer {
	return layer
}

// Mutates layers --> proceed with care
func passClean(layers []Layer) bool {
	myLoc, endDepth := -1, layers[len(layers)-1].myDepth;
	for ; myLoc <= endDepth; {
		myLoc++
		for i := range layers {
			if myLoc == layers[i].myDepth && layers[i].curLoc == 0 {
				return false
			}
			updateLayer(&layers[i])
		}
	}
	return true
}

func main() {
	firewall := readFile("assets/day13.in")
	part1(firewall)
	firewall = readFile("assets/day13.in")
	part2(firewall)
}
