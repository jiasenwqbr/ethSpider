package node

import (
	"ethSpider/chain"
	"ethSpider/conf"
	"ethSpider/logger"
	"github.com/urfave/cli/v2"
)

var log = logger.SetupLog("scan-block")

var NodeCmd = &cli.Command{
	Name:  "node",
	Usage: "Commands for remotely taking node related actions",
	Subcommands: []*cli.Command{
		startCmd,
	},
}

var startCmd = &cli.Command{
	Name:  "start",
	Usage: "Run eth-spider services",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "rpc-url",
			Usage: "Ethereum rpc url",
			Value: "",
		},
	},
	Action: func(cctx *cli.Context) error {
		rpcUrl := cctx.String("rpc-url")
		if rpcUrl == "" {
			log.Warnf("RPC parameter is empty, default value used")
			rpcUrl = conf.WSS_ETH_RPC
		}

		// 初始化spider对象
		spider, err := chain.New(rpcUrl)
		if err != nil {
			return err
		}

		// ETH client
		err = spider.Sync()
		if err != nil {
			return err
		}

		return nil

	},
}
