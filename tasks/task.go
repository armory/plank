package tasks

type ID string
type TaskStatus string

const (
	Succeeded  TaskStatus = "SUCCEEDED"
	NotStarted TaskStatus = "NOT_STARTED"
	Terminal   TaskStatus = "TERMINAL"
)
