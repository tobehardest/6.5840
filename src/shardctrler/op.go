package shardctrler

type Op struct {
	// Your data here.
	ClerkId int
	OpId    int
	OpType  string // "Join", "Leave", "Move", "Query".
	Args    interface{}
}
