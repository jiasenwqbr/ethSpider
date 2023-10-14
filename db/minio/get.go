package minio

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"github.com/pkg/errors"
	"io"
)

func GetJson(height string) ([]byte, error) {
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	//useSSL := true
	// Initialize minio client object.
	minioClient, err := minio.New(Endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		return nil, err
	}
	// Make a new bucket called mymusic.
	bucketName := "lotus-state"
	filename := fmt.Sprintf("ETH%s.json", height)

	objectCh := minioClient.ListObjects(bucketName, filename, false, nil)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return nil, object.Err
		}
		if object.Key == filename {
			o, err := minioClient.GetObject(bucketName, filename, minio.GetObjectOptions{})
			if err != nil {
				return nil, err
			}
			stat, err := o.Stat()
			if err != nil {
				return nil, err
			}
			data := make([]byte, stat.Size)
			_, err = o.Read(data)
			if err != nil && err != io.EOF {
				return nil, err
			}
			return data, nil
		}
	}
	return nil, errors.New("can't get height.json")
}
