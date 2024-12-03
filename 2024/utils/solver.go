package utils

import (
	"bufio"
	"os"
	"strconv"
)

type BaseSolver struct {
	puzzle int
}

func NewBaseSolver(puzzle int) BaseSolver {
	return BaseSolver{
		puzzle,
	}
}

func (solver BaseSolver) GetPuzzleNumber() int {
	return solver.puzzle
}

func (solver BaseSolver) GetPuzzleInput() ([]string, error) {
	path := "tasks/" + strconv.Itoa(solver.GetPuzzleNumber()) + "/input.txt"

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

func (solver BaseSolver) SolveFirst() string {
	return "N/A"
}

func (solver BaseSolver) SolveSecond() string {
	return "N/A"
}
