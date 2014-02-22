package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var commands map[string]func(ftpConn *FTPConn, p []string)

func initializeCommands() {
	commands = map[string]func(c *FTPConn, p []string){
		"ABOR": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"ACCT": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"ADAT": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"ALLO": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "Deprecated")
		},
		"APPE": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"AUTH": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"CCC": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"CONF": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"ENC": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"EPRT": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"EPSV": func(c *FTPConn, p []string) {
			socket, err := NewPassiveSocket()
			if err != nil {
				c.WriteMessage(getMessageFormat(425), "Data connection failed")
				return
			}
			c.data = socket
			msg := fmt.Sprintf("Entering Extended Passive Mode (|||%d|)", socket.Port())
			c.WriteMessage(getMessageFormat(229), msg)
		},
		"HELP": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"LANG": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"LPRT": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"LPSV": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"MDTM": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"MIC": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"MKD": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"MLSD": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"MLST": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"MODE": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"NOOP": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"OPTS": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"REIN": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"STOU": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"STRU": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"PBSZ": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"SITE": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"SMNT": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"RMD": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"STAT": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(202), "")
		},
		"FEAT": func(c *FTPConn, p []string) {
			c.WriteMessage("211-Extensions supported\r\n")
			c.WriteMessage(getMessageFormat(211), "End")
		},
		"SYST": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(215), "Go FTP awesome server")
		},
		"CWD": func(c *FTPConn, p []string) {
			// TODO: Make sure the specified directory is valid
			c.cwd = p[1]
			c.WriteMessage(getMessageFormat(250), `Directory changed to "`+c.cwd+`"`)
		},
		"PWD": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(257), `"`+c.cwd+`"`)
		},
		"TYPE": func(c *FTPConn, p []string) {
			dataEncodingParam := p[1]
			if dataEncodingParam == "A" || dataEncodingParam == "I" {
				if dataEncodingParam == "A" {
					dataEncoding = asciiEncoding
				} else {
					asciiEncoding = "binary"
				}
				c.WriteMessage(getMessageFormat(200), "")
			} else {
				c.WriteMessage(getMessageFormat(501), "")
			}
		},
		"USER": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(331), "User name OK, password required")
		},
		"PASS": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(230), "Password good to go, continue")
		},
		"PASV": func(c *FTPConn, p []string) {
			socket, err := NewPassiveSocket()
			if err != nil {
				c.WriteMessage(getMessageFormat(425), "Data connection failed")
				return
			}
			c.data = socket
			p1 := socket.Port() / 256
			p2 := socket.Port() - (p1 * 256)

			quads := strings.Split(socket.Host(), ".")
			target := fmt.Sprintf("(%s,%s,%s,%s,%d,%d)", quads[0], quads[1], quads[2], quads[3], p1, p2)
			msg := "Entering Passive Mode " + target
			c.WriteMessage(getMessageFormat(227), msg)
		},
		"PORT": func(c *FTPConn, p []string) {
			// TODOLATER: Implement active mode
			c.WriteMessage(getMessageFormat(202), "")
		},
		"LIST": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(150), "Opening ASCII mode data connection for file list")
			var output string
			files, _ := ioutil.ReadDir(c.cwd)
			for _, f := range files {
				output += FileString(f)
				output += "\r\n"
			}
			c.SendOutOfBandData(output)
		},
		"NLST": func(c *FTPConn, p []string) {
			// TODO: Just the list of file names
			c.WriteMessage(getMessageFormat(202), "")
		},
		"RETR": func(c *FTPConn, p []string) {
			path := p[1]
			bytes, err := ioutil.ReadFile(path)
			if err != nil {
				c.WriteMessage(getMessageFormat(551), "File not available.")
				return
			}

			c.WriteMessage(getMessageFormat(150), "Data transfer starting "+strconv.Itoa(len(bytes))+" bytes")
			c.SendOutOfBandData(string(bytes))
		},
		"STOR": func(c *FTPConn, p []string) {
			path := p[1]
			transferError := "Error during transfer"

			fmt.Println("it werks")

			c.WriteMessage(getMessageFormat(150), "Beginning data transfer...")
			tmpFile, err := ioutil.TempFile("", "stor")
			if err != nil {
				c.WriteMessage(getMessageFormat(450), transferError)
				return
			}
			bytes, err := io.Copy(tmpFile, c.data)
			if err != nil {
				c.WriteMessage(getMessageFormat(450), transferError)
			}
			tmpFile.Seek(0, 0)

			if err := os.Rename(tmpFile.Name(), path); err != nil {
				c.WriteMessage(getMessageFormat(550), "Action not taken. "+err.Error())
				return
			}

			c.WriteMessage(getMessageFormat(226), "All is well, received "+strconv.Itoa(int(bytes))+" bytes")
		},
		"DELE": func(c *FTPConn, p []string) {
			if err := os.Remove(p[1]); err != nil {
				c.WriteMessage(getMessageFormat(550), "Action not taken. "+err.Error())
				return
			}
			c.WriteMessage(getMessageFormat(250), "File deleted")
		},
		"RNFR": func(c *FTPConn, p []string) {
			// TODO: Rename from
			c.WriteMessage(getMessageFormat(202), "")
		},
		"RNTO": func(c *FTPConn, p []string) {
			// TODO: Rename to
			c.WriteMessage(getMessageFormat(202), "")
		},
		"REST": func(c *FTPConn, p []string) {
			// Restart the transfer from the specified point
			c.WriteMessage(getMessageFormat(202), "")
		},
		"QUIT": func(c *FTPConn, p []string) {
			c.WriteMessage(getMessageFormat(221), "")
			c.Close()
		},
	}
}
