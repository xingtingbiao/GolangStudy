package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	path := "test.txt"
	file, e := readFile(path)
	if e != nil {
		fmt.Printf("Error: %s\n", e)
	}
	fmt.Printf("The content of '%s':\n%s\n", path, file)
}

func readFile(path string) ([]byte, error) {
	dir, err := os.Getwd()
	if nil != err {
		return nil, err
	}
	fullPath := filepath.Join(dir, path)
	file, e := os.Open(fullPath)
	if nil != e {
		return nil, e
	}
	return read(file)
}

func read(file *os.File) ([]byte, error) {
	reader := bufio.NewReader(file)
	var writer bytes.Buffer
	for {
		line, isPrefix, err := reader.ReadLine()
		if nil != err {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		writer.Write(line)
		if !isPrefix {
			writer.WriteByte('\n')
		}
	}
	return writer.Bytes(), nil
}
