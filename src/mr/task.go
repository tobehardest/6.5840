package mr

import "time"

type TaskType string
type TaskStatus string

const (
	// Define task type
	Map    TaskType = "MapTask"
	Reduce TaskType = "ReduceTask"

	// Define task expired time
	TaskExpiredTime = 10

	// Define task status
	Ready    TaskStatus = "Ready"
	Running  TaskStatus = "Running"
	Finished TaskStatus = "Finished"
)

type Task struct {
	TaskType       TaskType   // the task type, map or reduce
	MapWorkerId    int        // worker id if in map phase, given my master
	ReduceWorkerId int        // worker id if in reduce phase, given my master
	InputFile      string     // if in map phase it should be a single file, if in reduce phase it should be a file pattern
	BeginTime      time.Time  // task begin time, given by master
	TaskStatus     TaskStatus // task status, ready, running or finished, for worker it should always be running
}
