package task_3

import (
	"strconv"

	"simon-lorenz.dev/advent-of-code/2023/utils"
)

type Solver struct{}

func (Solver) GetPuzzleNumber() int {
	return 3
}

func (solver Solver) GetPuzzleInput() ([]string, error) {
	return utils.GetPuzzleInput(solver.GetPuzzleNumber())
}

func (solver Solver) SolveFirst() string {
	lines, _ := solver.GetPuzzleInput()

	sumOfPartNumbers := 0

	for i, line := range lines {
		var upper string
		var lower string

		if i > 0 {
			upper = lines[i-1]
		}

		if i < len(lines)-1 {
			lower = lines[i+1]
		}

		currentNumber := ""
		neighbors := make([]byte, 0)

		for i, maybeNumber := range line {
			if isNumber(byte(maybeNumber)) {
				currentNumber += string(maybeNumber)
			}

			if len(currentNumber) > 0 {
				if upper != "" {
					appendIfSymbol(&neighbors, upper[i]) // top
				}

				if lower != "" {
					appendIfSymbol(&neighbors, lower[i]) // bottom
				}

				if i > 0 && len(currentNumber) == 1 {
					appendIfSymbol(&neighbors, line[i-1]) // left

					if lower != "" {
						appendIfSymbol(&neighbors, lower[i-1]) // bottom left
					}

					if upper != "" {
						appendIfSymbol(&neighbors, upper[i-1]) // upper left
					}
				}

				var rightNeighbor byte = 0

				if i < len(line)-1 {
					rightNeighbor = line[i+1]

					if !isNumber(rightNeighbor) {
						appendIfSymbol(&neighbors, rightNeighbor) // right

						if lower != "" {
							appendIfSymbol(&neighbors, lower[i+1]) // bottom right
						}

						if upper != "" {
							appendIfSymbol(&neighbors, upper[i+1]) // upper right
						}
					}
				}

				if !isNumber(rightNeighbor) {
					// number ends, check neighbors
					if len(neighbors) > 0 {
						partNumber, _ := strconv.Atoi(currentNumber)
						sumOfPartNumbers += partNumber
					}

					currentNumber = ""
					neighbors = make([]byte, 0)
				}
			}
		}
	}

	return strconv.Itoa(sumOfPartNumbers)
}

func (solver Solver) SolveSecond() string {
	return "N/A"
}

func isSymbol(char byte) bool {
	return !isNumber(char) && string(char) != "."
}

func isNumber(char byte) bool {
	_, err := strconv.Atoi(string(char))
	return err == nil
}

func appendIfSymbol(neighbors *[]byte, maybeSymbol byte) {
	if isSymbol(maybeSymbol) {
		*neighbors = append(*neighbors, maybeSymbol)
	}
}
