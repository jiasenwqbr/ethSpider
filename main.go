package main

import (
	"ethSpider/cmd/admin"
	"ethSpider/cmd/node"
	"ethSpider/logger"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
)

const VERSION = "0.0.1"

var log = logger.SetupLog("main-process")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := []*cli.Command{
		admin.AdminCmd,
		node.NodeCmd,
		
	}

	app := &cli.App{
		Name:     "ETH-Spider",
		Usage:    "Sync Ethereum node block data",
		Version:  VERSION,
		Commands: c,
	}

	if err := app.Run(os.Args); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
