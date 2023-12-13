package task_2

import (
	"regexp"
	"strconv"
	"strings"

	"simon-lorenz.dev/advent-of-code/2023/utils"
)

type Solver struct{}

func (Solver) GetPuzzleNumber() int {
	return 2
}

func (solver Solver) GetPuzzleInput() ([]string, error) {
	return utils.GetPuzzleInput(solver.GetPuzzleNumber())
}

func (solver Solver) SolveFirst() string {
	lines, _ := solver.GetPuzzleInput()

	solution := 0

	for _, line := range lines {
		var stats map[string]int = make(map[string]int)

		splitted := strings.Split(line, ": ")
		game := splitted[0]
		rounds := splitted[1]

		number, _ := strconv.Atoi(regexp.MustCompile(`\d+`).FindStringSubmatch(game)[0])

		for _, round := range strings.Split(rounds, "; ") {
			for _, cubesOfColor := range strings.Split(round, ", ") {
				countAndColor := strings.Split(cubesOfColor, " ")
				count, color := countAndColor[0], countAndColor[1]

				countAsInt, _ := strconv.Atoi(count)

				if stats[color] < countAsInt {
					stats[color] = countAsInt
				}
			}
		}

		if stats["red"] <= 12 && stats["green"] <= 13 && stats["blue"] <= 14 {
			solution += number
		}
	}

	return strconv.Itoa(solution)
}

func (solver Solver) SolveSecond() string {
	lines, _ := solver.GetPuzzleInput()

	solution := 0

	for _, line := range lines {
		var stats map[string]int = make(map[string]int)

		splitted := strings.Split(line, ": ")
		rounds := splitted[1]

		for _, round := range strings.Split(rounds, "; ") {
			for _, cubesOfColor := range strings.Split(round, ", ") {
				countAndColor := strings.Split(cubesOfColor, " ")
				count, color := countAndColor[0], countAndColor[1]

				countAsInt, _ := strconv.Atoi(count)

				if stats[color] < countAsInt {
					stats[color] = countAsInt
				}
			}
		}

		solution += stats["red"] * stats["green"] * stats["blue"]
	}

	return strconv.Itoa(solution)
}
