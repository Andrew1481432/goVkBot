package log

import (
	"github.com/fatih/color"
	"strings"
	"time"
)

const (
	ANSI_RESET  = "\u001B[0m"
	ANSI_BLACK  = "\u001B[30m"
	ANSI_RED    = "\u001B[31m"
	ANSI_GREEN  = "\u001B[32m"
	ANSI_YELLOW = "\u001B[33m"
	ANSI_BLUE   = "\u001B[34m"
	ANSI_PURPLE = "\u001B[35m"
	ANSI_CYAN   = "\u001B[36m"
	ANSI_WHITE  = "\u001B[37m"
)

func Create(format string) (l *Log) {
	l = &Log{}
	l.init(format)

	return
}

type Log struct {
	format string
}

func (l *Log) prepare(data ...string) (string, string, string) {
	return l.format, l.currentTime(), strings.Join(data, " ")
}

func (l *Log) init(format string) {
	l.format = format
}

func (l *Log) currentTime() string {
	return time.Now().Format("2006.01.02 15:04:05")
}

func (l *Log) Info(data ...string) {
	color.Cyan(l.prepare(data...))
}

func (l *Log) Log(data ...string) {
	color.White(l.prepare(data...))
}

func (l *Log) Warning(data ...string) {
	color.Yellow(l.prepare(data...))
}

func (l *Log) Error(data ...string) {
	color.Red(l.prepare(data...))
}
