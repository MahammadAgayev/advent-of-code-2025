package day3

import (
	"fmt"
	"strconv"

	"github.com/MahammadAgayev/advent-of-code2025/common"
)


func Main() error {

	scanner :=  common.ReadScanner()

	totalPart1 := 0
	totalPart2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		maxJoltage, err := findMaxJoltage(line, 2)
		if err != nil {
			return err
		}
		totalPart1 += maxJoltage

		maxJoltage12, err := findMaxJoltage(line, 12)
		if err != nil {
			return err
		}
		totalPart2 += maxJoltage12
	}

	fmt.Println("Part 1:", totalPart1)
	fmt.Println("Part 2:", totalPart2)

	return nil
}

func findMaxJoltage(line string, keep int) (int, error) {
	skip := len(line) - keep
	result := ""
	start := 0

	for i := 0; i < keep; i++ {
		// We need to pick 'keep - i' more digits from the remaining string
		// We can look ahead 'skip + 1' positions to find the best digit
		remaining := keep - i
		maxDigit := byte('0')
		maxPos := start

		// Find the maximum digit in the window where we can still pick enough digits after it
		end := len(line) - remaining + 1
		for j := start; j < end; j++ {
			if line[j] > maxDigit {
				maxDigit = line[j]
				maxPos = j
			}
		}

		result += string(maxDigit)
		skip -= (maxPos - start)
		start = maxPos + 1
	}

	return strconv.Atoi(result)
}
