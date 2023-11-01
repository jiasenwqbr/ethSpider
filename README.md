# Ethereum Block Field Summary
## Core technologies and components
golang 1.18
## Run
go run main.go node start
## Block layer
### data structure
```bash
type BlockMeta struct {
BlockHeight  string             `json:"blockHeight"`
Uncles       []*types.Header    `json:"uncles"`
Transactions []TransactionMeta  `json:"transactions"`
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
```

### Field Definition
```bash
字段	类型	解释	备注
blockHeight	string	区块高度
uncles	[]*types.Header 	叔块区块header
transactions	[]*TransactionMeta	Transaction元数据数组	TransactionMeta结构体
gasLimit	uint64	gas最大使用量
gasUsed	uint64	已使用gas数量
difficulty	string	挖矿难度
time	uint64	区块生成时间戳
mixDigest	common.Hash		以太坊标准类型
nonce	uint64		
bloom	types.Bloom		以太坊标准类型
coinbase	common.Address		以太坊标准类型
root	common.Hash		以太坊标准类型
parentHash	common.Hash	父块hash	以太坊标准类型
txHash	common.Hash		以太坊标准类型
receiptHash	common.Hash	收据hash	以太坊标准类型
uncleHash	common.Hash	叔块Hash	以太坊标准类型
extra	string		
baseFee	*big.Int		
header	*types.Header	区块头	以太坊标准类型
body	*types.Body	区块体	以太坊标准类型
size	common.StorageSize		以太坊标准类型
Ahash	common.Hash	区块hash	以太坊标准类型
receivedAt	time.Time	时间戳
receivedFrom	interface{}		
Transaction层
```

## Transaction layer

### data structure
```bash
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
Logs       []LogMeta          `json:"logs"`
Receipt    ReceiptMeta        `json:"receipt"`
}

```
### Field Definition
```azure
字段	类型	解释	备注
nonce	uint64		
value	string	代币数量	
to	string	地址	
accessList	types.AccessList		以太坊标准类型
type	uint8		
hash	string		
size	common.StorageSize		以太坊标准类型
data	string		
gasTipCap	string		
gas	uint64		
gasFeeCap	string		
gasPrice	string		
chainId	string	主网ID	
asMessage	MessageMeta	message元数据	messageMeta结构体
Cost	string		
Protected	bool		
Logs	[]*LogMeta	Log元数据数组	LogMeta结构体
Receipt	ReceiptMeta	Receipt元数据	ReceiptMeta结构体
```
## Receipt layer
### data structure
```bash
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
```
### Field Definition
```azure
字段	类型	解释	备注
type	uint8		
bloom	types.Bloom 		以太坊标准类型
txHash	common.Hash	事务hash	以太坊标准类型
gasUsed	uint64	已使用gas数量	
blockHash	common.Hash	区块hash	以太坊标准类型
blockNumber	*big.Int	区块高度	
status	uint64		
contractAddress	common.Address	合约地址	以太坊标准类型
cumulativeGasUsed	uint64		
postState	string		
transactionIndex	uint		
```

## Message layer
###  data structure
```base
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
```
### Field Definition
```azure
字段	类型	解释	备注
to	string	到达地址	
from	string	发送地址	
nonce	uint64		
amount	string		以太坊标准类型
gasLimit	uint64	gas使用限制	
gasPrice	string	gas价格	
gasFeeCap	string		
gasTipCap	string		
data	string		
accessList	types.AccessList		以太坊标准类型
isFake	bool
```

