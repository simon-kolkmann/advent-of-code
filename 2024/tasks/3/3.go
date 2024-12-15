package task_3

import (
	"regexp"
	"strconv"
	"strings"

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
	line := strings.Join(lines, "")
	sum := 0

	operations := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`).FindAllString(line, -1)

	for _, operation := range operations {
		numbers := regexp.MustCompile(`\d{1,3}`).FindAllString(operation, 2)
		sum += utils.UnsafeAtoi(numbers[0]) * utils.UnsafeAtoi(numbers[1])
	}

	return strconv.Itoa(sum)
}

func (solver Solver) SolveSecond() string {
	lines, _ := solver.GetPuzzleInput()
	line := strings.Join(lines, "")
	sum := 0

	code := ""
	enabled := true

	for i, char := range line {
		if enabled {
			code += string(char)
		}

		if i > len(line)-7 {
			continue
		}

		if enabled && line[i:i+7] == "don't()" {
			enabled = false
		}

		if !enabled && line[i:i+4] == "do()" {
			enabled = true
		}
	}

	operations := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`).FindAllString(code, -1)

	for _, operation := range operations {
		numbers := regexp.MustCompile(`\d{1,3}`).FindAllString(operation, 2)
		sum += utils.UnsafeAtoi(numbers[0]) * utils.UnsafeAtoi(numbers[1])
	}

	return strconv.Itoa(sum)
}
