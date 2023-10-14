package chain

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func ParsingBlockBody(ctx context.Context, client *ethclient.Client, blockHeight uint64) error {
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(blockHeight)))
	if err != nil {
		return err
	}

	body := block.Body()
	txs := body.Transactions

	for i := 0; i < len(txs); i++ {
		fmt.Printf("Type:%v\n", txs[i].Type())
		fmt.Printf("Inner Data:%v\n", hex.EncodeToString(txs[i].Data()))
	}

	return nil
}
