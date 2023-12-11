package tasks

import (
	task_1 "simon-lorenz.dev/advent-of-code/2023/tasks/1"
	task_2 "simon-lorenz.dev/advent-of-code/2023/tasks/2"
)

var solvers []ITaskSolver = make([]ITaskSolver, 0)

func init() {
	RegisterSolver(task_1.Solver{})
	RegisterSolver(task_2.Solver{})
}

func RegisterSolver(solver ITaskSolver) {
	solvers = append(solvers, solver)
}

func GetSolver(puzzle int) ITaskSolver {
	for _, solver := range solvers {
		if solver.GetPuzzleNumber() == puzzle {
			return solver
		}
	}

	return nil
}
