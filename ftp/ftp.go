package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	rootDir        = "/"
	welcomeMessage = "Welcome to the Go FTP Server"
	USER           = "USER"
	PASS           = "PASS"
	PWD            = "PWD"
	STAT           = "STAT"
	CWD            = "CWD"
	FEAT           = "FEAT"
)

var (
	dataEncoding  string = "A"
	asciiEncoding string = "binary"
)

func getMessageFormat(command int) (messageFormat string) {
	output := fmt.Sprintf("%d %%s\r\n", command)
	return output
}

type Array struct {
	container []interface{}
}

func (a *Array) Append(object interface{}) {
	if a.container == nil {
		a.container = make([]interface{}, 0)
	}
	newContainer := make([]interface{}, len(a.container)+1)
	copy(newContainer, a.container)
	newContainer[len(newContainer)-1] = object
	a.container = newContainer
}

func (a *Array) Remove(object interface{}) (result bool) {
	result = false
	newContainer := make([]interface{}, len(a.container)-1)
	i := 0
	for _, target := range a.container {
		if target != object {
			newContainer[i] = target
		} else {
			result = true
		}
		i++
	}
	return
}

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

type FTPConn struct {
	cwd           string
	control       *net.TCPConn
	controlReader *bufio.Reader
	controlWriter *bufio.Writer
	data          *net.TCPConn
}

func (ftpConn *FTPConn) WriteMessage(messageFormat string, v ...interface{}) (wrote int, err error) {
	message := fmt.Sprintf(messageFormat, v...)
	wrote, err = ftpConn.controlWriter.WriteString(message)
	ftpConn.controlWriter.Flush()
	log.Print(message)
	return
}

func (ftpConn *FTPConn) Serve(terminated chan bool) {
	log.Print("Connection Established")
	// send welcome
	ftpConn.WriteMessage(getMessageFormat(220), welcomeMessage)
	// read commands
	for {
		line, err := ftpConn.controlReader.ReadString('\n')
		if err != nil {
			break
		}
		log.Print(line)
		params := strings.Split(strings.Trim(line, "\r\n"), " ")
		if len(params) > 0 {
			command := params[0]
			commandFunc := commands[command]
			if commandFunc != nil {
				commandFunc(ftpConn, params)
			} else {
				ftpConn.WriteMessage(getMessageFormat(500), "Command not found")
			}
		}
	}
	terminated <- true
	log.Print("Connection Terminated")
}

func (ftpConn *FTPConn) Close() {
	ftpConn.control.Close()
	if ftpConn.data != nil {
		ftpConn.data.Close()
	}
}

func main() {
	initializeCommands()
	laddr, err := net.ResolveTCPAddr("tcp", "localhost:2021")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp4", laddr)
	if err != nil {
		log.Fatal(err)
	}
	ftpServer := &FTPServer{
		"Go FTP Server",
		listener,
		new(Array),
	}
	log.Fatal(ftpServer.Listen())
}
