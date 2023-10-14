package minio

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

func Gzip(data []byte) []byte {
	var res bytes.Buffer
	gz, _ := gzip.NewWriterLevel(&res, 7)
	_, err := gz.Write(data)
	if err != nil {
		fmt.Println(err)
	} else {
		gz.Close()
	}
	return res.Bytes()
}

func ParseGzip(height string) ([]byte, error) {
	data, err := GetJson(height)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, data)
	r, err := gzip.NewReader(b)
	if err != nil {
		fmt.Println(err)
		return data, err
	} else {
		defer r.Close()
		undatas, err := ioutil.ReadAll(r)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		return undatas, nil
	}
}
