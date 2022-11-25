package main

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	endpoint        string = "192.168.2.239:9000"
	accessKeyID     string = "Minio"
	secretAccessKey string = "minio123456"
	useSSL          bool   = false
)

var (
	client *minio.Client
	err    error
)

func main() {
	client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln("minio连接错误: ", err)
	}
	log.Printf("连接成功 %#v\n", client)
	err := createBucket()
	if err != nil {
		return
	}
	FileUploader()
}
