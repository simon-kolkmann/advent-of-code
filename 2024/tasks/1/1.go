package task_1

import (
	"slices"
	"strconv"
	"strings"

	"simon-lorenz.dev/advent-of-code/2024/utils"
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

	a := make([]int, 0)
	b := make([]int, 0)

	for _, line := range lines {
		numbers := strings.Split(line, "   ")

		a = append(a, utils.UnsafeAtoi(numbers[0]))
		b = append(b, utils.UnsafeAtoi(numbers[1]))
	}

	slices.Sort(a)
	slices.Sort(b)

	total := 0

	for i := 0; i < len(lines); i++ {
		distance := utils.Abs(a[i] - b[i])
		total += distance
	}

	return strconv.Itoa(total)
}

func (solver Solver) SolveSecond() string {
	lines, _ := solver.GetPuzzleInput()

	numbers := make([]int, 0)
	counters := make(map[int]int)

	for _, line := range lines {
		nums := strings.Split(line, "   ")

		a := utils.UnsafeAtoi(nums[0])
		b := utils.UnsafeAtoi(nums[1])

		numbers = append(numbers, a)
		counters[b]++
	}

	score := 0

	for _, number := range numbers {
		score += number * counters[number]
	}

	return strconv.Itoa(score)
}
