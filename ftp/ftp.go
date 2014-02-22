package main

import (
	"fmt"
	"log"
	"net"
)

const (
	rootDir        = "/"
	welcomeMessage = "Welcome to Depot."
)

var (
	dataEncoding  string = "A"
	asciiEncoding string = "binary"
	passive       bool
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
