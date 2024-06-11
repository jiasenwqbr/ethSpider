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
		//fmt.Println(msg.To())
		//to := ""
		//if msg.To() != nil {
		//	to = msg.To().String()
		//}
		//messageObj := MessageMeta{
		//	//To:         msg.To().String(),
		//	To:         to,
		//	From:       msg.From().String(),
		//	Nonce:      msg.Nonce(),
		//	Amount:     msg.Value().String(),
		//	GasLimit:   msg.Gas(),
		//	GasPrice:   msg.GasPrice().String(),
		//	GasFeeCap:  msg.GasFeeCap().String(),
		//	GasTipCap:  msg.GasTipCap().String(),
		//	Data:       hex.EncodeToString(msg.Data()),
		//	AccessList: msg.AccessList(),
		//	IsFake:     msg.IsFake(),
		//}

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
			//To:         tx.To().String(),
			//To:         to,
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
			//AsMessage:  messageObj,
			Cost:      tx.Cost().String(),
			Protected: tx.Protected(),
			Logs:      logsList,
			Receipt:   receiptObj,
		}

		txList = append(txList, &txObj)
	}

	return txList, nil
}
