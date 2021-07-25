package main

import (
	"flag"
	"go-demo/config/ini"
	"go-demo/internal/final/account/acountconf"
	acountmy "go-demo/internal/final/account/dao/mysql"
	"go-demo/internal/final/account/handler"
	acountpb "go-demo/internal/final/account/proto/acount"
	"go-demo/internal/final/account/server"
	"go-demo/pkg/log"
	zl "go-demo/pkg/log/zaplogger"
	"go-demo/pkg/mysql"
	"net"

	"google.golang.org/grpc"
)

var (
	iniPath = flag.String("iniPath", "../../../etc/acount.ini", "")
)

func main() {
	flag.Parse()

	parse := ini.NewConfig()

	conf := new(acountconf.AcountConf)
	err := parse.Load(*iniPath, conf)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Init(zl.NewLog(conf.LogConf))

	engine, err := acountmy.NewAngInitMyEngine(mysql.WithConf(&conf.MyEngine))
	if err != nil {
		log.Fatal(err)
		return
	}

	server := handler.NewPBServer(server.NewAcountServer(engine))

	gs := grpc.NewServer()

	acountpb.RegisterAcountSerServer(gs, server)

	listen, err := net.Listen("tcp", ":"+conf.ServerPort)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = gs.Serve(listen)
	if err != nil {
		log.Fatal(err)
		return
	}
}
