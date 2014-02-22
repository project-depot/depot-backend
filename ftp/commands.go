package main

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
			c.WriteMessage(getMessageFormat(202), "")
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
			c.WriteMessage(getMessageFormat(202), "")
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
			// TODO: Add real functionality
			c.WriteMessage(getMessageFormat(250), `Directory changed to "/"`)
		},
		"PWD": func(c *FTPConn, p []string) {
			// TODO: Add real functionality
			c.WriteMessage(getMessageFormat(257), `"/"`)
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
			passive = true
		},
		"PORT": func(c *FTPConn, p []string) {
			// TODOLATER: Implement active mode
			c.WriteMessage(getMessageFormat(202), "")
		},
		"LIST": func(c *FTPConn, p []string) {
			// TODO: Implement this
		},
		"NLST": func(c *FTPConn, p []string) {
			// TODO: Just the list of file names
			c.WriteMessage(getMessageFormat(202), "")
		},
		"RETR": func(c *FTPConn, p []string) {
			// TODO: Implement this
		},
		"STOR": func(c *FTPConn, p []string) {
			// TODO: Implement this
		},
		"DELE": func(c *FTPConn, p []string) {
			// TODO: Implement this
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
