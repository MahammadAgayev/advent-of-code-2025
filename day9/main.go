package day9

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

	data := parseInput(lines)

	// Part 1
	part1 := largestArea(&data)
	fmt.Println("Part 1:", part1)

	// Part 2
	part2 := largestAreaRedGreen(&data)
	fmt.Println("Part 2:", part2)

	return nil
}

type Position struct {
	X int
	Y int
}

func (p Position) area(other Position) int {
	width := max(p.X, other.X) - min(p.X, other.X) + 1
	height := max(p.Y, other.Y) - min(p.Y, other.Y) + 1
	return width * height
}

func (p Position) topLeft(other Position) Position {
	return Position{
		X: min(p.X, other.X),
		Y: min(p.Y, other.Y),
	}
}

func (p Position) bottomRight(other Position) Position {
	return Position{
		X: max(p.X, other.X),
		Y: max(p.Y, other.Y),
	}
}

type Data struct {
	Positions []Position
}

type Edge struct {
	From Position
	To   Position
}

func (e Edge) outside(a, b Position) bool {
	topLeft := a.topLeft(b)
	bottomRight := a.bottomRight(b)

	// Returns true if edge is NOT completely outside (i.e., it intersects or is inside)
	completelyOutside := (e.From.X <= topLeft.X && e.To.X <= topLeft.X) ||
		(e.From.Y <= topLeft.Y && e.To.Y <= topLeft.Y) ||
		(e.From.X+1 >= bottomRight.X && e.To.X+1 >= bottomRight.X) ||
		(e.From.Y+1 >= bottomRight.Y && e.To.Y+1 >= bottomRight.Y)

	return !completelyOutside
}

func parseInput(lines []string) Data {
	positions := []Position{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
			y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err1 == nil && err2 == nil {
				positions = append(positions, Position{X: x, Y: y})
			}
		}
	}
	return Data{Positions: positions}
}

func largestArea(data *Data) int {
	biggest := 0
	for i := 0; i < len(data.Positions); i++ {
		for j := i; j < len(data.Positions); j++ {
			area := data.Positions[i].area(data.Positions[j])
			if area > biggest {
				biggest = area
			}
		}
	}
	return biggest
}

func largestAreaRedGreen(data *Data) int {
	if len(data.Positions) == 0 {
		return 0
	}

	// Build edges connecting consecutive positions
	edges := []Edge{}
	for i := 0; i < len(data.Positions); i++ {
		a := data.Positions[i]
		b := data.Positions[(i+1)%len(data.Positions)]
		edges = append(edges, Edge{
			From: a.topLeft(b),
			To:   a.bottomRight(b),
		})
	}

	biggest := 0
	for i := 0; i < len(data.Positions); i++ {
		checkLoop:
		for j := 0; j < len(data.Positions); j++ {
			if i == j {
				continue
			}

			a := data.Positions[i]
			b := data.Positions[j]

			// Check if any edge is outside this rectangle
			for _, edge := range edges {
				if edge.outside(a, b) {
					continue checkLoop
				}
			}

			// All edges are inside, this is valid
			area := a.area(b)
			if area > biggest {
				biggest = area
			}
		}
	}
	return biggest
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
