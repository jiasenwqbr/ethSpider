package chain

import (
	"encoding/json"
	"ethSpider/conf"
	"fmt"
	"os"
	"time"
)

func (b *Spider) Sync() error {
	//var (
	//	currentBlock uint64
	//)

	client := b.Client
	ctx := b.Ctx
	logger := b.Logger
	number := uint64(20081713)

	for {
		maxNumber, err := client.BlockNumber(ctx)
		logger.Infof("the number is : %d\n", number)
		if err != nil {
			logger.Errorf("Get latest block number error:%s", err)

			return err
		}

		// logger.Infof("the number is : %d", number, ",The currentBlock is : %d", currentBlock)
		// 区块去重监测
		if number >= maxNumber {
			logger.Infof("Repeat block[%d], Nothing to do", number)

			time.Sleep(conf.SYNC_SLEEP_DELAY_TIME)

			continue
		}

		//logger.Infof("Now block number:%d", number)
		//logger.Infof("Client:%v", client)
		/*
			解析区块数据
		*/
		// 第一步提取区块消息
		blockData, err := ParsingBlock(ctx, client, number)
		if err != nil {
			logger.Errorf("Parse block meta data error:%s", err)
		}

		jsonData, err := json.Marshal(blockData)
		// logger.Infof("Block Data:%v\n", jsonData)
		// logger.Info(jsonData)
		fmt.Println(string(jsonData))
		if err != nil {
			log.Fatalf("Error marshalling to JSON: %v", err)
		}
		file, err := os.Create("records.json")
		if err != nil {
			log.Fatalf("Failed to create file: %s", err)
		}
		defer file.Close()
		// 写入 JSON 到文件
		_, err = file.Write(jsonData)
		if err != nil {
			log.Fatalf("Failed to write to file: %s", err)
		}

		fmt.Println("Records written to file successfully.")

		// 延迟同步
		// time.Sleep(conf.SYNC_ONE_SLEEP_DELAY_TIME)
		number++

	}
}
