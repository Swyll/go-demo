package server

import "fmt"

var log LogI

type LogI interface {
	Errorf(string, ...interface{})
	Warnf(string, ...interface{})
	Infof(string, ...interface{})
}

func init() {
	log = new(Log)
}

func GetLogger() LogI {
	return log
}

type Log struct{}

func (l *Log) Errorf(msg string, params ...interface{}) {
	fmt.Printf(msg, params...)
}

func (l *Log) Warnf(msg string, params ...interface{}) {
	fmt.Printf(msg, params...)
}

func (l *Log) Infof(msg string, params ...interface{}) {
	fmt.Printf(msg, params...)
}
