package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type FTPConn struct {
	cwd           string
	control       *net.TCPConn
	controlReader *bufio.Reader
	controlWriter *bufio.Writer
	data          FTPDataSocket
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

func (ftpConn *FTPConn) SendOutOfBandData(data string) {
	bytes := len(data)
	ftpConn.data.Write([]byte(data))
	ftpConn.data.Close()
	message := "Closing data connection, sent " + strconv.Itoa(bytes) + " bytes"
	ftpConn.WriteMessage(getMessageFormat(226), message)
}

func (ftpConn *FTPConn) Close() {
	ftpConn.control.Close()
	if ftpConn.data != nil {
		ftpConn.data.Close()
	}
}
