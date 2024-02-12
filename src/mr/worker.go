package mr

import (
	"errors"
	"fmt"
)
import "log"
import "net/rpc"
import "hash/fnv"

// Map functions return a slice of KeyValue.
type KeyValue struct {
	Key   string
	Value string
}

// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

// main/mrworker.go calls this function.
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {

	// Your worker implementation here.

	// uncomment to send the Example RPC to the coordinator.
	// CallExample()
	w := MrWorker{
		Mapf:    mapf,
		Reducef: reducef,
	}
	//w.register()
	w.run()
}

type MrWorker struct {
	WorkerId int64
	Mapf     func(string, string) []KeyValue
	Reducef  func(string, []string) string
	Done     bool
}

//func (w *worker) register() {
//	args := &RegisterArgs{}
//	reply := &RegisterReply{}
//	if ok := call("Master.RegWorker", args, reply); !ok {
//		log.Fatal("reg fail")
//	}
//	w.workId = reply.WorkerId
//}

func (w *MrWorker) run() {
	for !w.Done {
		task, err := w.reqTask()
		if err != nil {
			log.Println("work.reqTask err = ", err)
			// todo : sleep some time to retry
			continue
		}

		log.Printf("Worker with WorkerId = %d received task = %v", w.WorkerId, task)
		w.doTask(task)
	}
}

func (w *MrWorker) reqTask() (*Task, error) {
	args := TaskArgs{}
	reply := TaskReply{}
	ok := call("Coordinator.AssignTask", &args, &reply)
	if !ok {
		return nil, errors.New("req task fail")
	}
	if reply.Done {
		w.Done = true
		return nil, nil
	}
	return reply.Task, nil
}

func (w *MrWorker) doTask(t *Task) {

	switch t.TaskType {
	case Map:
		w.doMapTask(t)
	case Reduce:
		w.doReduceTask(t)
	default:
	}
}

func (w *MrWorker) doMapTask(t *Task) {

}

func (w *MrWorker) doReduceTask(t *Task) {

}

func (w *MrWorker) reportTask(t *Task, done bool, err error) {
	args := ReportTaskArgs{}
	reply := ReportTaskReply{}
	ok := call("Coordinator.ReportTask", &args, &reply)
	if !ok {
		log.Printf("report task fail:%+v", args)
	}
}

// example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
func CallExample() {

	// declare an argument structure.
	args := ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := ExampleReply{}

	// send the RPC request, wait for the reply.
	// the "Coordinator.Example" tells the
	// receiving server that we'd like to call
	// the Example() method of struct Coordinator.
	ok := call("Coordinator.Example", &args, &reply)
	if ok {
		// reply.Y should be 100.
		fmt.Printf("reply.Y %v\n", reply.Y)
	} else {
		fmt.Printf("call failed!\n")
	}
}

// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
