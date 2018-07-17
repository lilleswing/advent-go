package knothash

import "fmt"

func createLengths(s string) [] int {
	l := make([]int, len(s), len(s)+5)
	for i := range s {
		l[i] = int(s[i])
	}
	l = append(l, 17, 31, 73, 47, 23)
	return l
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

func hashIt(lengths []int, size int) string {
	spHash := knot(lengths, size, 64)
	dHash := denseHash(spHash)
	s := ""
	for i := range dHash {
		s += toHex(dHash[i])
	}
	return s
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

func toHex(i int) string {
	s := fmt.Sprintf("%x", i)
	if len(s) == 1 {
		return "0" + s
	}
	return s
}

func KnotHash(s string) string {
	lengths := createLengths(s)
	return hashIt(lengths, 256)
}
