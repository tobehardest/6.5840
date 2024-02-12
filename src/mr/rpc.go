package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import "os"
import "strconv"

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.
// worker request master for task
type TaskArgs struct {
	WorkerId int
}

// master reply worker a task(the task might be nil if no task available)
type TaskReply struct {
	Task       *Task
	ReducerNum int  // the number of reducer, so the mapper can seperate intermediate for different reducer
	Done       bool // true if all task done then the worker will exit, otherwise loop request master for task
}

type ReportTaskArgs struct {
	WorkerId int64
}

type ReportTaskReply struct {
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
