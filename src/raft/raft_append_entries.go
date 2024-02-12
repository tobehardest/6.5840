package raft

import "time"

const heartbeatTimeout = 150 * time.Millisecond

type AppendEntriesArgs struct {
	Term         int
	LeaderId     int
	PrevLogIndex int
	PervLogTerm  int
	Entries      []LogEntry
	LeaderCommit int
}

type AppendEntriesReply struct {
	Term      int
	Success   bool
	NextIndex int
}

func (rf *Raft) pastHeartbeatTimeout() bool {
	return time.Since(rf.lastHeartbeatTime) > rf.heartbeatTimeout
}

//func (rf *Raft) StartAppendEntriesTask() {
//	appendEntriesLoop := func(index int) {
//		for {
//			select {
//			case <-rf.stopCh:
//				return
//			case <-rf.appendEntriesTickers[index].C:
//				rf.appendEntriesToPeer(index)
//			}
//		}
//	}
//	for i, _ := range rf.peers {
//		if i == rf.me {
//			continue
//		}
//		appendEntriesLoop(i)
//	}
//}

func (rf *Raft) appendEntriesToPeer(peerIdx int) {

}
