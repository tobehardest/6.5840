package kvraft

import (
	"log"
)

func (kv *KVServer) propose(op *Op) bool {
	index, term, isLeader := kv.rf.Start(op)
	if !isLeader {
		return false
	}
	log.Printf("start cmd: index:%d, term:%d, op:%+v", index, term, op)
	return true
}

func (kv *KVServer) isLeader() bool {
	_, isLeader := kv.rf.GetState()
	return isLeader
}
