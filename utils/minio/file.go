package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"log"
	"os"
	"path"
)

// 上传文件
func FileUploader() {
	dir, _ := os.Getwd()
	filePath := path.Join(dir, "utils/minio", "pic.jpg")
	contextType := "image/jpeg"

	object, err := client.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contextType})
	if err != nil {
		log.Println("上传失败：", err)
		return
	}
	//  资源访问地址 http:127.0.0.1:9000/{bucketName}/{objectName}
	log.Printf("Successfully uploaded, path: %v", object)
}

// 删除文件
func FilesDelete() {
	ctx := context.Background()
	//删除一个文件
	_ = client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{GovernanceBypass: true})

	//批量删除文件
	objectsCh := make(chan minio.ObjectInfo)
	go func() {
		defer close(objectsCh)
		options := minio.ListObjectsOptions{Prefix: "test", Recursive: true}
		for object := range client.ListObjects(ctx, bucketName, options) {
			if object.Err != nil {
				log.Println(object.Err)
			}
			objectsCh <- object
		}
	}()
	client.RemoveObjects(ctx, objectName, objectsCh, minio.RemoveObjectsOptions{})
}
