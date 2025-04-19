package node

type INode interface {
	StartElection()
	RequestVote()
	SendHeartbeat()
}