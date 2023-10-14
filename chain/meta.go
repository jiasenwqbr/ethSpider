package chain

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

type BlockMeta struct {
	BlockHeight  string             `json:"blockHeight"`
	Uncles       []*types.Header    `json:"uncles"`
	Transactions []*TransactionMeta `json:"transactions"`
	GasLimit     uint64             `json:"gasLimit"`
	GasUsed      uint64             `json:"gasUsed"`
	Difficulty   string             `json:"difficulty"`
	Time         uint64             `json:"time"`
	MixDigest    common.Hash        `json:"mixDigest"`
	Nonce        uint64             `json:"nonce"`
	Bloom        types.Bloom        `json:"bloom"`
	Coinbase     common.Address     `json:"coinbase"`
	Root         common.Hash        `json:"root"`
	ParentHash   common.Hash        `json:"parentHash"`
	TxHash       common.Hash        `json:"txHash"`
	ReceiptHash  common.Hash        `json:"receiptHash"`
	UncleHash    common.Hash        `json:"uncleHash"`
	Extra        string             `json:"extra"`
	BaseFee      *big.Int           `json:"baseFee"`
	Header       *types.Header      `json:"header"`
	Body         *types.Body        `json:"body"`
	Size         common.StorageSize `json:"size"`
	Hash         common.Hash        `json:"hash"`
	ReceivedAt   time.Time          `json:"receivedAt"`
	ReceivedFrom interface{}        `json:"receivedFrom"`
}

type TransactionMeta struct {
	Nonce      uint64             `json:"nonce"`
	Value      string             `json:"value"`
	To         string             `json:"to"`
	AccessList types.AccessList   `json:"accessList"`
	Type       uint8              `json:"type"`
	Hash       string             `json:"hash"`
	Size       common.StorageSize `json:"size"`
	Data       string             `json:"data"`
	GasTipCap  string             `json:"gasTipCap"`
	Gas        uint64             `json:"gas"`
	GasFeeCap  string             `json:"gasFeeCap"`
	GasPrice   string             `json:"gasPrice"`
	ChainId    string             `json:"chainId"`
	AsMessage  MessageMeta        `json:"asMessage"`
	Cost       string             `json:"cost"`
	Protected  bool               `json:"protected"`
	Logs       []*LogMeta         `json:"logs"`
	Receipt    ReceiptMeta        `json:"receipt"`
}

type ReceiptMeta struct {
	Type              uint8          `json:"type"`
	Bloom             types.Bloom    `json:"bloom"`
	TxHash            common.Hash    `json:"txHash"`
	GasUsed           uint64         `json:"gasUsed"`
	BlockHash         common.Hash    `json:"blockHash"`
	BlockNumber       *big.Int       `json:"blockNumber"`
	Status            uint64         `json:"status"`
	ContractAddress   common.Address `json:"contractAddress"`
	CumulativeGasUsed uint64         `json:"cumulativeGasUsed"`
	PostState         string         `json:"postState"`
	TransactionIndex  uint           `json:"transactionIndex"`
}

type MessageMeta struct {
	To         string           `json:"to"`
	From       string           `json:"from"`
	Nonce      uint64           `json:"nonce"`
	Amount     string           `json:"amount"`
	GasLimit   uint64           `json:"gasLimit"`
	GasPrice   string           `json:"gasPrice"`
	GasFeeCap  string           `json:"gasFeeCap"`
	GasTipCap  string           `json:"gasTipCap"`
	Data       string           `json:"data"`
	AccessList types.AccessList `json:"accessList"`
	IsFake     bool             `json:"isFake"`
}

type LogMeta struct {
	Data        string         `json:"data"`
	Topics      []common.Hash  `json:"topics"`
	TxHash      common.Hash    `json:"txHash"`
	BlockNumber uint64         `json:"blockNumber"`
	BlockHash   common.Hash    `json:"blockHash"`
	Address     common.Address `json:"address"`
	Index       uint           `json:"index"`
	Removed     bool           `json:"removed"`
	TxIndex     uint           `json:"txIndex"`
}

type ERCToken struct {
	//转账人
	From common.Address `json:"from"`
	//接收人
	To common.Address `json:"to"`
	//交易代币金额
	Value string `json:"value"`
	//创建合约的地址 from
	ContractAddress string `json:"creator"`
	//代币交易哈希
	Hash common.Hash `json:"hash"`
	//当前区块from的余额
	FromBalance string `json:"fromBalance"`
	//当前区块to的余额
	ToBalance string `json:"toBalance"`
	//代币名全称
	Name string `json:"name"`
	//代币缩写
	Symbol string `json:"symbol"`

	//合约地址
	Address        common.Address `json:"address"`
	BlockHash      common.Hash    `json:"blockHash"`
	BlockNumber    *big.Int       `json:"blockNumber"`
	Decimals       uint8          `json:"decimals"`
	ERCType        string         `json:"ercType"`
	Time           uint64         `json:"timeStamp"`
	AddressBalance string         `json:"addressBalance"`
}

type Address struct {
	Address    common.Address `json:"address"`
	Balance    string         `json:"balance"`
	Code       abi.ABI        `json:"contractCode"`
	Nonce      uint64         `json:"nonce"`
	Number     *big.Int       `json:"blockHeight"`
	Time       uint64         `json:"timeStamp"`
	Storagekey string         `json:"stotageKey"`
}
