package node

type node struct {
	// unique identifier for the node
	id uint32
	// the current term of the master
	term uint32
	// timeout period before starting an election 
	// if it doesn't receive a heartbeat
	timeout uint32
	// if the given node is a master or not
	is_master bool

	// connected client connections
	connected_clients map[uint32]Connection
}

func NewNode() INode {
	return &node{
		id: 
	}
}