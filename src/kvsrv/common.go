package kvsrv

const (
	OK     = "OK"
	ErrNot = "Err"
	//ErrNoKey       = "ErrNoKey"
	//ErrWrongLeader = "ErrWrongLeader"
	//ErrTimeOut     = "ErrTimeOut"
)

type Err string

// Put or Append
type PutAppendArgs struct {
	Key   string
	Value string
	// You'll have to add definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.
	OpId   int
	OpType string // "Put" or "Append"
}

type PutAppendReply struct {
	Err Err
}

type GetArgs struct {
	Key string
	// You'll have to add definitions here.
	OpId int
}

type GetReply struct {
	Err   Err
	Value string
}
