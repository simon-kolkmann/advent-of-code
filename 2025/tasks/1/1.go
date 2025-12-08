package task_1

import (
	"strconv"

	"simon-lorenz.dev/advent-of-code/2025/utils"
)

type Solver struct {
	utils.BaseSolver
}

func NewSolver() Solver {
	return Solver{
		utils.NewBaseSolver(1),
	}
}

func (solver Solver) SolveFirst() string {
	lines, _ := solver.GetPuzzleInput()

	zeroes := 0
	pos := 50

	for _, line := range lines {
		direction := string(line[0])
		count := utils.UnsafeAtoi(line[1:])

		if direction == "R" {
			for count > 0 {
				pos++
				count--

				if pos == 100 {
					pos = 0
				}
			}
		} else {
			for count > 0 {
				pos--
				count--

				if pos == -1 {
					pos = 99
				}
			}
		}

		if pos == 0 {
			zeroes++
		}
	}

	return strconv.Itoa(zeroes)
}

func (solver Solver) SolveSecond() string {
	lines, _ := solver.GetPuzzleInput()

	zeroes := 0
	pos := 50

	for _, line := range lines {
		direction := string(line[0])
		count := utils.UnsafeAtoi(line[1:])

		if direction == "R" {
			for count > 0 {
				if pos == 0 {

					zeroes++
				}

				pos++
				count--

				if pos == 100 {
					pos = 0
				}
			}
		} else {
			for count > 0 {
				if pos == 0 {
					zeroes++
				}

				pos--
				count--

				if pos == -1 {
					pos = 99
				}
			}
		}
	}

	return strconv.Itoa(zeroes)
}
