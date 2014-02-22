package main

import "io"

type FTPDataSocket interface {
	io.Reader
	io.Writer
	io.Closer

	Host() string
	Port() int
}
