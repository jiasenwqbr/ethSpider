package chain

import (
	"ethSpider/conf"
	"time"
)

func (b *Spider) Sync() error {
	var (
		currentBlock uint64
	)

	client := b.Client
	ctx := b.Ctx
	logger := b.Logger
	number := uint64(20067610)

	for {
		maxNumber, err := client.BlockNumber(ctx)
		logger.Infof("the number is : %d", number, ",The currentBlock is : %d", currentBlock)
		if err != nil {
			logger.Errorf("Get latest block number error:%s", err)

			return err
		}

		logger.Infof("the number is : %d", number, ",The currentBlock is : %d", currentBlock)
		// 区块去重监测
		if number >= maxNumber {
			logger.Infof("Repeat block[%d], Nothing to do", number)

			time.Sleep(conf.SYNC_SLEEP_DELAY_TIME)

			continue
		}

		logger.Infof("Now block number:%d", number)
		logger.Infof("Client:%v", client)
		/*
			解析区块数据
		*/
		// 第一步提取区块消息
		blockData, err := ParsingBlock(ctx, client, number)
		if err != nil {
			logger.Errorf("Parse block meta data error:%s", err)
		}

		logger.Infof("Block Data:%v", blockData)

		//// 解析区块header
		//err = parse.ParsingBlockHeader(ctx, client, number)
		//if err != nil {
		//	logger.Errorf("Parse block header error:%s", err)
		//}

		//// 解析区块body
		//err = parse.ParsingBlockBody(ctx, client, number)
		//if err != nil {
		//	logger.Errorf("Parse block body error:%s", err)
		//}

		// 记录已同步区块高度
		//currentBlock = number
		// 延迟同步
		// time.Sleep(conf.SYNC_ONE_SLEEP_DELAY_TIME)
		number++
		logger.Infof("number:%v", number)
	}
}
