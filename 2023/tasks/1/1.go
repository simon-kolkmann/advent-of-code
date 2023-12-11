package task_1

import (
	"strconv"
	"strings"

	"simon-lorenz.dev/advent-of-code/2023/utils"
)

type Solver struct{}

func (Solver) GetPuzzleNumber() int {
	return 1
}

func (solver Solver) GetPuzzleInput() ([]string, error) {
	return utils.GetPuzzleInput(solver.GetPuzzleNumber())
}

func (solver Solver) SolveFirst() string {
	lines, _ := solver.GetPuzzleInput()

	solution := 0

	for _, line := range lines {
		numbersInLine := make([]string, 0)

		for _, char := range line {
			maybeNumber := string(char)

			if _, err := strconv.Atoi(maybeNumber); err == nil {
				numbersInLine = append(numbersInLine, maybeNumber)
			}
		}

		var firstNumber string = numbersInLine[0]
		var secondNumber string

		if len(numbersInLine) == 1 {
			secondNumber = firstNumber
		} else {
			secondNumber = numbersInLine[len(numbersInLine)-1]
		}

		result, _ := strconv.Atoi(firstNumber + "" + secondNumber)

		solution += result
	}

	return strconv.Itoa(solution)
}

func (solver Solver) SolveSecond() string {
	lines, _ := solver.GetPuzzleInput()
	solution := 0
	var tokens map[string]string = make(map[string]string)

	tokens["one"] = "1"
	tokens["two"] = "2"
	tokens["three"] = "3"
	tokens["four"] = "4"
	tokens["five"] = "5"
	tokens["six"] = "6"
	tokens["seven"] = "7"
	tokens["eight"] = "8"
	tokens["nine"] = "9"

	availableTokens := make([]string, 0, len(tokens))

	for token := range tokens {
		availableTokens = append(availableTokens, token)
	}

	for _, line := range lines {
		currentSubstring := ""
		numbersInLine := make([]string, 0)

		for _, char := range line {
			maybeNumber := string(char)

			if _, err := strconv.Atoi(maybeNumber); err == nil {
				numbersInLine = append(numbersInLine, maybeNumber)
				currentSubstring = ""
			} else {
				currentSubstring += maybeNumber

				for _, availableToken := range availableTokens {
					if strings.Contains(currentSubstring, availableToken) {
						numbersInLine = append(numbersInLine, tokens[availableToken])

						// We need to account for overlaps, e.g. "eighthree" (-> 83) or "sevenine" (-> 79).
						// Because of this, we'll reset the currentSubstring not to "" but the last character of the
						// previous substring.
						previousLastCharacter := currentSubstring[len(currentSubstring)-1]
						currentSubstring = string(previousLastCharacter)
						break
					}
				}
			}
		}

		var firstNumber string = numbersInLine[0]
		var secondNumber string

		if len(numbersInLine) == 1 {
			secondNumber = firstNumber
		} else {
			secondNumber = numbersInLine[len(numbersInLine)-1]
		}

		result, _ := strconv.Atoi(firstNumber + "" + secondNumber)

		solution += result
	}

	return strconv.Itoa(solution)
}
