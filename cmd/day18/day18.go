package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func readFile(fpath string) ([][]string) {

	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Print(err)
	}

	s := string(b)
	moves := strings.Split(s, "\n")
	lines := make([][]string, len(moves))
	for i := range lines {
		lines[i] = strings.Split(moves[i], " ")
	}
	return lines
}

func initializeRegisters() map[string]int {
	regs := make(map[string]int)
	for i := 0; i < 26; i++ {
		v := 'a' + i
		regs[string(v)] = 0
	}
	regs["snd"] = 0
	return regs
}

func toInt(regs map[string]int, s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return regs[s]
	}
	return v
}

func set(regs map[string]int, i string, i2 string) int {
	v := toInt(regs, i2)
	regs[i] = v
	return 1
}

func add(regs map[string]int, i string, i2 string) int {
	v := toInt(regs, i2)
	regs[i] += v
	return 1
}

func mul(regs map[string]int, i string, i2 string) int {
	v := toInt(regs, i2)
	regs[i] *= v
	return 1
}

func mod(regs map[string]int, i string, i2 string) int {
	v := toInt(regs, i2)
	regs[i] = regs[i] % v
	return 1
}

func snd(regs map[string]int, i string) int {
	regs["snd"] = regs[i]
	return 1
}

func rcv(regs map[string]int, i string) int {
	v, _ := regs[i]
	if v != 0 {
		regs["rcv"] = regs["snd"]
	}
	return 1
}

func jgz(regs map[string]int, i string, i2 string) int {
	v, _ := regs[i]
	y := toInt(regs, i2)
	if v > 0 {
		return y
	}
	return 1
}

func part1(moves [][]string) int {
	regs := initializeRegisters()
	index := 0
	for ; index >= 0 && index < len(moves); {
		move := moves[index]
		//fmt.Println(index, move[0])
		v, exists := regs["rcv"]
		if exists {
			return v
		}
		if move[0] == "snd" {
			index += snd(regs, move[1])
		} else if move[0] == "set" {
			index += set(regs, move[1], move[2])
		} else if move[0] == "add" {
			index += add(regs, move[1], move[2])
		} else if move[0] == "mul" {
			index += mul(regs, move[1], move[2])
		} else if move[0] == "mod" {
			index += mod(regs, move[1], move[2])
		} else if move[0] == "rcv" {
			index += rcv(regs, move[1])
		} else if move[0] == "jgz" {
			index += jgz(regs, move[1], move[2])
		} else {
			fmt.Println("No Command " + move[0])
		}
	}
	fmt.Println("Exixted Bounds")
	return -1
}

func main() {
	lines := readFile("assets/day18.in")
	fmt.Println(part1(lines))
	//messages := make(chan int)
	//regs := initializeRegisters()
	//set(regs, "i", "2")
	//fmt.Println(regs["i"])

}
