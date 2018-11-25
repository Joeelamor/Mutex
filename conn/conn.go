package conn

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Conn struct {
	Port string
}

func (conn *Conn) Dial(host string, port string) error {
	c, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(c, "Receive a Dial\n")
	status, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)
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
			status, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(status)
			fmt.Fprintf(c, "Dial success\n")
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
