package raft

import "log"

// the service says it has created a snapshot that has
// all info up to and including index. this means the
// service no longer needs the log through (and including)
// that index. Raft should now trim its log as much as possible.
func (rf *Raft) Snapshot(index int, snapshot []byte) {
	// Your code here (2D).
	rf.mu.Lock()
	log.Printf("savePs get logindex:%d", index)

	// must have commited
	if index > rf.commitIndex || index <= rf.lastIncludedIndex {
		log.Printf("")
		return
	}
	log.Printf("Server %v 上层传入快照, index %v ", rf.me, index)

	lastLog := rf.getLogByIndex(index)
	rf.logs = rf.logs[rf.getRealIdxByIndex(index):]
	rf.lastIncludedIndex = index
	rf.lastIncludedTerm = lastLog.Term
	persistData := rf.getPersistData()
	rf.persister.Save(persistData, snapshot)
}
