package tasks

import (
	task_1 "simon-lorenz.dev/advent-of-code/2023/tasks/1"
	task_2 "simon-lorenz.dev/advent-of-code/2023/tasks/2"
	task_3 "simon-lorenz.dev/advent-of-code/2023/tasks/3"
	task_4 "simon-lorenz.dev/advent-of-code/2023/tasks/4"
	task_5 "simon-lorenz.dev/advent-of-code/2023/tasks/5"
	task_6 "simon-lorenz.dev/advent-of-code/2023/tasks/6"
	task_7 "simon-lorenz.dev/advent-of-code/2023/tasks/7"
)

var solvers []ITaskSolver = make([]ITaskSolver, 0)

func init() {
	RegisterSolver(task_1.NewSolver())
	RegisterSolver(task_2.NewSolver())
	RegisterSolver(task_3.NewSolver())
	RegisterSolver(task_4.NewSolver())
	RegisterSolver(task_5.NewSolver())
	RegisterSolver(task_6.NewSolver())
	RegisterSolver(task_7.NewSolver())
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
