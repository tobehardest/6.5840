package mr

import (
	"log"
	"sync"
	"time"
)
import "net"
import "os"
import "net/rpc"
import "net/http"

const (
	MaxWaitTime = time.Second * 5
)

type Coordinator struct {
	// Your definitions here.
	files   []string
	nReduce int
	mu      sync.Mutex
	taskCh  chan Task
	tasks   []Task
}

// Your code here -- RPC handlers for the worker to call.

// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

// start a thread that listens for RPCs from worker.go
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.

	return ret
}

// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	// Your code here.
	c.files = files
	c.nReduce = nReduce
	c.mu = sync.Mutex{}
	c.taskCh = make(chan Task, nReduce)
	c.initMapTask()
	go c.run()
	c.server()
	return &c
}

func (c *Coordinator) initMapTask() {
	c.tasks = make([]Task, len(c.files))
}

func (c *Coordinator) run() {

}

func (c *Coordinator) selectTask() {
	if c.Done() {
		return
	}

}
