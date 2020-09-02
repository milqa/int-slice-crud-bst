package bst

import (
	"fmt"
	"time"
)

type Printf func(format string, args ...interface{})

type PrintfLogger struct {
	printf Printf
}

func NewPrintfLogger(printf Printf) *PrintfLogger {
	return &PrintfLogger{printf}
}

func (pl *PrintfLogger) beforeFind(address int) {

	pl.printf(">>> FIND VALUE FOR ADDRESS = %v", address)
}

func (pl *PrintfLogger) beforeDelete(address int) {

	pl.printf(">>> DELETE VALUE FOR ADDRESS = %v", address)
}

func (pl *PrintfLogger) beforeInsert(address int, value interface{}) {

	pl.printf(">>> INSERT VALUE [%s] FOR ADDRESS = %v", value, address)
}

func (pl *PrintfLogger) afterFind(address int, value interface{}, d time.Duration, err error) {

	msg := fmt.Sprintf("FIND VALUE [%s] FOR ADDRESS = %d | %s", value, address, d)
	if err != nil {
		msg += ": " + err.Error()
	}
	pl.printf("<<< %s", msg)
}

func (pl *PrintfLogger) afterDelete(address int, d time.Duration, err error) {

	msg := fmt.Sprintf("DELETE VALUE FOR ADDRESS = %d | %s", address, d)
	if err != nil {
		msg += ": " + err.Error()
	}
	pl.printf("<<< %s", msg)
}

func (pl *PrintfLogger) afterInsert(address int, value interface{}, d time.Duration, err error) {

	msg := fmt.Sprintf("INSERT VALUE [%s] FOR ADDRESS = %d | %s", value, address, d)
	if err != nil {
		msg += ": " + err.Error()
	}
	pl.printf("<<< %s", msg)
}
