package log

import "fmt"

type BaseLog interface {
	Warn(msg any)
	Error(msg any)
	Info(msg any)
}

type Log struct{}

func (l Log) Warn(msg ...any) {
	fmt.Printf("[warn] %v \n", msg...)
}
func (l Log) Error(msg ...any) {
	fmt.Printf("[error] %v \n", msg...)
}
func (l Log) Info(msg ...any) {
	fmt.Printf("[info] %v \n", msg...)
}

var Logger = new(Log)
