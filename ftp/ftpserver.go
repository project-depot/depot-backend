package main

import (
	"bufio"
	"net"
)

type FTPServer struct {
	name        string
	listener    *net.TCPListener
	connections *Array
}

func (ftpServer *FTPServer) Listen() (err error) {
	for {
		ftpConn, err := ftpServer.Accept()
		if err != nil {
			break
		}
		ftpServer.connections.Append(ftpConn)
		terminated := make(chan bool)
		go ftpConn.Serve(terminated)
		<-terminated
		ftpServer.connections.Remove(ftpConn)
		ftpConn.Close()
	}
	return
}

func (ftpServer *FTPServer) Accept() (ftpConn *FTPConn, err error) {
	tcpConn, err := ftpServer.listener.AcceptTCP()
	if err == nil {
		controlReader := bufio.NewReader(tcpConn)
		controlWriter := bufio.NewWriter(tcpConn)
		ftpConn = &FTPConn{
			rootDir,
			tcpConn,
			controlReader,
			controlWriter,
			nil,
		}
	}
	return
}
