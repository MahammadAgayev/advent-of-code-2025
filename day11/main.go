package day11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput(filename string) (map[string][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graph := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		source := strings.TrimSpace(parts[0])
		destinations := strings.Fields(strings.TrimSpace(parts[1]))

		graph[source] = destinations
	}

	return graph, scanner.Err()
}

func countPaths(graph map[string][]string, current string, target string, memo map[string]int) int {
	rs, ok := memo[current]

	if ok {
		return rs
	}

	if current == target {
		return 1
	}

	totalPaths := 0
	for _, neighbor := range graph[current] {
		totalPaths += countPaths(graph, neighbor, target, memo)
	}

	memo[current] = totalPaths
	return totalPaths
}

func solve(filename string) (int, error) {
	graph, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	memo := make(map[string]int)
	paths := countPaths(graph, "you", "out", memo)

	return paths, nil
}

func solve_dac(filename string) (int, error) {
	graph, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	// Case 1: Paths that visit fft first, then dac: svr → fft → dac → out
	memo1 := make(map[string]int)
	memo2 := make(map[string]int)
	memo3 := make(map[string]int)
	case1 := countPaths(graph, "svr", "fft", memo1) *
		countPaths(graph, "fft", "dac", memo2) *
		countPaths(graph, "dac", "out", memo3)

	// Case 2: Paths that visit dac first, then fft: svr → dac → fft → out
	memo4 := make(map[string]int)
	memo5 := make(map[string]int)
	memo6 := make(map[string]int)
	case2 := countPaths(graph, "svr", "dac", memo4) *
		countPaths(graph, "dac", "fft", memo5) *
		countPaths(graph, "fft", "out", memo6)

	return case1 + case2, nil
}

func Main() error {
	// Test with example
	testResult, err := solve_dac("day11/test.txt")
	if err != nil {
		fmt.Printf("Error with test input: %v\n", err)
	} else {
		fmt.Printf("Test result: %d paths\n", testResult)
	}

	// Actual input
	result, err := solve_dac("day11/input.txt")
	if err != nil {
		return err
	}
	fmt.Printf("Part 2 result: %d paths\n", result)

	return nil
}
