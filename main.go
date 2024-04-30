package main

import (
	"flag"
	"fmt"
	"github.com/signmem/falcon-plus/common/redisdb"
	"github.com/signmem/raftleader/g"
	"github.com/signmem/raftleader/selector"
	"github.com/signmem/raftleader/http"
	"os"
)

func init() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		version := g.Version
		fmt.Printf("%s", version)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	redisdb.Server = g.Config().Redis.Server + ":" + g.Config().Redis.Port
	redisdb.MaxIdle = g.Config().Redis.MaxIdle
	redisdb.MaxActive = g.Config().Redis.MaxActive
	redisdb.IdleTimeOut = g.Config().Redis.IdleTimeOut
	redisdb.Pool = redisdb.NewPool(redisdb.MaxIdle, redisdb.MaxActive,
		redisdb.IdleTimeOut, redisdb.Server)
	redisdb.CleanupHook()

}

func main() {

	g.Logger = g.InitLog()
	g.Logger.Info("start ....")


	go selector.Start()
	go selector.RoleCheck()
	http.Start()

	select {}

}