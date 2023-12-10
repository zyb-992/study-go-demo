package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func newReader() strings.Reader {
	var r = strings.Reader{}
	return r
}

func readFrom(r *os.File) (string, error) {
	var data = make([]byte, 1024)
	_, err := r.Read(data)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("read err:%v", err)
	}

	return string(data), nil
}
