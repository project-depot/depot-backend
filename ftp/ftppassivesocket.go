package main

import (
	"errors"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type FTPPassiveSocket struct {
	conn    *net.TCPConn
	port    int
	ingress chan []byte
	egress  chan []byte
}

func (socket *FTPPassiveSocket) Host() string {
	return "0.0.0.0"
}

func (socket *FTPPassiveSocket) Port() int {
	return socket.port
}

func (socket *FTPPassiveSocket) Read(p []byte) (n int, err error) {
	if socket.waitForOpenSocket() == false {
		return 0, errors.New("Data socket unavailable")
	}
	return socket.conn.Read(p)
}

func (socket *FTPPassiveSocket) Write(p []byte) (n int, err error) {
	if socket.waitForOpenSocket() == false {
		return 0, errors.New("Data socket unavailable")
	}
	return socket.conn.Write(p)
}

func (socket *FTPPassiveSocket) Close() error {
	log.Print("Closing passive data socket")
	return socket.conn.Close()
}

func (socket *FTPPassiveSocket) ListenAndServe() {
	laddr, err := net.ResolveTCPAddr("tcp", socket.Host()+":0")
	if err != nil {
		log.Print(err)
		return
	}
	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		log.Print(err)
		return
	}
	add := listener.Addr()
	parts := strings.Split(add.String(), ":")
	port, err := strconv.Atoi(parts[3])
	if err == nil {
		socket.port = port
	}
	tcpConn, err := listener.AcceptTCP()
	if err != nil {
		log.Print(err)
		return
	}
	socket.conn = tcpConn
}

func (socket *FTPPassiveSocket) waitForOpenSocket() bool {
	retries := 0
	for {
		if socket.conn != nil {
			break
		}
		if retries > 3 {
			return false
		}
		log.Print("Sleeping, socket isn't open")
		time.Sleep(500 * time.Millisecond)
		retries++
	}
	return true
}

func NewPassiveSocket() (FTPDataSocket, error) {
	socket := new(FTPPassiveSocket)
	socket.ingress = make(chan []byte)
	socket.egress = make(chan []byte)
	go socket.ListenAndServe()
	for {
		if socket.Port() > 0 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	return socket, nil
}
