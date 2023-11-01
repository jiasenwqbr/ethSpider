package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"math"
	"math/big"
	"strings"
)

// IntToBytes 整型转byte数组
func IntToBytes(n int) ([]byte, error) {
	data := int64(n)
	byteBuff := bytes.NewBuffer([]byte{})

	err := binary.Write(byteBuff, binary.BigEndian, data)
	if err != nil {
		return nil, err
	}

	return byteBuff.Bytes(), nil
}

// BytesToInt 数字byte数组转整型
func BytesToInt(bys []byte) (int, error) {
	byteBuff := bytes.NewBuffer(bys)
	var num int64
	err := binary.Read(byteBuff, binary.BigEndian, &num)
	if err != nil {
		return 0, err
	}

	return int(num), nil
}

func NewBytes32ID() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func NewSecret(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.Errorf("generating secret failed:%v", err)
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// Uint64ConvertBytes uint64转bytes数组
func Uint64ConvertBytes(num uint64) []byte {
	byteArr := make([]byte, 8)
	binary.BigEndian.PutUint64(byteArr, num)

	return byteArr
}

// BytesConvertUint64 bytes数组转成uint64
func BytesConvertUint64(byteNum []byte) uint64 {
	num := binary.BigEndian.Uint64(byteNum)

	return num
}

// CleanPrecisionEighteenInt 大整数去除精度并返回浮点数(精度18)
func CleanPrecisionEighteenInt(num *big.Int) *big.Float {
	n := big.NewFloat(0).SetInt(num)
	amountBigFloat := new(big.Float).Quo(n, big.NewFloat(math.Pow10(18)))

	return amountBigFloat
}

// CleanPrecisionNineInt 大整数去除精度并返回浮点数(精度9)
func CleanPrecisionNineInt(num *big.Int) *big.Float {
	n := big.NewFloat(0).SetInt(num)
	amountBigFloat := new(big.Float).Quo(n, big.NewFloat(math.Pow10(9)))

	return amountBigFloat
}

// CleanPrecisionEighteenFloat 浮点数去除精度并返回浮点数(精度18)
func CleanPrecisionEighteenFloat(num *big.Float) *big.Float {
	amountBigFloat := new(big.Float).Quo(num, big.NewFloat(math.Pow10(18)))

	return amountBigFloat
}

// CleanPrecisionNineFloat 浮点数去除精度并返回浮点数(精度9)
func CleanPrecisionNineFloat(num *big.Float) *big.Float {
	amountBigFloat := new(big.Float).Quo(num, big.NewFloat(math.Pow10(9)))

	return amountBigFloat
}

func StrToBigFloat(num string) *big.Float {
	newNum, status := big.NewFloat(0).SetString(num)
	if !status {
		return big.NewFloat(0)
	}

	return newNum
}

func StrToBigInt(num string) *big.Int {
	newNum, status := big.NewInt(0).SetString(num, 10)
	if !status {
		return big.NewInt(0)
	}

	return newNum
}

func StrToCommonAddress(addr string) common.Address {
	newAddr := []byte(addr)

	return common.BytesToAddress(newAddr)
}
