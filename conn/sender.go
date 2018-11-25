package conn

import (
	"Mutex/proto"
	"github.com/golang/glog"
	"github.com/matttproud/golang_protobuf_extensions/pbutil"
	"net"
)

type sender struct {
	outbox chan *mutex_info.MutexInfo
	stop   chan bool
	socket net.Conn
}

func NewSender(socket net.Conn) *sender {
	return &sender{
		outbox: make(chan *mutex_info.MutexInfo, 100),
		stop:   make(chan bool),
		socket: socket,
	}
}

func (s *sender) Send(msg *mutex_info.MutexInfo) {
	s.outbox <- msg
}

func (s *sender) Start() {
	for {
		select {
		case msg := <-s.outbox:
			glog.Info("new message to send")
			glog.Info(msg)
			n, err := pbutil.WriteDelimited(s.socket, msg)
			if err != nil {
				glog.Error(err.Error())
				continue
			}
			glog.Info("new message sent")
			glog.Info(n)
		case <-s.stop:
			return
		}
	}
}

func (s *sender) Stop() {
	s.stop <- true
}
