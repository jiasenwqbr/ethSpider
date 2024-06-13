package chain

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ParsingTransaction(txs types.Transactions, client *ethclient.Client) ([]*TransactionMeta, error) {
	txList := make([]*TransactionMeta, 0)
	for i := 0; i < txs.Len(); i++ {
		ctx := context.Background()
		// transaction对象
		tx := txs[i]
		//log.Infof("tx is : %v\n", tx)

		//chainId := tx.ChainId()
		logsList := make([]*LogMeta, 0)

		//msg, err := tx.AsMessage(types.LatestSignerForChainID(chainId), tx.GasPrice())

		//if err != nil {
		//	log.Errorf("Extract data error %s\n", err.Error())
		//}

		txReceipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Errorf("Sync transaction receipt error:%v", err)

			continue
		}

		logs := txReceipt.Logs
		for i := 0; i < len(logs); i++ {
			log := logs[i]
			logObj := LogMeta{
				Data:        hex.EncodeToString(log.Data),
				Topics:      log.Topics,
				TxHash:      log.TxHash,
				BlockNumber: log.BlockNumber,
				BlockHash:   log.BlockHash,
				Address:     log.Address,
				Index:       log.Index,
				Removed:     log.Removed,
				TxIndex:     log.TxIndex,
			}

			logsList = append(logsList, &logObj)
		}

		// 获取签名者对象
		chainId := tx.ChainId()
		signer := types.LatestSignerForChainID(chainId)

		// 获取交易签名者
		from, err := types.Sender(signer, tx)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(msg.To())
		to := ""
		if tx.To() != nil {
			to = tx.To().String()
		}
		messageObj := MessageMeta{
			To:         to,
			From:       from.Hex(),
			Nonce:      tx.Nonce(),
			Amount:     tx.Value().String(),
			GasLimit:   tx.Gas(),
			GasPrice:   tx.GasPrice().String(),
			GasFeeCap:  tx.GasFeeCap().String(),
			GasTipCap:  tx.GasTipCap().String(),
			Data:       hex.EncodeToString(tx.Data()),
			AccessList: tx.AccessList(),
			// IsFake:     tx.IsFake(),
		}

		receiptObj := ReceiptMeta{
			Type:              txReceipt.Type,
			Bloom:             txReceipt.Bloom,
			TxHash:            txReceipt.TxHash,
			GasUsed:           txReceipt.GasUsed,
			BlockHash:         txReceipt.BlockHash,
			BlockNumber:       txReceipt.BlockNumber,
			Status:            txReceipt.Status,
			ContractAddress:   txReceipt.ContractAddress,
			CumulativeGasUsed: txReceipt.CumulativeGasUsed,
			PostState:         hex.EncodeToString(txReceipt.PostState),
			TransactionIndex:  txReceipt.TransactionIndex,
		}

		txObj := TransactionMeta{
			Nonce: tx.Nonce(),
			Value: tx.Value().String(),
			//To:    tx.To(),
			To:         to,
			AccessList: tx.AccessList(),
			Type:       tx.Type(),
			Hash:       tx.Hash().String(),
			Size:       tx.Size(),
			Data:       hex.EncodeToString(tx.Data()),
			GasTipCap:  tx.GasTipCap().String(),
			Gas:        tx.Gas(),
			GasFeeCap:  tx.GasFeeCap().String(),
			GasPrice:   tx.GasPrice().String(),
			ChainId:    tx.ChainId().String(),
			AsMessage:  messageObj,
			Cost:       tx.Cost().String(),
			Protected:  tx.Protected(),
			Logs:       logsList,
			Receipt:    receiptObj,
		}

		//log.Infof("tx.Data() is %s\n", hex.EncodeToString(tx.Data()))

		txList = append(txList, &txObj)
	}

	return txList, nil
}
