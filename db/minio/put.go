package minio

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/minio/minio-go/v6"
)

const Endpoint string = "172.26.230.51:9000"

func PutJson(data []byte, height string) error {
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	minioClient, err := minio.New(Endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		fmt.Println(err)
		return err
	}
	bucketName := "lotus-state"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			//log.Printf("We already own %s\n", bucketName)
		} else {
			return err
		}
	}
	// Upload the zip file
	contentType := "application/octet-stream"
	param := Gzip(data)
	buf := bytes.NewBuffer(param)
	filename := fmt.Sprintf("ETH%s.json", height)

	n, err := minioClient.PutObject(bucketName, filename, buf, int64(len(param)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Info(fmt.Sprintf("upload %s size %d", filename, n))
	return nil
}
