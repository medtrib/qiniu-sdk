// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 设置存储空间的 cache-control: max-age 响应头
package set_bucket_max_age

import credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	Bucket      string                          // 空间名称
	MaxAge      int64                           // maxAge 为 0 或者负数表示为默认值（31536000）
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct{}
