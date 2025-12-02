package common

import (
	"bufio"
	"os"
)

func ReadScanner() *bufio.Scanner {
	file := os.Args[1]

	reader, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(reader)

	return scanner
}
