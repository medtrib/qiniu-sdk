package main

import (
	"fmt"
	"os"

	"github.com/medtrib/qiniu-sdk/v7/auth"
	"github.com/medtrib/qiniu-sdk/v7/storage"
)

var (
	accessKey = os.Getenv("QINIU_ACCESS_KEY")
	secretKey = os.Getenv("QINIU_SECRET_KEY")
	bucket    = os.Getenv("QINIU_TEST_BUCKET")
)

func main() {
	cfg := storage.Config{}
	mac := auth.New(accessKey, secretKey)
	bucketManger := storage.NewBucketManager(mac, &cfg)
	siteURL := "http://devtools.qiniu.com"

	// 设置镜像存储
	err := bucketManger.SetImage(siteURL, bucket)
	if err != nil {
		fmt.Println(err)
	}

	// 取消设置镜像存储
	err = bucketManger.UnsetImage(bucket)
	if err != nil {
		fmt.Println(err)
	}

}
