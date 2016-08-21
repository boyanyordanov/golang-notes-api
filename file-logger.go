package main

import (
	"bufio"
	"os"
	"time"
)

// FileLogger: File based logger
type FileLogger struct {
	path string
}

func (log FileLogger) Info(text string) {
	time.Sleep(2 * time.Second)
	file, err := os.OpenFile(log.path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	writer.WriteString("[INFO] " + text + "\n")
	writer.Flush()
}

func (log FileLogger) Error(text string) {
	time.Sleep(2 * time.Second)
	file, err := os.OpenFile(log.path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	writer.WriteString("[ERROR] " + text + "\n")
	writer.Flush()
}
