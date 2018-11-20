package main

import (
	"Mutex/node"
	"os"
)

func main() {

	Hostname, _ := os.Hostname()
	NodeNum, InterReqDelay, CsExecTime, ReqNum, HostList := parse("config.txt")
	curNode := node.Node{
		Hostname:      Hostname,
		NodeNum:       NodeNum,
		InterReqDelay: InterReqDelay,
		CsExecTime:    CsExecTime,
		ReqNum:        ReqNum,
		HostList:      HostList,
	}
	curNode.Init()
	curNode.Start()
}
