package node

import (
	"Mutex/conn"
	"Mutex/util"
)

type Node struct {
	Hostname      string
	NodeNum       int
	InterReqDelay int
	CsExecTime    int
	ReqNum        int
	HostList      []util.HostInfo

	nodeId int32
	port   string
}

func (n *Node) Init() {
	for _, host := range n.HostList {
		if host.HostName == n.Hostname {
			n.nodeId = host.Id
			n.port = host.Port
		}
	}
	Conn := conn.NewConn(n.port)
	for _, host := range n.HostList {
		if host.Id < n.nodeId {
			Conn.Dial(host.Id, host.HostName, host.Port)
		}
	}
}

func (n *Node) Start() {
	ch := make(chan bool)
	<-ch
}
