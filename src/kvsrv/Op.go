package kvsrv

type Op struct {
	// Your definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.
	ClerkId int64
	OpId    int
	OpType  string // "Get", "Put", "Append", "NoOp".
	Key     string
	Value   string
}
