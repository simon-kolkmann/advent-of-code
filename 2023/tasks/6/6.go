package task_6

import (
	"regexp"
	"strconv"
	"strings"

	"simon-lorenz.dev/advent-of-code/2023/utils"
)

type Solver struct {
	utils.BaseSolver
}

func NewSolver() Solver {
	return Solver{
		utils.NewBaseSolver(6),
	}
}

func (solver Solver) SolveFirst() string {
	lines, _ := solver.GetPuzzleInput()

	regex := regexp.MustCompile(`\d+`)

	times := regex.FindAllString(lines[0], -1)
	distances := regex.FindAllString(lines[1], -1)

	possibilities := make([]int, 0)

	for i, time := range times {
		time, _ := strconv.Atoi(time)
		distanceToBeat, _ := strconv.Atoi(distances[i])

		possibilities = append(possibilities, 0)

		for j := 1; j < time; j++ {
			distance := j * (time - j)

			if distance > distanceToBeat {
				possibilities[i]++
			}
		}
	}

	solution := 1

	for _, possibility := range possibilities {
		solution *= possibility
	}

	return strconv.Itoa(solution)
}

func (solver Solver) SolveSecond() string {
	lines, _ := solver.GetPuzzleInput()

	regex := regexp.MustCompile(`\d+`)

	time, _ := strconv.Atoi(strings.Join(regex.FindAllString(lines[0], -1), ""))
	distanceToBeat, _ := strconv.Atoi(strings.Join(regex.FindAllString(lines[1], -1), ""))
	possibilities := 0

	for j := 1; j < time; j++ {
		distance := j * (time - j)

		if distance > distanceToBeat {
			possibilities++
		}
	}
	return strconv.Itoa(possibilities)
}
