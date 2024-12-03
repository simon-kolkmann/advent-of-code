package task_2

import (
	"strconv"
	"strings"

	"simon-lorenz.dev/advent-of-code/2024/utils"
)

type Solver struct {
	utils.BaseSolver
}

func NewSolver() Solver {
	return Solver{
		utils.NewBaseSolver(2),
	}
}

func (solver Solver) SolveFirst() string {
	lines, _ := solver.GetPuzzleInput()

	safe := 0

	for _, report := range lines {
		levels := strings.Split(report, " ")

		isAscending := true
		isDescending := true
		hasPlateau := false
		hasTooLargeDistance := false

		for i, level := range levels {
			if i > 0 {
				curr := utils.UnsafeAtoi(level)
				prev := utils.UnsafeAtoi(levels[i-1])
				distance := curr - prev

				if distance < 0 {
					isAscending = false
				} else if distance > 0 {
					isDescending = false
				} else {
					hasPlateau = true
					break
				}

				if utils.Abs(distance) > 3 {
					hasTooLargeDistance = true
					break
				}
			}
		}

		if (!isAscending && !isDescending) || hasPlateau || hasTooLargeDistance {
			continue
		}

		safe++
	}

	return strconv.Itoa(safe)
}

func (solver Solver) SolveSecond() string {
	// lines, _ := solver.GetPuzzleInput()

	return ""
}
