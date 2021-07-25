package main

import (
	"flag"
	"go-demo/config/ini"
	acountpb "go-demo/internal/final/bff/acount/proto"
	commentpb "go-demo/internal/final/bff/comment/proto"
	"go-demo/internal/final/bff/handler"
	"go-demo/internal/final/bff/server"
	"go-demo/internal/final/comment/commentconf"
	"go-demo/pkg/log"
	zl "go-demo/pkg/log/zaplogger"
	_ "net/http/pprof"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
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
	conn, err := grpc.Dial("127.0.0.1:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	cclient := commentpb.NewCommentSerClient(conn)

	conn2, err := grpc.Dial("127.0.0.1:50002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	aclient := acountpb.NewAcountSerClient(conn2)

	e := gin.Default()
	handler.Init(e, handler.NewHandler(server.NewBffServer(aclient, cclient)))

	pprof.Register(e)
	err = e.Run(":9090")
	if err != nil {
		log.Fatal(err)
		return
	}
}
