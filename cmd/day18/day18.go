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
	regs["index"] = 0
	return regs
}

func toInt(regs map[string]int, s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return regs[s]
	}
	return v
}

func set(regs map[string]int, i string, v int) {
	regs[i] = v
	regs["index"] += 1
}

func add(regs map[string]int, i string, v int) {
	regs[i] += v
	regs["index"] += 1
}

func mul(regs map[string]int, i string, v int) {
	regs[i] *= v
	regs["index"] += 1
}

func mod(regs map[string]int, i string, v int) {
	regs[i] = regs[i] % v
	regs["index"] += 1
}

func snd(regs map[string]int, v int) {
	regs["snd"] = v
	regs["index"] += 1
}

func rcv(regs map[string]int, i string) {
	v, _ := regs[i]
	if v != 0 {
		regs["rcv"] = regs["snd"]
	}
	regs["index"] += 1
}

func jgz(regs map[string]int, i string, v int) {
	regValue := regs[i]
	if regValue > 0 {
		regs["index"] += v
		return
	}
	regs["index"] += 1
}

func part1(moves [][]string) int {
	regs := initializeRegisters()
	for ; regs["index"] >= 0 && regs["index"] < len(moves); {
		move := moves[regs["index"]]
		fmt.Println(regs)
		fmt.Println(move)
		v, exists := regs["rcv"]
		if exists {
			return v
		}
		if move[0] == "snd" {
			v := toInt(regs, move[1])
			snd(regs, v)
		} else if move[0] == "set" {
			v := toInt(regs, move[2])
			set(regs, move[1], v)
		} else if move[0] == "add" {
			v := toInt(regs, move[2])
			add(regs, move[1], v)
		} else if move[0] == "mul" {
			v := toInt(regs, move[2])
			mul(regs, move[1], v)
		} else if move[0] == "mod" {
			v := toInt(regs, move[2])
			mod(regs, move[1], v)
		} else if move[0] == "rcv" {
			rcv(regs, move[1])
		} else if move[0] == "jgz" {
			v := toInt(regs, move[2])
			jgz(regs, move[1], v)
		} else {
			fmt.Println("No Command " + move[0])
		}
	}
	fmt.Println("Exited Bounds")
	return -1
}

func part2(moves [][]string) int {
	messages0 := make(chan int, 10000)
	regs0 := initializeRegisters()
	regs0["cnt"] = 0
	messages1 := make(chan int, 10000)
	regs1 := initializeRegisters()
	regs1["cnt"] = 0
	regs1["p"] = 1

	mut1, mut2 := true, true
	steps := 0
	for ; mut1 || mut2; {
		mut1 = duet(moves, regs0, messages0, messages1)
		mut2 = duet(moves, regs1, messages1, messages0)
		steps += 1
	}
	v, _ := regs1["cnt"]
	fmt.Println(v)
	return 0
}

func duet(moves [][]string, regs map[string]int,
	send chan int, receive chan int) (bool) {
	mutated := false
	for ; regs["index"] >= 0 && regs["index"] < len(moves); {
		move := moves[regs["index"]]
		if move[0] == "snd" {
			v := toInt(regs, move[1])
			ok := snd2(regs, v, send)
			if !ok {
				return mutated
			}
		} else if move[0] == "set" {
			v := toInt(regs, move[2])
			set(regs, move[1], v)
		} else if move[0] == "add" {
			v := toInt(regs, move[2])
			add(regs, move[1], v)
		} else if move[0] == "mul" {
			v := toInt(regs, move[2])
			mul(regs, move[1], v)
		} else if move[0] == "mod" {
			v := toInt(regs, move[2])
			mod(regs, move[1], v)
		} else if move[0] == "rcv" {
			ok := rcv2(regs, move[1], receive)
			if !ok {
				return mutated
			}
		} else if move[0] == "jgz" {
			v := toInt(regs, move[2])
			jgz(regs, move[1], v)
		} else {
			fmt.Println("No Command " + move[0])
		}
		mutated = true
	}
	return mutated
}

func rcv2(regs map[string]int, reg string, channel chan int) (bool) {
	select {
	case regs[reg] = <-channel:
		regs["index"] += 1
		return true
	default:
		return false
	}
}

func snd2(regs map[string]int, v int, channel chan int) (bool) {
	select {
	case channel <- v:
		regs["cnt"] += 1
		regs["index"] += 1
		return true
	default:
		return false
	}
}

func main() {
	lines := readFile("assets/day18.in")
	//fmt.Println(part1(lines))
	part2(lines)
}
