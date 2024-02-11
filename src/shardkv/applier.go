package shardkv

func (kv *ShardKV) execute() {

}

func (kv *ShardKV) pullConfig() {

}

func (kv *ShardKV) pullShards() {
	for {
		select {
		case <-kv.stopCh:
			return
		case <-kv.pullShardsTimer.C:
			isLeader := kv.isLeader()
			if isLeader {
				//kv.lock("pullShards")
				//for shardId, _ := range kv.waitShardIds {
				//	go kv.pullShard(shardId, kv.oldConfig)
				//}
				//kv.unlock("pullShards")
			}
		}
	}
}
