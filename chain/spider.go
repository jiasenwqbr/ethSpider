package chain

import (
	"context"
	"ethSpider/logger"
	"github.com/ethereum/go-ethereum/ethclient"
	logging "github.com/ipfs/go-log/v2"
)

var log = logger.SetupLog("chain")

type Spider struct {
	Client *ethclient.Client
	Ctx    context.Context
	Logger logging.StandardLogger // 添加Logger对象
}

func New(rpc string) (*Spider, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	return &Spider{
		Client: client,
		Ctx:    context.Background(),
		Logger: logger.SetupLog("scan-block"),
	}, nil
}
