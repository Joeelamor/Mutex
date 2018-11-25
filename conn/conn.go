package conn

import (
	"Mutex/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
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
		log.Fatal(err)
	}
	s := NewSender(c)
	go s.Start()
	conn.m.Store(id, s)
	ini := &mutex_info.MutexInfo{
		Type:          mutex_info.MutexInfo_INI,
		SenderId:      conn.id,
		DestinationId: id,
	}
	s.Send(ini)
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
		log.Fatal(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			pb := &mutex_info.MutexInfo{}
			b := make([]byte, 2048)
			n, _ := c.Read(b)
			err := proto.Unmarshal(b[:n], pb)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(pb)
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
