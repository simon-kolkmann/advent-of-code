package tasks

type ITaskSolver interface {
	GetPuzzleNumber() int
	SolveFirst() string
	SolveSecond() string
}
