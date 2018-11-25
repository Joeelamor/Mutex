package conn

import (
	"Mutex/proto"
	"github.com/golang/protobuf/proto"
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
			b, err := proto.Marshal(msg)
			if err != nil {
				// handle error
				continue
			}
			_, err = s.socket.Write(b)
			if err != nil {
				continue
			}
		case <-s.stop:
			return
		}
	}
}

func (s *sender) Stop() {
	s.stop <- true
}
