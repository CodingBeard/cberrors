package iowriterprovider

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strings"
)

type errorProvider struct {
	writer io.Writer
}

func New(writer io.Writer) *errorProvider {
	return &errorProvider{
		writer: writer,
	}
}

func (p *errorProvider) Error(e error) {
	if p.writer != nil {
		_, _ = p.writer.Write([]byte(fmt.Sprintf("ERROR: %s\n%s", e.Error(), stack())))
	}
}

func (p *errorProvider) Recover(e interface{}) {
	var message string
	err, ok := e.(error)
	if ok {
		message = fmt.Sprintf("PANIC ERROR: %s\n%s", err.Error(), stack())
	} else {
		message = fmt.Sprintf("PANIC: %s\n%s", e, stack())
	}
	if p.writer != nil {
		_, _ = p.writer.Write([]byte(message))
	}
}

func (p *errorProvider) Defer() {}

func stack() string {
	buf := make([]byte, 100000)
	runtime.Stack(buf, false)
	buf = bytes.Trim(buf, "\x00")
	stackParts := strings.Split(string(buf), "\n")
	newStackParts := []string{stackParts[0]}
	newStackParts = append(newStackParts, stackParts[7:]...)
	return strings.Join(newStackParts, "\n")
}
