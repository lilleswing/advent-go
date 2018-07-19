package main

import "fmt"

func nextA(i int) int {
	return (i * 16807) % 2147483647

}

func nextB(i int) int {
	return (i * 48271) % 2147483647
}

func p2A(i int) int {
	i = nextA(i)
	for ; i%4 != 0; {
		i = nextA(i)
	}
	return i
}

func p2B(i int) int {
	i = nextB(i)
	for ; i%8 != 0; {
		i = nextB(i)
	}
	return i
}

func maskEq(a int, b int, mask int) bool {
	a = a & mask
	b = b & mask
	if a == b {
		return true
	}
	return false

}

func bit16() int {
	start := 1
	total := 0
	for i := 0; i < 16; i++ {
		total += start
		start *= 2
	}
	return total
}

func part1(a int, b int) {
	total := 0
	mask := bit16()
	for i := 0; i < 40000000; i++ {
		a = nextA(a)
		b = nextB(b)
		if maskEq(a, b, mask) {
			total += 1
		}

	}
	fmt.Println(total)
}

func part2(a int, b int) {
	total := 0
	mask := bit16()
	for i := 0; i < 5000000; i++ {
		a = p2A(a)
		b = p2B(b)
		if maskEq(a, b, mask) {
			total += 1
		}

	}
	fmt.Println(total)
}

//Generator A starts with 618
//Generator B starts with 814
func main() {
	a, b := 618, 814
	part1(a, b)
	part2(a, b)

}
