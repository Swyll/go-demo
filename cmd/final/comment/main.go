package main

import (
	"flag"
	"go-demo/config/ini"
	"go-demo/internal/final/comment/commentconf"
	commentmy "go-demo/internal/final/comment/dao/mysql"
	"go-demo/internal/final/comment/handler"
	commentpb "go-demo/internal/final/comment/proto/comment"
	"go-demo/internal/final/comment/server"
	"go-demo/pkg/log"
	zl "go-demo/pkg/log/zaplogger"
	"go-demo/pkg/mysql"
	"net"

	"google.golang.org/grpc"
)

var (
	iniPath = flag.String("iniPath", "../../../etc/comment.ini", "")
)

func main() {
	flag.Parse()

	parse := ini.NewConfig()

	conf := new(commentconf.CommentConf)
	err := parse.Load(*iniPath, conf)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Init(zl.NewLog(conf.LogConf))

	engine, err := commentmy.NewAngInitMyEngine(mysql.WithConf(&conf.MyEngine))
	if err != nil {
		log.Fatal(err)
		return
	}

	server := handler.NewPBServer(server.NewCommentServer(engine))

	gs := grpc.NewServer()

	commentpb.RegisterCommentSerServer(gs, server)

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
