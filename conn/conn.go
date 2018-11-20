package conn

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

type Conn struct {
	Port string
}

func (conn *Conn) Dial(host string, port string) error {
	c, err := net.Dial("tcp", host + ":" + port)
	if err != nil {
		// handle error
	}
	fmt.Fprintf(c, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)
}

func (conn *Conn) Listen() {
	ln, err := net.Listen("tcp", ":" + strconv.Itoa(conn.Port))
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
		}()
	}
}

func NewConn(port string) *Conn {
	conn := &Conn{
		Port:port,
	}
	go conn.Listen()
	return conn
}
