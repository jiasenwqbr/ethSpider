package chain

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func ParsingBlock(ctx context.Context, client *ethclient.Client, blockHeight uint64) (map[*big.Int]BlockMeta, error) {
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(blockHeight)))

	//log.Infof("big.NewInt(int64(blockHeight)):%v", big.NewInt(int64(blockHeight)))
	if err != nil {
		log.Errorf("----------%s\n", err.Error())
		return nil, err
	}
	//log.Infof("block.Transactions():%v", block.Transactions())

	txList, err := ParsingTransaction(block.Transactions(), client)
	if err != nil {
		return nil, err
	}

	blockData := BlockMeta{
		BlockHeight:  block.Number().String(),
		Uncles:       block.Uncles(),
		Transactions: txList,
		GasLimit:     block.GasLimit(),
		GasUsed:      block.GasUsed(),
		Difficulty:   block.Difficulty().String(),
		Time:         block.Time(),
		MixDigest:    block.MixDigest(),
		Nonce:        block.Nonce(),
		Bloom:        block.Bloom(),
		Coinbase:     block.Coinbase(),
		Root:         block.Root(),
		ParentHash:   block.ParentHash(),
		TxHash:       block.TxHash(),
		ReceiptHash:  block.ReceiptHash(),
		UncleHash:    block.UncleHash(),
		Extra:        hex.EncodeToString(block.Extra()),
		BaseFee:      block.BaseFee(),
		Header:       block.Header(),
		Body:         block.Body(),
		Size:         block.Size(),
		Hash:         block.Hash(),
		ReceivedAt:   block.ReceivedAt,
		ReceivedFrom: block.ReceivedFrom,
	}

	blockJson := map[*big.Int]BlockMeta{
		block.Number(): blockData,
	}

	return blockJson, nil
}
