package kvraft

import "sync"

type Notifier struct {
	cond sync.Cond
	//maxRegisteredOpId int
}
