package shardctrler

func (sc *ShardCtrler) execute() {
	for {
		select {
		case <-sc.stopCh:
			return
		case msg := <-sc.applyCh:
			if !msg.CommandValid {
				continue
			}
			op := msg.Command.(Op)
			switch op.OpType {
			//case "Join":
			//	sc.Join(op.Args.(JoinArgs))
			//case "Leave":
			//	sc.leave(op.Args.(LeaveArgs))
			//case "Move":
			//	sc.move(op.Args.(MoveArgs))
			case "Query":
			default:
				panic("unknown method")
			}
		}
	}
}
