package task_5

import (
	"regexp"
	"strconv"
	"strings"

	"simon-lorenz.dev/advent-of-code/2023/utils"
)

type Solver struct{}

func (Solver) GetPuzzleNumber() int {
	return 5
}

func (solver Solver) GetPuzzleInput() ([]string, error) {
	return utils.GetPuzzleInput(solver.GetPuzzleNumber())
}

type ResourceMap struct {
	source      string
	destination string
	entries     []ResourceMapEntry
}

func (resourceMap ResourceMap) convertNumber(source int) int {
	for _, entry := range resourceMap.entries {
		if (entry.sourceRangeStart <= source) && (source <= entry.sourceRangeStart+entry.rangeLength-1) {
			return entry.destinationRangeStart + source - entry.sourceRangeStart
		}
	}

	return source
}

type ResourceMapEntry struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func (solver Solver) SolveFirst() string {
	lines, _ := solver.GetPuzzleInput()

	var seeds []int

	for _, seed := range regexp.MustCompile(`\d+`).FindAllString(lines[0], -1) {
		seed, _ := strconv.Atoi(seed)
		seeds = append(seeds, seed)
	}

	maps := make([]ResourceMap, 0)

	for i := 1; i < len(lines); i++ {
		line := lines[i]

		if strings.Contains(line, "map:") {
			mapRegex := regexp.MustCompile(`(\w*)-to-(\w*) map:`).FindAllStringSubmatch(line, -1)

			maps = append(maps, ResourceMap{
				source:      mapRegex[0][1],
				destination: mapRegex[0][2],
				entries:     make([]ResourceMapEntry, 0),
			})

			continue
		}

		currentMap := &maps[len(maps)-1]
		mapEntryRegex := regexp.MustCompile(`\d+`).FindAllString(line, -1)

		destinationRangeStart, _ := strconv.Atoi(mapEntryRegex[0])
		sourceRangeStart, _ := strconv.Atoi(mapEntryRegex[1])
		rangeLength, _ := strconv.Atoi(mapEntryRegex[2])

		currentMap.entries = append(currentMap.entries, ResourceMapEntry{
			destinationRangeStart,
			sourceRangeStart,
			rangeLength,
		})
	}

	for i := 0; i < len(seeds); i++ {
		for _, resourceMap := range maps {
			mapped := resourceMap.convertNumber(seeds[i])
			seeds[i] = mapped
		}
	}

	return strconv.Itoa(Min(seeds))
}

func (solver Solver) SolveSecond() string {
	// lines, _ := solver.GetPuzzleInput()

	// var seeds []int

	// pairs := regexp.MustCompile(`\d+`).FindAllString(lines[0], -1)

	// for i := 0; i < len(pairs); i = i + 2 {
	// 	start, _ := strconv.Atoi(pairs[i])
	// 	r, _ := strconv.Atoi(pairs[i+1])

	// 	for i := start; i <= start+r; i++ {
	// 		seeds = append(seeds, i)
	// 	}
	// }

	return "N/A"
}

func Min(values []int) int {
	var min int

	for i, value := range values {
		if i == 0 || value < min {
			min = value
		}
	}

	return min
}
