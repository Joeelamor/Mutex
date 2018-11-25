package main

import (
	"Mutex/node"
	"flag"
	"github.com/golang/glog"
	"os"
)

func main() {

	flag.Parse()
	glog.Infoln("Prepare to repel boarders")
	glog.Flush()
	Hostname, _ := os.Hostname()
	NodeNum, InterReqDelay, CsExecTime, ReqNum, HostList := parse("config1.txt")
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
	//fmt.Println(Hostname)
}
