package raft

import (
	"math/rand"
	"time"
)

const baseElectionTimeout = 300

const None = -1 // to indicate a peer has not voted to anyone at the current term.

// example RequestVote RPC arguments structure.
// field names must start with capital letters!
type RequestVoteArgs struct {
	// Your data here (2A, 2B).
	Term         int
	CandidateId  int
	LastLogIndex int
	LastLogTerm  int
}

// example RequestVote RPC reply structure.
// field names must start with capital letters!
type RequestVoteReply struct {
	// Your data here (2A).
	Term        int
	VoteGranted bool
}

// example RequestVote RPC handler.
func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	// Your code here (2A, 2B).

	reply.Term = rf.currentTerm
	if args.Term < rf.currentTerm {
		return
	} else if args.Term == rf.currentTerm {

	} else {
		// todo fix
	}
}

func (rf *Raft) resetElectionTimer() {
	electionTimeout := baseElectionTimeout + (rand.Int63() % baseElectionTimeout)
	rf.electionTimeout = time.Duration(electionTimeout) * time.Millisecond
	rf.lastElectionTime = time.Now()
}

func (rf *Raft) pastElectionTimeout() bool {
	return time.Since(rf.lastElectionTime) > rf.electionTimeout
}

func (rf *Raft) becomeCandidate() {
	rf.state = Candidate
	rf.currentTerm++
	rf.voteFor = rf.me
}

func (rf *Raft) becomeFollower() {
	rf.state = Follower
}

func (rf *Raft) becomeLeader() {
	rf.state = Leader

}

func (rf *Raft) broadcastRequestVote() {
	for i := range rf.peers {
		if i != rf.me {
			args := rf.makeRequestVoteArgs(i)
			reply := &RequestVoteReply{}
			go rf.sendRequestVote(i, args, reply)
			// todo collect vote
		}
	}
}

func (rf *Raft) makeRequestVoteArgs(to int) *RequestVoteArgs {
	// todo fix log index
	lastLogIndex := len(rf.logs)
	index := rf.getLogByIndex(lastLogIndex)
	args := &RequestVoteArgs{
		Term:         rf.currentTerm,
		CandidateId:  rf.me,
		LastLogIndex: lastLogIndex,
		LastLogTerm:  index.Term,
	}
	return args
}

func (rf *Raft) collectVote() {

}

func (rf *Raft) changeRole(state StateType) {
	rf.state = state
	switch state {
	case Follower:
		// todo reset ticker
	case Candidate:
		rf.currentTerm += 1
		rf.voteFor = rf.me
		//rf.resetElectionTimer()
	case Leader:
		//_, lastLogIndex := rf.lastLogTermIndex()
		//rf.nextIndex = make([]int, len(rf.peers))
		//for i := 0; i < len(rf.peers); i++ {
		//	rf.nextIndex[i] = lastLogIndex + 1
		//}
		//rf.matchIndex = make([]int, len(rf.peers))
		//rf.matchIndex[rf.me] = lastLogIndex
		//rf.resetElectionTimer()
	default:
		panic("unknown role")
	}
}
