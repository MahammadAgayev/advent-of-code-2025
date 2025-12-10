package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MahammadAgayev/advent-of-code2025/common"
)

func Main() error {
	scanner := common.ReadScanner()

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		return fmt.Errorf("no input found")
	}

	// Part 1
	problems := parseProblems(lines)
	grandTotal := 0
	for _, problem := range problems {
		result := calculateProblem(problem)
		grandTotal += result
	}
	fmt.Println("Part 1:", grandTotal)

	// Part 2
	problemsPart2 := parseProblemsPart2(lines)
	grandTotalPart2 := 0
	for _, problem := range problemsPart2 {
		result := calculateProblem(problem)
		grandTotalPart2 += result
	}
	fmt.Println("Part 2:", grandTotalPart2)

	return nil
}

type Problem struct {
	Numbers []int
	Operation byte
}

func parseProblems(lines []string) []Problem {
	if len(lines) == 0 {
		return []Problem{}
	}

	width := len(lines[0])
	problems := []Problem{}

	col := 0
	for col < width {
		// Skip spaces between problems
		if col < width && isEmptyColumn(lines, col) {
			col++
			continue
		}

		// Found start of a problem - collect all columns until we hit empty column
		problemCols := []int{}
		for col < width && !isEmptyColumn(lines, col) {
			problemCols = append(problemCols, col)
			col++
		}

		if len(problemCols) > 0 {
			problem := extractProblem(lines, problemCols)
			problems = append(problems, problem)
		}
	}

	return problems
}

func isEmptyColumn(lines []string, col int) bool {
	for _, line := range lines {
		if col < len(line) && line[col] != ' ' {
			return false
		}
	}
	return true
}

func extractProblem(lines []string, cols []int) Problem {
	numbers := []int{}
	var operation byte

	// Read rows - last row has the operation
	for i := 0; i < len(lines)-1; i++ {
		numStr := ""
		for _, col := range cols {
			if col < len(lines[i]) {
				ch := lines[i][col]
				if ch != ' ' {
					numStr += string(ch)
				}
			}
		}

		numStr = strings.TrimSpace(numStr)
		if numStr != "" {
			num, err := strconv.Atoi(numStr)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	// Get operation from last row
	lastRow := lines[len(lines)-1]
	for _, col := range cols {
		if col < len(lastRow) && lastRow[col] != ' ' {
			operation = lastRow[col]
			break
		}
	}

	return Problem{
		Numbers: numbers,
		Operation: operation,
	}
}

func calculateProblem(problem Problem) int {
	if len(problem.Numbers) == 0 {
		return 0
	}

	result := problem.Numbers[0]

	for i := 1; i < len(problem.Numbers); i++ {
		if problem.Operation == '+' {
			result += problem.Numbers[i]
		} else if problem.Operation == '*' {
			result *= problem.Numbers[i]
		}
	}

	return result
}

// Part 2: Read problems right-to-left, each column forms a number top-to-bottom
func parseProblemsPart2(lines []string) []Problem {
	if len(lines) == 0 {
		return []Problem{}
	}

	// Add space to end of each line for consistent processing
	paddedLines := make([]string, len(lines))
	for i, line := range lines {
		paddedLines[i] = line + " "
	}

	// Get operator line and reverse it
	opLine := paddedLines[len(paddedLines)-1]
	reversed := reverseString(opLine)

	// Split by operators (keeping operator with each segment)
	segments := splitInclusiveByOperators(reversed)

	// Reverse segments to process left-to-right in original order
	for i, j := 0, len(segments)-1; i < j; i, j = i+1, j-1 {
		segments[i], segments[j] = segments[j], segments[i]
	}

	problems := []Problem{}
	position := 0

	for _, segment := range segments {
		// Get operator (last char of segment)
		op := segment[len(segment)-1]

		// For each column in this segment, read top-to-bottom to form numbers
		numbers := []int{}
		for i := 0; i < len(segment); i++ {
			acc := 0
			hasDigit := false

			// Read vertically through all number rows (not operator row)
			for row := 0; row < len(paddedLines)-1; row++ {
				colPos := position + i
				if colPos < len(paddedLines[row]) {
					ch := paddedLines[row][colPos]
					if ch >= '0' && ch <= '9' {
						acc = acc*10 + int(ch-'0')
						hasDigit = true
					}
				}
			}

			if hasDigit {
				numbers = append(numbers, acc)
			}
		}

		// Reverse numbers to get right-to-left order
		for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}

		problems = append(problems, Problem{
			Numbers:   numbers,
			Operation: op,
		})

		position += len(segment)
	}

	return problems
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func splitInclusiveByOperators(s string) []string {
	result := []string{}
	current := ""

	for _, ch := range s {
		current += string(ch)
		if ch == '+' || ch == '*' {
			result = append(result, current)
			current = ""
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}
