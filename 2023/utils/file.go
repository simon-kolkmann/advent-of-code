package utils

import (
	"bufio"
	"os"
	"strconv"
)

func GetPuzzleInput(puzzleNumber int) ([]string, error) {
	path := "tasks/" + strconv.Itoa(puzzleNumber) + "/input.txt"

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			lines = append(lines, scanner.Text())
		}
	}

	return lines, scanner.Err()
}
