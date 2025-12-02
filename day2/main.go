package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/MahammadAgayev/advent-of-code2025/common"
)

type Range struct {
	Start int
	End int
}

func Main() error {
	ranges, err := readAndParse()

	if err != nil {
		return err
	}

	sum := 0

	for _,rg := range ranges {
		for i := rg.Start; i <= rg.End; i++ {
			// we can do better by not converting, but nevermind
			if invalidSequential(i) {
				fmt.Println(i)
				sum += i
			}
		}
	}

	fmt.Println(sum)

	return nil
}

func readAndParse() ([]Range, error) {
	scanner := common.ReadScanner()
	if !scanner.Scan() {
		return nil, fmt.Errorf("no line found")
	}

	line := scanner.Text()
	ranges := make([]Range, 0)

	for _, r := range strings.Split(line, ",") {
		rangeInfo := strings.Split(r, "-")

		start, err := strconv.Atoi(rangeInfo[0])

		if err != nil {
			return nil, err
		}

		end, err := strconv.Atoi(rangeInfo[1])

		if err != nil {
			return nil, err
		}

		rg := Range {
			Start: start,
			End: end,
		}

		ranges = append(ranges, rg)
	}

	return ranges, nil
}

func invalidNum(num string) bool {

	if len(num) % 2 != 0  {
		return false
	}

	mid := len(num) / 2
	for i:=0; i<mid; i++ {
		if num[i] != num[mid + i] {
			return false
		}
	}

	return true
}

func invalidSequential(num int) bool {
	ln := getLn(num)

	sequence, sequencei := findSequence(num, ln, 0);

	for sequencei != -1 {
		cycleFound :=0
		tmpSequence := 0

		for sequencei < ln {
			val := getNth(num, ln, sequencei)

			tmpSequence = tmpSequence * 10 + val

			if sequence == tmpSequence {
				tmpSequence = 0
				cycleFound++
			} else if getLn(tmpSequence) > getLn(sequence) {
				// tmpSequence grew larger than sequence without matching
				break
			}

			sequencei++
		}

		// Only check if we completed the full loop AND didn't break early
		if sequencei == ln && tmpSequence == 0 && getLn(sequence) * (cycleFound + 1) == ln {
			return true
		}

		// Try next sequence length (current sequence length + 1)
		sequence, sequencei = findSequence(num, ln, getLn(sequence))
	}

	return false
}

//Returns sequence and first index after sequence
func findSequence(num int, ln int, minLen int) (int, int) {
	sequence := getNth(num, ln, 0)
	sequencei := 1
	for sequencei < ln {
		val := getNth(num, ln, sequencei)
		if val == getNth(num, ln, 0) && getLn(sequence) > minLen {
			//found a sequence
			break
		}

		sequence = sequence * 10 + val
		sequencei++
	}

	if sequencei == ln {
		return  -1, -1
	}

	return sequence, sequencei
}


func getLn(num int) int {
	ln := int(math.Log10(float64(num))) + 1
	return ln
}

func getNth(num int, len_ int, idx int) int {

	//num = 654321
	//len_ = 6
	// idx = 3
	//output: 3

	//div = 10^2 = 100
	// if n = 0, then 10^5 = 100000
	div := int(math.Pow10(len_ - idx - 1))

	// 654321 / 100 = 6543
	// 6543 % 10 = 3
	num /= div
	num %= 10

	return num
}
