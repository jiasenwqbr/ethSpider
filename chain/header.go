package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func ParsingBlockHeader(ctx context.Context, client *ethclient.Client, blockHeight uint64) (*types.Header, error) {
	headers, err := client.HeaderByNumber(ctx, big.NewInt(int64(blockHeight)))
	if err != nil {
		return nil, err
	}

	return headers, nil
}
