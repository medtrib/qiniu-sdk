// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 设置源站镜像回源
package set_bucket_image

import credentials "github.com/sulwan/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	Bucket      string                          // 存储空间名称
	Url         string                          // 回源 URL
	Host        string                          // 从指定源站下载数据时使用的 Host
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct{}
