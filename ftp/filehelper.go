package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/jehiah/go-strftime"
)

func FileString(file os.FileInfo) (fileString string) {
	fileString += file.Mode().String()
	fileString += " 1 owner group "

	fileSize := strconv.Itoa(int(file.Size()))
	paddingAmount := 12
	var paddedFileSize string
	if len(fileSize) < paddingAmount {
		paddedFileSize = strings.Repeat(" ", paddingAmount-len(fileSize)) + fileSize
	} else if len(fileSize) == paddingAmount {
		paddedFileSize = fileSize
	} else {
		paddedFileSize = fileSize[0:paddingAmount]
	}

	fileString += paddedFileSize
	fileString += " " + strftime.Format("%b %d %H:%M", file.ModTime())
	fileString += " " + file.Name()

	return fileString
}
