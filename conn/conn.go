package conn

import (
	"Mutex/proto"
	"github.com/golang/glog"
	"github.com/matttproud/golang_protobuf_extensions/pbutil"
	"net"
	"sync"
)

type Conn struct {
	Port string

	id int32
	m  sync.Map
}

func (conn *Conn) Dial(id int32, host string, port string) error {
	c, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		glog.Fatal(err)
		return err
	}
	s := NewSender(c)
	go s.Start()
	conn.m.Store(id, s)
	ini := &mutex_info.MutexInfo{
		Type:          mutex_info.MutexInfo_INI,
		SenderId:      1,
		DestinationId: 2,
		SourceId:      111,
		Timestamp:     123456,
	}
	glog.Infoln(ini.String())
	for {
		s.Send(ini)
	}
	//fmt.Fprintf(c, "Receive a Dial\n")
	//status, err := bufio.NewReader(c).ReadString('\n')
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(status)
	return err
}

func (conn *Conn) Listen() {
	ln, err := net.Listen("tcp", ":"+conn.Port)
	if err != nil {
		glog.Fatal(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			glog.Fatal(err)
		}
		glog.Info("new connection")
		go func() {
			for {
				pb := &mutex_info.MutexInfo{}
				n, err := pbutil.ReadDelimited(c, pb)
				if err != nil {
					glog.Infoln(err.Error())
					continue
				}
				glog.Info(n)
				glog.Info(pb)
			}
		}()
	}
}

func NewConn(port string) *Conn {
	conn := &Conn{
		Port: port,
	}
	go conn.Listen()
	return conn
}
