package minio

import (
	"assemble/config"
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var putObjectOptions = minio.PutObjectOptions{ContentType: "application/octet-stream"}

const BucketName = "easm"

func GetFile(remotePath string) ([]byte, error) {
	client, err := minioClient()
	if err != nil {
		return nil, err
	}
	object, err := client.GetObject(context.Background(), BucketName, remotePath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return io.ReadAll(object)
}

func SaveFile(filename string, buf []byte, size int) error {
	client, err := minioClient()
	if err != nil {
		return err
	}
	_, err = client.PutObject(context.Background(), BucketName, filename, bytes.NewReader(buf), int64(size), putObjectOptions)
	return err
}

func SaveFiles(filename string, file multipart.File, size int64) error {
	client, err := minioClient()
	if err != nil {
		return err
	}

	_, err = client.PutObject(context.Background(), BucketName, filename, file, size, putObjectOptions)
	return err
}

func DelFile(filename string) error {
	client, err := minioClient()
	if err != nil {
		return err
	}
	return client.RemoveObject(context.Background(), BucketName, filename, minio.RemoveObjectOptions{})
}

func minioClient() (*minio.Client, error) {
	minioConfig := config.GetMinio()
	return minio.New(net.JoinHostPort(minioConfig.Endpoint, strconv.Itoa(minioConfig.Port)),
		&minio.Options{Creds: credentials.NewStaticV4(minioConfig.AccessKeyId, minioConfig.SecretAccessKey, "")})
}
