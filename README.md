# raft-consensus-algo
an implementation of raft consensus algorithm in golang

## resources
1. [Raft paper](https://raft.github.io/raft.pdf)

### architecture
+ overall, we have N nodes/instances. one of the nodes is elected a leader, while other nodes and read replicas. 
+ we'll only deal with electing a leader in this sample.

### how things work
+ every node knows who the leader is.
+ each of them sends a heartbeat to the master node, in a fixed interval.
+ each node has a different timeout. when they don't receive heartbeat within time.
+ the start an election, and asks every other node for casting vote.
+ each node maintains a value called `term`. it indicates which term of leader is currently.
+ a node can vote to other nodes, but it cannot vote to a node which has term lesser than itself.
+ those nodes which are in election only can be voted, otherwise they vote themselves.
+ once a node, receives a majority of the votes, it is elected as a leader.
+ every node then updates its WAL + `term`.
