package task_3

import (
	"strconv"

	"simon-lorenz.dev/advent-of-code/2025/utils"
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

	for _, bank := range lines {
		m1 := 0
		m2 := 0

		for i, battery := range bank {
			battery := utils.UnsafeAtoi(string(battery))

			if m1 < battery && i < len(bank)-1 {
				m2 = 0
				m1 = battery
				continue
			}

			if m2 < battery {
				m2 = battery
			}
		}

		joltage := utils.UnsafeAtoi(strconv.Itoa(m1) + strconv.Itoa(m2))
		sum += joltage
	}

	return strconv.Itoa(sum)
}
