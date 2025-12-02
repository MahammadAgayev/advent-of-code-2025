package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Main() {
	file := os.Args[1]

	reader, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(reader)
	pointer := 50
	pass := 0

	for scanner.Scan() {
		line := scanner.Text()

		val, err := strconv.Atoi(line[1:])

		if err != nil {
			panic(err)
		}

		if line[0] == 'L' {
			val*=-1
		}

		pointerNew, zeros := circulate(pointer, val)
		pointer = pointerNew

		pass+=zeros
	}

	fmt.Println(pass)
}

// floorDiv performs floor division (rounds toward negative infinity)
func floorDiv(a, b int) int {
	q := a / b
	r := a % b

	// Adjust if remainder is non-zero and signs differ
	if r != 0 && (a < 0) != (b < 0) {
		q--
	}
	return q
}

// Returns pointer and hit count to 0
func circulate(pointer int, val int) (int, int) {
	zeros := 0

	if val > 0 {
		zeros = floorDiv(pointer+val, 100) - 0
	} else if val < 0 {
		// escaping from the zero
		zeros = floorDiv(pointer-1, 100) - floorDiv(pointer+val-1, 100)
	}

	newPointer := ((pointer + val) % 100 + 100) % 100

	return newPointer, zeros
}
