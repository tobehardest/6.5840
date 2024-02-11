package kvraft

import "log"

func (kv *KVServer) execute() {
	for {
		select {
		case <-kv.stopCh:
			log.Printf("stop ch get")
			return
		case apply := <-kv.applyCh:
			op := apply.Command.(*Op)
			kv.ApplyClientOp(op)
		}
	}
}

func (kv *KVServer) ApplyClientOp(op *Op) {
	kv.applyClientOp(op)
}

func (kv *KVServer) applyClientOp(op *Op) {

	switch op.OpType {
	case "Get":
		// only write ops are applied to the database.
	case "Put":
		kv.db[op.Key] = op.Value
	case "Append":
		// note: the default value is returned if the key does not exist.
		kv.db[op.Key] += op.Value

	default:
		log.Fatalf("unexpected client op type %v", op.OpType)
	}
}

func (kv *KVServer) waitUntilAppliedOrTimeout(op *Op) (Err, string) {
	if !kv.propose(op) {
		return ErrWrongLeader, ""
	}

	return "", ""
}
