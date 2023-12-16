package task_3

import (
	"slices"
	"strconv"

	"k8s.io/apimachinery/pkg/util/sets"
	"simon-lorenz.dev/advent-of-code/2023/utils"
)

type Solver struct{}

type Neighbor struct {
	value    byte
	line     string
	index    int
	kind     SymbolKind
	position NeighborPosition
}

func NewNeighbor(value byte, line string, index int, position NeighborPosition) Neighbor {
	var kind SymbolKind

	if isNumber(value) {
		kind = skNumber
	} else if isGear(value) {
		kind = skGear
	} else if isSymbol(value) {
		kind = skSymbol
	} else {
		kind = skDot
	}

	return Neighbor{value, line, index, kind, position}
}

func (n Neighbor) is(kind SymbolKind) bool {
	return n.kind == kind
}

func (n Neighbor) isOneOf(kinds []SymbolKind) bool {
	return slices.Contains(kinds, n.kind)
}

type SymbolKind uint

const (
	skSymbol = iota
	skNumber = iota
	skGear   = iota
	skDot    = iota
)

type NeighborPosition uint

const (
	npUpperLeft  = iota
	npUpper      = iota
	npUpperRight = iota
	npLeft       = iota
	npRight      = iota
	npLowerLeft  = iota
	npLower      = iota
	npLowerRight = iota
)

func (Solver) GetPuzzleNumber() int {
	return 3
}

func (solver Solver) GetPuzzleInput() ([]string, error) {
	return utils.GetPuzzleInput(solver.GetPuzzleNumber())
}

func (solver Solver) SolveFirst() string {
	lines, _ := solver.GetPuzzleInput()

	sumOfPartNumbers := 0

	for i, line := range lines {
		var upper string
		var lower string

		if i > 0 {
			upper = lines[i-1]
		}

		if i < len(lines)-1 {
			lower = lines[i+1]
		}

		currentNumber := ""
		neighbors := make([]Neighbor, 0)

		for i, maybeNumber := range line {
			if isNumber(byte(maybeNumber)) {
				currentNumber += string(maybeNumber)
			}

			if len(currentNumber) > 0 {
				neighbors = append(neighbors, getNeighbors(line, upper, lower, i, []SymbolKind{skSymbol, skGear})...)

				if i == len(lines)-1 || !isNumber(line[i+1]) {
					// number ends, check neighbors
					if len(neighbors) > 0 {
						partNumber, _ := strconv.Atoi(currentNumber)
						sumOfPartNumbers += partNumber
					}

					currentNumber = ""
					neighbors = make([]Neighbor, 0)
				}
			}
		}
	}

	return strconv.Itoa(sumOfPartNumbers)
}

func (solver Solver) SolveSecond() string {
	lines, _ := solver.GetPuzzleInput()

	solution := 0

	for i, line := range lines {
		var upper string
		var lower string

		if i > 0 {
			upper = lines[i-1]
		}

		if i < len(lines)-1 {
			lower = lines[i+1]
		}

		for i, maybeGear := range line {
			if isGear(byte(maybeGear)) {
				neighbors := getNeighbors(line, upper, lower, i, []SymbolKind{skNumber})
				numbers := make(sets.Set[int], 0)

				for _, neighbor := range neighbors {
					if neighbor.kind == skNumber {
						numbers.Insert(getFullNumber(neighbor.line, neighbor.index))
					}
				}

				if numbers.Len() == 2 {
					firstPartNumber, _ := numbers.PopAny()
					secondPartNumber, _ := numbers.PopAny()

					solution += firstPartNumber * secondPartNumber
				}
			}
		}
	}

	return strconv.Itoa(solution)
}

func getNeighbors(line, upper, lower string, i int, allowed []SymbolKind) []Neighbor {
	neighbors := make([]Neighbor, 0)

	if upper != "" {
		appendNeighborIfAllowed(&neighbors, NewNeighbor(upper[i], upper, i, npUpper), allowed)
	}

	if lower != "" {
		appendNeighborIfAllowed(&neighbors, NewNeighbor(lower[i], lower, i, npLower), allowed)
	}

	if i > 0 {
		appendNeighborIfAllowed(&neighbors, NewNeighbor(line[i-1], line, i-1, npLeft), allowed)

		if lower != "" {
			appendNeighborIfAllowed(&neighbors, NewNeighbor(lower[i-1], lower, i-1, npLowerLeft), allowed)
		}

		if upper != "" {
			appendNeighborIfAllowed(&neighbors, NewNeighbor(upper[i-1], upper, i-1, npUpperLeft), allowed)
		}
	}

	if i < len(line)-1 {
		appendNeighborIfAllowed(&neighbors, NewNeighbor(line[i+1], line, i+1, npRight), allowed)

		if lower != "" {
			appendNeighborIfAllowed(&neighbors, NewNeighbor(lower[i+1], lower, i+1, npLowerRight), allowed)
		}

		if upper != "" {
			appendNeighborIfAllowed(&neighbors, NewNeighbor(upper[i+1], upper, i+1, npUpperRight), allowed)
		}
	}

	return neighbors
}

func getFullNumber(line string, i int) int {
	number := string(line[i])

	for j := i + 1; j < len(line); j++ {
		maybeNumber := line[j]

		if isNumber(maybeNumber) {
			number = number + string(maybeNumber)
		} else {
			break
		}
	}

	for j := i - 1; j >= 0; j-- {
		maybeNumber := line[j]

		if isNumber(maybeNumber) {
			number = string(maybeNumber) + number
		} else {
			break
		}
	}

	result, _ := strconv.Atoi(number)

	return result
}

func isSymbol(char byte) bool {
	return !isNumber(char) && !isGear(char) && !isDot(char)
}

func isNumber(char byte) bool {
	_, err := strconv.Atoi(string(char))
	return err == nil
}

func isGear(char byte) bool {
	return string(char) == "*"
}

func isDot(char byte) bool {
	return string(char) == "."
}

func appendNeighborIfAllowed(neighbors *[]Neighbor, neighbor Neighbor, allowed []SymbolKind) {
	if neighbor.isOneOf(allowed) {
		*neighbors = append(*neighbors, neighbor)
	}
}
