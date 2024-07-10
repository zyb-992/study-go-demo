package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func newBuilder() strings.Builder {
	var b = strings.Builder{}
	return b
}

const writingData = `1. hello, people before window
2. using strings.builder to make writing data to file efficiently`

func write2File(f *os.File, s string) error {
	data := s
	if len(data) == 0 {
		data = writingData
	}

	builder := newBuilder()
	_, err := builder.WriteString(data)
	if err != nil {
		return fmt.Errorf("writing to buf err:%v", err)
	}

	_, err = f.WriteString(builder.String())
	if err != nil {
		return fmt.Errorf("writing to disk err:%v", err)
	}

	return nil
}

func createFile(name string) (*os.File, error) {
	if _, err := os.Stat(name); errors.Is(err, os.ErrExist) {
		return nil, fmt.Errorf("create file err:%v", err)
	}

	f, err := os.Create(name)
	if err != nil {
		return nil, fmt.Errorf("create file err:%v", err)
	}

	return f, nil
}
