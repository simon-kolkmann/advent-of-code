package task_4

import (
	"math"
	"regexp"
	"strconv"

	"k8s.io/apimachinery/pkg/util/sets"
	"simon-lorenz.dev/advent-of-code/2023/utils"
)

type Solver struct{}

func (Solver) GetPuzzleNumber() int {
	return 4
}

func (solver Solver) GetPuzzleInput() ([]string, error) {
	return utils.GetPuzzleInput(solver.GetPuzzleNumber())
}

func (solver Solver) SolveFirst() string {
	lines, _ := solver.GetPuzzleInput()
	solution := 0

	for _, line := range lines {
		matches := regexp.MustCompile(`Card .*:((?: \s*\d+)+) \|((?: \s*\d+)+)`).FindAllStringSubmatch(line, -1)
		winningNumbers := make(sets.Set[int])
		numbers := make(sets.Set[int])

		for _, num := range regexp.MustCompile(`\d+`).FindAllString(matches[0][1], -1) {
			if num, err := strconv.Atoi(num); err == nil {
				winningNumbers.Insert(num)
			}
		}

		for _, num := range regexp.MustCompile(`\d+`).FindAllString(matches[0][2], -1) {
			if num, err := strconv.Atoi(num); err == nil {
				numbers.Insert(num)
			}
		}

		myWinningNumbers := winningNumbers.Intersection(numbers)
		solution += int(math.Pow(2, float64(myWinningNumbers.Len()-1)))
	}

	return strconv.Itoa(solution)
}

func (solver Solver) SolveSecond() string {
	lines, _ := solver.GetPuzzleInput()
	won := make([]int, 0)
	solution := 0

	for i, line := range lines {
		matches := regexp.MustCompile(`Card\s+(\d+):((?: \s*\d+)+) \|((?: \s*\d+)+)`).FindAllStringSubmatch(line, -1)
		winningNumbers := make(sets.Set[int])
		numbers := make(sets.Set[int])

		card, _ := strconv.Atoi(matches[0][1])

		for _, num := range regexp.MustCompile(`\d+`).FindAllString(matches[0][2], -1) {
			if num, err := strconv.Atoi(num); err == nil {
				winningNumbers.Insert(num)
			}
		}

		for _, num := range regexp.MustCompile(`\d+`).FindAllString(matches[0][3], -1) {
			if num, err := strconv.Atoi(num); err == nil {
				numbers.Insert(num)
			}
		}

		myWinningNumbers := winningNumbers.Intersection(numbers)
		sumOfOriginalAndClones := 0

		for _, won := range won {
			if won == card {
				sumOfOriginalAndClones++
			}
		}

		for j := i + 1; j <= i+myWinningNumbers.Len(); j++ {
			for k := sumOfOriginalAndClones; k >= 0; k-- {
				won = append(won, j+1)
			}
		}
	}

	solution = len(won) + len(lines)

	return strconv.Itoa(solution)
}
