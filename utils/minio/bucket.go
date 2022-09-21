package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
)

var bucketName = "picture"
var objectName = "pic.jpg"

// 创建桶
func createBucket() error {
	ctx := context.Background()
	exists, _ := client.BucketExists(ctx, bucketName)
	if exists {
		log.Printf("bucket: %s已经存在", bucketName)
		return err
	}
	if err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "cn-south-1", ObjectLocking: false}); err != nil {
		log.Println("创建bucket错误: ", err)
		return err
	} else {
		log.Printf("Successfully created %s\n", bucketName)
		return nil
	}
}

// 查看桶
func listBucket() {
	buckets, _ := client.ListBuckets(context.Background())
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}
