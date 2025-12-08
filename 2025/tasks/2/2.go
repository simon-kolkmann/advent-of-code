package task_2

import (
	"strconv"
	"strings"

	"simon-lorenz.dev/advent-of-code/2025/utils"
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
	ranges := strings.Split(lines[0], ",")

	result := 0

	for _, r := range ranges {
		start, end, _ := strings.Cut(r, "-")

		for i := utils.UnsafeAtoi(start); i <= utils.UnsafeAtoi(end); i++ {
			str := strconv.Itoa(i)

			if len(str)%2 != 0 {
				continue
			}

			first := str[:len(str)/2]
			second := str[len(str)/2:]

			if first == second {
				result += i
			}
		}
	}

	return strconv.Itoa(result)
}

func (solver Solver) SolveSecond() string {
	lines, _ := solver.GetPuzzleInput()
	ranges := strings.Split(lines[0], ",")

	result := 0

	for _, r := range ranges {
		start, end, _ := strings.Cut(r, "-")

		for i := utils.UnsafeAtoi(start); i <= utils.UnsafeAtoi(end); i++ {
			str := strconv.Itoa(i)

			for size := 1; size <= len(str)/2; size++ {
				if sequences := SplitStringIntoSequences(str, size); sequences != nil {
					if SequencesAreEqual(sequences) {
						result += i
						break
					}
				}
			}
		}
	}

	return strconv.Itoa(result)
}

func SplitStringIntoSequences(str string, size int) []string {
	sequences := make([]string, 0)

	if size == 0 || len(str)%size != 0 {
		return nil
	}

	for i := 0; i < len(str); i += size {
		sequences = append(sequences, str[i:i+size])
	}

	return sequences
}

func SequencesAreEqual(sequences []string) bool {
	first := sequences[0]

	for _, v := range sequences {
		if v != first {
			return false
		}
	}

	return true
}
