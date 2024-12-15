package task_3

import (
	"regexp"
	"strconv"

	"simon-lorenz.dev/advent-of-code/2024/utils"
)

type Solver struct {
	utils.BaseSolver
}

func NewSolver() Solver {
	return Solver{
		utils.NewBaseSolver(3),
	}
}

func (solver Solver) SolveFirst() string {
	lines, _ := solver.GetPuzzleInput()
	sum := 0

	for _, line := range lines {
		operations := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`).FindAllString(line, -1)

		for _, operation := range operations {
			numbers := regexp.MustCompile(`\d{1,3}`).FindAllString(operation, 2)
			sum += utils.UnsafeAtoi(numbers[0]) * utils.UnsafeAtoi(numbers[1])
		}
	}

	return strconv.Itoa(sum)
}

func (solver Solver) SolveSecond() string {
	return ""
}
