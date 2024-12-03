package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"simon-lorenz.dev/advent-of-code/2024/tasks"
)

func main() {
	puzzle := getRequestedPuzzleOrExit()
	solver := tasks.GetSolver(puzzle)

	if solver != nil {
		solveAndPrint(solver)
	} else {
		fmt.Println("The requested puzzle (" + strconv.Itoa(puzzle) + ") has no solution (yet).")
	}
}

func solveAndPrint(solver tasks.ITaskSolver) {
	fmt.Println("Solving Task " + strconv.Itoa(solver.GetPuzzleNumber()) + ":")
	fmt.Println("First: " + solver.SolveFirst())
	fmt.Println("Second: " + solver.SolveSecond())
}

func getRequestedPuzzleOrExit() int {
	var puzzle int

	flag.IntVar(&puzzle, "p", 0, "which puzzle to solve (shorthand)")
	flag.IntVar(&puzzle, "puzzle", 0, "which puzzle to solve")
	flag.Parse()

	if puzzle < 1 || puzzle > 24 {
		fmt.Println("You didn't specify a valid --puzzle (1-24) to solve.")
		os.Exit(1)
	}

	return puzzle
}
