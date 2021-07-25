package acountconf

import (
	z1 "go-demo/pkg/log/zaplogger"
	"go-demo/pkg/mysql"
)

type AcountConf struct {
	mysql.MyEngine
	ServerConf
	LogConf *z1.LoggerConf `ini:"log"`
}

type ServerConf struct {
	ServerName string `ini:"server_name"`
	ServerPort string `ini:"port"`
}
