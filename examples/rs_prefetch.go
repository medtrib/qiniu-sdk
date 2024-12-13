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
	mac := auth.New(accessKey, secretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	key := "qiniu.png"
	err := bucketManager.Prefetch(bucket, key)
	if err != nil {
		fmt.Println("fetch error,", err)
	}
}
