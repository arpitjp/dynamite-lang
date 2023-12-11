package logger

import (
	"fmt"
)

const (
	red = "\x1b[31m"
	reset = "\x1b[0m"
	yellow = "\x1b[33m"
	green = "\x1b[32m"
)

func Error(s string) string {
	return fmt.Sprintf("%s%s%s", red, s, reset)
}


func Info(s string) string {
	return fmt.Sprintf("%s%s%s", green, s, reset)
}

func Warn(s string) string {
	return fmt.Sprintf("%s%s%s", yellow, s, reset)
}